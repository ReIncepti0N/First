package main

import "fmt"

func main() {

	massege := sayhello("Юрий!")
	fmt.Println(massege)
}
func sayhello(name string) string {
	return "Привет" + name
}
