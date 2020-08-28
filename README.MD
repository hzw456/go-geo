go-geo
=========
这是一个基于go的基础gis库，提供了基本的空间计算能力
Package go-geo is a general golang libary work for gis, it provide several basic function of spacial calculation.

# 项目结构：project structure #
```
├── element/          # 包含所有基础地理元素的定义和相关操作 defination of basic geometry
├── calculation/      # 所有空间计算相关，包括空间要素间的距离、角度等的计算 spacial calculation ,including distance,angle
├── relation/         # 空间关系相关，包含、相离、相交等的空间关系判断，基本空间关系使用九交模型（还未实现）spacial relation,including contain,intesect,etc. based on                           9-intersection model(not archieve yet)
├── buffer/           # 缓冲区相关，计算元素的缓冲区 calculate buffer of element
├── simplify/         # 简化相关类，对元素进行简化 class of simplify
├── box/              # 外包框相关类 calculate box of geometry
└── convert/          # 转换类，可将基础地理要素转换为wkt等的标准格式 class of conversion,covert basic elmnent to standard format,such as wkt
```

# Features #
* define standard geometry
* output geometry to wkt,geojson
* provide basic function of spacial calculation

# Install #
Use go get to install this package.
go get -u github.com/hzw456/go-geo

# Usage #

## define geometry ##
```go
    //define 3 points 
	newPoint1 := *NewPoint(0, 0)
	newPoint2 := *NewPoint(0.5, 0.5)
	newPoint3 := *NewPoint(1, 1)
    //define linestring by points
	line := *NewLine(newPoint1, newPoint2, newPoint3)
    //change linestirng to ring
    lr := *NewLinearRing(line)
    //define a polygon by linearring
    poly := *NewPolygon(*lr)
```    

## calculation ##
```go
    //calculate area
    calculation.Area(poly)
```   
## convert ##
```go
    //convert geometry to wkt
    convert.GeoToWkt(poly)
``` 