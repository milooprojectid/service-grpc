protoc ./morbius/morbius.proto --go_out=plugins=grpc:.

python3 -m grpc_tools.protoc -I $PWD/morbius --python_out=./morbius --grpc_python_out=./morbius $PWD/morbius/morbius.proto