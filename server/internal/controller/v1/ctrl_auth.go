package v1

import (
	"net/http"
	"strconv"

	"app/internal/model"
	"app/internal/service"
	"app/pkg/logger"

	"github.com/gin-gonic/gin"
)

// RegisterHandler
// @Summary Register
// @Description register
// @ID register
// @Tags auth
// @Accept json
// @Produce json
// @Param profile body service.ProfileBody true "Register profile"
// @Success 201
// @Failure 422,500 {object} Msg
// @Router /auth/register [post]
func (ctrl *Controller) RegisterHandler(c *gin.Context) {
	var profileBody service.ProfileBody
	if err := c.ShouldBindJSON(&profileBody); err != nil {
		logger.Err(err)
		c.JSON(http.StatusUnprocessableEntity, Msg{"Invalid json provided"})
		return
	}

	err := ctrl.services.Profile.Register(profileBody)
	jsonResponse(c, err,
		Response{Code: http.StatusCreated},
		ErrResponse{Code: http.StatusInternalServerError})
}

// LoginHandler
// @BasePath /auth
// @Summary Login
// @Description login
// @ID login
// @Tags auth
// @Accept json
// @Produce json
// @Param profile body service.LoginBody true "Login profile"
// @Success 200 {object} loginResponse
// @Failure 401,422 {object} Msg
// Failure 401 {string} string "Email or password is invalid"
// @Router /auth/login [post]
func (ctrl *Controller) LoginHandler(c *gin.Context) {
	var user model.Profile
	id := 0

	token, err := c.Cookie("token")
	if err == nil {
		id, err = ctrl.auth.ParseTokenId(token)
		if err == nil && id > 0 {
			c.Set("ID", id)
			user, _ = ctrl.services.Profile.Get(id)
		}
	}

	if err != nil {
		u := service.LoginBody{}

		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusUnprocessableEntity, Msg{"Invalid json provided"})
			return
		}

		user, _ = ctrl.services.Profile.GetByEmail(u.Email)

		if user.Email != u.Email ||
			!ctrl.auth.ComparePassword(u.Password, user.Salt, user.Hash) {
			c.JSON(http.StatusUnauthorized, Msg{"Email or password is invalid"})
			return
		}

		token, err = ctrl.auth.CreateToken(strconv.Itoa(user.Id))
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, Msg{err.Error()})
			return
		}
	}

	c.SetCookie("token", token, 60*60*24, "/", ctrl.conf.Host, true, true)
	c.JSON(http.StatusOK, loginResponse{token, toProfileResponse(user)})
}

// LogoutHandler
// @BasePath /auth
// @Summary Logout
// @Description logout
// @ID logout
// @Tags auth
// @Security ApiKeyAuth
// @Success 200
// @Router /auth/logout [delete]
func (ctrl *Controller) LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", ctrl.conf.Host, true, true)
	c.Status(http.StatusOK)
}
