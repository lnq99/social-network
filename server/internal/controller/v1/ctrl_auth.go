package v1

import (
	"net/http"
	"strconv"

	"app/internal/model"
	"app/internal/service"
	"app/pkg/logger"

	"github.com/gin-gonic/gin"
)

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

func (ctrl *Controller) LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", ctrl.conf.Host, true, true)
	c.Status(http.StatusOK)
}
