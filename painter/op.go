package painter

import (
	"golang.org/x/image/draw"
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

type CrossShape struct {
	CenterCoordinates image.Point
}

func (op CrossShape) Do(t screen.Texture) bool {
	//TODO fix shape proportions
	c := color.RGBA{B: 255, A: 255}
	crossSize := 400
	t.Fill(image.Rect(op.CenterCoordinates.X-crossSize/2, op.CenterCoordinates.Y-70, op.CenterCoordinates.X+crossSize/2, op.CenterCoordinates.Y+70), c, draw.Src)
	t.Fill(image.Rect(op.CenterCoordinates.X-70, op.CenterCoordinates.Y-crossSize/2, op.CenterCoordinates.X+70, op.CenterCoordinates.Y+crossSize/2), c, draw.Src)
	return false
}

type BackgroundRectangle struct {
	LeftTop     image.Point
	RightBottom image.Point
}

func (op *BackgroundRectangle) Do(t screen.Texture) bool {
	c := color.Black
	t.Fill(image.Rect(op.LeftTop.X, op.LeftTop.Y, op.RightBottom.X, op.RightBottom.Y), c, screen.Src)
	return false
}

func Reset(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, screen.Src)
}

type MoveOperation struct {
	X           int
	Y           int
	ShapesArray []*CrossShape
}

func (op *MoveOperation) Do(t screen.Texture) bool {
	for i := range op.ShapesArray {
		op.ShapesArray[i].CenterCoordinates.X += op.X
		op.ShapesArray[i].CenterCoordinates.Y += op.Y
	}
	return false
}
