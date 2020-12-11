#!/bin/bash

APP=go
GROUP=$APP
USER=$APP
SYSD=/etc/systemd/system
SERFILE=godoc.service

initialize() {
  if [[ ! -s $SYSD/$SERFILE ]]; then
    ln -s /opt/$APP/setup/$SERFILE $SYSD/$SERFILE
    echo "($APP) create symlink: $SYSD/$SERFILE --> /opt/$APP/setup/$SERFILE"
  fi
}

deinitialize() {
  if [[ -s $SYSD/$SERFILE ]]; then
    rm -rf $SYSD/$SERFILE
    echo "($APP) delete symlink: $SYSD/$SERFILE"
  fi
}

godoc_start() {
  pgrep -x godoc >/dev/null
  if [[ $? != 0 ]]; then
	/opt/$APP/app/bin/godoc -http :8009 &
    echo "($APP) godoc start!"
  fi
  godoc_show
}

godoc_stop() {
  pgrep -x godoc >/dev/null
  if [[ $? == 0 ]]; then
    kill -9 `pidof godoc`
	echo "($APP) godoc stop! (auto restart when systemd)"
  fi
  godoc_show
}

godoc_show() {
  ps -ef | grep godoc | grep -v 'grep'
}

case "$1" in
  init)
    initialize
    ;;
  deinit)
    deinitialize
    ;;
  start)
    godoc_start
    ;;
  stop)
    godoc_stop
    ;;
  show)
	godoc_show
	;;
  *)
    SCRIPTNAME="${0##*/}"
    echo "Usage: $SCRIPTNAME {init|deinit|start|stop|show}"
    exit 3
    ;;
esac

exit 0

# vim: syntax=sh ts=4 sw=4 sts=4 sr noet
