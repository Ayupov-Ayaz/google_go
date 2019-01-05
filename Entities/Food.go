package Entities
import(
	"fmt"
)
type Food struct {
	Name string
}

func (this *Food) Eating() {
	fmt.Printf("(type Food) I eating %s \n", this.Name)
}