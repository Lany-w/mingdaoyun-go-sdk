package mingdaoyun

import (
	"encoding/json"
	"math"
	"strconv"

	"github.com/Lany-w/mingdaoyun-go-sdk/params"
	"github.com/Lany-w/mingdaoyun-go-sdk/request"
)

var _spliceType = 1

func processList(mdy *MingDaoYun) params.MdyListResponse {
	result := params.MdyListResponse{}
	if mdy.PageSize == 0 {
		//获取总行数
		resultCount := &params.MdyResponse{}
		data := request.Do(Host+RowsTotalUri, mdy)
		json.Unmarshal(data, resultCount)

		total, _ := strconv.Atoi(resultCount.Data)
		//fmt.Println(total, "total")
		totalPage := int(math.Floor(float64(total)/1000)) + 1
		mdy.NotGetTotal = true

		for i := 1; i <= totalPage; i++ {
			mdy.PageSize = 1000
			mdy.PageIndex = i
			dataList := request.Do(Host+ListUri, mdy)
			if result.Data.Total == 0 {
				json.Unmarshal(dataList, &result)
				result.Data.Total = total
			}
			listItem := params.MdyListResponse{}
			json.Unmarshal(dataList, &listItem)
			result.Data.Rows = append(result.Data.Rows, listItem.Data.Rows...)
		}

		result.Success = true
		result.ErrorCode = 1
	} else {
		data := request.Do(Host+ListUri, mdy)
		json.Unmarshal(data, &result)
	}
	mdy.RequestParams = &params.RequestParams{}
	return result
}

func processTable(md *MingDaoYun) {
	result := request.Do(Host+WorkSheetMapUri, md)

	mdyMap := params.MdyMapResponse{}
	json.Unmarshal(result, &mdyMap)
	md.WorksheetMap = mdyMap
}

func processFilter(w []params.Filter, md *MingDaoYun) {

	for _, val := range w {
		var filterType int
		item := params.FilterItem{}
		switch val.Operate {
		case "contains":
			filterType = 1
		case "=":
			filterType = 2
		case "startWith":
			filterType = 3
		case "endWith":
			filterType = 4
		case "notContain":
			filterType = 5
		case "!=":
			filterType = 6
		case ">":
			filterType = 13
		case ">=":
			filterType = 14
		case "<":
			filterType = 15
		case "<=":
			filterType = 16
		case "DateEnum":
			filterType = 17
		case "NDateEnum":
			filterType = 18
		case "RCEq":
			filterType = 24
		case "RCNe":
			filterType = 25
		default:
			panic("不支持的条件符号!")
		}
		item.ControlId = val.Field
		item.FilterType = filterType
		item.DataType = getFieldDataType(md, val.Field)
		item.SpliceType = _spliceType
		item.Value = val.Value

		md.Filters = append(md.Filters, item)
		//fmt.Printf("%+v", md.Filters)
	}
}

func getFieldDataType(md *MingDaoYun, field string) float64 {
	for _, item := range md.WorksheetMap.Data.Controls {
		if item["controlId"] == field || item["alias"] == field {
			return item["type"].(float64)
		}
	}
	return 0
}

func processAdd(md *MingDaoYun, url string) params.MdyResponse {
	result := request.Do(url, md)
	data := params.MdyResponse{}
	json.Unmarshal(result, &data)
	md.RequestParams = &params.RequestParams{}
	return data
}

func processRowDetail(md *MingDaoYun) params.MdyRowDetailResponse {
	result := request.Do(Host+RowDetailUri, md)
	data := params.MdyRowDetailResponse{}
	json.Unmarshal(result, &data)
	md.RequestParams = &params.RequestParams{}
	return data
}

func processUpdate(md *MingDaoYun, url string) params.MdyResponse {
	result := request.Do(url, md)
	data := params.MdyResponse{}
	json.Unmarshal(result, &data)
	md.RequestParams = &params.RequestParams{}
	return data
}
