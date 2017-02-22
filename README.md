# DCMP 
[![Build Status](https://travis-ci.org/silenceper/dcmp.svg?branch=master)](https://travis-ci.org/silenceper/dcmp)


Distributed Configuration Management Platform

提供了一个etcd的管理界面，可通过界面修改配置信息，借助confd可实现配置文件的同步。 

###  安装 && 启动

```go
> go get github.com/silenceper/dcmp
> ./service.sh

```

### 配置

```

listen: "0.0.0.0:8000"  # 监听的IP，端口

base_path: "/config"    #etcd读取的根目录

endpoints:              # etcd 接入地址
   - "http://127.0.0.1:2379"

# etcd ssl 配置
#
# ca_file: "/path/to/ca.crt"
# cert_file: "/path/to/client.crt"
# key_file: "/path/to/client.key"



```

### 界面预览
访问 `http://127.0.0.1:8000/`


![snapshot](./docs/image.png)
