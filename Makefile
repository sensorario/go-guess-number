build:
	env GO111MODULE=on go build -o /tmp/ggn
	cp -f /tmp/ggn /usr/local/bin/ggn

upgrade:
	go get -u ./...
