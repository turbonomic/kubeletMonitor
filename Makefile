OUTPUT_DIR=./_output

product: clean
	env GOOS=linux GOARCH=amd64 go build -o ${OUTPUT_DIR}/kubelet.linux ./

build: clean
	go build -o ${OUTPUT_DIR}/kubelet ./

test: clean
	@go test -v -race ./pkg/...
	
clean:
	rm -rf ${OUTPUT_DIR}/*
