#!/bin/bash
#Install oapi generator
#Un-comment the below line to install the generator
#go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
#Generate api code
oapi-codegen -package api -generate "types,server,spec" -o api.gen.go cerberus-api.yaml
#Generate test api code
#oapi-codegen -package api -generate "client" -o api.gen_test.go cerberus-api.yaml
#Add file to git stage
git add api.gen.go
#git add api.gen_test.go
