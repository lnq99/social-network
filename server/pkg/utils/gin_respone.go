package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int
	Obj  interface{}
}

type ErrResponse struct {
	Code int    `json:"-"`
	Msg  string `json:"message"`
}

type Msg struct {
	Msg string `json:"message"`
}

func JsonResponse(c *gin.Context, err error, res Response, errRes ErrResponse) {
	if err != nil {
		if errRes.Msg == "" {
			errRes.Msg = err.Error()
		}
		//logger.Err(err)
		c.JSON(errRes.Code, errRes)
	} else if res.Obj != nil {
		c.JSON(res.Code, res.Obj)
	} else {
		c.Status(res.Code)
	}
}

func StatusResponse(c *gin.Context, err error, code int, errCode int) {
	if err != nil {
		//logger.Err(err)
		c.Status(errCode)
	} else {
		c.Status(code)
	}
}
