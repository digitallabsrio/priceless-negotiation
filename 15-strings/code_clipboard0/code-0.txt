package main
import "fmt"

func main(){
  s := "Hel" + "lo "
  s += "World!"
  fmt.Println(s) // prints out “Hello World!”
}