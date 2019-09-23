package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
	//INVALID_PARAMS : "请求参数错误",
	//ERROR_EXIST_TAG : "已存在该标签名称",
	//ERROR_NOT_EXIST_TAG : "该标签不存在",
	//ERROR_NOT_EXIST_ARTICLE : "该文章不存在",
	//ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
	//ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
	//ERROR_AUTH_TOKEN : "Token生成失败",
	//ERROR_AUTH : "Token错误",
	//SUCCESS                    :   "服务器成功返回请求的数据。",
	SUCCESS_CREATED:            "新建或修改数据成功。",
	SUCCESS_QUEUE:              "一个请求已经进入后台排队（异步任务）。",
	SUCCESS_DELETE:             "删除数据成功。",
	ERROR_BAD_REQUEST:          "发出的请求有错误，服务器没有进行新建或修改数据的操作。",
	ERROR_UNAUTHORIZED:         "用户没有权限（令牌、用户名、密码错误）。",
	ERROR_FORBIDDEN:            "用户得到授权，但是访问是被禁止的。",
	ERROR_NOT_FOUND:            "发出的请求针对的是不存在的记录，服务器没有进行操作。",
	ERROR_NOT_ACCEPTABLE:       "请求的格式不可得。",
	ERROR_GONE:                 "请求的资源被永久删除，且不会再得到的。",
	ERROR_UNPROCESSABLE_ENTITY: "当创建一个对象时，发生一个验证错误。",
	//ERROR                      :   "服务器发生错误，请检查服务器。",
	ERROR_BAD_GATEWAY:         "网关错误。",
	ERROR_SERVICE_UNAVAILABLE: "服务不可用，服务器暂时过载或维护。",
	ERROR_GATEWAY_TIMEOUT:     "网关超时。",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
