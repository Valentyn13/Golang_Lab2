package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixSolver(t *testing.T) {

	assert := assert.New(t)
	testTable := []struct {
		expression string
		expected   float64
	}{
		{"+ 1 2", 3},
		{"- 10 4", 6},
		{"* 2 3", 6},
		{"/ 10 5", 2},
		{"+ + 2 7 - 3 1", 11},
		{"+ 5 * - 4 2 3", 11},
	}
	for _, testestCase := range testTable {
		res, err := PrefixSolver(testestCase.expression)
		if assert.Nil(err) {
			assert.Equal(res, testestCase.expected)
		}

	}
}

func ExamplePrefixSolver() {
	res, _ := PrefixSolver("+ 2 2")
	fmt.Println(res)
}
