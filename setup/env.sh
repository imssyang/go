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
# GOTOOLDIR="/opt/go/pkg/tool/linux_amd64"
# GOROOT="/opt/go"
# GOPATH=$GOROOT/app
# GOBIN=$GOPATH/bin
# GOOS=linux
# GOARCH=amd64
# GOROOT_BOOTSTRAP=DIR14
# go env -w GOBIN=/somewhere/else/bin
# go env -u GOBIN

export GOENV_ROOT="/opt/go/goenv"
export GOENV_GOPATH_PREFIX="/opt/go/gomod"
export GOENV_GOMOD_VERSION_ENABLE=1
if [[ -f "${GOENV_ROOT:-"$HOMEBREW_PREFIX"}/bin/goenv" ]]; then
  eval "$(goenv init - zsh)"
  path=(${GOENV_ROOT}/shims $GOROOT/bin $GOPATH/bin $path)
fi
export GO111MODULE=on
export GOPROXY=https://goproxy.cn # QiNiuYun
#export GOPROXY=https://mirrors.aliyun.com/goproxy/

