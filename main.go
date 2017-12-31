package main

import "fmt"

func main() {
	cryptocoin := Blockchain{}
	cryptocoin.Init()

	fmt.Printf("%#v", cryptocoin)
}
