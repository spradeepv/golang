package main

import (
	"fmt"
	"github.com/spradeepv/golang/json"
)

func main() {
	// This compact representation contains all the information but it’s hard to read
	fmt.Println("Hard to read...")
	json.Marshal()

	fmt.Println()

	// json.MarshalIndent produces neatly indented output
	fmt.Println("Neatly indented output...")
	json.MarshalIndent()

	json.Unmarshal()
}
