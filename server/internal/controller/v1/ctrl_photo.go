package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetPhoto(c *gin.Context) {
	photo, err := ctrl.services.Photo.GetPhoto(toInt(c.Param("id")))
	jsonResponse(c, err,
		Response{http.StatusCreated, photo},
		ErrResponse{Code: http.StatusInternalServerError})

}

func (ctrl *Controller) GetPhotoByUserId(c *gin.Context) {
	photo, err := ctrl.services.Photo.GetPhotoByUserId(toInt(c.Param("id")))
	jsonResponse(c, err,
		Response{http.StatusCreated, photo},
		ErrResponse{Code: http.StatusInternalServerError})
}
