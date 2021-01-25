#!/bin/sh
protoc --go_out=plugins=grpc:. pb/serialize.proto