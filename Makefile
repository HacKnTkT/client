default: build
all: build

AVRO=java -jar jar/avro-tools-1.7.7.jar idl
ICED=node_modules/.bin/iced

json/%.json: avdl/%.avdl
	$(AVRO) $< $@~ && mv $@~ $@

config:
	npm install -d

build-stamp: \
	json/config.json \
	json/identify.json \
	json/identify_ui.json \
	json/login.json \
	json/login_ui.json \
	json/signup.json
	@mkdir -p json
	date > $@

go/keybase_v1.go: build-stamp
	@mkdir -p go/
	$(ICED) ./bin/compile.iced -d json -t go -o $@
	gofmt -w $@

objc-build-stamp: build-stamp
	ruby ./bin/objc.rb
	date > $@

clean:
	rm -rf json/*.json go/*.go objc/*

build: build-stamp go/keybase_v1.go objc-build-stamp

.PHONY: test setup config

