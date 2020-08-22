# toushi-api
自分用の投資API

## 開発環境
Visual Studio Code

## デバッグ方法
TBC

## デプロイ方法
### buy-notification
cd buy-notification
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main #これをlambdaにデプロイ

## 機能
設定した有価証券が値下がりしたら通知を送る