cp ../../*.proto .

protoc -I=. *.proto \
  --js_out=import_style=commonjs,binary:. \
  --grpc-web_out=import_style=typescript,mode=grpcweb:.
