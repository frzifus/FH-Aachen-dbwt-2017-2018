BINARY = dbwt
DATE = $(shell date +%FT%T%Z)
BUILD_DIR = build/bin
LOG_DIR= build/log

LDFLAGS =

.PHONY: test clean arm amd64 run install uninstall docker-cp

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

docker-build:
	docker-compose up -d --build
	docker ps

docker-cp:
	docker-compose start
	docker cp ./ dbwt_app_1:/go/src/github.com/frzifus/dbwt
	docker exec -it dbwt_app_1 /bin/bash -c "cd /go/src/github.com/frzifus/dbwt; go install"
	docker-compose restart
	docker ps
install:
	#$(shell cp ./build/bin/dbwt /usr/bin/)
	#$(shell cp ./systemd/dbwt.service /etc/systemd/system)

uninstall:
	$(shell rm /etc/systemd/system/dbwt.service)
test:
	echo "Write testlog..."
	@mkdir -p build
	@mkdir -p build/log
	@$(shell go test -v ./test... > ${LOG_DIR}/test_${DATE}.log)
viz:
	go-callvis -minlen 3 -focus ./ -group pkg,type ./ | dot -Tpng -o dbwt-overview.png

copyResources:
	@echo Copy resources
	cp -r static ${BUILD_DIR}/
	cp -r view ${BUILD_DIR}/
	cp -r config ${BUILD_DIR}/
clean:
	-rm -f ${BUILD_DIR}/${BINARY}-*


