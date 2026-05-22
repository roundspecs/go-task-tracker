build:
	rm -rf out
	mkdir out
	go fmt .
	go fmt ./...
	go build -o out/task .
	chmod +x out/task