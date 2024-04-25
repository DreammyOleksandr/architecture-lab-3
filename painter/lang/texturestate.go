package lang

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
	"image"
)

type Texturestate struct {
	backgroundColor     painter.Operation
	backgroundRectangle *painter.BackgroundRectangle
	shapesSlice         []*painter.CrossShape
	moveOperations      []painter.Operation
	updateOperation     painter.Operation
}

func (ts *Texturestate) GetOperations() []painter.Operation {
	var ops []painter.Operation

	if ts.backgroundColor != nil {
		ops = append(ops, ts.backgroundColor)
	}
	if ts.backgroundRectangle != nil {
		ops = append(ops, ts.backgroundRectangle)
	}
	if len(ts.moveOperations) != 0 {
		ops = append(ops, ts.moveOperations...)
		ts.moveOperations = nil
	}
	if len(ts.shapesSlice) != 0 {
		for _, shape := range ts.shapesSlice {
			ops = append(ops, shape)
		}
	}
	if ts.updateOperation != nil {
		ops = append(ops, ts.updateOperation)
	}

	return ops
}

func (ts *Texturestate) GreenBackground() {
	ts.backgroundColor = painter.OperationFunc(painter.GreenFill)
}

func (ts *Texturestate) WhiteBackground() {
	ts.backgroundColor = painter.OperationFunc(painter.WhiteFill)
}

func (ts *Texturestate) UpdateOperation() {
	ts.updateOperation = painter.UpdateOp
}

func (ts *Texturestate) BackgroundRectangle(leftTop image.Point, rightBottom image.Point) {
	ts.backgroundRectangle = &painter.BackgroundRectangle{
		LeftTop:     leftTop,
		RightBottom: rightBottom,
	}
}

func (ts *Texturestate) Figure(centralPoint image.Point) {
	figure := painter.CrossShape{
		CenterCoordinates: centralPoint,
	}
	ts.shapesSlice = append(ts.shapesSlice, &figure)
}

func (ts *Texturestate) MoveOperation(x int, y int) {
	moveOp := painter.MoveOperation{X: x, Y: y, ShapesArray: ts.shapesSlice}
	ts.moveOperations = append(ts.moveOperations, &moveOp)
}

func (ts *Texturestate) ResetOperations() {
	if ts.backgroundColor == nil {
		ts.backgroundColor = painter.OperationFunc(painter.Reset)
	}
	if ts.updateOperation != nil {
		ts.updateOperation = nil
	}
}

func (ts *Texturestate) Reset() {
	ts.backgroundColor = nil
	ts.backgroundRectangle = nil
	ts.shapesSlice = nil
	ts.moveOperations = nil
	ts.updateOperation = nil
}
