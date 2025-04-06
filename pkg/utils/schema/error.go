package schema

type commonAPIErrorResponseJSON struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

var InvalidRequestErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "Invalid Request",
	Message: "リクエスト形式が正しくありません",
}
var InternalServerErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "Internal Server Error",
	Message: "エラーが発生しました",
}
var UnAuthorizedRequestErrorResponseJSON = commonAPIErrorResponseJSON{
	Type:    "Unauthorized",
	Message: "ログインしていません",
}
