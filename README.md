# DCMP 
[![Build Status](https://travis-ci.org/silenceper/dcmp.svg?branch=master)](https://travis-ci.org/silenceper/dcmp)


Distributed Configuration Management Platform

提供了一个etcd的管理界面，可通过界面修改配置信息，借助confd可实现配置文件的同步。 

###  安装 && 启动

```go
> go get github.com/silenceper/dcmp
> ./service.sh

```

### 如果./service.sh报错，出现类似有些包找不到的问题，需要在本地添加一些依赖

依赖的项目有：gin、gopkg.in和confd。

confd的安装方法：
```
$ mkdir -p $GOPATH/src/github.com/kelseyhightower
$ git clone https://github.com/kelseyhightower/confd.git $GOPATH/src/github.com/kelseyhightower/confd
$ cd $GOPATH/src/github.com/kelseyhightower/confd
$ ./build
```

gopkg.in的安装方法：
```
go get gopkg.in/yaml.v1
```

gin的安装方法：
```go
//👆上面一步我们已经安装了gopkg,如果安装gopkg失败,这步无法执行
//另外墙内用户这一步可能会比较卡,建议多执行几次。
go get gopkg.in/gin-gonic/gin.v1
```

### 界面预览
访问 `http://127.0.0.1:8000/`


![snapshot](./docs/image.png)
