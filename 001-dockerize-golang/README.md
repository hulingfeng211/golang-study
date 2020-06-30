# golang docker 容器化

- 初始化go mod init
```cgo
go mod init github.com/hulingfeng211/golang-study/001-dockerize-golang
```

- 创建入口文件main.go

- 编译文件&打包镜像（ 模式1）
> 通过golang的镜像进行编译打包工作
```text
docker build -t example-scratch -f Dockerfile.multistage  .
```

运行
```
docker build --rm -it example-scratch
```
- 编译文件&打包镜像（ 模式2）
>通过本地编译，然后copy编译后的文件到image中

编译
cd /模块所在目录下
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```
打包
```
docker build -t example-scratch -f Dockerfile.scratch  .
```
 

运行
```
docker build --rm -it example-scratch
```
 

