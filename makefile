APP=vt

all: build

ui:
	npm run build --prefix ui 1>/dev/null

build: ui
	go build -o ${APP}

install: build
	mv ${APP} ~/.go/bin/

.PHONY: ui build install
