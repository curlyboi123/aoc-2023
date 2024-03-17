package main

import (
	"fmt"

	"example.com/utils"
)

func main() {
	lines := utils.GetFileContentByLine()
	fmt.Println(lines)
}
