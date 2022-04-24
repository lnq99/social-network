package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetNotifications(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	notif, err := ctrl.services.Notification.Get(ID)
	jsonResponse(c, err,
		Response{http.StatusOK, notif},
		ErrResponse{Code: http.StatusInternalServerError})

}
