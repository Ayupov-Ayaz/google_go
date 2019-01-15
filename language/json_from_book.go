package language

import (
	"auction/errors"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

const (
	fileVersion = 1 // Используется всеми форматами.
	dateFormat = "02.01.2006" // Дату не менять! можно менять только формат
	fileType = "INVOICES" // Используется текстовыми форматами
	magicNumber = 0x125D // Используется двоичными форматами
)

type InvoicesMarshaler interface {
	MarshalInvoices(writer io.Writer, invoices []*Invoice) error
}

type InvoicesUnmarshaler interface {
	UnmarshalInvoices(reader io.Reader) ([]*Invoice, error)
}

// Пустая структура, для реализации в ней методов интерфесов
type JSONMarshaller struct{}

// накладная
type Invoice struct {
	Id int
	CustomerId int
	Raised time.Time
	Due time.Time
	Paid bool
	Note string
	Items []*Item
}
// в каком формате данные будут парситься в JSOn
type JSONInvoice struct {
	Id int
	CustomerId int
	Raised string // Преобразование time.Time в строку
	Due string // Преобразование time.Time в строку
	Paid bool
	Note string
	Items []*Item
}

// Данные накладной
type Item struct {
	Id int
	Price float64
	Quantity int
	Note string
}

func (JSONMarshaller) MarshalInvoices(writer io.Writer, invoices []*Invoice) error {
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(fileType); err != nil {
		return err
	}
	if err := encoder.Encode(fileVersion); err != nil {
		return err
	}
	return encoder.Encode(invoices) // Запись данных
}

func (jm *JSONMarshaller) UnmarshalInvoices(reader io.Reader) ([]*Invoice, error) {
	decoder := json.NewDecoder(reader)
	var kind string
	if err := decoder.Decode(&kind); err != nil {
		return nil, err
	}
	if kind != fileType {
		return nil, errors.New("Не может прочитать файл который не является json")
	}
	var version int
	if err := decoder.Decode(&version); err != nil {
		return nil, err
	}
	if version > fileVersion {
		return nil, fmt.Errorf("Версия программы %d слишком новая для чтения", version)
	}
	var invoices []*Invoice
	err := decoder.Decode(&invoices)
	return invoices, err
}

// Преобразование Invoice в JSON
func (i Invoice) MarshalJSON() ([]byte, error) {
	jsonInvoice := JSONInvoice{
		i.Id,
		i.CustomerId,
		i.Raised.Format(dateFormat),
		i.Due.Format(dateFormat),
		i.Paid,
		i.Note,
		i.Items,
	}
	return json.Marshal(jsonInvoice)
}

func (i *Invoice) UnmarshalJSON(data []byte) (err error) {
	var jsonInvoice JSONInvoice
	if err := json.Unmarshal(data, &jsonInvoice); err != nil {
		return err
	}
	var raised, due time.Time
	if raised, err = time.Parse(dateFormat, jsonInvoice.Raised); err != nil {
		return err
	}
	if due, err = time.Parse(dateFormat, jsonInvoice.Due); err != nil {
		return err
	}
	* i = Invoice{
		jsonInvoice.Id,
		jsonInvoice.CustomerId,
		raised,
		due,
		jsonInvoice.Paid,
		jsonInvoice.Note,
		jsonInvoice.Items,
	}
	return nil

}

func JsonFromBookStart() {
	//
	items := []*Item {
		{Id:1, Price: 2344, Quantity: 233, Note: "Note for first Item" },
		{Id:2, Price: 4344, Quantity: 123, Note: "Note for second Item" },
	}
	invoice := Invoice{
		Id: 12,
		Due: time.Now(),
		Note: "Бла бла бла",
		Paid: true,
		CustomerId: 23,
		Raised: time.Now().AddDate(0,0,1),
		Items: items,
	}
	fmt.Printf("| Как выглядят накладные: | \n %v \n", invoice)
	jsonInvoice, err := invoice.MarshalJSON()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("| Те же накладные в формате json : | \n%v \n", string(jsonInvoice))
	err = invoice.UnmarshalJSON(jsonInvoice)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("| После преобразования из JSON | \n %v", invoice)
}