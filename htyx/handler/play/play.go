package play

import (
	"htyx/handler"
	"htyx/lib/errno"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func PlayStream(c *gin.Context) {
	streamname := c.Param("streamname")
	streamtype := c.Param("streamtype")
	filepath := "./" + streamtype + "/" + streamname
	stream, err := os.Open(filepath)
	if err != nil {
		handler.SendResponse(c, errno.ErrStream, nil)

		return
	}

	http.ServeContent(c.Writer, c.Request, "", time.Now(), stream)

	defer stream.Close()
}
