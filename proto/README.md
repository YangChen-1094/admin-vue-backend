##1����ȡprotoc-gen-go
go get github.com/golang/protobuf/protoc-gen-go

�����������±�������ʱ��

```
protoc --proto_path=src --go_out=build/gen src/foo.proto src/bar/baz.proto
```
���������ȡ�ļ�src/foo.proto��src/bar/baz.proto���⽫��������������ļ�build/gen/foo.pb.go��build/gen/bar/baz.pb.go

##2��proto�﷨ ������������ɵ�pb�ļ�Ŀ¼
```
option go_package = "aaa;bbb";
aaa ��ʾ���ɵ�go�ļ��Ĵ�ŵ�ַ�����Զ�����Ŀ¼�ġ�
bbb ��ʾ���ɵ�go�ļ������İ���
```

##3��ֱ��ִ��protoGo.sh�ļ�
```
./protoGo.sh
#�������ɱ��ļ����µ�����pbЭ��
```