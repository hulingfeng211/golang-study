package room

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

func main() {
	//demo1()
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.GET("/html", func(context *gin.Context) {
		context.HTML(200, "index.html", "hello,world")
	})
	r.Run(":8080")
}

func demo1() {
	r := gin.Default()

	r.GET("/html", func(c *gin.Context) {
		c.Status(200)
		const templateText = `george {{ printf "%s" .}}`
		tmpl, err := template.New("htmlTest").Parse(templateText)
		if err != nil {
			log.Fatal("parsing: %s", err)
		}
		tmpl.Execute(c.Writer, "hello,world")
	})
	r.Run(":8080")
}
