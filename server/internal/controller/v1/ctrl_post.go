package v1

import (
	"net/http"

	"app/internal/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetPost(c *gin.Context) {
	post, err := ctrl.services.Post.Get(toInt(c.Param("id")))

	jsonResponse(c, err,
		Response{http.StatusOK, post},
		ErrResponse{Code: http.StatusInternalServerError})
}

func (ctrl *Controller) GetPostByUserId(c *gin.Context) {
	postsId, err := ctrl.services.Post.GetByUserId(toInt(c.Param("id")))

	jsonResponse(c, err,
		Response{http.StatusOK, postsId},
		ErrResponse{Code: http.StatusInternalServerError})
}

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

func (ctrl *Controller) DeletePost(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	id := toInt(c.Param("id"))
	err := ctrl.services.Post.Delete(ID, id)
	statusResponse(c, err, http.StatusOK, http.StatusInternalServerError)
}
