package main
import (
	"fmt"
)
type vertes struct{
	X int
	Y int
}
func main(){
	v := vertes{1,31}
	p := &v
	p.X = 10000000
	fmt.Println(v)
}
