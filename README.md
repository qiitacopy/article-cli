# article-cli

## 起動コマンド
~~~sh
go run main.go
~~~

## 追加方法
gRPCサーバのrpc定義を追加、変更した場合には、次の手順で動作確認する

1. protoファイルからソース自動生成
    ~~~sh
    protoc --go_opt=paths=source_relative --go_out=plugins=grpc:./grpc --proto_path=./grpcspec ./grpcspec/article.proto
    ~~~
1. main.goに呼び出し処理を追加

1. 動作確認
    ~~~sh
    go run main.go
    ~~~