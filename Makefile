

deploy:
	make clean
	make debug
	scp bin/wsniffer  openwrt:/root/wsniffer
	scp config.sh openwrt:/root/wsniffer

debug:
	GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -a -o bin/wsniffer *.go

local:
	GOOS=linux GOARCH=amd64 go build -o bin/wsniffer *.go

wsniffer:
	GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "-s -w" -a -o bin/wsniffer *.go

clean:
	ssh openwrt "killall wsniffer ; rm /root/wsniffer/wsniffer /root/wsniffer/config.sh" &
	rm -rf bin/*

.PHONY: debug client clean deploy local