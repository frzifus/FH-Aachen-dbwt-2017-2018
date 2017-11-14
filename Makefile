BINARY = dbwt
DATE = $(shell date +%FT%T%Z)
BUILD_DIR = build/bin
LOG_DIR= build/log
PWD = $(shell pwd)

LDFLAGS =

.PHONY: test clean arm amd64 run install uninstall

# Build the project
all: clean test amd64 copyResources

run:
	$(shell build/bin/dbwt-linux-amd64)

amd64:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-linux-amd64 -v
arm:
	GOOS=linux GOARCH=arm go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-linux-arm -v
win64:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-win64.exe -v

install:
	#$(shell cp ./build/bin/dbwt /usr/bin/)
	#$(shell cp ./systemd/dbwt.service /etc/systemd/system)

uninstall:
	$(shell rm /etc/systemd/system/dbwt.service)
test:
	@echo "Write testlog..."
	$(shell go test -v ./test... > ${LOG_DIR}/test_${DATE}.log)
viz:
	go-callvis -minlen 3 -focus ./ -group pkg,type ./ | dot -Tpng -o dbwt-overview.png
copyResources:
	@echo Copy resources
	cp -r static ${BUILD_DIR}/
	cp -r view ${BUILD_DIR}/
	cp -r config ${BUILD_DIR}/
clean:
	-rm -f ${BUILD_DIR}/${BINARY}-*


