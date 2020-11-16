#!/bin/bash
export LANG=zh_CN.UTF-8

ENVARG=GO111MODULE=on
LINUXARG=CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
BUILDARG=-ldflags " -s -X main.buildTime=`date '+%Y-%m-%dT%H:%M:%S'` -X main.gitHash=`git symbolic-ref --short -q HEAD`:`git rev-parse HEAD`"

dep:
	cd src; ${ENVARG} go get ./...; cd -

update_dep:
	cd src; ${ENVARG} go get -u ./...; go mod tidy; cd -

finance:
	cd src/finance; ${ENVARG} go build ${BUILDARG} -o ../../bin/fin main.go;

currency:
	cd src/currency; ${ENVARG} go build ${BUILDARG} -o ../../bin/currency main.go;

youdao:
	cd src/youdao; ${ENVARG} go build ${BUILDARG} -o ../../bin/youdao main.go;

clean:
	rm -fr bin/*
	chmod -R 766 pkg/*
	\rm -r pkg/*
	rm src/go.sum
