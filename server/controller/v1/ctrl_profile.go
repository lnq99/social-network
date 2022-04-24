package v1

import (
	"net/http"

	"app/internal/model"
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

// GetProfile
// @Summary Get profile
// @Description get profile by ID
// @ID get-profile-by-id
// @Tags profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Profile ID"
// @Success 200 {object} ProfileResponse
// @Failure 500 {object} Msg
// @Router /profile/{id} [get]
func (ctrl *Controller) GetProfile(c *gin.Context) {
	id := toInt(c.Param("id"))
	profile, err := ctrl.services.Profile.Get(id)

	jsonResponse(c, err,
		Response{http.StatusOK, toProfileResponse(profile)},
		ErrResponse{Code: http.StatusInternalServerError})
}

// GetShortProfile
// @Summary Get brief of profile
// @Description get brief profile by ID
// @ID get-brief-profile-by-id
// @Tags profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Profile ID"
// @Success 200 {object} model.ShortInfo
// @Failure 500 {object} Msg
// @Router /profile/short/{id} [get]
func (ctrl *Controller) GetShortProfile(c *gin.Context) {
	id := toInt(c.Param("id"))
	profile, err := ctrl.services.Profile.Get(id)

	jsonResponse(c, err,
		Response{http.StatusOK, model.ShortInfo{id, profile.Name, profile.AvatarS}},
		ErrResponse{Code: http.StatusInternalServerError})
}

// ChangeIntro
// @Summary Change intro of profile
// @Description change intro profile by ID
// @ID change-intro-profile-by-id
// @Tags profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param  account body service.IntroBody true "Intro body"
// @Success 200
// @Failure 500 {object} Msg
// @Router /profile/intro [patch]
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
