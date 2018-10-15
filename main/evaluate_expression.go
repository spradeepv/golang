package main

import (
	"bufio"
	"fmt"
	"github.com/spradeepv/golang/expression"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// fmt.Println("TEXT :: ", text)
	fmt.Println("Result: ", expression.Eval(text))
}
