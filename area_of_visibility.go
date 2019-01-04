package main

import(
	"fmt"
)

// Области видимости переменных
var a,b,c string
// заметьте, переменная "c" не использовалась ниразу, а компилятор не ругается. 
// TODO: Выяснить почему) 

var emptyStr = "var \"%s\" is empty string\n"
var numbStr = "var \"%s\" = %d \n"

func main(){
	// переменные назначенные внутри блока if видны только в теле блока if. 
	// переменная a просто перезапишется
	if a := "var \"a\" is not empty"; a != "" {
		fmt.Println(a)
	}
	// тут мы получим уже глобальную переменную
	if a == "" {
		fmt.Printf(emptyStr, "a")
	}

	// точно так же работают переменные и внутри блока цикла
	for b := 0; b < 4; b++ {
		// более того, переменные можно назначить другого типа, он никак не будет конфликтовать
		// так как находится в поле видимости только цикла
		fmt.Printf(numbStr, "b", b)
	}

	func () {
		if(b == "") {
			fmt.Printf(emptyStr, "b")
		}  
	}()

	// можно создать отдельную область видимости и все переменные заданные внутри нее будут доступны только там
	{
		a := 256
		func() {
			fmt.Printf(numbStr, "a", a)
		}()

		{
			// еще одна область видимости
			if a == 256 {
				a := "Мы тут даже можем переопределить тип у переменной!!!! \n"
				fmt.Printf("var \"a\" = %s",a)
			}
		}
	}
	// снаружи компилятор даже не позволит проверить a == 256, так как у a тип уже string
	if a == "" {
		fmt.Printf(emptyStr, "a")
	}

}