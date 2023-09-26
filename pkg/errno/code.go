package errno

var (
	// OK 代表请求成功.
	OK = &Errno{HTTP: 0, Code: "", Message: "Request Success."}

	// InternalServerError 表示所有未知的服务器端错误.
	InternalServerError = &Errno{HTTP: 500, Code: "InternalError", Message: "Internal server errno."}

	// ErrPageNotFound 表示路由不匹配错误.
	ErrPageNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.PageNotFound", Message: "Page not found."}

	// ErrBind 表示参数绑定错误.
	ErrBind = &Errno{HTTP: 400, Code: "InvalidParameter.BindError", Message: "Error occurred while binding the request body to the struct."}

	// ErrInvalidParameter 表示所有验证失败的错误.
	ErrInvalidParameter = &Errno{HTTP: 400, Code: "InvalidParameter", Message: "Parameter verification failed."}

	// ErrSignToken 表示签发 JWT Token 时出错
	ErrSignToken = &Errno{HTTP: 401, Code: "AuthFailure.SignTokenError", Message: "Error occurred while signing the JSON web token."}

	// ErrTokenInvalid 表示 JWT Token 格式错误.
	ErrTokenInvalid = &Errno{HTTP: 401, Code: "AuthFailure.TokenInvalid", Message: "Token was invalid."}

	// ErrUnauthorized 表示请求没有被授权.
	ErrUnauthorized = &Errno{HTTP: 401, Code: "AuthFailure.Unauthorized", Message: "Unauthorized."}

	// ErrUserAlredyExist 表示用户已经存在.
	ErrUserAlredyExist = &Errno{HTTP: 409, Code: "UserConflict.UserAlreadyExist",
		Message: "User already exist."}

	// ErrUserNotExist 表示用户不存在的错误.
	ErrUserNotExist = &Errno{HTTP: 404, Code: "UserNotFound", Message: "User does not exist."}

	// ErrUserPassword 表示用户密码错误.
	ErrUserPassword = &Errno{HTTP: 401, Code: "UserPassword", Message: "User password error."}

	// ErrAlreadyFollow  表示用户已经关注.
	ErrAlreadyFollow = &Errno{HTTP: 409,
		Code:    "UserConflict.UserAlreadyFollow",
		Message: "Already follow the user."}

	// ErrNotFollow 表示用户没有关注.
	ErrNotFollow = &Errno{HTTP: 409, Code: "UserConflict.UserNotFollow",
		Message: "Not following the user."}
	// ErrInvalidUpdate 表示更新数据库时出错.
	ErrInvalidUpdate = &Errno{HTTP: 400, Code: "InvalidParameter." +
		"UpdateError", Message: "Error occurred while updating the database."}
	// ErrRedisNotInitialized ErrInvalidDelete 表示redis数据库未初始化连接.
	ErrRedisNotInitialized = &Errno{HTTP: 500, Code: "RedisNotInitialized", Message: "Redis not initialized."}

	ErrInvalidDeleteComment = &Errno{HTTP: 400, Code: "InvalidParameter",
		Message: "you can't delete this comment."}
)
