# twirp-sample
twirp sample

# goおよびprotobufのインストール
```
brew install go
brew install protobuf
```

# twirp、protoc-gen-toおよびprotoc-gen-twirpのインストール
```
go get github.com/twitchtv/twirp
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/twitchtv/twirp/protoc-gen-twirp
```

# ソースコードの取得
```
mkdir -p $GOPATH/src/github.com/yoshikazu-ooba
cd $GOPATH/src/github.com/yoshikazu-ooba
git clone git@github.com:yoshikazu-ooba/twirp-sample.git
```

# サーバの実行
```
cd $GOPATH/src/github.com/yoshikazu-ooba/twirp-sample
go run cmd/server/main.go
```

# クライアントの実行
```
cd $GOPATH/src/github.com/yoshikazu-ooba/twirp-sample
go run cmd/client/main.go
```
