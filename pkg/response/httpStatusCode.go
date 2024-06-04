package response

const (
	ErrCodeSuccess      = 2001
	ErrCodeParamInvalid = 2003
)

var ErrCodeMsg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Param invalid",
}
