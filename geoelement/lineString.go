package geoelement

import (
	"errors"
)

type LineString []Point
type Points []Point

func newLineByPoints(points Points) (LineString, error) {
	if len(points) < 2 {
		err := errors.New("line cant be less than two points")
		return nil, err
	}
	return points, nil
}

func newLine(point ...Point) (LineString, error) {
	var points Points
	for _, point := range point {
		points = append(points, point)
	}
	return newLineByPoints(points)
}

func (line LineString) getPointCount() int {
	return len(line)
}

func (line LineString) GetfirstPoint() (Point, error) {
	if line.getPointCount() == 0 {
		err := errors.New("line has no point")
		return Point{0, 0}, err
	}
	if line.getPointCount() == 1 {
		err := errors.New("line has only one point")
		return Point{0, 0}, err
	}
	return line[0], nil
}

func (line LineString) getEndPoint() (Point, error) {
	if line.getPointCount() == 0 {
		err := errors.New("line has no point")
		return Point{0, 0}, err
	}
	if line.getPointCount() == 1 {
		err := errors.New("line has only one point")
		return Point{0, 0}, err
	}
	return line[line.getPointCount()-1], nil
}
