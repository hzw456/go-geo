package gogeo

import "math"

// 二维向量：(x,y,z)
type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func ConvertPointToVector(p1, p2 Point) *Vector2 {
	return &Vector2{p2.X - p1.X, p2.Y - p1.Y}
}

func ConvertSegToVector(seg LineSegment) *Vector2 {
	return ConvertPointToVector(seg.Start, seg.End)
}

// 返回：新向量
func NewVector2(x, y float64) *Vector2 {
	return &Vector2{X: x, Y: y}
}

func (vec1 *Vector2) Equal(vec2 Vector2) bool {
	return vec1.X == vec2.X && vec1.Y == vec2.Y
}

// 二维向量：拷贝
func (vec1 *Vector2) Clone() *Vector2 {
	return NewVector2(vec1.X, vec1.Y)
}

// 二维向量：加上
// vec1 = vec1 + vec2
func (vec1 Vector2) Add(vec2 Vector2) (res Vector2) {
	res.X = vec1.X + vec2.X
	res.Y = vec1.Y + vec2.Y
	return
}

// 二维向量：减去
// vec1 = vec1 - vec2
func (vec1 Vector2) Sub(vec2 Vector2) (res Vector2) {
	res.X = vec1.X - vec2.X
	res.Y = vec1.Y - vec2.Y
	return
}

// 二维向量：数乘
func (vec1 *Vector2) Multiply(scalar float64) {
	vec1.X *= scalar
	vec1.Y *= scalar
}

func (vec1 *Vector2) Divide(scalar float64) {
	if scalar == 0 {
		panic("分母不能为零！")
	}
	vec1.Multiply(1 / scalar)
}

// 二维向量：点积
func (vec1 *Vector2) Dot(vec2 Vector2) float64 {
	return vec1.X*vec2.X + vec1.Y*vec2.Y
}

// 二维向量：叉积
func (vec1 *Vector2) Cross(vec2 Vector2) float64 {
	return vec1.X*vec2.Y - vec1.Y*vec2.X
}

// 二维向量：长度
func (vec1 *Vector2) Length() float64 {
	return math.Sqrt(vec1.X*vec1.X + vec1.Y*vec1.Y)
}

// 二维向量：长度平方
func (vec1 *Vector2) LengthSq() float64 {
	return vec1.X*vec1.X + vec1.Y*vec1.Y
}

// 二维向量：单位化
func (vec1 *Vector2) Normalize() {
	vec1.Divide(vec1.Length())
}

// 返回：a + b 向量
func Add(a, b Vector2) Vector2 {
	return Vector2{X: a.X + b.X, Y: a.Y + b.Y}
}

// 返回：单位化向量
func Normalize(a Vector2) *Vector2 {
	b := a.Clone()
	b.Normalize()
	return b
}
