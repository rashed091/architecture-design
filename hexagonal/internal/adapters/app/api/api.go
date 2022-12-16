package api

import (
	"github.com/rashed091/architecture-design/hexagonal/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPort
}

func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func	(apia Adapter) GetAddtion(lhs, rhs int) (int, error) {
	result, err := apia.arith.Addition(lhs, rhs)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func	(apia Adapter) GetSubtraction(lhs, rhs int) (int, error) {
	result, err := apia.arith.Subtraction(lhs, rhs)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func	(apia Adapter) GetMultiplecation(lhs, rhs int) (int, error) {
	result, err := apia.arith.Multiplication(lhs, rhs)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func	(apia Adapter) GetDivision(lhs, rhs int) (int, error) {
	result, err := apia.arith.Division(lhs, rhs)
	if err != nil {
		return 0, err
	}
	return result, nil
}
