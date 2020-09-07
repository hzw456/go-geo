package geo

import (
	"errors"
)

type LineString []Point

type LineSegment struct {
	Start Point
	End   Point
}

func NewLineString(point ...Point) *LineString {
	var points []Point
	for _, point := range point {
		points = append(points, point)
	}
	line := LineString(points)
	return &line
}

func (line LineString) GetPointCount() int {
	return len(line)
}

func (line LineString) Verify() error {
	if line.GetPointCount() == 0 {
		err := errors.New("line has no point")
		return err
	}
	if line.GetPointCount() == 1 {
		err := errors.New("line has only one point")
		return err
	}
	return nil
}

//取首点
func (line LineString) GetFirstPoint() Point {
	if len(line) < 2 {
		return Point{0, 0}
	}
	return line[0]
}

//取尾点
func (line LineString) GetEndPoint() Point {
	if len(line) < 2 {
		return Point{0, 0}
	}
	return line[line.GetPointCount()-1]
}

func (line *LineString) Append(point Point) error {
	*line = append(*line, point)
	return nil
}

//在指定位置改变linestring的点
func (line *LineString) SetPoint(position int, point Point) error {
	if line.GetPointCount() < position {
		return errors.New("line has no such position")
	}
	if line.GetPointCount() == position {
		line.Append(point)
	}
	(*line)[position] = point
	return nil
}

//在指定位置插入点
func (line *LineString) InsertPoint(position int, point Point) error {
	if line.GetPointCount() < position {
		return errors.New("line has no such position")
	}
	if line.GetPointCount() == position {
		line.Append(point)
	}
	*line = append((*line)[:position+1], (*line)[position:]...)
	(*line)[position] = point
	return nil
}

//在指定位置删除点
func (line *LineString) DelPoint(position int) error {
	if line.GetPointCount() <= position {
		return errors.New("line has no such position")
	}
	(*line) = append((*line)[:position], (*line)[position+1:]...)
	return nil
}

func (line LineString) Length() float64 {
	var dis float64
	for i := 0; i < line.GetPointCount()-1; i++ {
		dis += line[i].Distance(line[i+1])
	}
	return dis
}

//判断两条线是不是相同
func (line1 LineString) Equal(line2 LineString) bool {
	for i := 0; i < line1.GetPointCount(); i++ {
		if !line1[i].Equal(line2[i]) {
			return false
		}
	}
	return true
}

func (line *LineString) Reverse() {
	count := line.GetPointCount()
	mid := count / 2
	for i := 0; i < mid; i++ {
		(*line)[i], (*line)[line.GetPointCount()-1-i] = (*line)[line.GetPointCount()-1-i], (*line)[i]
	}
}

func (line LineString) ToRing() LinearRing {
	return *NewRingFromLine(line)
}

func (line LineString) GetPointSet() []Point {
	return line
}
