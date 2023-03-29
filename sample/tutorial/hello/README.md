
# 创建go.mod
go mod init example/hello

# 获取依赖包
go get rsc.io/quote

# 指定本地路径
go mod edit -replace example.com/greetings=../greetings

# 同步模块依赖
go mod tidy

# 运行
go run .

# 编译
go build

# 查看安装目录
go list -f '{{.Target}}'

# 安装到GOBIN
go install
