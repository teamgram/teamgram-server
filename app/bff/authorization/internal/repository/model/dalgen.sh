#!/bin/bash

tgctl model dalgen datasource -url "root:@tcp(127.0.0.1:3306)/teamgram?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -out ./ -xml ./tables/
goimports -w *.go
