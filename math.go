package main
import "fmt"

func main() {
	fmt.Println(Soma(10, 222))
}

// Soma é uma função que soma dois números inteiros
func Soma(a int, b int) int {
	return a + b
}