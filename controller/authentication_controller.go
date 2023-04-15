package controller

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hy-reza/mygram-api/data/request"
	"github.com/hy-reza/mygram-api/data/response"
	"github.com/hy-reza/mygram-api/helper"
	"github.com/hy-reza/mygram-api/service"

	"github.com/gin-gonic/gin"
)

var validate = validator.New()

type AuthenticationController struct {
	authenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: service}
}

func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)

	token, err_token := controller.authenticationService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	// ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUserRequest{}
	ctx.ShouldBindJSON(&createUsersRequest)
	err := validate.Struct(createUsersRequest)
	if err != nil {
		helper.ErrorBinding(ctx, err, http.StatusBadRequest, "Registration Failed!")
		return
	}

	savedUser, err := controller.authenticationService.Register(createUsersRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Registration Failed!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    savedUser,
	})
}
