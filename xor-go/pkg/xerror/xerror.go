package xerror

import (
	"fmt"
	"net/http"
	"xor-go/pkg/xapp"
)

const (
	appUnknown = "unknown_app"
	verUnknown = "unknown_version"
	ErrUnknown = "unknown_error"
)

var appConfig *xapp.Config

func init() {
	appConfig = &xapp.Config{
		Service:     appUnknown,
		Environment: xapp.Production,
		Version:     verUnknown,
	}
}

func InitAppError(cfg *app.Config) error {
	if cfg == nil {
		appConfig = &app.Config{
			Name:        appUnknown,
			Environment: app.Production,
			Version:     verUnknown,
		}
		return errors.New("application config is nil")
	}
	cfgCopy := *cfg
	appConfig = &cfgCopy
	return nil
}

type AppError struct {
	Code       int    // HTTP status code
	Message    string // Safe message for HTTP response
	DevMessage string // Developer/debugging message
	Err        error  // Underlying error from the repository or other layers
}

func NewAppError(code int, message string, devMessage string, err error) *AppError {
	return &AppError{Code: code, Message: message, DevMessage: devMessage, Err: err}
}

func (a AppError) Unwrap() error {
	return a.Err
}

func (a AppError) Error() string {
	return fmt.Sprintf("[%d %s]: %s: %s", a.Code, a.Message, a.DevMessage, a.Err.Error())
}

func GetLastMessage(err error) string {
	var myErr AppError
	if errors.As(err, &myErr) {
		if appConfig.IsProduction() {
			return myErr.Message
		} else if appConfig.IsLocal() {
			return myErr.DevMessage
		}
		return myErr.Message
	} else {
		if appConfig.IsLocal() {
			return err.Error()
		}
		return ErrUnknown
	}
}

func GetCode(err error) int {
	var myErr AppError
	if errors.As(err, &myErr) {
		return myErr.Code
	}
	return http.StatusInternalServerError
}
