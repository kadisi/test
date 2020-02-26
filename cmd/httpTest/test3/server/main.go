package main

import (
	"fmt"
	"net/http"
	"time"

	rest "github.com/emicklei/go-restful/v3"
)

func main() {
	web := new(rest.WebService)
	web.Path("/rest")
	web.Route(web.GET("/test").To(func(request *rest.Request, response *rest.Response) {
		cookies := request.Request.Cookies()
		fmt.Println("#########")
		for _, c := range cookies {
			fmt.Printf("Get Cookie:%v\n", c.String())
		}
		respC := &http.Cookie{
			Name:    "zhangjie",
			Value:   "hshshshshshssh",
			Expires: time.Now().Add(time.Hour * 24),
		}
		http.SetCookie(response.ResponseWriter, respC)
		response.Write([]byte("hello world\n"))
	}))
	web.Route(web.GET("/check").To(func(request *rest.Request, response *rest.Response) {
		//http.Redirect(response, request.Request, "localhost:9999/rest/test", http.StatusTemporaryRedirect)
		http.Redirect(response, request.Request, "/rest/test", http.StatusTemporaryRedirect)
	}))
	container := rest.NewContainer()
	container.Add(web)

	server := http.Server{
		Addr:    ":9999",
		Handler: container,
	}
	server.ListenAndServe()
}
