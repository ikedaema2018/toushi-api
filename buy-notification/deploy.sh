#!/bin/bash -eu
GOOS=linux GOARCH=amd64 go build -o main main.go #ビルド
zip main.zip main #zipに固める

#zipをlambdaにアップロード
aws lambda update-function-code \
--region ap-northeast-1 \
--function-name buy-notification \
--zip-file fileb://./main.zip --publish