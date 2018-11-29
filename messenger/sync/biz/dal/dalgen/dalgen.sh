#!/bin/bash
./nebula-dal-generator --xml=$1
gofmt -w ../dao/mysql_dao/*.go
gofmt -w ../dataobject/*.go
