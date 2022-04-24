package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetReaction(c *gin.Context) {
	react, err := ctrl.services.Post.GetReaction(toInt(c.Param("post_id")))
	jsonResponse(c, err,
		Response{http.StatusOK, react},
		ErrResponse{Code: http.StatusInternalServerError})
}

func (ctrl *Controller) GetReactionByUserPost(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	react, err := ctrl.services.Reaction.GetByUserPost(ID, toInt(c.Param("u_id")))
	jsonResponse(c, err,
		Response{http.StatusOK, dataResponse{react}},
		ErrResponse{Code: http.StatusInternalServerError})
}

func (ctrl *Controller) PutReaction(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	postId := toInt(c.Param("post_id"))
	t := c.Param("type")
	err := ctrl.services.Reaction.UpdateReaction(ID, postId, t)
	jsonResponse(c, err,
		Response{Code: http.StatusOK},
		ErrResponse{Code: http.StatusInternalServerError})
}
