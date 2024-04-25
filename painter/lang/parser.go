package lang

import (
	"bufio"
	"fmt"
	"github.com/roman-mazur/architecture-lab-3/painter"
	"image"
	"io"
	"strconv"
	"strings"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	texturestate Texturestate
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	p.texturestate.ResetOperations()

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines := scanner.Text()

		err := p.parse(lines)
		if err != nil {
			return nil, err
		}
	}
	res := p.texturestate.GetOperationsList()

	return res, nil
}

func (p *Parser) parse(lines string) error {
	params := strings.Split(lines, " ")
	switch params[0] {
	case "white":
		if len(params) != 1 {
			return fmt.Errorf("'white' instruction should have no other parameters")
		}
		p.texturestate.WhiteBackground()
	case "green":
		if len(params) != 1 {
			return fmt.Errorf("'green' instruction should have no other parameters")
		}
		p.texturestate.GreenBackground()
	case "update":
		if len(params) != 1 {
			return fmt.Errorf("'update' instruction should have no other parameters")
		}
		p.texturestate.UpdateOperation()
	case "bgrect":
		parsedParams, err := parseAndConvertParams(params[1:], 4, params[0])
		if err != nil {
			return fmt.Errorf("error while processing 'bgrect' parameters: %v", err)
		}
		p.texturestate.BackgroundRectangle(
			image.Point{X: int(parsedParams[0]), Y: int(parsedParams[1])},
			image.Point{X: int(parsedParams[2]), Y: int(parsedParams[3])})
	case "figure":
		parsedParams, err := parseAndConvertParams(params[1:], 2, params[0])
		if err != nil {
			return fmt.Errorf("error while processing 'figure' parameters: %v", err)
		}
		p.texturestate.Figure(image.Point{X: int(parsedParams[0]), Y: int(parsedParams[1])})
	case "move":
		parsedParams, err := parseAndConvertParams(params[1:], 2, params[0])
		if err != nil {
			return fmt.Errorf("error while processing 'move' parameters: %v", err)
		}
		p.texturestate.MoveOperation(image.Point{X: int(parsedParams[0]), Y: int(parsedParams[1])})
	case "reset":
		if len(params) != 1 {
			return fmt.Errorf("'reset' instruction should have no other parameters")
		}
		p.texturestate.ResetStateAndBackground()
	default:
		return fmt.Errorf("command does not exist")
	}

	return nil
}

func parseAndConvertParams(params []string, expectedSize int, instrType string) ([]float64, error) {
	if len(params) != expectedSize {
		return nil, fmt.Errorf("'%s' instruction should have %d parameters", instrType, expectedSize)
	}
	var converted []float64
	for _, str := range params {
		i, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse parameter '%s' as a float. Err: %v", str, err)
		}

		converted = append(converted, i*800)
	}
	return converted, nil
}
