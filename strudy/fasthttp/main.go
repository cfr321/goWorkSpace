package main

import (
	"encoding/base64"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"strings"
)

func basicAuth(ctx *fasthttp.RequestCtx) (username, password string, ok bool) {
	auth := ctx.Request.Header.Peek("Authorization")
	if auth == nil {
		return
	}
	return parseBasicAuth(string(auth))
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

func BasicAuth(h fasthttp.RequestHandler, requiredUser, requiredPassword string) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := basicAuth(ctx)
		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			h(ctx)
			return
		}
		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
		ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
	})
}

// Protected is the Protected handler
func Protected(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Protected!\n")
}

func main() {
	router := fasthttprouter.New()
	router.GET("/Hello/:name", hello())

	err := fasthttp.ListenAndServe(":8080", router.Handler)
	if err != nil {
		fmt.Println("error to start")
	}
}

func hello() func(ctx *fasthttp.RequestCtx) {
	//engine := gin.Default()
	//engine.Run()
	return func(ctx *fasthttp.RequestCtx) {
		_, err := fmt.Fprint(ctx, "hello")
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(ctx.UserValue("name"))
	}
}
