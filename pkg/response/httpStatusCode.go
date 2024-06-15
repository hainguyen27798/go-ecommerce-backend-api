package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 2003

	ErrInvalidToken = 30001
)

var ErrCodeMsg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Param invalid",
	ErrInvalidToken:     "Token is invalid",
}
