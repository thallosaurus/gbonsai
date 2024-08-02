package main

import "fmt"

func main() {
	i := "Hello World!"
	fmt.Println(i[3 : len(i)-5])
}
