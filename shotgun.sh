#!/bin/bash

go-shotgun -u http://0.0.0.0:8080 -p 8088 \
  -buildCmd="go build" -runCmd="./go-tomtelizer-server" \
  -checkCmd="exit \`find . -name '*.go' -newer ./go-tomtelizer-server|wc -l\`"
