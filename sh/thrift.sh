#!/bin/bash
thrift -strict -out ../thrift-go  --gen go ../idl/authSer.thrift
thrift -strict -out ../thrift-js  --gen js ../idl/authSer.thrift