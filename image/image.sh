#!/usr/bin/env bash

go run image.go | sed 's/IMAGE://' | base64 -d > image.png
