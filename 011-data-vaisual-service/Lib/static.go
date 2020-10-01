package Lib

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
)

//DEFAULT INDEX FILE
const INDEX = "index.html"

type ServeFileSystem interface {
	http.FileSystem
	Exists(prefix, path string) bool
}

type localFileSystem struct {
	http.FileSystem
	root    string
	indexes bool
}

func LocalFile(root string, indexes bool) *localFileSystem {
	return &localFileSystem{
		FileSystem: gin.Dir(root, indexes),
		root:       root,
		indexes:    indexes,
	}
}

func (l *localFileSystem) Exists(prefix string, filePath string) bool {
	if p := strings.TrimPrefix(filePath, prefix); len(p) < len(filePath) {
		name := path.Join(l.root, p)
		status, err := os.Stat(name)
		if err != nil {
			return false
		}
		if status.IsDir() {
			if !l.indexes {
				index := path.Join(name, INDEX)
				_, err := os.Stat(index)
				if err != nil {
					return false
				}
			}
		}
		return true
	}
	return false
}
func StaticServe(urlPrefix string, fs ServeFileSystem) gin.HandlerFunc {

	fileServer := http.FileServer(fs)
	if urlPrefix != "" {
		fileServer = http.StripPrefix(urlPrefix, fileServer)
	}
	return func(context *gin.Context) {
		if fs.Exists(urlPrefix, context.Request.URL.Path) {
			fileServer.ServeHTTP(context.Writer, context.Request)
			context.Abort()
		}
	}

}

func ServerRoot(urlPrefix, root string) gin.HandlerFunc {
	return StaticServe(urlPrefix, LocalFile(root, false))
}
