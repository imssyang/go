#!/bin/bash
# https://golang.org/doc/install/source#introduction
# https://go-zh.org/doc/
#
# ./go env
# GOCACHE="/root/.cache/go-build"
# GOENV="/root/.config/go/env"
# GOMODCACHE="/root/go/pkg/mod"
# GOPATH="/root/go"
# GOPROXY="https://proxy.golang.org,direct"
# GOROOT="/opt/go"
# GOTOOLDIR="/opt/go/pkg/tool/linux_amd64"

#export GOROOT_BOOTSTRAP=DIR14
#export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
#go env -w GOBIN=/somewhere/else/bin
#go env -u GOBIN

export GO_DISABLE_ENV=yes
export GOROOT=/opt/go
export GOPATH=$GOROOT/app
export GOBIN=$GOPATH/bin
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

export GO111MODULE=auto
export GOPROXY=https://goproxy.cn,direct
#export GOPROXY=https://mirrors.aliyun.com/goproxy/

