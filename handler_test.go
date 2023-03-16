package lab2

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Error(msg string) error {
	return fmt.Errorf(msg)
}

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      string
		expectedError error
	}{
		{
			name:     "valid expression",
			input:    "+ 3 4",
			expected: "7",
		},
		{
			name:          "invalid expression",
			input:         "/ 4 0",
			expectedError: Error("division by zero"),
		},
		{
			name:          "empty input",
			input:         "",
			expectedError: Error("EOF"),
		},
		{
			name:          "invalid input",
			input:         "hello world",
			expectedError: Error("invalid expression"),
		},
		{
			name:          "missing operand",
			input:         "+ 1",
			expectedError: Error("missing operand:stack index out of range"),
		},
		{
			name:          "missing operator",
			input:         "1 2",
			expectedError: Error("missing operator:stack index out of range"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			output := new(bytes.Buffer)
			handler := &ComputeHandler{
				Input:  input,
				Output: output,
			}

			err := handler.Compute()

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, output.String())
		})
	}
}
