# Releasing `rsc`

This how to describes how to release a new version of rsc. Obviously the first thing to do is to
make sure that all tests pass. Also the [CHANGELOG](https://github.com/rightscale/rsc/blob/master/CHANGELOG.md)
and [README](https://github.com/rightscale/rsc/blob/master/README.md) must be kept up-to-date.

### Make a release branch

If making a major release version, change `GOPKG_VERS` in Makefile.

Make sure you switch the go version to the same version travis uses, e.g 1.10. gofmt sometimes
hash minor changes which can break the build.

In order to cut a release branch from master, the steps are:
```
git checkout -b v2.3.4
rm -rf vendor/
make govers
git commit -a -m 'add import constraints for release v2.3.4'
git push origin v2.3.4
```
This will trigger a CI build which will upload the new version binaries. Also creating the remote
branch means that `gopkg.in` will get the newly released code for the major version. We remove the
`vendor` directory in order to avoid a bug in `govers` that tries to look at files in the directory
causing it to miss the top level directory `.go` files. You can run `make depend` to restore `vendor`
after releasing.

### Test a release

Once the release branch has been pushed and the CI job completes:
* Download the binary, run `rsc --version` and make sure the correct version is displayed, on Linux:
```
curl https://binaries.rightscale.com/rsbin/rsc/v8.1/rsc-linux-amd64.tgz | tar -zxf - -O rsc/rsc > rsc
chmod +x ./rsc
./rsc --version
```
* Create a temporary go workspace and make sure go get retrieves the correct version:
```
mkdir tmp
export SAVED_GOPATH=$GOPATH
export GOPATH=`pwd`/tmp
go get gopkg.in/rightscale/rsc.v8
cd tmp/pkg/mod/gopkg.in/rightscale/rsc.v8
git log -2
export GOPATH=$SAVED_GOPATH
cd ../../../../../..
rm -rf tmp
```

