$DST = "$(go env GOPATH)/src"

protoc.exe `
    --proto_path="$(Get-Location)" `
    --go_out="plugins=grpc:$DST" `
    $(Get-Item ./*.proto)
    