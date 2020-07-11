package main

type matrix struct {
	rows int
	cols int
	data []array
}

func newMatrix(rows, cols int) matrix {
	output := matrix{
		rows: rows,
		cols: cols,
		data: make([]array, rows),
	}
	for i := range output.data {
		output.data[i] = newArray(cols)
	}
	return output
}

func copyMatrix(m1 matrix) matrix {
	output := newMatrix(m1.rows, m1.cols)
	for i := 0; i < m1.rows; i++ {
		for j := 0; j < m1.cols; j++ {
			output.data[i].data[j] = m1.data[i].data[j]
		}
	}
	return output
}

func randomMatrix(rows, cols int, min, max float64) matrix {
	output := matrix{
		rows: rows,
		cols: cols,
		data: make([]array, rows),
	}
	for i := range output.data {
		output.data[i] = randomArray(cols, min, max)
	}
	return output
}

func multiplyMatrix(m1, m2 matrix) matrix {
	output := newMatrix(m1.rows, m2.cols)
	m3 := m2.transpose()
	for x := 0; x < output.rows; x++ {
		for y := 0; y < output.cols; y++ {
			output.data[x].data[y] = multiplyArray(m1.data[x], m3.data[y])
		}
	}
	return output
}

func multiplyMatrixElem(m1, m2 matrix) matrix {
	output := newMatrix(m1.rows, m2.cols)
	for x := 0; x < output.rows; x++ {
		for y := 0; y < output.cols; y++ {
			output.data[x].data[y] = m1.data[x].data[y] * m2.data[x].data[y]
		}
	}
	return output
}

func (m1 *matrix) transpose() matrix {
	output := newMatrix(m1.cols, m1.rows)
	for x := 0; x < output.rows; x++ {
		for y := 0; y < output.cols; y++ {
			output.data[x].data[y] = m1.data[y].data[x]
		}
	}
	return output
}

func addMatrix(m1, m2 matrix) matrix {
	output := newMatrix(m1.rows, m1.cols)
	for x := 0; x < m1.rows; x++ {
		for y := 0; y < m1.cols; y++ {
			output.data[x].data[y] = m1.data[x].data[y] + m2.data[x].data[y]
		}
	}
	return output
}

func subMatrix(m1, m2 matrix) matrix {
	output := newMatrix(m1.rows, m1.cols)
	for x := 0; x < m1.rows; x++ {
		for y := 0; y < m1.cols; y++ {
			output.data[x].data[y] = m1.data[x].data[y] - m2.data[x].data[y]
		}
	}
	return output
}

func (m1 *matrix) diff(m2 matrix) {
	for x := 0; x < m1.rows; x++ {
		for y := 0; y < m1.cols; y++ {
			m1.data[x].data[y] = (m1.data[x].data[y] - m2.data[x].data[y]) * (m1.data[x].data[y] - m2.data[x].data[y])
		}
	}
}

func (m1 *matrix) div(value float64) {
	for x := 0; x < m1.rows; x++ {
		for y := 0; y < m1.cols; y++ {
			m1.data[x].data[y] = m1.data[x].data[y] / value
		}
	}
}

func (m1 *matrix) mult(value float64) {
	for x := 0; x < m1.rows; x++ {
		for y := 0; y < m1.cols; y++ {
			m1.data[x].data[y] = m1.data[x].data[y] * value
		}
	}
}

func (m1 *matrix) print() {
	for _, a := range m1.data {
		a.print()
	}
}
