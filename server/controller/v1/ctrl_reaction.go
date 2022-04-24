package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReaction
// @Summary Get reaction of post
// @Description get reaction by post id
// @ID get-reaction-by-post-id
// @Tags reaction
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param post_id path int true "Post ID"
// @Success 200 {array} int64
// @Failure 500 {object} Msg
// @Router /react/{post_id} [get]
func (ctrl *Controller) GetReaction(c *gin.Context) {
	react, err := ctrl.services.Post.GetReaction(toInt(c.Param("post_id")))
	jsonResponse(c, err,
		Response{http.StatusOK, react},
		ErrResponse{Code: http.StatusInternalServerError})
}

// GetReactionByUserPost
// @Summary Get user's reaction of post
// @Description get reaction by user id and post id
// @ID get-reaction-by-user-post-id
// @Tags reaction
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param post_id path int true "Post ID"
// @Success 200 {object} dataResponse
// @Failure 500 {object} Msg
// @Router /react/u/{post_id} [get]
func (ctrl *Controller) GetReactionByUserPost(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	react, err := ctrl.services.Reaction.GetByUserPost(ID, toInt(c.Param("u_id")))
	jsonResponse(c, err,
		Response{http.StatusOK, dataResponse{react}},
		ErrResponse{Code: http.StatusInternalServerError})
}

// PutReaction
// @Summary Get reaction of post
// @Description get reaction by post id
// @ID put-reaction
// @Tags reaction
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param post_id path int true "Post ID"
// @Param type path int true "Reaction type"
// @Success 200
// @Failure 500 {object} Msg
// @Router /react/{post_id}/{type} [put]
func (ctrl *Controller) PutReaction(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	postId := toInt(c.Param("post_id"))
	t := c.Param("type")
	err := ctrl.services.Reaction.UpdateReaction(ID, postId, t)
	jsonResponse(c, err,
		Response{Code: http.StatusOK},
		ErrResponse{Code: http.StatusInternalServerError})
}
