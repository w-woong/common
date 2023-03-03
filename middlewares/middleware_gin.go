package middlewares

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/w-woong/common/logger/core"
	"github.com/w-woong/common/utils"
)

func NoCacheGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.SetNoCache(c.Writer)
		c.Next()
	}
}

func LoggerGin(log core.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		gw := &ginWriter{c.Writer, &bytes.Buffer{}}
		c.Writer = gw
		r := c.Request
		b, err := utils.RequestDump(r)
		if err != nil {
			log.Error(err.Error())
		}

		c.Next()

		log.Debug("transaction", core.WithBytesField("request", b), core.WithBytesField("response", gw.body.Bytes()))

	}
}

type ginWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *ginWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w *ginWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
