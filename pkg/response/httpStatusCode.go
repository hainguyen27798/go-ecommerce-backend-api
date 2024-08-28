package response

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParamInvalid  = 2003  // Param invalid
	ErrInvalidToken      = 30001 // Token is invalid
	ErrCodeUserHasExists = 50001 // User has already exists
)

var ErrCodeMsg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeParamInvalid:  "Param invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrCodeUserHasExists: "User has already exists",
}
