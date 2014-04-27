package renderer

type Matrix4 [16]float32

var identity = Matrix4{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
}

func (m *Matrix4) Ortho(left, right, bottom, top, near, far float32) {
	dx := left - right
	dy := bottom - top
	dz := near - far

	*m = Matrix4{
		-2 / dx, 0, 0, 0,
		0, -2 / dy, 0, 0,
		0, 0, 2 / dz, 0,
		(left + right) / dx, (top + bottom) / dy, (far + near) / dz, 1,
	}
}
