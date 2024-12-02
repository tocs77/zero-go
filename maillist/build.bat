docker build -f Dockerfile.go.build -t go-protobuf .

setlocal
FOR /F "tokens=*" %%i in ('type .env') do SET %%i
docker build -f Dockerfile.go.build --build-arg PROTOC_VERSION=%PROTOC_VERSION%  ^
--build-arg BUF_VERSION=%BUF_VERSION% ^
--build-arg GO_VERSION=%GO_VERSION% ^
--build-arg PROTOC_GEN_GO_VERSION=%PROTOC_GEN_GO_VERSION% ^
--build-arg PROTOC_GEN_GO_GRPC_VERSION=%PROTOC_GEN_GO_GRPC_VERSION% ^
 -t %IMAGE_NAME% .
endlocal