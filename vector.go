package gogeo

import "math"

// 三维向量：(x,y,z)
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

// func (vec1 *Vector3) Equal(vec2 Vector3) bool {
// 	return vec1.X == vec2.X && vec1.Y == vec2.Y && vec1.Z == vec2.Z
// }

// // 三维向量：设值
// func (vec1 *Vector3) Set(x, y, z float64) {
// 	vec1.X = x
// 	vec1.Y = y
// 	vec1.Z = z
// }

// // 三维向量：拷贝
// func (vec1 *Vector3) Clone() Vector3 {
// 	return NewVector3(vec1.X, vec1.Y, vec1.Z)
// }

// 三维向量：加上
// vec1 = vec1 + vec2
func (vec1 Vector2) Add(vec2 Vector2) (res Vector2) {
	res.X = vec1.X + vec2.X
	res.Y = vec1.Y + vec2.Y
	return
}

// 三维向量：减去
// vec1 = vec1 - vec2
func (vec1 Vector2) Sub(vec2 Vector2) (res Vector2) {
	res.X = vec1.X - vec2.X
	res.Y = vec1.Y - vec2.Y
	return
}

// 三维向量：数乘
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

// 三维向量：点积
func (vec1 Vector2) Dot(vec2 Vector2) float64 {
	return vec1.X*vec2.X + vec1.Y*vec2.Y
}

// 三维向量：叉积
func (vec1 *Vector2) Cross(vec2 Vector2) float64 {
	return vec1.X*vec2.Y - vec1.Y*vec2.X
}

// 三维向量：长度
func (vec1 *Vector2) Length() float64 {
	return math.Sqrt(vec1.X*vec1.X + vec1.Y*vec1.Y)
}

// 三维向量：长度平方
func (vec1 *Vector2) LengthSq() float64 {
	return vec1.X*vec1.X + vec1.Y*vec1.Y
}

// 三维向量：单位化
func (vec1 *Vector2) Normalize() {
	vec1.Divide(vec1.Length())
}

// // 返回：新向量
// func NewVector3(x, y, z float64) Vector3 {
// 	return Vector3{X: x, Y: y, Z: z}
// }

// // 返回：零向量(0,0,0)
// func Zero3() Vector3 {
// 	return Vector3{X: 0, Y: 0, Z: 0}
// }

// // X 轴 单位向量
// func XAxis3() Vector3 {
// 	return Vector3{X: 1, Y: 0, Z: 0}
// }

// // Y 轴 单位向量
// func YAxis3() Vector3 {
// 	return Vector3{X: 0, Y: 1, Z: 0}
// }

// // Z 轴 单位向量
// func ZAxis3() Vector3 {
// 	return Vector3{X: 0, Y: 0, Z: 1}
// }
// func XYAxis3() Vector3 {
// 	return Vector3{X: 1, Y: 1, Z: 0}
// }
// func XZAxis3() Vector3 {
// 	return Vector3{X: 1, Y: 0, Z: 1}
// }
// func YZAxis3() Vector3 {
// 	return Vector3{X: 0, Y: 1, Z: 1}
// }
// func XYZAxis3() Vector3 {
// 	return Vector3{X: 1, Y: 1, Z: 1}
// }

// // 返回：a + b 向量
// func Add3(a, b Vector3) Vector3 {
// 	return Vector3{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
// }

// // 返回：a - b 向量
// func Sub3(a, b Vector3) Vector3 {
// 	return Vector3{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
// }

// // 返回：a X b 向量 (X 叉乘)
// func Cross3(a, b Vector3) Vector3 {
// 	return Vector3{X: a.Y*b.Z - a.Z*b.Y, Y: a.Z*b.X - a.X*b.Z, Z: a.X*b.Y - a.Y*b.X}
// }

// func AddArray3(vs []Vector3, dv Vector3) []Vector3 {
// 	for i, _ := range vs {
// 		vs[i].Add(dv)
// 	}
// 	return vs
// }

// func Multiply3(vec2 Vector3, scalars []float64) []Vector3 {
// 	vs := []Vector3{}
// 	for _, value := range scalars {
// 		vector := vec2.Clone()
// 		vector.Multiply(value)
// 		vs = append(vs, vector)
// 	}
// 	return vs
// }

// // 返回：单位化向量
// func Normalize(a Vector2) Vector2 {
// 	b := a.Clone()
// 	b.Normalize()
// 	return b
// }
