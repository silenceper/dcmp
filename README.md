# DCMP 
[![Build Status](https://travis-ci.org/silenceper/dcmp.svg?branch=master)](https://travis-ci.org/silenceper/dcmp)


Distributed Configuration Management Platform

提供了一个etcd的管理界面，可通过界面修改配置信息，借助confd可实现配置文件的同步。 

###  安装 && 启动

```go
> go get github.com/silenceper/dcmp
> ./service.sh

```

### 界面预览
访问 `http://127.0.0.1:8000/`


![snapshot](./docs/image.png)
