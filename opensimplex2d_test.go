package opensimplex2d

import (
	"math"
	"testing"

	//"github.com/ojrac/opensimplex-go"
)

func TestConstants(t *testing.T) {
	trueStretch := (1.0/math.Sqrt(2.0+1.0) - 1.0) / 2.0
	trueSquish := (math.Sqrt(2.0+1.0) - 1.0) / 2.0
	tolerance := math.Pow10(-14)

	if math.Abs(StretchConstant2-trueStretch) > tolerance {
		t.Errorf("Stretch constant incorrect; got %f, expected %f",
			StretchConstant2, trueStretch)
	}

	if math.Abs(SquishConstant2-trueSquish) > tolerance {
		t.Errorf("Squish constant incorrect; got %f, expected %f",
			SquishConstant2, trueSquish)
	}
}

//func TestNoise2D(t *testing.T) {
//	var seed int64 = 298610
//	x, y := 14.0, 12.0
//	tolerance := math.Pow10(-14)
//
//	testNoise := NewNoise(seed)
//	testValue := testNoise.Noise2D(x, y)
//
//	libNoise := opensimplex.New(seed)
//	libValue := libNoise.Eval2(x, y)
//
//	if math.Abs(testValue-libValue) > tolerance {
//		t.Errorf("Inconsistent noise result; got %f, expected %f",
//			testValue, libValue)
//	}
//
//}
