.PHONY: all api15 api15gen api15json api16 praxisgen ss

all: api15 api16 ss
	go test && go install

api15: api15gen
	cd rsapi15 && go generate && go test

api15gen:
	cd gen/api15gen && go test && go install

api15json:
	curl -s -o rsapi15/api_data.json http://reference.rightscale.com/api1.5/api_data.json

api16: praxisgen
	cd rsapi16 && go generate && go test

ss: praxisgen
	cd ss && go generate && go test

praxisgen:
	cd gen/praxisgen && go test && go install
