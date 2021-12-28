go-geo
=========
这是一个基于go的基础计算几何库，提供多种空间计算能力
Package go-geo is a general golang libary work for gis, it provide several basic function of spacial calculation.

定义了若干种通用的几何类型
如Point,LineString,Polygon,MultiPoint,MultiLineString,PointZ,LineStringZ等

# 项目结构：project structure #
```
geojson/       #使用geojson的编解码相关功能
geos/          #支持使用使用c编写的geos库进行更多计算几何方法的扩展（需安装geos库）
wkt/           #使用wkt进行编解码相关功能
test/          #功能测试
```

# 基础地理元素的定义和相关操作 defination of basic geometry
## define geometry ##
```go
    //define 3 points 
	newPoint1 := geo.NewPoint(0, 0)
	newPoint2 := geo.NewPoint(0.5, 0.5)
	newPoint3 := geo.NewPoint(1, 1)
    //define linestring by points
	line := geo.NewLineString(*newPoint1, *newPoint2, *newPoint3)
    //define a polygon by points
    poly1 := geo.NewPolygonFromPois(*newPoint1, *newPoint2, *newPoint3)
    //Geometry is interface of all type
    geom := Geometry(*poly1)
    //convert geom to Polygon
    poly1 = geom.(geo.Polygon)
```    

# 数据转换，可将基础地理要素转换为wkt、geojson等的标准格式 conversion,covert basic elmnent to standard format,such as wkt
```go
    //convert geometry to wkt string
	wktStr := wkt.Encode(*poly1))
    //transform wkt to geometry
    geom,_ := wkt.Decode(wktStr)
    //convert geometry to geojson string
    fea := geojson.NewFeature(geom)
    geojsonStr := MarshalJSON(fea)
    // transform geojson to feature
    f, err := geojson.UnmarshalFeature([]byte(geojsonStr))
``` 

# 所有空间计算相关，包括空间要素间的距离、角度等的计算 spacial calculation ,including distance,angle
```go
    //calculate area
    calculation.Area(poly)
```   

# 简化相关类，对元素进行简化 class of simplify

# 外包框相关类 calculate box of geometry

# 缓冲区相关，计算元素的缓冲区 calculate buffer of element

# 空间关系相关，包含、相离、相交等的空间关系判断，基本空间关系使用九交模型 spacial relation,including contain,intesect,etc. based on 9-intersection model

# Features #
* define standard geometry
* output geometry to wkt,geojson
* provide basic function of spacial calculation

# Install #
Use go get to install this package.
go get -u github.com/hzw456/go-geo

从v1.1.0起，建议使用经纬度坐标时，将x定义为经度，y定义为纬度
