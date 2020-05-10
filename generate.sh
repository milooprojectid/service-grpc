protoc ./storm/storm.proto --go_out=plugins=grpc:.

python3 -m grpc_tools.protoc -I $PWD/storm --python_out=./storm --grpc_python_out=./storm $PWD/storm/storm.proto