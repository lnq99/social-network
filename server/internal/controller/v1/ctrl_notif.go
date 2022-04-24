package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotifications
// @Summary Get notification
// @Description get notification
// @ID get-notif
// @Tags notification
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []model.Notification
// @Failure 500 {object} Msg
// @Router /notif [get]
func (ctrl *Controller) GetNotifications(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	notif, err := ctrl.services.Notification.Get(ID)
	jsonResponse(c, err,
		Response{http.StatusOK, notif},
		ErrResponse{Code: http.StatusInternalServerError})

}
