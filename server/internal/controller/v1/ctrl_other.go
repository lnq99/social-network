package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"app/pkg/logger"

	"github.com/gin-gonic/gin"
)

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
