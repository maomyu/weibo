Build:
	go build  ./cmd/main/main.go
	chmod +x main
	docker build -t weibo .
	docker run --name weibov1 -p 9090:9090 -d weibo
	