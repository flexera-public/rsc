.PHONY: all api15 api15gen api15json

all: api15
	go build && go test

api15: api15gen
	cd rsapi15 && go build

api15gen:
	cd rsapi15 && go generate

api15json:
	curl -s -o rsapi15/api_data.json http://reference.rightscale.com/api1.5/api_data.json
