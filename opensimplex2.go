package opensimplex2

import "math"

type Noise interface {
	Noise2D(x, y float64) float64
}

type noise struct {
	perm []int16
}

func NewNoise(seed int64) Noise {
	n := &noise{perm: make([]int16, 256)}

	var source [256]int16

	for i := range source {
		source[i] = int16(i)
	}

	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407
	seed = seed*6364136223846793005 + 1442695040888963407

	var i int32
	for i = 255; i >= 0; i-- {
		seed = seed*6364136223846793005 + 1442695040888963407
		r := int32((seed + 31) % int64(i+1))
		if r < 0 {
			r += i + 1
		}

		n.perm[i] = source[r]
		source[r] = source[i]
	}

	return n
}

func (n *noise) Noise2D(x, y float64) float64 {

	var noisyValue float64

	stretchOffset := (x + y) * StretchConstant2
	xs := x + stretchOffset
	ys := y + stretchOffset

	xsb := math.Floor(xs)
	ysb := math.Floor(ys)

	squishOffset := (xsb + ysb) * SquishConstant2
	xb := xsb + squishOffset
	yb := ysb + squishOffset

	xins := xs - xsb
	yins := ys - ysb

	inSum := xins + yins

	dx0 := x - xb
	dy0 := y - yb

	var value float64 = 0.0

	dx1 := dx0 - 1 - SquishConstant2
	dy1 := dy0 - 0 - SquishConstant2
	attn1 := 2 - dx1*dx1 - dy1*dy1
	if attn1 > 0 {
		attn1 *= attn1
		value += attn1 * attn1 * extrapolate2(n.perm[:], xsb+1, ysb+0, dx1, dy1)
	}

	dx2 := dx0 - 0 - SquishConstant2
	dy2 := dy0 - 1 - SquishConstant2
	attn2 := 2 - dx2*dx2 - dy2*dy2
	if attn2 > 0 {
		attn2 *= attn2
		value += attn2 * attn2 * extrapolate2(n.perm[:], xsb+0, ysb+1, dx2, dy2)
	}

	var xsvExt, ysvExt, dxExt, dyExt float64
	if inSum <= 1 {
		zins := 1 - inSum
		if zins > xins || zins > yins {
			if xins > yins {
				xsvExt = xsb + 1
				ysvExt = ysb - 1
				dxExt = dx0 - 1
				dyExt = dy0 + 1
			} else {
				xsvExt = xsb - 1
				ysvExt = ysb + 1
				dxExt = dx0 + 1
				dyExt = dy0 - 1
			}
		} else {
			xsvExt = xb + 1
			ysvExt = ysb + 1
			dxExt = dx0 - 1 - 2*SquishConstant2
			dyExt = dy0 - 1 - 2*SquishConstant2
		}
	} else {
		zins := 2 - inSum
		if zins < xins || zins < yins {
			if xins > yins {
				xsvExt = xsb + 2
				ysvExt = ysb + 0
				dxExt = dx0 - 2 - 2*SquishConstant2
				dyExt = dy0 + 0 - 2*SquishConstant2
			} else {
				xsvExt = xsb + 0
				ysvExt = ysb + 2
				dxExt = dx0 + 0 - 2*SquishConstant2
				dyExt = dy0 - 2 - 2*SquishConstant2
			}
		} else {
			xsvExt = xsb
			ysvExt = ysb
			dxExt = dx0
			dyExt = dy0
		}

		xsb += 1
		ysb += 1
		dx0 = dx0 - 1 - 2*SquishConstant2
		dy0 = dy0 - 1 - 2*SquishConstant2
	}

	attn0 := 2 - dx0*dx0 - dy0*dy0
	if attn0 > 0 {
		attn0 *= attn0
		value += attn0 * attn0 * extrapolate2(n.perm[:], xsb, ysb, dx0, dy0)
	}

	attnExt := 2 - dxExt*dxExt - dyExt*dyExt
	if attnExt > 0 {
		attnExt *= attnExt
		value += attnExt * attnExt * extrapolate2(n.perm[:], xsvExt, ysvExt, dxExt, dyExt)
	}

	noisyValue = value / NormConstant2

	return noisyValue
}

func extrapolate2(perm []int16, xsb, ysb, dx, dy float64) float64 {
	index := perm[(int32(perm[int32(xsb)&0xff])+int32(ysb))&0xff] & 0x0e
	g1 := float64(Gradients2[index])
	g2 := float64(Gradients2[index+1])

	return g1*dx + g2*dy
}
