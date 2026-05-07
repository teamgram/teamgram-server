#!/bin/bash
set -euo pipefail

DALGEN_DSN="${DALGEN_DSN:-root:@tcp(127.0.0.1:3306)/teamgooo?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai}"

tgctl model dalgen datasource -url "${DALGEN_DSN}" -out ./ -xml ./tables/
goimports -w *.go
