package expression

import (
	"github.com/spradeepv/golang/ds"
	"strconv"
)

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

func Eval(text string) int {
	var valStack = &ds.Stack{}
	var operatorStack = &ds.Stack{}
	for i := 0; i < len(text); i++ {
		token := string(text[i])
		// Skip space
		if token == " " {
			continue
		} else if _, isDigit := isNumber(token); isDigit {
			val := 0
			for i < len(text) {
				if v, isDigit := isNumber(string(text[i])); isDigit {
					val = (val * 10) + v
					i++
				} else {
					break
				}
			}
			valStack.Push(val)
			i--
		} else if token == "[" {
			operatorStack.Push(token)
		} else if token == "]" {
			for {
				val, err := operatorStack.Peek()
				if err == nil {
					if val.(string) != "[" {
						applyOperation(operatorStack, valStack)
					} else {
						operatorStack.Pop()
						break
					}
				}
			}
		} else if isOperator(token) {
			for {
				if operatorStack.Len() == 0 {
					operatorStack.Push(token)
					break
				}
				val, _ := operatorStack.Peek()

				if operatorStack.Len() != 0 && (precedence(val.(string))) >= precedence(token) {
					applyOperation(operatorStack, valStack)
				} else {
					operatorStack.Push(token)
					break
				}
			}
		}
	}
	//fmt.Println("Val_Stack: ", *val_stack)
	//fmt.Println("Ops_Stack: ", *operator_stack)
	for {
		if operatorStack.Len() == 0 {
			break
		}
		applyOperation(operatorStack, valStack)
	}
	rVal, _ := valStack.Peek()
	return rVal.(int)
}

func applyOperation(operatorStack *ds.Stack, valStack *ds.Stack) {
	op, _ := operatorStack.Pop()
	val2Str, _ := valStack.Pop()
	val1Str, _ := valStack.Pop()
	val2, _ := val2Str.(int)
	val1, _ := val1Str.(int)
	op = op.(string)
	var val int
	switch op {
	case "+":
		val = val1 + val2
	case "-":
		val = val1 - val2
	case "*":
		val = val1 * val2
	case "/":
		val = val1 / val2
	}
	valStack.Push(val)
}
