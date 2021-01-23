# 
```
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

# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](doc/gopher/fiveyears.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 3.0 Attributions license][cc3-by].*

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Unless otherwise noted, the Go source files are distributed under the
BSD-style license found in the LICENSE file.

### Download and Install

#### Binary Distributions

Official binary distributions are available at https://golang.org/dl/.

After downloading a binary release, visit https://golang.org/doc/install
or load [doc/install.html](./doc/install.html) in your web browser for installation
instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://golang.org/doc/install/source or load [doc/install-source.html](./doc/install-source.html)
in your web browser for source installation instructions.

### Contributing

Go is the work of thousands of contributors. We appreciate your help!

To contribute, please read the contribution guidelines:
	https://golang.org/doc/contribute.html

Note that the Go project uses the issue tracker for bug reports and
proposals only. See https://golang.org/wiki/Questions for a list of
places to ask questions about the Go language.

[rf]: https://reneefrench.blogspot.com/
[cc3-by]: https://creativecommons.org/licenses/by/3.0/

### Workspace
git clone https://github.com/imssyang/mygo.git my

