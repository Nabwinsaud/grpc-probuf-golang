#!/bin/bash

protoc \
  --plugin=protoc-gen-ts_proto=$(which protoc-gen-ts_proto) \
  --ts_proto_out=. \
  --ts_proto_opt=outputServices=false \
  *.proto