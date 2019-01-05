package Entities
import "fmt"

type Duck struct{}

func (this *Duck) Swim(){
	fmt.Println("Плавать")
}
func (this *Duck) Fly(){
	fmt.Println("Летать")
}
func (this *Duck) Quack(){
	fmt.Println("Крякать")
}
