package main
import "fmt"

func main(){
    b3 := 10 > 5 && 7 < 15     // AND operator
    fmt.Println(b3)
    b3 = 10 < 5 || 2 > 7       // OR operator
    fmt.Println(b3)
    b3 = !b3                   // NOT operator
    fmt.Println(b3)
}