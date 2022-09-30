BINARY_NAME=upreq
BUILD_DIR=build

build:
	GOARCH=amd64 GOOS=linux go build -o ${BUILD_DIR}/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=darwin go build -o ${BUILD_DIR}/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=windows go build -o ${BUILD_DIR}/${BINARY_NAME}-windows.exe main.go

clean:
	rm -rf ${BUILD_DIR}