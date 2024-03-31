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

export GOENV_ROOT="/opt/go/goenv"
export GOENV_GOPATH_PREFIX="/opt/go/gomod"
export GOENV_GOMOD_VERSION_ENABLE=1
eval "$(goenv init - zsh)"
path=($GOROOT/bin $path $GOPATH/bin)
export GO111MODULE=on
export GOPROXY=https://goproxy.cn # QiNiuYun
#export GOPROXY=https://mirrors.aliyun.com/goproxy/

#export GOROOT=/opt/go
#export GOPATH=$GOROOT/app
#export GOOS=linux
#export GOARCH=amd64
#export GOBIN=$GOPATH/bin
#export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
#export GOROOT_BOOTSTRAP=DIR14
#export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
#go env -w GOBIN=/somewhere/else/bin
#go env -u GOBIN

