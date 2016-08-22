#! /usr/bin/make
#
# Makefile for Golang projects, v0.9.0
#
# Features:
# - runs ginkgo tests recursively, computes code coverage report
# - code coverage ready for travis-ci to upload and produce badges for README.md
# - build for linux/amd64, linux/arm, darwin/amd64, windows/amd64
# - just 'make' builds for local OS/arch
# - produces .tgz/.zip build output
# - bundles *.sh files in ./script subdirectory
# - produces version.go for each build with string in global variable VV, please
#   print this using a --version option in the executable
# - to include the build status and code coverage badge in CI use (replace NAME by what
#   you set $(NAME) to further down, and also replace magnum.travis-ci.com by travis-ci.org for
#   publicly accessible repos [sigh]):
#   [![Build Status](https://magnum.travis-ci.com/rightscale/NAME.svg?token=4Q13wQTY4zqXgU7Edw3B&branch=master)](https://magnum.travis-ci.com/rightscale/NAME
#   ![Code Coverage](https://s3.amazonaws.com/rs-code-coverage/NAME/cc_badge_master.svg)
#
# Top-level targets:
# default: compile the program, you can thus use make && ./NAME -options ...
# build: builds binaries for linux and darwin
# test: runs unit tests recursively and produces code coverage stats and shows them
# travis-test: just runs unit tests recursively
# clean: removes build stuff
#
# the upload target is used in the .travis.yml file and pushes binary archives to
# https://$(BUCKET).s3.amazonaws.com/rsbin/$(NAME)/$(BRANCH)/$(NAME)-$(GOOS)-$(GOARCH).tgz
# (.zip for windows)
#

#NAME=$(shell basename $$PWD)
NAME=rsc
BUCKET=rightscale-binaries
ACL=public-read
# version for gopkg.in, e.g. v1, v2, ...
GOPKG_VERS=v6
GLIDE_VERSION?=v0.11.1
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
ifeq (windows,$(GOOS))
GLIDE_EXEC=glide-$(GLIDE_VERSION).exe
else
GLIDE_EXEC=glide-$(GLIDE_VERSION)
endif
# Dependencies not handled by glide, i.e. that are used to build/test/upload this puppy
DEPEND=golang.org/x/tools/cmd/cover \
			github.com/onsi/ginkgo \
			github.com/onsi/ginkgo/ginkgo \
			github.com/golang/protobuf/proto \
			github.com/rlmcpherson/s3gof3r/gof3r \
			github.com/rogpeppe/govers

#=== below this line ideally remains unchanged, add new targets at the end  ===

TRAVIS_BRANCH?=dev
DATE=$(shell date '+%F %T')
SECONDS=$(shell date '+%s')
TRAVIS_COMMIT?=$(shell git symbolic-ref HEAD | cut -d"/" -f 3)
GIT_BRANCH:=$(shell git symbolic-ref --short -q HEAD || echo "master")
SHELL:=$(shell which bash)
# on Mac OS X and other platforms, sed might not be GNU sed and the -i flag needs an extension argument,
# but that argument can be an empty string to indicate no backup
ifeq ($(shell sed --version 2>/dev/null | grep 'GNU sed'),)
SED_I=''
else
SED_I=
endif

# the default target builds a binary in the top-level dir for whatever the local OS is
default: $(NAME)
$(NAME): *.go version generate
	go build -o $(NAME) .

install: $(NAME)
	go install

# the standard build produces a "local" executable, a linux tgz, and a darwin (macos) tgz
build: depend generate $(NAME) build/$(NAME)-linux-amd64.tgz build/$(NAME)-darwin-amd64.tgz build/$(NAME)-linux-arm.tgz build/$(NAME)-windows-amd64.zip

# create a tgz with the binary and any artifacts that are necessary
# note the hack to allow for various GOOS & GOARCH combos, sigh
build/$(NAME)-%.tgz: *.go version depend
	rm -rf build/$(NAME)
	mkdir -p build/$(NAME)
	tgt=$*; GOOS=$${tgt%-*} GOARCH=$${tgt#*-} go build -tags make -o build/$(NAME)/$(NAME) .
	chmod +x build/$(NAME)/$(NAME)
	tar -zcf $@ -C build $(NAME)
	rm -r build/$(NAME)

# create a zip with the binary and any artifacts that are necessary
# note the hack to allow for various GOOS & GOARCH combos, sigh
build/$(NAME)-%.zip: *.go version depend
	rm -rf build/$(NAME)
	mkdir -p build/$(NAME)
	tgt=$*; GOOS=$${tgt%-*} GOARCH=$${tgt#*-} go build -tags make -o build/$(NAME)/$(NAME).exe .
	cd build; zip -r $(notdir $@) $(NAME)
	rm -r build/$(NAME)

# upload assumes you have AWS_ACCESS_KEY_ID and AWS_SECRET_KEY env variables set,
# which happens in the .travis.yml for CI
upload:
	@which gof3r >/dev/null || (echo 'Please "go get github.com/rlmcpherson/s3gof3r/gof3r"'; false)
	(cd build; set -ex; \
	  for f in *.tgz *.zip; do \
	    gof3r put --no-md5 --acl=$(ACL) -b ${BUCKET} -k rsbin/$(NAME)/$(TRAVIS_COMMIT)/$$f <$$f; \
	    if [ "$(TRAVIS_PULL_REQUEST)" = "false" ]; then \
	      gof3r put --no-md5 --acl=$(ACL) -b ${BUCKET} -k rsbin/$(NAME)/$(TRAVIS_BRANCH)/$$f <$$f; \
	      re='^(v[0-9]+)\.[0-9]+\.[0-9]+$$' ;\
	      if [[ "$(TRAVIS_BRANCH)" =~ $$re ]]; then \
	        gof3r put --no-md5 --acl=$(ACL) -b ${BUCKET} -k rsbin/$(NAME)/$${BASH_REMATCH[1]}/$$f <$$f; \
	      fi; \
	    fi; \
	  done)

# produce a version string that is embedded into the binary that captures the branch, the date
# and the commit we're building
version:
	@echo -e "// +build make\n\npackage main\n\nconst VV = \"$(NAME) $(TRAVIS_BRANCH) - $(DATE) - $(TRAVIS_COMMIT)\"" \
	  >version.go
	@echo -e "// +build make\n\npackage httpclient\n\nconst UA = \"$(NAME)/$(TRAVIS_BRANCH)-$(SECONDS)-$(TRAVIS_COMMIT)\"" \
	  >httpclient/user_agent.go
	@echo "version.go: `tail -1 version.go`"

# descend into go hell and change/add import statements to suit gopkg.in versioning
# it forces import via gopkg.in/rightscale/$(NAME).$(GOPKG_VERS)
# - runs govers to change imports of rsc packages to rsc.v1
# - runs sed to add import comments to all package statements to force gopkg.in
# - runs sed to change import lines in codegen writers
# - runs gofmt because the order of imports changes in some files (sigh)
govers:
	govers -d -m 'g[^/]*/rightscale/rsc[^/]*' gopkg.in/rightscale/rsc.$(GOPKG_VERS)
	@echo "adding package import comments"
	@for f in `find . -path './[a-z]*' -path ./\*/\*.go \! -name \*_test.go`; do \
		dir=`dirname $${f#./}` ;\
		sed -E -i $(SED_I) \
		  -e '1,10 s;^(package +[_a-z0-9]+).*;\1 // import "gopkg.in/rightscale/$(NAME).$(GOPKG_VERS)/'"$${dir}"'";' \
			$$f;\
	done
	@echo "fixing code gen templates"
	@for f in gen/writers/*.go; do \
		dir=`dirname $${f#./}` ;\
	  sed -E -i $(SED_I) \
		  -e 's;g[a-z.]+/rightscale/rsc[-.a-z0-9]*;gopkg.in/rightscale/$(NAME).$(GOPKG_VERS);' $$f ;\
	done
	gofmt -w *.go */*.go

# revert govers, i.e. remove import constraints and use github.com/rightscale/$(NAME)
unvers:
	@echo "changing import statements"
	@for f in `find . -path './[a-z]*' -name \*.go`; do \
		sed -E -i $(SED_I) -e 's;g[a-z.]+/rightscale/$(NAME)[-.a-z0-9]*;github.com/rightscale/$(NAME);' $$f ;\
	done
	@echo "removing package import comments"
	@for f in `find . -path './[a-z]*' -path ./\*/\*.go \! -name \*_test.go`; do \
		sed -E -i $(SED_I) -e '1,10 s;^(package +[a-z][^ /]*).*;\1;' $$f; \
	done
	gofmt -w *.go */*.go

# for release branches, i.e. having names like v1, v1.2, v1.2.3, check that gopkg.in import
# constraints are in place, i.e. that 'make govers' has been run
ifneq ($(findstring $(GOPKG_VERS),$(GIT_BRANCH)),)
check-govers:
	@if ! govers -d -n gopkg.in/rightscale/$(NAME).$(GOPKG_VERS); then \
		echo "   check failed, run 'make govers'"; exit 1; fi
	@echo "checking package statements"
	@files=`find . -path './[a-z]*' -path ./\*/\*.go \! -name \*_test.go \! -name user_agent.go`; \
	if egrep '^package\s+[a-z]' $$files | \
	   egrep -v codegen_ | \
		 egrep -v "import \"gopkg.in/rightscale/$(NAME).$(GOPKG_VERS)"; then \
		echo "   check failed, run 'make govers'"; exit 1; fi
	@echo "checking code gen templates"
	@if egrep 'rightscale/rsc' gen/writers/*.go | egrep -v "gopkg.in/rightscale/rsc.$(GOPKG_VERS)"; then \
		echo "   check failed, run 'make govers'"; exit 1; fi
else
check-govers:
	@echo "not a release branch: not checking import constraints"
endif

bin/$(GLIDE_EXEC):
	mkdir -p tmp bin
	curl -sSfL --tlsv1 --connect-timeout 30 --max-time 180 --retry 3 \
		-o tmp/glide.tar.gz https://github.com/Masterminds/glide/releases/download/$(GLIDE_VERSION)/glide-$(GLIDE_VERSION)-$(GOOS)-$(GOARCH).tar.gz
	cd tmp && tar zxvf glide.tar.gz
	mv tmp/$(GOOS)-$(GOARCH)/glide* bin/$(GLIDE_EXEC)
	rm -rf tmp

# Installing build dependencies is a bit of a mess. Don't want to spend lots of time in
# Travis doing this. The following just relies on go get no reinstalling when it's already
# there, like your laptop.
depend: bin/$(GLIDE_EXEC)
	go get $(DEPEND)
	./bin/$(GLIDE_EXEC) install --cache


clean:
	rm -rf build
	rm -f version.go httpclient/user_agent.go

lint: check-govers
	@if gofmt -l *.go */*.go 2>&1 | grep .go; then \
	  echo "^- Repo contains improperly formatted go files; run gofmt -w *.go */*.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	go tool vet -composites=false *.go
	go tool vet -composites=false **/*.go

travis-test: lint generate
	ginkgo -r -cover -skipPackage vendor

# running ginkgo twice, sadly, the problem is that -cover modifies the source code with the effect
# that if there are errors the output of gingko refers to incorrect line numbers
# tip: if you don't like colors use gingkgo -r -noColor
test: lint generate
	@test "$PWD" != `/bin/pwd` && echo "*** Please cd `/bin/pwd` if compilation fails"
	ginkgo -r -skipPackage vendor
	ginkgo -r -cover -skipPackage vendor
	go tool cover -func=`basename $$PWD`.coverprofile

#===== SPECIAL TARGETS FOR RSC =====

.PHONY: rsc test generate api15gen praxisgen api15json

generate: api15gen praxisgen
	go generate

api15gen:
	cd gen/api15gen && go install

praxisgen:
	cd gen/praxisgen && go install
	@if ! which praxisgen >/dev/null; then \
	  echo '*** Praxisgen got installed in a location that is not in your PATH ***'; \
	  echo GOPATH=$$GOPATH ;\
	  echo PATH=$$PATH ;\
	fi

api15json:
	mkdir -p rsapi15
	curl -s -o rsapi15/api_data.json http://reference.rightscale.com/api1.5/api_data.json
