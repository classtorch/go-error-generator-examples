#!/bin/bash

rm -r golang/*

protoc --go-error-generator_out=:. \
 --go-error-generator_opt descriptor_file=errors/errors.proto \
 --go-error-generator_opt merge_error=false \
 --go-error-generator_opt merge_error_path=golang/errors \
 --go_out=. -I . protobuf/**/errors.proto
