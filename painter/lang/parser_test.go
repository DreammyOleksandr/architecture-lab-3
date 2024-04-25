package lang

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
	"github.com/stretchr/testify/assert"
	"image"
	"strings"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
		err      bool
	}{
		{
			name:     "White",
			input:    "white\n",
			expected: painter.OperationFunc(nil),
			err:      false,
		},
		{
			name:     "Green",
			input:    "green\n",
			expected: painter.OperationFunc(nil),
			err:      false,
		},
		{
			name:     "Update",
			input:    "update\n",
			expected: painter.UpdateOp,
			err:      false,
		},
		{
			name:  "Background Rectangle",
			input: "bgrect 0.25 0.25 0.75 0.75\n",
			expected: &painter.BackgroundRectangle{
				LeftTop:     image.Point{X: 200, Y: 200},
				RightBottom: image.Point{X: 600, Y: 600},
			},
			err: false,
		},
		{
			name:     "Figure",
			input:    "figure 0.2 0.2\n",
			expected: &painter.CrossShape{CenterCoordinates: image.Point{X: 160, Y: 160}},
			err:      false,
		},
		{
			name:     "Move",
			input:    "move 0.5 0.5\n",
			expected: &painter.MoveOperation{X: 400, Y: 400, ShapesArray: []*painter.CrossShape(nil)},
			err:      false,
		},
		{
			name:     "Reset",
			input:    "reset\n",
			expected: painter.OperationFunc(nil),
			err:      false,
		},
		{
			name:     "Unknown Command",
			input:    "unknownCommand 0.2\n",
			expected: nil,
			err:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := strings.NewReader(test.input)
			parser := &Parser{}

			operations, err := parser.Parse(reader)

			if (err != nil) != test.err {
				t.Errorf("Unexpected error: %v", err)
			}

			if !test.err && operations != nil && len(operations) > 1 {
				if !assert.IsType(t, test.expected, operations[1]) {
					t.Errorf("Expected operation of type %T, but got %T", test.expected, operations[1])
				}

				if !assert.Equal(t, test.expected, operations[1]) {
					t.Errorf("Expected operation %v, but got %v", test.expected, operations[1])
				}
			}
		})
	}
}
