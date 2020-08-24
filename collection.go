package geo

type Collection []Geometry

func (c Collection) TypeString() string {
	return "GeometryCollection"
}
