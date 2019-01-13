package language

import (
	"sync"
)
/**
	Основная проблема использования Mutex-ов возникает, если у нас попадется рекурсивная блокировка,
	Например: Выполнение кода без неэкспортируемых копий функций приведет к дедлоку,
	так как метод Get() заблокируется и метод Count() также заблокируется перед тем как Get()
	разблокирует мьютекс

    Обратите внимание в этом коде, что для каждого не-экспортированного метода есть аналогичный
	экспортированный. Эти методы работают как публичный API, и заботятся о блокировках на этом уровне.
	Далее они вызывают неэкспортированные методы, которые вообще не заботятся о блокировках.
	Это гарантирует, что все вызовы ваших методов извне будут блокироваться лишь раз и лишены проблемы
	рекурсивной блокировки.
 */
type DataStore struct {
	sync.Mutex // мы должны указывать Mutex перед элементом который он охраняет
	cache map[string]string
}

func NewDataStore() *DataStore{
	return &DataStore{
		cache: make(map[string]string),
	}
}

func (ds *DataStore) set(key, value string) {
	ds.cache[key] = value
}

func (ds *DataStore) get(key string) string {
	return ds.cache[key]
}

func (ds *DataStore) count() int {
	return len(ds.cache)
}

func (ds *DataStore) Set(key, value string) {
	ds.Lock()
	defer ds.Unlock()
	ds.set(key, value)
}

func (ds *DataStore) Get(key string) string{
	ds.Lock()
	defer ds.Unlock()
	return ds.get(key)
}

func (ds *DataStore) Count() int {
	ds.Lock()
	defer ds.Unlock()
	return ds.count()
}
