package main

import "fmt"

func main() {

	massege := sayhello("Сергей!")
	fmt.Println(massege)
	fmt.Println("Меня зовут Юрий.")
}
func sayhello(name string) string {
	return "Привет " + name
}
