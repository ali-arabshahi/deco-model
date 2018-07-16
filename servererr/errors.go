package servererr

// ErrCode enum type
type ErrCode int

const (
	// Unknown error enum
	Unknown ErrCode = -1
	// Success error enum
	Success ErrCode = 0
	// Internal error enum
	Internal ErrCode = 1
	// DBValueExists error enum
	DBValueExists ErrCode = 100
	// DBValueNotFound error enum
	DBValueNotFound ErrCode = 101
	// DBValueExpired error enum
	DBValueExpired ErrCode = 102
	// TooMuchRequests error enum
	TooMuchRequests ErrCode = 200
	// ParametersMissing error enum
	ParametersMissing ErrCode = 201
	// InvalidParameter error enum
	InvalidParameter ErrCode = 202
	// AppNotSupported error enum
	AppNotSupported = 300
	// UserUnauthorized error enum
	UserUnauthorized = 400
	// UserAuthenticationFailed error enum
	UserAuthenticationFailed = 401
)

// ServerError enum.
type ServerError struct {
	code    ErrCode
	message string
	details string
}

func (err ServerError) Error() string {
	return err.message
}

// Code returns code value
func (err ServerError) Code() ErrCode {
	return err.code
}

// Message returns message value
func (err ServerError) Message() string {
	return err.message
}

// Details returns details value
func (err ServerError) Details() string {
	return err.details
}

// Make an immutable error object
func Make(code ErrCode, details string) ServerError {
	message := ""
	switch code {
	case Success:
		message = "success"
	case Internal:
		message = "internal server error"
	case DBValueExists:
		message = "the value is already exist"
	case DBValueNotFound:
		message = "the value was not found"
	case DBValueExpired:
		message = "the value is expired"
	case TooMuchRequests:
		message = "too much requests, you are blocked now"
	case ParametersMissing:
		message = "some parameters are missing"
	case InvalidParameter:
		message = "invalid or malformatted parameter value"
	case AppNotSupported:
		message = "the app type or version is not supported"
	case UserUnauthorized:
		message = "the user is not authorized to access this resource"
	case UserAuthenticationFailed:
		message = "the authentication process failed"
	default:
		message = "unknown error type"
		code = Unknown
	}

	return ServerError{
		code:    code,
		message: message,
		details: details,
	}
}
