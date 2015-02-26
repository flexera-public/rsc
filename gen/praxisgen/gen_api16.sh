#!/bin/bash -e

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
pushd $DIR >/dev/null
go install
praxisgen -metadata="../../rsapi16/api_docs" -output="../../rsapi16" -pkg="rsapi16" -target="1.6" -client="Api16"
popd >/dev/null
