package v1

import (
	"net/http"

	"app/internal/service"

	"github.com/gin-gonic/gin"
)

// GetPost
// @Summary Get post
// @Description get post by ID
// @ID get-post-by-id
// @Tags post
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Post ID"
// @Success 200 {object} model.Post
// @Failure 500 {object} Msg
// @Router /post/{id} [get]
func (ctrl *Controller) GetPost(c *gin.Context) {
	post, err := ctrl.services.Post.Get(toInt(c.Param("id")))

	jsonResponse(c, err,
		Response{http.StatusOK, post},
		ErrResponse{Code: http.StatusInternalServerError})
}

// GetPostByUserId
// @Summary Get id of posts by user id
// @Description get id of posts by user id
// @ID get-post-by-user-id
// @Tags post
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Success 200 {array} int64
// @Failure 500 {object} Msg
// @Router /post/u/{id} [get]
func (ctrl *Controller) GetPostByUserId(c *gin.Context) {
	postsId, err := ctrl.services.Post.GetByUserId(toInt(c.Param("id")))

	jsonResponse(c, err,
		Response{http.StatusOK, postsId},
		ErrResponse{Code: http.StatusInternalServerError})
}

// PostPost
// @Summary Post a post
// @Description post a post
// @ID post-post
// @Tags post
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param  account body service.PostBody true "Post body"
// @Success 200
// @Failure 422,500 {object} Msg
// @Router /post [post]
func (ctrl *Controller) PostPost(c *gin.Context) {
	var postBody service.PostBody
	ID := c.MustGet("ID").(int)
	if err := c.ShouldBindJSON(&postBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, Msg{"Invalid json provided"})
		return
	}
	err := ctrl.services.Post.Post(ID, postBody)

	jsonResponse(c, err,
		Response{Code: http.StatusCreated},
		ErrResponse{Code: http.StatusInternalServerError})
}

// DeletePost
// @Summary Delete a post
// @Description delete a post
// @ID delete-post
// @Tags post
// @Security ApiKeyAuth
// @Param id path int true "Post ID"
// @Success 200
// @Failure 500 {object} Msg
// @Router /post/{id} [delete]
func (ctrl *Controller) DeletePost(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := toInt(c.Param("id"))
	err := ctrl.services.Post.Delete(ID, id)
	statusResponse(c, err, http.StatusOK, http.StatusInternalServerError)
}
