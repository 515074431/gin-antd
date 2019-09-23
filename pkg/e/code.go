package e

const (
	SUCCESS                    = 200 // '服务器成功返回请求的数据。',
	SUCCESS_CREATED            = 201 // '新建或修改数据成功。',
	SUCCESS_QUEUE              = 202 // '一个请求已经进入后台排队（异步任务）。',
	SUCCESS_DELETE             = 204 // '删除数据成功。',
	ERROR_BAD_REQUEST          = 400 // '发出的请求有错误，服务器没有进行新建或修改数据的操作。',
	ERROR_UNAUTHORIZED         = 401 // '用户没有权限（令牌、用户名、密码错误）。',
	ERROR_FORBIDDEN            = 403 // '用户得到授权，但是访问是被禁止的。',
	ERROR_NOT_FOUND            = 404 // '发出的请求针对的是不存在的记录，服务器没有进行操作。',
	ERROR_NOT_ACCEPTABLE       = 406 // '请求的格式不可得。',
	ERROR_GONE                 = 410 // '请求的资源被永久删除，且不会再得到的。',
	ERROR_UNPROCESSABLE_ENTITY = 422 // '当创建一个对象时，发生一个验证错误。',
	ERROR                      = 500 // '服务器发生错误，请检查服务器。',
	ERROR_BAD_GATEWAY          = 502 // '网关错误。',
	ERROR_SERVICE_UNAVAILABLE  = 503 // '服务不可用，服务器暂时过载或维护。',
	ERROR_GATEWAY_TIMEOUT      = 504 // '网关超时。',
)
