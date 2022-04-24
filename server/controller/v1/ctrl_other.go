package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"app/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Feed
// @Summary Feed
// @Description feed
// @ID feed
// @Tags feed
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param lim query int true "Limit"
// @Param off query int true "Offset"
// @Success 200 {object} []int64
// @Failure 500 {object} Msg
// @Router /feed [get]
func (ctrl *Controller) Feed(c *gin.Context) {
	//id := toInt(c.Param("id"))
	ID := c.MustGet("ID").(int)
	limit := toInt(c.Query("lim"))
	offset := toInt(c.Query("off"))
	feed, err := ctrl.services.Feed.GetFeed(ID, limit, offset)
	log.Println(ID, limit, offset, feed, err)
	jsonResponse(c, err,
		Response{http.StatusOK, feed},
		ErrResponse{Code: http.StatusInternalServerError})
}

// Search
// @Summary Search by username
// @Description search by username
// @ID search
// @Tags search
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query int true "Key"
// @Success 200 {object} []SearchResponse
// @Failure 500 {object} Msg
// @Router /search [get]
func (ctrl *Controller) Search(c *gin.Context) {
	ID := c.MustGet("ID").(int)
	key := c.Query("k")
	res, err := ctrl.services.Profile.SearchName(ID, key)
	if err != nil {
		logger.Err(err)
		c.JSON(http.StatusInternalServerError, Msg{err.Error()})
	}

	var s interface{}
	err = json.Unmarshal([]byte(res), &s)
	jsonResponse(c, err,
		Response{http.StatusOK, s},
		ErrResponse{Code: http.StatusInternalServerError})
}

func (ctrl *Controller) HandleNoRoute(c *gin.Context) {
	path := c.Request.URL.Path
	if strings.HasPrefix(path, "/api") {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.FileFromFS("/", http.Dir(ctrl.conf.StaticRoot))
	}
}
