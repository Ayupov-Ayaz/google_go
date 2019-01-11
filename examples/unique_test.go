package examples

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)
// значения для тестирования
var testOk = `1
2
3
4
4
4
5`
// значения для сравнения с результатом
var testOkResult = `1
2
3
4
5
`
// Проверка уникальности
func TestUnique(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := Uniq(in, out)
	if err != nil {
		t.Errorf("Test for unique Failed")
	}
	result := out.String()
	if result != testOkResult {
		t.Errorf("Test for unique Failed: results not match \n [your:] \n %v \n need: \n%v ",result, testOkResult )
	}
}

var testFail = `1
2
1`
// проверка на реакцию программы при получении неотсортированного ввода
func TestSorted(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := Uniq(in, out)
	if err == nil {
		t.Errorf("Test for sorted Failed")
	}

}