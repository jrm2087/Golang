package main

import "fmt"

func main() {
	//colors := map[string]string{
	//	"red":   "#FF0000",
	//	"green": "#1D8348",
	//}

	//var colors map[string]string

	colors := make(map[int]string)

	colors[10] = "#FFFFFF"

	delete(colors, 10)

	fmt.Println(colors)
}
