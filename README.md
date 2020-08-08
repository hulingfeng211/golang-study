# golang 学习项目

- 启动nacos

下载
[https://github.com/alibaba/nacos/releases/download/1.3.1/nacos-server-1.3.1.zip](https://github.com/alibaba/nacos/releases/download/1.3.1/nacos-server-1.3.1.zip)

启动

```
C:\Users\15921\Desktop\nacos-server-1.3.1\nacos\bin\startup.cmd
```

- 设置go中国代理
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct

# 设置不走 proxy 的私有仓库，多个用逗号相隔（可选）
go env -w GOPRIVATE=*.corp.example.com

# 设置不走 proxy 的私有组织（可选）
go env -w GOPRIVATE=example.com/org_name

```

