#!/usr/bin/env bash

#
# USER VARIABLES
#
image_ref='lokomotes/metro-server:latest-linux-amd64'
my_pkg='github.com/lokomotes/metro/cmd/metro-server'

#
# DO NOT TOUCH
#
go_path=$(go env GOPATH)

if [ ! -f  "./linux.dockerfile" ]; then
    echo "Dockerfile is not provided"
    exit 1
fi

tmp_path=$(mktemp -d)

deps=$(go list -f '{{.Deps}}' | tr "[" " " | tr "]" " " | xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}')
deps="${deps}
${my_pkg}"

while read -r element; do
    mkdir -p $tmp_path/src/$element/
    rsync -a $go_path/src/$element/ $tmp_path/src/$element/
    echo $element" done"
done <<< "$deps"

cp ./linux.dockerfile $tmp_path/Dockerfile

docker build \
    -t $image_ref \
    $tmp_path

rm -rf $tmp_path
