#!/bin/bash

SOURCE="$0"
while [ -h "$SOURCE"  ]; do # resolve $SOURCE until the file is no longer a symlink
    DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /*  ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"

#build
go build -o "$DIR/../_out/thrift_ser" github.com/asyoume/Auth/pkg/thrift_ser
#run server
"$DIR/../_out/thrift_ser"  -P $1 -transport $2  -addr $3 -framed  $4 -buffered $5