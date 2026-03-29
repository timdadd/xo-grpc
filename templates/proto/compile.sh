#!/bin/sh
set -u
set -e

Compile () {
    rm -rf $1
    mkdir -p $1
    echo "Compiling $1.proto..."
    protoc -I. -I3rd-party --go_out $1 --go_opt paths=source_relative --go-grpc_out $1 --go-grpc_opt paths=source_relative $1.proto
    echo "Generating reverse proxy (grpc-gateway) $1.proto..."
    protoc -I. -I3rd-party --grpc-gateway_out $1 --grpc-gateway_opt logtostderr=true,paths=source_relative,generate_unbound_methods=true $1.proto
}

fn="*"
if [ "$#" -eq "1" ]; then
  fn="$1"
fi

protos=""
for i in $fn.proto; do
    protos="$protos $i"
    pkg=$(echo "$i" | cut -f 1 -d '.')
    Compile $pkg
done

echo "Generating OpenAPIv2 specs"
protoc -I. -I3rd-party --openapiv2_out . --openapiv2_opt logtostderr=true,generate_unbound_methods=true,allow_merge=true $protos

echo "Finished!"
