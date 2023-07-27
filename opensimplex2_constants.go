package opensimplex2

const (
	StretchConstant2 float64 = -0.21132486540518708
	SquishConstant2  float64 = 0.3660254037844386
	NormConstant2    float64 = 47
)

var Gradients2 [16]int8 = [16]int8{
	5, 2, 2, 5,
	-5, 2, -2, 5,
	5, -2, 2, -5,
	-5, -2, -2, -5,
}
