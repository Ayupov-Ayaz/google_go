package language

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	// Если данное поле будет пустым, то в json не попадет дефолтное значение, то есть тоже будет пустым
	ID int `json:"id, omitemty"`
	Username string
	Phone string
}

var jsonString = `{"id":13, "username":"tommy", "phone":"+7898898898"}`

func StartJsonWithStructs() {
	// Для того, что бы распоковать json нам необходимо его преобразовать в slice byte
	data := []byte(jsonString)
	// Создаем структуру в которую будем распаковывать
	user := User{}
	// за распаковку отвечает функция json.Unmarshal
	err := json.Unmarshal(data, &user)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(user.ID, user.Username, user.Phone)
	user.Username = "Ray"

	result, _ := json.Marshal(&user)
	fmt.Printf("json string: %s \n", result)
}

/**
 	При преобразовании json где у нас несколько структур можно воспользоваться
	пустым интерфейсов
 */
var jsonDynamicData = `[
	{"username":"Axcel", "phone":"+7897898723"},
	{"id":11, "address":"none", "company":"Nefis Group"}
]`

func StartJsonWithDynamicData() {
	var i interface{}
	err := json.Unmarshal([]byte(jsonDynamicData),  &i)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Распаковка в slice пустых интерфейсов \n %#v", i)
	// как извлекать дальше?
}

func handlerFunc(rw http.ResponseWriter, r * http.Request) {
	// отправка
	var user = User{
		Username: "Alex",
		Phone: "+78989898989",
	}
	json.NewEncoder(rw).Encode(user) // отправит сразу же ответ в формате json

}