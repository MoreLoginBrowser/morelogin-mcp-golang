package types

// CommonResponse 通用响应结构
type CommonResponse[T any] struct {
	Code      int    `json:"code"`      // 返回结果编码 0:正常 其他编码都是异常
	Msg       string `json:"msg"`       // 错误信息
	Data      T      `json:"data"`      // 响应数据
	RequestId string `json:"requestId"` // 操作请求ID
}

type EndpointConfig struct {
	UrlTemplate string
	PathParams  []string
}

type CommonPageRes[T any] struct {
	Current  string `json:"current"`  // current page number
	Pages    string `json:"pages"`    // page number
	Total    string `json:"total"`    // total number
	DataList []T    `json:"dataList"` // Webdriver Path
}
