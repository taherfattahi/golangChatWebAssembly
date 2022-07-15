SET GOOS=js
SET GOARCH=wasm

go build -o ./html/main.wasm .

SET GOOS=windows
SET GOARCH=amd64
