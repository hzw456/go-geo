package geo

import "errors"

type Polygon []LinearRing

func NewPolygon(lrs ...LinearRing) *Polygon {
	var rings []LinearRing
	for _, v := range lrs {
		rings = append(rings, v)
	}
	poly := Polygon(rings)
	return &poly
}

func NewPolygonFromPois(pts ...Point) *Polygon {
	ring := NewLinearRing(pts...)
	return NewPolygon(*ring)
}

func (g Polygon) SetSrid(srid uint64) {
	SridMap[&g] = srid
}

func (p Polygon) Buffer(width float64) Polygon {
	return polyBuffer(p, width)
}

//多边形的外环
func (poly Polygon) GetExteriorRing() LinearRing {
	if len(poly) == 0 {
		return nil
	}
	return poly[0]
}

func (poly *Polygon) SetExteriorRing(line LinearRing) {
	if len(*poly) == 0 {
		*poly = append(*poly, line)
	} else {
		(*poly)[0] = line
	}
}

func (poly Polygon) GetInteriorRing() []LinearRing {
	if len(poly) == 0 {
		return nil
	}
	return poly[1:]
}

func (poly Polygon) SetSRID(srid int) {
	return
}

//多边形内的洞
func (poly *Polygon) AddInteriorRing(lines LinearRing) {
	*poly = append(*poly, lines)
}

func (multipoly *MultiPolygon) AddPolygon(poly Polygon) {
	*multipoly = append(*multipoly, poly)
}

func (p *Polygon) Verify() error {
	extring := p.GetExteriorRing()
	ptCount := extring.GetPointCount() - 1
	if ptCount < 3 {
		return errors.New("polygon invaild, point less than 3")
	}
	for i := 0; i < ptCount; i++ {
		if extring[i].X > 90.0 || extring[i].X < -90 || extring[i].Y > 180.0 || extring[i].Y < -180.0 {
			return errors.New("lnglat invaild, lat[-90,90] lng[-180,180]")
		}
	}
	//每次计算面积判断是否合法会有点复杂，考虑采用其他的方式判断
	if GetArea(*p) < 0.0000000001 {
		return errors.New("polygon invaild, area equal 0")
	}
	if p.SelfIntersect() {
		return errors.New("polygon self-intersect")
	}
	return nil
}

func (poly Polygon) SelfIntersect() bool {
	exRing := poly.GetExteriorRing()
	pointCount := exRing.GetPointCount()
	for i := 0; i < pointCount-1; i++ {
		srcV0 := exRing[i]
		srcV1 := exRing[(i+1)%pointCount]
		for j := i + 1; j < pointCount-2; j++ {
			dstV0 := exRing[j]
			dstV1 := exRing[j+1]
			relation := SegmentRelation(LineSegment{srcV0, srcV1}, LineSegment{dstV0, dstV1})
			if relation == RELA_INTERSECT {
				return true
			}
		}
	}
	return false
}

func (poly Polygon) GetExteriorPoints() []Point {
	return poly.GetExteriorRing()
}

func (poly Polygon) IsCCW() bool {
	return relativeArea(poly) > 0
}
