# 土豆片控制面板

基于腾讯云API实现的轻量级云资源控制面板

##  功能列表

支持的功能和开发进度，请参阅 [Issues #1](https://github.com/tdp-resource/tdp-cloud/issues/1)

WebUI界面请查看文档 [界面预览](https://github.com/tdp-resource/tdp-cloud/blob/main/docs/界面预览.md)

## 开发说明

### 启动开发服务

在项目目录运行  `serve.bat` 或 `./serve.sh`

### 编译为二进制

在项目目录运行 `build.bat` 或 `./build.sh`。你还可以下载 [稳定版](https://github.com/tdp-resource/tdp-cloud/releases) 或 [午夜构建版](http://curl.rpc.im/?dir=/tdp-cloud)

### 额外参数设置

如果项目无法运行或编译，请尝试设置系统环境变量或临时环境变量

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## 部署说明

### 安装服务端

https://github.com/tdp-resource/tdp-cloud/blob/main/docs/部署服务端.md

### 添加子节点

https://github.com/tdp-resource/tdp-cloud/blob/main/docs/添加子节点.md

### 添加腾讯云账号

https://github.com/tdp-resource/tdp-cloud/blob/main/docs/添加腾讯云账号.md

## 其他

License [GPL-3.0](https://opensource.org/licenses/GPL-3.0)

Copyright (c) 2022 - 2023 TDP Cloud
