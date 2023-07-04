package params

type MingDaoRequest interface {
	do()
}

type MingDaoResponse interface {
	response()
}

type BaseRequestParams struct {
	AppKey      string `json:"appKey"`
	Sign        string `json:"sign"`
	WorkSheetId string `json:"worksheetId"`
}

type RequestParams struct {
	ViewId           string       `json:"viewId,omitempty"`
	PageIndex        int          `json:"pageIndex,omitempty"`
	PageSize         int          `json:"pageSize,omitempty"`
	SortId           string       `json:"sortId,omitempty"`
	IsAsc            bool         `json:"isAsc,omitempty"`
	Filters          []FilterItem `json:"filters,omitempty"`
	NotGetTotal      bool         `json:"notGetTotal,omitempty"`
	UseControlId     bool         `json:"useControlId,omitempty"`
	GetSystemControl bool         `json:"getSystemControl,omitempty"`
	KeyWords         string       `json:"keywords,omitempty"`
	TriggerWorkflow  bool         `json:"triggerWorkflow,omitempty"`
	Controls         []Control    `json:"controls,omitempty"`
	Rows             [][]Control  `json:"rows,omitempty"`
	RowId            string       `json:"rowId,omitempty"`
}

type FilterItem struct {
	ControlId  string      `json:"controlId,omitempty"`
	DataType   float64     `json:"dataType,omitempty"`
	FilterType int         `json:"filterType,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	SpliceType int         `json:"spliceType,omitempty"`
}

type Control struct {
	ControlId    string       `json:"controlId,omitempty"`
	Value        interface{}  `json:"value,omitempty"`
	ValueType    int          `json:"valueType,omitempty"`
	EditType     int          `json:"editType,omitempty"`
	ControlFiles []Base64File `json:"controlFiles,omitempty"`
}

type Base64File struct {
	BaseFile string `json:"baseFile,omitempty"`
	FileName string `json:"fileName,omitempty"`
}

type MdyResponse struct {
	Data      string `json:"data"`
	Success   bool   `json:"success"`
	ErrorCode int    `json:"error_code"`
}
type MdyRowDetailResponse struct {
	Data      map[string]interface{} `json:"data"`
	Success   bool                   `json:"success"`
	ErrorCode int                    `json:"error_code"`
}

type MdyMapResponse struct {
	Data      MdyMapItemResponse `json:"data"`
	Success   bool               `json:"success"`
	ErrorCode int                `json:"error_code"`
}
type MdyMapItemResponse struct {
	WorksheetId string                   `json:"worksheetId"`
	Controls    []map[string]interface{} `json:"controls"`
}

type MdyListResponse struct {
	Data      MdyListItemResponse `json:"data"`
	Success   bool                `json:"success"`
	ErrorCode int                 `json:"error_code"`
}

type MdyListItemResponse struct {
	Rows  []map[string]interface{} `json:"rows"`
	Total int                      `json:"total"`
}

type Filter struct {
	Field   string
	Operate string //支持 =,!=,contains,startWith,endWith,notContain,>,>=,<,<=,DateEnum,NDateEnum,RCEq,RCNe
	Value   interface{}
}

func (*RequestParams) do()         {}
func (*MdyResponse) response()     {}
func (*MdyListResponse) response() {}
