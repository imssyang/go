#!/bin/bash
# https://golang.org/doc/install/source#introduction
# https://go-zh.org/doc/
#
# ./go env
# GO111MODULE=""
# GOARCH="amd64"
# GOBIN=""
# GOCACHE="/root/.cache/go-build"
# GOENV="/root/.config/go/env"
# GOEXE=""
# GOFLAGS=""
# GOHOSTARCH="amd64"
# GOHOSTOS="linux"
# GOINSECURE=""
# GOMODCACHE="/root/go/pkg/mod"
# GONOPROXY=""
# GONOSUMDB=""
# GOOS="linux"
# GOPATH="/root/go"
# GOPRIVATE=""
# GOPROXY="https://proxy.golang.org,direct"
# GOROOT="/opt/go"
# GOSUMDB="sum.golang.org"
# GOTMPDIR=""
# GOTOOLDIR="/opt/go/pkg/tool/linux_amd64"
# GCCGO="gccgo"
# AR="ar"
# CC="gcc"
# CXX="g++"
# CGO_ENABLED="1"
# GOMOD=""
# CGO_CFLAGS="-g -O2"
# CGO_CPPFLAGS=""
# CGO_CXXFLAGS="-g -O2"
# CGO_FFLAGS="-g -O2"
# CGO_LDFLAGS="-g -O2"
# PKG_CONFIG="pkg-config"
# GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build669611671=/tmp/go-build -gno-record-gcc-switches"

#export GOROOT_BOOTSTRAP=DIR14
#export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
#go env -w GOBIN=/somewhere/else/bin
#go env -u GOBIN

export GO_DISABLE_ENV=yes
export GOROOT=/opt/go
export GOPATH=$GOROOT/app
export GOBIN=$GOPATH/bin
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

#export GO111MODULE=on
#export GOPROXY=https://mirrors.aliyun.com/goproxy/

