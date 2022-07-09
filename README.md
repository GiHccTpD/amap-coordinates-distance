# 高德地图计算两坐标点距离

可通过高得地图 [两点间距离](https://lbs.amap.com/demo/javascript-api/example/calcutation/calculate-distance-between-two-markers)在线验证结果

## 演示

```go
package main

import (
	"fmt"

	amapDisc "github.com/GiHccTpD/amap-coordinates-distance"
)

func main() {
	pointA := amapDisc.Coordinates{116.368904, 39.923423}
	pointB := amapDisc.Coordinates{116.387271, 39.922501}

	distance := pointA.Distance(pointB)
	fmt.Println("distance: ", distance)
}
```

输出

```bash
➜ Proving Ground go run amap/amap.go
distance:  1569.6213922676518
```

