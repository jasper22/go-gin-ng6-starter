#!/usr/bin/env bash

protoc --go_out=../backend/models models.proto
protoc-go-inject-tag -input=../backend/models/models.pb.go
../frontend/node_modules/protobufjs/bin/pbjs -t static-module -w commonjs -o ../frontend/src/app/shared/types/proto-types.js models.proto
../frontend/node_modules/protobufjs/bin/pbts -o ../frontend/src/app/shared/types/proto-types.d.ts ../frontend/src/app/shared/types/proto-types.js
