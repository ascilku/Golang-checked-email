package member

import (
	"go-5/errors"
	"go-5/respons"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	services Services
}

func NewHandler(services Services) *handler {
	return &handler{services}
}

func (h *handler) SaveHandler(g *gin.Context) {
	var inputMember InputMember
	err := g.ShouldBindJSON(&inputMember)
	if err != nil {
		// var errors []string
		// for _, e := range err.(validator.ValidationErrors) {
		// 	errors = append(errors, e.Error())
		// }
		ErrorMessage := errors.ErrorMessage(err)
		errorMessage := gin.H{"errors": ErrorMessage}
		APIRespons := respons.APIRespons("Filed Input", http.StatusUnprocessableEntity, "Filed", errorMessage)
		g.JSON(http.StatusBadRequest, APIRespons)
	} else {
		newHandler, err := h.services.SaveServices(inputMember)
		if err != nil {
			g.JSON(http.StatusBadRequest, err)
		} else {
			formatter := Formatter(newHandler, "token token")
			APIRespons := respons.APIRespons("Input Success", http.StatusOK, "Success", formatter)
			g.JSON(http.StatusOK, APIRespons)
		}
	}
}

func (h *handler) LoginHandler(g *gin.Context) {
	var loginMember LoginMember
	err := g.ShouldBindJSON(&loginMember)
	if err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
		errorMessage := gin.H{"errors": errors}
		newAPIRespons := respons.APIRespons("Filed Login", http.StatusUnprocessableEntity, "Filed", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, newAPIRespons)
	} else {
		newFindByEmailService, err := h.services.FindByEmailService(loginMember)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			newAPIRespons := respons.APIRespons("filed Login", http.StatusBadRequest, "filed", errorMessage)
			g.JSON(http.StatusBadRequest, newAPIRespons)
		} else {
			formatter := Formatter(newFindByEmailService, "Token Token")
			newAPIRespons := respons.APIRespons("Success Login", http.StatusOK, "Success", formatter)
			g.JSON(http.StatusOK, newAPIRespons)
		}
	}
}

func (h *handler) CheckEmailAvailability(g *gin.Context) {
	var checkEmailInput CheckEmailInput
	err := g.ShouldBindJSON(&checkEmailInput)

	if err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
		errorMessage := gin.H{"errors": errors}
		APIRespons := respons.APIRespons("Email Checking Error", http.StatusBadRequest, "error", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, APIRespons)
	} else {
		newIsEmailAvailable, err := h.services.IsEmailAvailable(checkEmailInput)

		if err != nil {
			errorMessage := gin.H{"errors": "Server Error"}
			APIRespons := respons.APIRespons("Email Checking Error", http.StatusBadRequest, "error", errorMessage)
			g.JSON(http.StatusBadRequest, APIRespons)
		} else {
			data := gin.H{"is_available": newIsEmailAvailable}

			// cara pertama
			var metaMessage string
			if newIsEmailAvailable {
				metaMessage = "Email Is Available "
			} else {
				metaMessage = "Email Has Been Registered"
			}

			// cara kedua
			// metaMessage := "Email Has Been Registered"
			// if newIsEmailAvailable {
			// 	metaMessage = "Email Is Available"
			// }
			APIRespons := respons.APIRespons(metaMessage, http.StatusOK, "Success", data)
			g.JSON(http.StatusOK, APIRespons)
		}
	}
}
