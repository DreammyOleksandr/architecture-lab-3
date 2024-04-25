package lang

import (
	"image"
	"io"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	texturestate Texturestate
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation

	// TODO: Реалізувати парсинг команд.
	res = append(res, painter.OperationFunc(painter.GreenFill))
	res = append(res, painter.CrossShape{
		CenterCoordinates: image.Point{X: 30, Y: 30},
	})
	res = append(res, painter.UpdateOp)

	return res, nil
}
