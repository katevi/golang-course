package coordinate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCoordinate(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    Coordinate
		expectedErr string
	}{
		{
			name:     "valid input",
			input:    "(1.5, 2.3, 4.7)",
			expected: Coordinate{x: 1.5, y: 2.3, z: 4.7},
		},
		{
			name:        "missing parentheses",
			input:       "1.5, 2.3, 4.7",
			expectedErr: "invalid format: input must be in the form '(x, y, z)'",
		},
		{
			name:        "too few components",
			input:       "(1.5, 2.3)",
			expectedErr: "invalid format: expected exactly 3 components",
		},
		{
			name:        "too many components",
			input:       "(1.5, 2.3, 4.7, 5.6)",
			expectedErr: "invalid format: expected exactly 3 components",
		},
		{
			name:        "invalid x value",
			input:       "(abc, 2.3, 4.7)",
			expectedErr: "invalid x value",
		},
		{
			name:        "invalid y value",
			input:       "(1.5, def, 4.7)",
			expectedErr: "invalid y value",
		},
		{
			name:        "invalid z value",
			input:       "(1.5, 2.3, ghi)",
			expectedErr: "invalid z value",
		},
		{
			name:        "empty input",
			input:       "",
			expectedErr: "invalid format: input must be in the form '(x, y, z)'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if tt.expectedErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, result)
			}
		})
	}
}
