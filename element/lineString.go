package element

import (
	"errors"
)

type LineString []Point

func NewLine2(points []Point) *LineString {
	if len(points) < 2 {
		return nil
	}
	line := LineString(points)
	return &line
}

func NewLine(point ...Point) *LineString {
	var points []Point
	for _, point := range point {
		points = append(points, point)
	}
	return NewLine2(points)
}

func (line LineString) GetPointCount() int {
	return len(line)
}

//取首点
func (line LineString) GetFirstPoint() (Point, error) {
	if line.GetPointCount() == 0 {
		err := errors.New("line has no point")
		return Point{0, 0}, err
	}
	if line.GetPointCount() == 1 {
		err := errors.New("line has only one point")
		return Point{0, 0}, err
	}
	return line[0], nil
}

//取尾点
func (line LineString) GetEndPoint() (Point, error) {
	if line.GetPointCount() == 0 {
		err := errors.New("line has no point")
		return Point{0, 0}, err
	}
	if line.GetPointCount() == 1 {
		err := errors.New("line has only one point")
		return Point{0, 0}, err
	}
	return line[line.GetPointCount()-1], nil
}

func (line *LineString) AppendPoint(point Point) error {
	*line = append(*line, point)
	return nil
}

//在指定位置改变linestring的点
func (line *LineString) SetPoint(position int, point Point) error {
	if line == nil {
		return errors.New("line is nil")
	}
	if line.GetPointCount() < position {
		return errors.New("line has no such position")
	}
	if line.GetPointCount() == position {
		line.AppendPoint(point)
	}
	(*line)[position] = point
	return nil
}

//在指定位置插入点
func (line *LineString) InsertPoint(position int, point Point) error {
	if line == nil {
		return errors.New("line is nil")
	}
	if line.GetPointCount() < position {
		return errors.New("line has no such position")
	}
	if line.GetPointCount() == position {
		line.AppendPoint(point)
	}
	*line = append((*line)[:position+1], (*line)[position:]...)
	(*line)[position] = point
	return nil
}

//在指定位置删除点
func (line *LineString) DelPoint(position int) error {
	if line == nil {
		return errors.New("line is nil")
	}
	if line.GetPointCount() <= position {
		return errors.New("line has no such position")
	}
	(*line) = append((*line)[:position], (*line)[position+1:]...)
	return nil
}

func (line LineString) Length() float64 {
	var dis float64
	for i := 0; i < line.GetPointCount()-1; i++ {
		dis += line[i].PointDistance(line[i+1])
	}
	return dis
}

//判断两条线是不是相同 TODO：是否开gorountine
func (line1 LineString) Equal(line2 LineString) bool {
	for i := 0; i < line1.GetPointCount(); i++ {
		if !line1[i].Equal(line2[i]) {
			return false
		}
	}
	return true
}
