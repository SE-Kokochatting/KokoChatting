PROJECT_NAME="koko"


build() {
  go mod tidy
  go build -o koko ./main.go
}

stop() {
  pid=$(ps -ef | grep $PROJECT_NAME | grep -v grep | awk '{print $2}')
  if [ "$pid" ]; then
    kill $pid
    ret=0
    for ((i = 0; i < 15; i++)); do
      sleep 1
      pid=$(ps -ef | grep $PROJECT_NAME | grep -v grep | awk '{print $2}')
      if [ "$pid" ]; then
        ret=0
      else
        ret=1
        break
      fi
    done
    if [ "$ret" ]; then
      echo "ok"
    else
      echo "fail"
    fi
  fi
}

start() {
  nohup ./koko > ./server.log 2>&1 &
}

restart() {
  stop
  sleep 1
  start
}


case "$1" in
build)
  build
  ;;
start)
  start
  ;;
stop)
  stop
  ;;
restart)
  restart
  ;;
*)
  @echo $"Usage: $0 [start|stop|restart|build]"
  exit 2
  ;;
esac