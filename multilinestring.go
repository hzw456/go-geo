/*
 * @Author: haozhiwei@baidu.com
 * @Date: 2020-08-27 10:51:20
 * @LastEditors: haozhiwei@baidu.com
 * @LastEditTime: 2020-09-04 17:54:15
 */
package geo

type MultiLineString []LineString

func NewMultiLineString(ls ...LineString) *MultiLineString {
	var ml MultiLineString
	for _, v := range ls {
		ml = append(ml, v)
	}
	return &ml
}

func (l MultiLineString) SetSrid(srid uint64) {
	SridMap[&l] = srid
}
