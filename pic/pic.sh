#!/usr/bin/env bash

go build pic.go
./pic | sed 's/IMAGE://' | base64 -d > image.png
