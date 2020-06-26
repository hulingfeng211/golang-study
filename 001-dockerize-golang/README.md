# golang docker 容器化

- 初始化go mod init
```cgo
go mod init github.com/hulingfeng211/golang-study/001-dockerize-golang
```

- 创建入口文件main.go

- 编译文件&打包镜像
```text
go build -o main .
docker build -t example-scratch -f Dockerfile.scratch .
```