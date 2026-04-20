#!/bin/bash

set -euo pipefail

shopt -s nullglob
xmlFiles=(./tables/*.xml)

if [ ${#xmlFiles[@]} -eq 0 ]; then
	echo "no xml files found in ./tables"
	exit 0
fi

for xmlFile in "${xmlFiles[@]}"; do
	echo "tgctl dalgen with ${xmlFile}"
	tgctl model dalgen datasource -url "root:@tcp(127.0.0.1:3306)/teamgram?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai" -out ./ -xml "${xmlFile}"
done

goimports -w *.go