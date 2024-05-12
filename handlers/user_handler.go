package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nahdukesaba/be-assignment/helpers"
	"github.com/nahdukesaba/be-assignment/repo"
	"github.com/nahdukesaba/be-assignment/services"
	"github.com/nahdukesaba/be-assignment/services/request"
)

type UserHandler struct {
	userService services.Users
}

func NewUserHandler(db *repo.DB) *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(db),
	}
}

// Login Login the user
//
//	@Summary		Login user
//	@Description	login the user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		request.UserRequest	true	"Update account"
//	@Success		200		{object}	response.UserLoginUser
//	@Failure		400		{object}	helpers.Message
//	@Failure		404		{object}	helpers.Message
//	@Failure		500		{object}	helpers.Message
//	@Router			/api/user/login [post]
func (uh *UserHandler) Login(c *gin.Context) {
	var form *request.UserRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&form); err != nil {
		log.Printf("Error reading body request: %v\n", err)
		c.JSON(http.StatusBadRequest, helpers.BadRequesetMessage)
		return
	}

	res, err := uh.userService.LoginUser(c, form)
	if err != nil {
		log.Printf("Bad request: %v\n", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// Register Register the user
//
//	@Summary		Register user
//	@Description	register the user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		request.UserRequest	true	"Update account"
//	@Success		200		{object}	helpers.Message
//	@Failure		400		{object}	helpers.Message
//	@Failure		404		{object}	helpers.Message
//	@Failure		500		{object}	helpers.Message
//	@Router			/api/user/register [post]
func (uh *UserHandler) Register(c *gin.Context) {
	var form *request.UserRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&form); err != nil {
		log.Printf("Error reading body request: %v\n", err)
		c.JSON(http.StatusBadRequest, helpers.BadRequesetMessage)
		return
	}

	if err := uh.userService.RegisterUser(c, form); err != nil {
		log.Printf("Bad request: %v\n", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, helpers.SuccessMessage)
}
