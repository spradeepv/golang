package main

import (
	"fmt"
	"errors"
	"bufio"
	"os"
	"strconv"
)

type stack []interface{}

func (s *stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *stack) Pop() (interface{}, error) {
	temp, err := s.Peek()
	if err != nil {
		return 0.0, err
	}

	*s = (*s)[0 : len(*s)-1]

	return temp, nil
}

func (s *stack) Peek() (interface{}, error) {
	if len(*s) == 0 {
		return 0.0, errors.New("stack is empty")
	}

	return (*s)[len(*s)-1], nil
}

func (s *stack) Len() int {
	return len(*s)
}

func isNumber(token string) (int, bool) {
	i, err := strconv.Atoi(token)
	if err != nil {
		return 0, false
	}
	return i, true
}

func isOperator(token string) bool {
	operatorList := [4]string{"+", "-", "*", "/"}
	for _, op := range operatorList {
		if op == token {
			return true
		}
	}
	return false
}

func precedence(op string) int {
	if op == "+" || op == "-" {
		return 1
	}
	if op == "*" || op == "/" {
		return 2
	}
	return 0
}

func eval(text string) int {
	var val_stack = &stack{}
	var operator_stack = &stack{}
	for i := 0; i < len(text); i++ {
		token := string(text[i])
		// fmt.Println("Token: ", token)

		// Skip space
		if token == " " {
			continue
		} else if _, isDigit := isNumber(token); isDigit {
			val := 0
			for (i < len(text)) {
				if v, isDigit := isNumber(string(text[i])); isDigit {
					val = (val * 10) + v
					i++
				} else {
					break
				}
			}
			val_stack.Push(val)
			i--
		} else if token == "[" {
			operator_stack.Push(token)
		} else if token == "]" {
			for {
				val, err := operator_stack.Peek()
				if err == nil {
					if val.(string) != "[" {
						op, _ := operator_stack.Pop()
						val2, _ := val_stack.Pop()
						val1, _ := val_stack.Pop()
						val_2, _ := val2.(int)
						val_1, _ := val1.(int)
						op = op.(string)
						var val int
						switch (op) {
						case "+":
							val = val_1 + val_2
						case "-":
							val = val_1 - val_2
						case "*":
							val = val_1 * val_2
						case "/":
							val = val_1 / val_2
						}
						val_stack.Push(val)
					} else {
						operator_stack.Pop()
						break
					}
				}
			}
		} else if isOperator(token) {
			for {
				if operator_stack.Len() == 0 {
					operator_stack.Push(token)
					break
				}
				val, _ := operator_stack.Peek()

				if operator_stack.Len() != 0 && (precedence(val.(string))) >= precedence(token) {
					op, _ := operator_stack.Pop()
					val2, _ := val_stack.Pop()
					val1, _ := val_stack.Pop()
					val_2, _ := val2.(int)
					val_1, _ := val1.(int)
					op = op.(string)
					var val int
					switch (op) {
					case "+":
						val = val_1 + val_2
					case "-":
						val = val_1 - val_2
					case "*":
						val = val_1 * val_2
					case "/":
						val = val_1 / val_2
					}
					val_stack.Push(val)
				} else {
					operator_stack.Push(token)
					break
				}
			}
		}
	}
	// fmt.Println("Val_Stack: ", *val_stack)
	// fmt.Println("Ops_Stack: ", *operator_stack)
	for {
		if operator_stack.Len() == 0 {
			break
		}
		op, _ := operator_stack.Pop()
		val2, _ := val_stack.Pop()
		val1, _ := val_stack.Pop()
		val_2, _ := val2.(int)
		val_1, _ := val1.(int)
		op = op.(string)
		var val int
		switch (op) {
		case "+":
			val = val_1 + val_2
		case "-":
			val = val_1- val_2
		case "*":
			val = val_1 * val_2
		case "/":
			val = val_1 / val_2
		}
		val_stack.Push(val)
	}
	rVal, _ := val_stack.Peek()
	return rVal.(int)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// fmt.Println("TEXT :: ", text)
	fmt.Println("Result: ", eval(text))
}
