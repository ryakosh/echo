package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Any("/echo/*path", echo)

	r.Run(":8080")
}

func echo(c *gin.Context) {
	resp := Echo{
		Version: c.Request.Proto,
		Request: Request{
			Method: c.Request.Method,
			URL: URL{
				Raw:  c.Request.URL.String(),
				Path: c.Param("path"),
				Query: Query{
					Raw:    c.Request.URL.RawQuery,
					Parsed: c.Request.URL.Query(),
				},
			},
			Headers: c.Request.Header,
		},
		Response: Response{
			Status: Status{
				Code:    c.Writer.Status(),
				Message: http.StatusText(c.Writer.Status()),
			},
			Headers: c.Writer.Header(),
		},
	}

	c.JSON(http.StatusOK, resp)
}
