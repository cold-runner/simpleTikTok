package errno

func init() {
	// 此处的message是返回给前端看的错误消息，不能暴露业务错误信息
	register(ErrSuccess, 200, "OK")
	register(ErrDatabase, 500, "database errno")
	register(ErrRpcCall, 500, "internal error")
}
