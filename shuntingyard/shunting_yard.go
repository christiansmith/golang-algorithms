package shuntingyard

import (
	"bytes"
	"strconv"
)

// http://en.wikipedia.org/wiki/Shunting-yard_algorithm
// http://www.engr.mun.ca/~theo/Misc/exp_parsing.htm
// http://www.cs.utexas.edu/~EWD/MCReps/MR35.PDF
// http://en.wikipedia.org/wiki/Recursive_descent_parser
// http://en.wikipedia.org/wiki/Operator-precedence_parser
// http://en.literateprograms.org/Shunting_yard_algorithm_(C)

type Tokens []interface{}

func (t *Tokens) Push(token interface{}) {
	*t = append(*t, token)
}

func (t *Tokens) Pop() interface{} {
	tokens := *t
	length := len(tokens)

	if length == 0 {
		return 0
	}

	value := tokens[length-1]
	*t = tokens[:length-1]
	return value
}

func EvaluateInfix(expr string) int64 {
	tokens := bytes.NewBufferString(expr)

	var operator string
	var left, right, result int64
	var operators, operands Tokens

	token, eof := tokens.ReadByte()

	for eof == nil {
		switch string(token) {
		case "(":
			// do nothing

		case "+", "-", "*", "/":
			operators.Push(string(token))

		case ")":
			operator = operators.Pop().(string)
			left = operands.Pop().(int64)
			right = operands.Pop().(int64)

			switch string(operator) {
			case "+":
				result = left + right
			case "-":
				result = left - right
			case "*":
				result = left * right
			case "/":
				result = left / right
			}

			operands.Push(result)

		default:
			v, _ := strconv.ParseInt(string(token), 0, 64)
			operands.Push(v)
		}

		token, eof = tokens.ReadByte()
	}

	return operands.Pop().(int64)
}
