package ports

type ArithmeticPort interface{
	Addition(lhs int, rhs int) (int, error)
	Subtraction(lhs int, rhs int) (int, error)
	Multiplication(lhs int, rhs int) (int, error)
	Division(lhs int, rhs int) (int, error)
}
