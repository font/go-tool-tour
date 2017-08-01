#!/usr/bin/env bash

go run pic.go | sed 's/IMAGE://' | base64 -d > image.png
