#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
pushd $DIR >/dev/null
go install
api15gen -metadata="../../rsapi15" -output="../../rsapi15"
popd >/dev/null
