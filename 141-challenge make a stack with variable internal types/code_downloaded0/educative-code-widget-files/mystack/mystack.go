package mystack
import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return 0
}

func (stack Stack) Cap() int {
	return 0
}

func (stack Stack) IsEmpty() bool {
	return true
}

func (stack *Stack) Push(e interface{}) {
	return
}

func (stack Stack) Top() (interface{}, error) {
	return nil, nil
}

func (stack *Stack) Pop() (interface{}, error) {
	return nil, nil
}