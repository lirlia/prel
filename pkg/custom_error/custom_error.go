package custom_error

type CustomError struct {
	HttpStatusCode int
	Title          string
	Message        string
}

type ErrorCode string

const (
	DefaultTitle               string = "Error"
	DefaultMessage             string = "Sorry, something went wrong. <br>Please try again later."
	DefaultMessageBadRequest   string = "Sorry, there was a problem with your request. <br>Please check your input and try again."
	DefaultMessageUnauthorized string = "Sorry, you don't have access to this page. <br>Please sign in first."
	DefaultMessageForbidden    string = "Sorry, you don't have permission to access this page."
	DefaultNotFound            string = "Sorry, the page you are looking for could not be found."
	DefaultInternalServerError string = "Sorry, something went wrong. <br>Please try again later."
)

const (
	InvalidArgument   ErrorCode = "InvalidArgument"
	NotSignIn         ErrorCode = "NotSignIn"
	NotAllowed        ErrorCode = "NotAllowed"
	ResourceNotFound  ErrorCode = "ResourceNotFound"
	Unknown           ErrorCode = "Unknown"
	UnavailableUser   ErrorCode = "UnavailableUser"
	OnlyAdmin         ErrorCode = "OnlyAdmin"
	SessionExpired    ErrorCode = "SessionExpired"
	NotInvited        ErrorCode = "NotInvited"
	AlreadyInvited    ErrorCode = "AlreadyInvited"
	AlreadyRegistered ErrorCode = "AlreadyRegistered"
)

var errMap = map[ErrorCode]CustomError{
	// common
	InvalidArgument:  {400, "Invalid Argument", DefaultMessageBadRequest},
	NotSignIn:        {401, "Not Signed In", DefaultMessageUnauthorized},
	SessionExpired:   {401, "Session Expired", "Sorry, your session has expired. <br>Please sign in again."},
	NotAllowed:       {403, "Not Allowed", DefaultMessageForbidden},
	ResourceNotFound: {404, "404 Not Found", DefaultNotFound},
	Unknown:          {500, "Error", DefaultInternalServerError},

	// user permission
	UnavailableUser: {403, "Account Unavailable", "Sorry, your account is not available. <br>Please contact the administrator."},
	OnlyAdmin:       {403, "Only Admin", "Sorry, you don't have permission to access this page. <br>Please contact the administrator."},

	// invitation
	NotInvited:        {403, "Not Invited", "Sorry, you are not invited. <br>Please contact the administrator."},
	AlreadyInvited:    {400, "Already Invited", "Sorry, requested user is already invited. <br>Please check email."},
	AlreadyRegistered: {400, "Already Registered", "Sorry, requested user is already registered. <br>Please check email."},
}

func GetError(c ErrorCode) (CustomError, bool) {
	e, ok := errMap[c]
	return e, ok
}
