package main
import (
"bufio"
"fmt"
"os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Please enters some input: ")
  if scanner.Scan() {
    fmt.Println("The input was", scanner.Text())
  }
}