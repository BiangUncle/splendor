package main

import "fmt"

func main() {
	c := ConstructClient()

	var content string
	var err error

	content, err = c.JoinGame()
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	content, err = c.Alive()
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	content, err = c.Leave()
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	content, err = c.Alive()
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

}
