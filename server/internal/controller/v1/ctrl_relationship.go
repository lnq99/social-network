package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetFriendsDetail(c *gin.Context) {
	id := toInt(c.Param("id"))
	friends, err := ctrl.services.Relationship.FriendsDetail(id)
	var s interface{}
	json.Unmarshal([]byte(friends), &s)
	jsonResponse(c, err,
		Response{http.StatusOK, s},
		ErrResponse{Code: http.StatusInternalServerError})

}

func (ctrl *Controller) GetMutualFriends(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := toInt(c.Param("id"))
	mf, err := ctrl.services.Relationship.MutualFriends(ID, id)
	jsonResponse(c, err,
		Response{http.StatusOK, dataResponse{mf}},
		ErrResponse{Code: http.StatusInternalServerError})
}

func (ctrl *Controller) ChangeType(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := toInt(c.Param("id"))
	t := c.Param("type")
	err := ctrl.services.Relationship.ChangeType(ID, id, t)
	jsonResponse(c, err,
		Response{Code: http.StatusOK},
		ErrResponse{Code: http.StatusInternalServerError})
}

func (ctrl *Controller) GetMutualAndType(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := toInt(c.Param("id"))
	t := ctrl.services.Relationship.GetRelationshipWith(ID, id)
	m, err := ctrl.services.Relationship.MutualFriends(ID, id)

	jsonResponse(c, err,
		Response{http.StatusOK, GetMutualAndTypeResponse{t, m}},
		ErrResponse{Code: http.StatusInternalServerError})
}
