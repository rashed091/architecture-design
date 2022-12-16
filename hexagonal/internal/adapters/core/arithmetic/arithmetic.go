package arithmetic

type Adapter struct {

}

func	NewAdapter() * Adapter {
	return &Adapter{};
}

func (arith Adapter) Addition(lhs int, rhs int) (int, error) {
	return lhs + rhs, nil
}

func (arith Adapter) Subtraction(lhs int, rhs int) (int, error) {
	return lhs - rhs, nil
}

func (arith Adapter) Multiplication(lhs int, rhs int) (int, error) {
	return lhs * rhs, nil
}

func (arith Adapter) Division(lhs int, rhs int) (int, error) {
	return lhs / rhs, nil
}
