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
# HACKS - a couple of things here are unconventional in order to keep travis-ci fast:
# - use 'godep save' on your laptop if you add dependencies, but we don't use godep in the
#   makefile, instead, we simply add the godep workspace to the GOPATH

#NAME=$(shell basename $$PWD)
NAME=rsc
BUCKET=rightscale-binaries
ACL=public-read
# Dependencies not handled by Godep, i.e. that are used to build/test/upload this puppy
DEPEND=golang.org/x/tools/cmd/cover github.com/onsi/ginkgo/ginkgo \
			 github.com/rlmcpherson/s3gof3r/gof3r github.com/tools/godep \
			 github.com/rogpeppe/govers

#=== below this line ideally remains unchanged, add new targets at the end  ===

TRAVIS_BRANCH?=dev
DATE=$(shell date '+%F %T')
SECONDS=$(shell date '+%s')
TRAVIS_COMMIT?=$(shell git symbolic-ref HEAD | cut -d"/" -f 3)
GIT_BRANCH:=$(shell git symbolic-ref --short -q HEAD || echo "v1-unstable")
# by manually adding the godep workspace to the path we don't need to run godep itself
ifeq ($(OS),Windows_NT)
	SHELL:=/bin/dash
	GOPATH:=$(shell cygpath --windows $(PWD))/Godeps/_workspace;$(GOPATH)
else
	GOPATH:=$(PWD)/Godeps/_workspace:$(GOPATH)
endif
# because of the Godep path we build ginkgo into the godep workspace
PATH:=$(PWD)/Godeps/_workspace/bin:$(PATH)

# the default target builds a binary in the top-level dir for whatever the local OS is
default: $(NAME)
$(NAME): *.go version govers generate
	go build -o $(NAME) .

install: $(NAME)
	go install

# the standard build produces a "local" executable, a linux tgz, and a darwin (macos) tgz
build: $(NAME) generate build/$(NAME)-linux-amd64.tgz build/$(NAME)-darwin-amd64.tgz build/$(NAME)-linux-arm.tgz build/$(NAME)-windows-amd64.zip

# create a tgz with the binary and any artifacts that are necessary
# note the hack to allow for various GOOS & GOARCH combos, sigh
build/$(NAME)-%.tgz: *.go version depend govers
	rm -rf build/$(NAME)
	mkdir -p build/$(NAME)
	tgt=$*; GOOS=$${tgt%-*} GOARCH=$${tgt#*-} go build -o build/$(NAME)/$(NAME) .
	chmod +x build/$(NAME)/$(NAME)
	for d in script init; do if [ -d $$d ]; then cp -r $$d build/$(NAME); fi; done
	if [ "build/*/*.sh" != 'build/*/*.sh' ]; then \
	  sed -i -e "s/BRANCH/$(TRAVIS_BRANCH)/" build/*/*.sh; \
	  chmod +x build/*/*.sh; \
	fi
	tar -zcf $@ -C build $(NAME)
	rm -r build/$(NAME)

# create a zip with the binary and any artifacts that are necessary
# note the hack to allow for various GOOS & GOARCH combos, sigh
build/$(NAME)-%.zip: *.go version depend govers
	rm -rf build/$(NAME)
	mkdir -p build/$(NAME)
	tgt=$*; GOOS=$${tgt%-*} GOARCH=$${tgt#*-} go build -o build/$(NAME)/$(NAME).exe .
	cd build; zip -r $(notdir $@) $(NAME)
	rm -r build/$(NAME)

# upload assumes you have AWS_ACCESS_KEY_ID and AWS_SECRET_KEY env variables set,
# which happens in the .travis.yml for CI
upload: depend
	@which gof3r >/dev/null || (echo 'Please "go get github.com/rlmcpherson/s3gof3r/gof3r"'; false)
	(cd build; set -ex; \
	  for f in *.tgz *.zip; do \
	    gof3r put --no-md5 --acl=$(ACL) -b ${BUCKET} -k rsbin/$(NAME)/$(TRAVIS_COMMIT)/$$f <$$f; \
	    if [ "$(TRAVIS_PULL_REQUEST)" = "false" ]; then \
	      gof3r put --no-md5 --acl=$(ACL) -b ${BUCKET} -k rsbin/$(NAME)/$(TRAVIS_BRANCH)/$$f <$$f; \
	    fi; \
	  done)

# produce a version string that is embedded into the binary that captures the branch, the date
# and the commit we're building
version:
	@echo "// +build make\n\npackage main\n\nconst VV = \"$(NAME) $(TRAVIS_BRANCH) - $(DATE) - $(TRAVIS_COMMIT)\"" \
	  >version.go
	@echo "// +build make\n\npackage rsapi\n\nconst UA = \"$(NAME)/$(TRAVIS_BRANCH)-$(SECONDS)-$(TRAVIS_COMMIT)\"" \
	  >rsapi/user_agent.go
	@echo "version.go: `tail -1 version.go`"

# descend into go hell and fix import statements
# runs govers to change imports of rsc packages to current branch
# runs sed to add import comments to all package statements
# runs sed to change import lines in godegen writers
govers: depend
	govers -d gopkg.in/rightscale/rsc.$(GIT_BRANCH)
	@echo "adding import comments"
	@for f in `find [a-z]* -mindepth 2 -name \*.go \! -name \*_test.go`; do \
		dir=`dirname $${f#./}` ;\
		sed -i -r \
		  -e '1,10 s;^(package\s+\S+).*;\1 // import "gopkg.in/rightscale/$(NAME).$(GIT_BRANCH)/'"$${dir}"'";' \
			$$f;\
	done
	@echo "fixing code gen templates"
	@for f in gen/writers/*.go; do \
	  sed -i -re 's;g[a-z.]+/rightscale/rsc[-.a-z0-9]*;gopkg.in/rightscale/rsc.$(GIT_BRANCH);' $$f ;\
	done

# Installing build dependencies is a bit of a mess. Don't want to spend lots of time in
# Travis doing this. The following just relies on go get no reinstalling when it's already
# there, like your laptop.
depend:
	go get $(DEPEND)
	godep restore

clean:
	rm -rf build
	git checkout -f version.go rsapi/user_agent.go

# gofmt uses the awkward *.go */*.go because gofmt -l . descends into the Godeps workspace
# and then pointlessly complains about bad formatting in imported packages, sigh
lint:
	@if gofmt -l *.go */*.go 2>&1 | grep .go; then \
	  echo "^- Repo contains improperly formatted go files; run gofmt -w *.go */*.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	go tool vet -composites=false *.go
	go tool vet -composites=false **/*.go

travis-test: lint generate
	ginkgo -r -cover

# running ginkgo twice, sadly, the problem is that -cover modifies the source code with the effect
# that if there are errors the output of gingko refers to incorrect line numbers
# tip: if you don't like colors use gingkgo -r -noColor
test: lint generate
	@test "$PWD" != `/bin/pwd` && echo "*** Please cd `/bin/pwd` if compilation fails"
	ginkgo -r
	ginkgo -r -cover
	go tool cover -func=`basename $$PWD`.coverprofile

#===== SPECIAL TARGETS FOR RSC =====

.PHONY: rsc test generate api15gen praxisgen api15json 

generate: api15gen praxisgen
	go generate

api15gen:
	cd gen/api15gen && go install

praxisgen:
	cd gen/praxisgen && go test && go install
	which praxisgen

api15json:
	curl -s -o rsapi15/api_data.json http://reference.rightscale.com/api1.5/api_data.json
