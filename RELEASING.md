# Releasing `rsc`

This how to describes how to release a new version of rsc. Obviously the first thing to do is to
make sure that all tests pass. Also the [CHANGELOG](https://github.com/rightscale/rsc/blob/master/CHANGELOG.md)
must be kept up-to-date.

### Make a release branch

In order to cut a release branch from master, the steps are:
```
git checkout -b v2.3.4
make govers
git commit -a -m 'add import constraints for release'
git push origin v2.3.4
```
This will trigger a CI build which will upload the new version binaries. Also creating the remote
branch means that `gopkg.in` will get the newly released code for the major version.

### Test a release

Once the release branch has been pushed and the CI job completes:
* Download the binary, run `rsc --version` and make sure the correct version is displayed, on Linux:
```
curl https://binaries.rightscale.com/rsbin/rsc/v1/rsc-linux-amd64.tgz | tar -zxf - -O rsc/rsc > rsc
chmod +x ./rsc
./rsc --version
```
* Create a temporary go workspace and make sure go get retrieves the correct version:
```
mkdir tmp
export SAVED_GOPATH=$GOPATH
export GOPATH=`pwd`/tmp
go get gopkg.in/rightscale/rsc.v2
cd tmp/src/gopkg.in/rightscale/rsc.v2
git log -2
export GOPATH=$SAVED_GOPATH
cd ../../../../..
rm -rf tmp
```

