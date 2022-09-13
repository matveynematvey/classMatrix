package matrix

import (
	"strconv"
	"strings"
)

type matrix struct {
	fileName string
	width    int
	height   int
	matrix   [][]int
}

func NewMatrix(fileName string) matrix {
	return matrix{
		fileName: fileName,
	}
}

func (m *matrix) ParseMatrixSize(buf string) {
	sliceString := strings.Fields(buf)

	m.width, _ = strconv.Atoi(sliceString[0])
	m.height, _ = strconv.Atoi(sliceString[1])
}

func (m *matrix) ParseMatrix(buf string) {
	sliceString := strings.Fields(buf)

	m.matrix = make([][]int, m.height)
	for ind := range m.matrix {
		m.matrix[ind] = make([]int, m.width)
	}

	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			m.matrix[i][j], _ = strconv.Atoi(sliceString[i+j])
		}
	}
}

func (m matrix) MultiplyMatrix() {
	factor := 5

	for _, line := range m.matrix {
		for ind := range line {
			line[ind] *= factor
		}
	}
}
