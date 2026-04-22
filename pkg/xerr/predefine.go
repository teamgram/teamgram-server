package xerr

const (
	// General error codes.
	ServerInternalError = 500  // Server internal error
	ArgsError           = 1001 // Input parameter error
	NoPermissionError   = 1002 // Insufficient permission
	DuplicateKeyError   = 1003
	RecordNotFoundError = 1004 // Record does not exist

	TokenExpiredError     = 1501
	TokenInvalidError     = 1502
	TokenMalformedError   = 1503
	TokenNotValidYetError = 1504
	TokenUnknownError     = 1505
	TokenKickedError      = 1506
	TokenNotExistError    = 1507

	OrgUserNoPermissionError = 1520
)

var (
	ErrArgs                     = NewCodeError(ArgsError, "ArgsError")
	ErrNoPermission             = NewCodeError(NoPermissionError, "NoPermissionError")
	ErrInternalServer           = NewCodeError(ServerInternalError, "ServerInternalError")
	ErrRecordNotFound           = NewCodeError(RecordNotFoundError, "RecordNotFoundError")
	ErrDuplicateKey             = NewCodeError(DuplicateKeyError, "DuplicateKeyError")
	ErrTokenExpired             = NewCodeError(TokenExpiredError, "TokenExpiredError")
	ErrTokenInvalid             = NewCodeError(TokenInvalidError, "TokenInvalidError")
	ErrTokenMalformed           = NewCodeError(TokenMalformedError, "TokenMalformedError")
	ErrTokenNotValidYet         = NewCodeError(TokenNotValidYetError, "TokenNotValidYetError")
	ErrTokenUnknown             = NewCodeError(TokenUnknownError, "TokenUnknownError")
	ErrTokenKicked              = NewCodeError(TokenKickedError, "TokenKickedError")
	ErrTokenNotExist            = NewCodeError(TokenNotExistError, "TokenNotExistError")
	ErrOrgUserNoPermissionError = NewCodeError(OrgUserNoPermissionError, "OrgUserNoPermissionError")
)
