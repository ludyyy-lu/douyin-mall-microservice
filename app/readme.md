# 业务代码开发目录


使用cwgo根据idl生成RPC代码
新建业务代码目录，
例如/app/auth，终端进入后执行：
```
cwgo server -I ../../idl/ --type RPC --module douyinmall --service douyinmall --idl ../../idl/auth.proto
```


