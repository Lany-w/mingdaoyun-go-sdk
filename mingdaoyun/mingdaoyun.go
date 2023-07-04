package mingdaoyun

import (
	"github.com/Lany-w/mingdaoyun-go-sdk/params"
)

type MingDaoYun struct {
	*params.BaseRequestParams
	*params.RequestParams
	WorksheetMap params.MdyMapResponse `json:"-"`
}

func (*MingDaoYun) do() {}

var Host string
var _appKey string
var _sign string

const (
	ListUri         = "/api/v2/open/worksheet/getFilterRows"         //获取列表
	WorkSheetMapUri = "/api/v2/open/worksheet/getWorksheetInfo"      //获取工作表结构
	RowDetailUri    = "/api/v2/open/worksheet/getRowByIdPost"        //获取行记录详情
	RowsTotalUri    = "/api/v2/open/worksheet/getFilterRowsTotalNum" //获取总行数
	AddRowUri       = "/api/v2/open/worksheet/addRow"                //新增行记录
	DeleteRowUri    = "/api/v2/open/worksheet/deleteRow"             //删除行记录
	EditRowUri      = "/api/v2/open/worksheet/editRow"               //更新行记录
	AddRowsUri      = "/api/v2/open/worksheet/addRows"               //批量新增行记录
)

func Init(host string, appkey string, sign string) {
	Host = host
	_appKey = appkey
	_sign = sign
}

func Client() *MingDaoYun {
	return &MingDaoYun{
		BaseRequestParams: &params.BaseRequestParams{
			AppKey: _appKey,
			Sign:   _sign,
		},
		RequestParams: &params.RequestParams{},
	}
}

//设置要操作的worksheet
func (md *MingDaoYun) Table(tableName string) *MingDaoYun {
	if Host == "" || _appKey == "" || _sign == "" {
		panic("请正确设置明道云参数!")
	}
	md.WorkSheetId = tableName
	if md.WorksheetMap.Data.WorksheetId != tableName {
		processTable(md)
	}

	return md
}

func (md *MingDaoYun) Where(w []params.Filter) *MingDaoYun {
	_spliceType = 1
	processFilter(w, md)
	return md
}

func (md *MingDaoYun) WhereOr(w []params.Filter) *MingDaoYun {
	_spliceType = 2
	processFilter(w, md)
	return md
}

//设置获取数量 最高为1000
func (md *MingDaoYun) Limit(num int) *MingDaoYun {
	md.PageSize = num
	return md
}

//设置页码
func (md *MingDaoYun) Page(num int) *MingDaoYun {
	md.PageIndex = num
	return md
}

//设置查询视图
func (md *MingDaoYun) View(viewName string) *MingDaoYun {
	md.ViewId = viewName
	return md
}

//关键词
func (md *MingDaoYun) KeyWord(keyword string) *MingDaoYun {
	md.KeyWords = keyword
	return md
}

//排序字段
func (md *MingDaoYun) Sort(field string) *MingDaoYun {
	md.SortId = field
	return md
}

//是否升序
func (md *MingDaoYun) Asc(asc bool) *MingDaoYun {
	md.IsAsc = asc
	return md
}

//是否不统计总行数以提高性能
func (md *MingDaoYun) NotTotal(notTotal bool) *MingDaoYun {
	md.NotGetTotal = notTotal
	return md
}

//是否只返回controlId
func (md *MingDaoYun) OnlyControlId(onlyControlId bool) *MingDaoYun {
	md.UseControlId = onlyControlId
	return md
}

//是否获取系统字段
func (md *MingDaoYun) SystemControl(systemControl bool) *MingDaoYun {
	md.GetSystemControl = systemControl
	return md
}

//获取数据
func (md *MingDaoYun) Get() params.MdyListResponse {
	return processList(md)
}

//是否触发工作流
func (md *MingDaoYun) IsTriggerWorkflow(isTrigger bool) *MingDaoYun {
	md.TriggerWorkflow = isTrigger
	return md
}

//新增行记录
func (md *MingDaoYun) Insert(insertData []params.Control) params.MdyResponse {
	md.Controls = insertData
	return processAdd(md, Host+AddRowUri)
}

//批量新建行记录
func (md *MingDaoYun) Create(insertData [][]params.Control) params.MdyResponse {
	md.Rows = insertData
	return processAdd(md, Host+AddRowsUri)
}

//行记录详情
func (md *MingDaoYun) Find(rowId string) params.MdyRowDetailResponse {
	md.RowId = rowId
	return processRowDetail(md)
}

//更新行记录
func (md *MingDaoYun) Update(rowId string, updateData []params.Control) params.MdyResponse {
	md.Controls = updateData
	md.RowId = rowId
	return processUpdate(md, Host+EditRowUri)
}

//删除行记录
func (md *MingDaoYun) Delete(rowId string) params.MdyResponse {
	md.RowId = rowId
	return processUpdate(md, Host+DeleteRowUri)
}
