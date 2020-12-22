##1、获取protoc-gen-go
go get github.com/golang/protobuf/protoc-gen-go

当你运行如下编译命令时：

```
protoc --proto_path=src --go_out=build/gen src/foo.proto src/bar/baz.proto
```
编译器会读取文件src/foo.proto和src/bar/baz.proto，这将会生成两个输出文件build/gen/foo.pb.go和build/gen/bar/baz.pb.go

##2、proto语法 定义包名及生成的pb文件目录
```
option go_package = "aaa;bbb";
aaa 表示生成的go文件的存放地址，会自动生成目录的。
bbb 表示生成的go文件所属的包名
```

##3、直接执行protoGo.sh文件
```
./protoGo.sh
#他会生成本文件夹下的所有pb协议
```