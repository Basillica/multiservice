package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/x/multiservice/gateway/types/auth"
)

func Gateway(c *gin.Context) {
	accessToken, cErr := c.Cookie("access_token")
	if cErr != nil {
		c.String(http.StatusUnauthorized, "you do not have the access rights to access this endpoint")
		c.Abort()
		return
	}
	t := auth.Token{
		Token: accessToken,
	}
	tErr := t.ValidateToken()
	if tErr != nil {
		c.String(http.StatusForbidden, "you not have valid access rights to this endpoint")
		c.Abort()
		return
	}

	switch c.Request.URL.Path {
	case "/parsejson":
		proxy(c, "/parsejson", "http://localhost:8081/parsejson")
		return
	case "/algorithm":
		proxy(c, "/algorithm", "http://localhost:8082/algorithm")
		return
	default:
		// c.JSON(http.StatusBadGateway, gin.H{"status": "authorzed"})
		c.JSON(http.StatusOK, gin.H{"status": "authorzed"})
		return
	}
}

func proxy(c *gin.Context, path, scheme string) {
	remote, err := url.Parse(scheme)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = path
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
