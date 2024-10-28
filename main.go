package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SwanHtetAungPhyo/interpreter/internal/environment"
)

func main() {
	fmt.Println("Tree work Interpreter")
	fmt.Println("Type 'exit' to quit")

	interpreter := environment.NewInterpreter()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\033[35m>> \033[0m")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		result, err := interpreter.Interpret(input)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("Result: %v\n", result)
		}
	}
}
