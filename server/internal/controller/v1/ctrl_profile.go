package v1

import (
	"net/http"

	"app/internal/model"
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetProfile(c *gin.Context) {
	id := toInt(c.Param("id"))
	profile, err := ctrl.services.Profile.Get(id)

	jsonResponse(c, err,
		Response{http.StatusOK, toProfileResponse(profile)},
		ErrResponse{Code: http.StatusInternalServerError})
}

func (ctrl *Controller) GetShortProfile(c *gin.Context) {
	id := toInt(c.Param("id"))
	profile, err := ctrl.services.Profile.Get(id)

	jsonResponse(c, err,
		Response{http.StatusOK, model.ShortInfo{id, profile.Name, profile.AvatarS}},
		ErrResponse{Code: http.StatusInternalServerError})
}

func (ctrl *Controller) ChangeIntro(c *gin.Context) {
	var introBody service.IntroBody
	ID := c.MustGet("ID").(int)
	if err := c.ShouldBindJSON(&introBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Msg{"Invalid json provided"})
		return
	}
	err := ctrl.services.Profile.ChangeIntro(ID, introBody)
	jsonResponse(c, err,
		Response{Code: http.StatusOK},
		ErrResponse{Code: http.StatusInternalServerError})
}
