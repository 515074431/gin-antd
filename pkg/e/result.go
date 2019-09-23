package e
//返回值
type Result struct {
	Status bool  `json:"status"`
	Error interface{} `json:"error"`
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}
//实例化一个默认返回值
func NewDefaultResult() Result {
	return Result{
		Status:true,
		Code:SUCCESS,
	}
}
//实例化一个返回值
func NewResult(status bool, code int) Result {
	return Result{
		Status:status,
		Code:code,
	}
}
