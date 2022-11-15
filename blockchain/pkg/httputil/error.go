package httputil

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewError sets error code and status to response.
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

func ProcessHttpError(ctx *gin.Context, e HTTPError) {
	ctx.JSON(e.Code, e)
}

// HTTPError represents a general http error.
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message" example:"status bad request"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func PasswordGenerationError() HTTPError {
	return HTTPError{
		Code:    http.StatusUnprocessableEntity,
		Message: "fail to generate user password hash",
	}
}

func PasswordValidationError() HTTPError {
	return HTTPError{
		Code:    http.StatusUnprocessableEntity,
		Message: "fail to validate user password",
	}
}

func FailToCreateAccessTokenError() HTTPError {
	return HTTPError{
		Code:    http.StatusUnprocessableEntity,
		Message: "fail to create access token",
	}
}

func FailToCreateRefreshTokenError() HTTPError {
	return HTTPError{
		Code:    http.StatusUnprocessableEntity,
		Message: "fail to create refresh token",
	}
}

func FailToCreateNewUserError(err error) HTTPError {
	return HTTPError{
		Code:    http.StatusUnprocessableEntity,
		Message: fmt.Sprintf("fail to create new user because %s", err.Error()),
	}
}

func UserNotFoundWithUsernameError(username string) HTTPError {
	return HTTPError{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("fail to find user with username %s", username),
	}
}
