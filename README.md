# Aniya

## 免责声明

该工具仅用于安全研究，禁止使用工具发起非法攻击等违法行为，造成的后果使用者负责

## 介绍

`Golang`免杀马生成工具，最重要的loader部分参考其他作者。

相较其他免杀工具具备以下优势：

1. GUI界面，不算难看，简单易懂
2. 可自定义多种反沙箱，其中检查微信的适合钓鱼
3. 可自定义多种编译选项，额外支持garble编译环境

## 生成免杀马

在生成免杀马之前请注意以下四件事

1. 确保安装`Golang`且环境变量中包含`go`否则无法编译
2. 请在当前目录先执行`go env -w GO111MODULE=on`然后`go mod download`命令下载依赖
3. 生成木马时需将杀软关闭，go产生的中间文件会被查杀
4. 如果下载依赖过慢配置镜像`go env -w GOPROXY=https://mirrors.aliyun.com/goproxy`

一切就绪后就可以开始生成了

## 参考

- https://github.com/safe6Sec/GolangBypassAV
- https://github.com/Ne0nd0g/go-shellcode
