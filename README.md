# goenv

```bash
# Ignore untracked content of goenv
git config -f .gitmodules submodule.goenv.ignore untracked
```

### Go

```
https://github.com/go-nv/goenv.git
ln -s /opt/go/goenv ~/.goenv

go install -v golang.org/x/tools/gopls@latest  # 安装语言服务器，与VSCode通信

go env                    查看环境变量
go version                查看版本
go build helloworld.go    编译源码
go install helloworld.go  安装到GOBIN
gofmt -w helloworld.go    代码风格格式化
go doc fmt                查看fmt包的文档注释
go doc container/list     查看container/list子包的文档注释
go doc fmt Printf         查看fmt包中Printf函数的文档注释
godoc -http=:6060         浏览器打开http://localhost:6060查看文档
go fix helloworld.go      将go代码从旧版本迁移至新版本
go test helloworld.go     单元测试

go get gopl.io/ch1/helloworld  自动拉取gopl.io包，然后构建并安装子包helloworld到$GOPATH
go run helloworld.go           编译helloworld.go源文件，链接库文件，最终生成临时的可执行文件
```

### Go Modules

[Why is GO111MODULE everywhere, and everything about Go Modules](https://dev.to/maelvls/why-is-go111module-everywhere-and-everything-about-go-modules-24k)

```
go mod init example.com/hello        创建模块，生成go.mod
go get golang.org/x/text             更新模块到最新版
go get rsc.io/sampler                更新模块到最新版
go get rsc.io/sampler@v1.3.1         切换模块的版本
go list -m all                       列出当前模块以及所有依赖，生成go.sum
go list -m -versions rsc.io/sampler  列出模块的可用版本
go list -m rsc.io/q...               列出模块（匹配一部分）
go test                              测试模块
go doc rsc.io/quote/v3               查看模块的文档
go mod tidy                          清理不再使用的模块
```

### containerd-cnitool

```
go get github.com/containernetworking/cni
go install github.com/containernetworking/cni/cnitool
```

### checktool

```bash
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install -v honnef.co/go/tools/cmd/staticcheck@latest
```

