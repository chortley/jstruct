package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chortley/jstruct/generator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter JSON: ")
	jsonInput, _ := reader.ReadString('\n')
	jsonInput = strings.TrimSpace(jsonInput)

	fmt.Print("Enter structName: ")
	structName, _ := reader.ReadString('\n')
	structName = strings.TrimSpace(structName)

	result, err := generator.GenerateStruct(jsonInput, structName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\nGenerated Struct:\n\n%s\n", result)
}

