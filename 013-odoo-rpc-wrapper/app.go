package main

import (
	"alexejk.io/go-xmlrpc"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	client, _ := xmlrpc.NewClient("http://erp.gaowei.com")
	client.Call()
	result := &struct {
		Bugzilla struct {
			Version string
		}
	}{}
	_ = client.Call("Bugzilla.version", nil, result)
	fmt.Printf("Version: %s\n", result.Bugzilla.Version)
	r := gin.Default()

	print("abc")
	r.Run(":8080")

	
}
