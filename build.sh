#!/bin/bash

hugo

mkdir -p functions
# go build -o functions/goget ./goget

go run utils/goget_redirects.go
