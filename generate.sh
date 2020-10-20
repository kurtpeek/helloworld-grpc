#!/bin/bash

buf check lint

for protofile in `buf ls-files`; do
    protoc $protofile --go_out=.
done
