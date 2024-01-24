APP=vt
PW_DIR=ui/tests

all: build

build: build-ui
	go build -o ${APP}

install: build
	mv ${APP} ~/.go/bin/

test-ui: update-ui
	go run . ${PW_DIR}/sample-video.mp4 -s ${PW_DIR}/sample-subtitle.vtt -p 1201

build-ui:
	npm run hard-build --prefix ui 1>/dev/null

update-ui:
	npm run soft-build --prefix ui 1>/dev/null

.PHONY: build install test-ui build-ui update-ui
