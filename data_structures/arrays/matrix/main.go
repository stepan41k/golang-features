package main

import (
	"errors"
	"fmt"
)

type Matrix struct {
	rows int
	cols int
	data []float64
}

func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		rows: rows,
		cols: cols,
		data: make([]float64, rows*cols),
	}
}

func NewMatrixFromSlice(slice [][]float64) (*Matrix, error) {
	rows := len(slice)
	if rows == 0 {
		return nil, errors.New("matrix cant be empty")
	}

	cols := len(slice[0])
	m := NewMatrix(rows, cols)
	for r := 0; r < rows; r++ {
		if len(slice[r]) != cols {
			return nil, errors.New("different lengths")
		}
		for c := 0; c < cols; c++ {
			m.Set(r, c, slice[r][c])
		}
	}
	return m, nil
}

func (m *Matrix) At(r, c int) float64 {
	return m.data[r*m.cols+c]
}

func (m *Matrix) Set(r, c int, val float64) {
	m.data[r*m.cols+c] = val
}

func (m *Matrix) Add(other *Matrix) (*Matrix, error) {
	if m.rows != other.rows || m.cols != other.cols {
		return nil, errors.New("sizes of matrixs must be equal")
	}
	res := NewMatrix(m.rows, m.cols)
	for i := range m.data {
		res.data[i] = m.data[i] + other.data[i]
	}
	return res, nil
}

func (m *Matrix) Multiply(other *Matrix) (*Matrix, error) {
	if m.cols != other.rows {
		return nil, errors.New("cols1 must be equal cols2")
	}
	res := NewMatrix(m.rows, other.cols)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < other.cols; c++ {
			var sum float64
			for k := 0; k < m.cols; k++ {
				sum += m.At(r, k) * other.At(k, c)
			}
			res.Set(r, c, sum)
		}
	}
	return res, nil
}

func (m *Matrix) Transpose() *Matrix {
	res := NewMatrix(m.cols, m.rows)
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			res.Set(c, r, m.At(r, c))
		}
	}
	return res
}

func (m *Matrix) String() string {
	s := ""
	for r := 0; r < m.rows; r++ {
		s += fmt.Sprintf("%v\n", m.data[r*m.cols : (r+1)*m.cols])
	}
	return s
}

func main() {
	m1, _ := NewMatrixFromSlice([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	m2, _ := NewMatrixFromSlice([][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	})

	fmt.Println("Matrix A:")
	fmt.Print(m1)

	fmt.Println("\nMatrix B:")
	fmt.Print(m2)

	result, err := m1.Multiply(m2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\result A * B:")
		fmt.Print(result)
	}

	fmt.Println("\ntransposed matrix A:")
	fmt.Print(m1.Transpose())
}