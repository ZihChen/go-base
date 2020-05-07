# goformat go 語言基底

## 安裝

### Docker

安裝教學　[連結](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-18-04)

### Golang

1. 點此 [連結](https://golang.org/dl/) 取得 golang 版本

2. 解壓縮檔案
   
```
$ sudo tar -C /usr/local -xzf ~/Downloads/go1.14.linux-amd64.tar.gz
```

3. 確認 __PATH__ 有加上 `/usr/local/go/bin`
   
```
$ echo $PATH | grep "/usr/local/go/bin"
```

### Swag
安裝 swagger 文件　[連結](https://github.com/swaggo/swag)

```
$ go get -u github.com/swaggo/swag/cmd/swag
```

### GRPC
Proto 安裝教學 點此　[連結](https://yami.io/protobuf/)

```
$ protoc --proto_path=protobuf/ protobuf/*.proto --go_out=plugins=grpc:library/rpc/rpcpitaya/.
```

---
## Git Submodule
下載共用函式庫 

- errorcode 
- rpc

```
$ git submodule update --init --recursive
```  

更新 Submodule
```
$ cd submodule/ && git pull origin master
```

---
## 啟動服務
執行腳本 `RunService.sh`

```
$ sh RunService.sh
```