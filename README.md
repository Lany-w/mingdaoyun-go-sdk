Mingdaoyun-go-sdk

针对mingdaoyun的API封装的包  

Python 版本：[https://github.com/ghostlitao/mingdaoyun-python-sdk](https://github.com/ghostlitao/mingdaoyun-python-sdk)

PHP 版本：[https://github.com/Lany-w/mingdaoyun-php-sdk](https://github.com/Lany-w/mingdaoyun-php-sdk)

### Installing

```shell
go get -u github.com/Lany-w/mingdaoyun-go-sdk
```



### Usage

```go
package main

import (
	"fmt"

	"github.com/Lany-w/mingdaoyun-go-sdk/mingdaoyun"
	"github.com/Lany-w/mingdaoyun-go-sdk/params"
)

func main() {
	mingdaoyun.Init("https://xxxxxxx.com", "appkey", "sign")
	mdy := mingdaoyun.Client()
  
  data := mdy.Table("60f631095d106d99c054e0bd").Where([]params.Filter{
		{Field: "user", Operate: "=", Value: "Frank"},
		{Field: "60f633935d106d99c054e13a", Operate: "=", Value: "Erwin"},
	}).Get()

	result := mdy.Table("60f631095d106d99c054e0bd").WhereOr([]params.Filter{
		{Field: "user", Operate: "=", Value: "Snow"},
		{Field: "user", Operate: "=", Value: "Lany"},
	}).Get()
  
  resultInsert := mdy.Table("60f631095d106d99c054e0bd").Insert([]params.Insert{
		{ControlId: "60f631095d106d99c054e0be", Value: "Lany"},
		{ControlId: "60f631095d106d99c054e0bf", Value: "test"},
		{ControlId: "60f631095d106d99c054e0c0", Value: "http://img.crcz.com/allimg/202002/11/1581398317535315.jpg"},
	})
  
  var data [][]params.Insert
	row1 := []params.Insert{
		{ControlId: "60f631095d106d99c054e0be", Value: "Lany"},
		{ControlId: "60f631095d106d99c054e0bf", Value: "test-create1"},
	}
	data = append(data, row1)
	row2 := []params.Insert{
		{ControlId: "60f631095d106d99c054e0be", Value: "Lany"},
		{ControlId: "60f631095d106d99c054e0bf", Value: "test-create2"},
	}
	data = append(data, row2)
	resultCreate := mdy.Table("60f631095d106d99c054e0bd").Create(data)
	
  
  result := mdy.Table("60f631095d106d99c054e0bd").Find("382224da-2219-4026-83e2-c6278415ee7b")
  
  mdy.Table("60f631095d106d99c054e0bd").Update("382224da-2219-4026-83e2-c6278415ee7b", []params.Control{
		{ControlId: "60f631095d106d99c054e0bf", Value: "test-create-update"},
	})
  mdy.Table("60f631095d106d99c054e0bd").Delete("382224da-2219-4026-83e2-c6278415ee7b")
}

```

