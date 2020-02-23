package e

// 錯誤代碼
const (
	SUCCESS           = 200
	ERROR             = 500
	INVALID_PARAMS    = 400
	NOT_FOUNT_USER    = 20001
	WRONG_PASSWORD    = 20002
	NOT_FOUNT_ARTICLE = 30001

	JWT_NOT_FOUND = 40001
	JWT_NOT_ERROR = 40002
	JWT_TIME_OUT  = 40003
)

// MsgFlags 錯誤訊息
var MsgFlags = map[int]string{
	SUCCESS:           "ok",
	ERROR:             "fail",
	INVALID_PARAMS:    "請求參數錯誤",
	NOT_FOUNT_USER:    "沒有找到使用者",
	WRONG_PASSWORD:    "密碼錯誤",
	NOT_FOUNT_ARTICLE: "沒有找到文章",
	JWT_NOT_FOUND:     "沒有JWT TOKEN",
	JWT_NOT_ERROR:     "JWT TOKEN 錯誤",
	JWT_TIME_OUT:      "JWT TOKEN 過期",
}

// GetMsg 取得錯誤訊息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
