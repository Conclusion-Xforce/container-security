from ubuntu:20.04
arg go_version=1.20

run export DEBIAN_FRONTEND=noninteractive && apt-get update && apt-get install -y \
    redis-server \
    curl \
    golang-${go_version}

env PATH=$PATH:/usr/lib/go-${go_version}/bin

run mkdir /app

add src/* /app/

run chmod +x /app/start.sh

run cd /app/ && go mod tidy && go build -o /app/server

expose 80

cmd ["/app/start.sh"]