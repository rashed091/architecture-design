package ports

type APIPort interface {
	GetAddtion(lhs, rhs int) (int, error)
	GetSubtraction(lhs, rhs int) (int, error)
	GetMultiplecation(lhs, rhs int) (int, error)
	GetDivision(lhs, rhs int) (int, error)
}
