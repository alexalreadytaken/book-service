echo "Generate grpc services"

protoc \
    --go_out=plugins=grpc:. \
    --go_opt=paths=source_relative \
    ./proto/*.proto