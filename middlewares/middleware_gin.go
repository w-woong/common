package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-wonk/si/v2"
	"github.com/go-wonk/si/v2/sicore"
	"github.com/w-woong/common/logger"
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

func AuthBearerTokenGin(log core.Logger, token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		r := c.Request

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Error(http.StatusText(http.StatusBadRequest), logger.UrlField(r.URL.String()))
			c.Abort()
			return
		}

		authVal := strings.Split(authHeader, " ")
		if len(authVal) != 2 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Error(http.StatusText(http.StatusBadRequest), logger.UrlField(r.URL.String()))
			c.Abort()
			return
		}

		if authVal[1] != token {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			log.Error(http.StatusText(http.StatusUnauthorized), logger.UrlField(r.URL.String()))
			c.Abort()
			return
		}

		c.Next()
	}
}

// AuthHmacSha256Gin to verify the request with hmacsha256
func AuthHmacSha256Gin(hmacHeader string, hmacKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		r := c.Request

		var reqBytes []byte
		var err error

		if r.Body != nil {
			reqBytes, err = si.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				logger.Error(err.Error(), logger.UrlField(r.URL.String()))
				c.Abort()
				return
			}
		}

		var msg []byte
		if len(reqBytes) == 0 {
			msg = []byte(r.URL.Path)
		} else {
			msg = reqBytes
		}

		hmacHexStr, err := sicore.HmacSha256HexEncoded(string(hmacKey), msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			logger.Error(err.Error(), logger.UrlField(r.URL.String()))
			c.Abort()
			return
		}

		if hmacHexStr != r.Header.Get(hmacHeader) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			logger.Error(fmt.Sprintf("generated hmac value %v is invalid(expected: %v)", r.Header.Get(hmacHeader), hmacHexStr),
				logger.UrlField(r.URL.String()), logger.ReqBodyField(reqBytes))
			c.Abort()
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(reqBytes))
		c.Next()
	}
}
