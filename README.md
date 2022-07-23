# AniYa-GUI免杀框架

## 免责声明

该工具仅用于安全研究，禁止使用工具发起非法攻击等违法行为，造成的后果使用者负责。免杀具有时效性，免杀效果无法保证。

## 介绍

`Golang`免杀马生成工具，在重复造轮子的基础上尽可能多一点自己的东西，最重要的loader部分参考其他作者。

相较其他免杀工具具备以下优势：

1. 使用[fyne](https://github.com/fyne-io/fyne)的GUI界面，不算难看，简单易懂，还有个炫酷的进度条！wakuwaku(*^▽^*)
2. 可自定义多种反沙箱，其中检查微信的适合钓鱼
3. 可自定义多种编译选项，支持garble编译环境
4. 分离免杀(本地/HTTP)
5. 支持打包PE文件（如`mimikatz`）

## 使用

![image-20220711114540147](https://github.com/piiperxyz/AniYa/blob/main/img/Snipaste_2022-07-22_18-00-40.jpg)

1. 后缀支持bin/exe/dll，可输入绝对路径或相对路径或点击按钮选择。默认beacon.bin。（必选）
2. 生成木马的名称。默认result.exe。（必选）
3. 本地分离免杀，可输入绝对路径或相对路径，但生成的文件（默认code.txt）是固定在当前目录生成，木马会去读取目标路径下的分离shellcode
4. 远程分离免杀，木马去请求网络地址下载shellcode，加密的shellcode为当前目录的code.txt
6. 选择shellcode加密算法（必选）
7. 选择loader（必选）
8. 反沙箱
9. 编译选项

| 反沙箱参数     | 说明                                                         |
| -------------- | ------------------------------------------------------------ |
| timestart      | 添加启动参数，参数为运行木马的机器的系统时间格式为HHMMDD,如9日11:08，则启动参数为1189，不填充0 |
| ramcheck       | 检查内存，小于4G则退出                                       |
| cpunumbercheck | 检查CPU核数，小于4则退出                                     |
| wechatcheck    | 检查微信进程，不存在则退出                                   |
| disksizecheck  | 检查C盘硬盘大小，小于60G则退出                               |

loader的说明搬一下[4ra1n](https://github.com/4ra1n)的介绍。

| 模块名                   | 简介                                                     |
| ------------------------ | -------------------------------------------------------- |
| CreateFiber              | 利用Windows CreateFiber函数                              |
| CreateProcess            | 利用Windows CreateProcess函数在挂起状态下创建进程        |
| CreateRemoteThread       | 远程进程注入ShellCode（注入explorer.exe）                |
| CreateRemoteThreadNative | 和上一条区别在于使用更底层的方式（注入explorer.exe）     |
| CreateThread             | 利用Windows CreateThread函数                             |
| CreateThreadNative       | 和上一条区别在于使用更底层的方式                         |
| EarlyBird                | 注入的代码在进程主线程的入口点之前运行                   |
| EtwpCreateEtwThread      | 利用Windows EtwpCreateEtwThread函数在进程中执行ShellCode |
| HeapAlloc                | 创建一个可供调用进程使用的堆并分配内存写入ShellCode      |
| NtQueueApcThreadEx       | 在当前进程的当前线程中创建一个特殊用户APC来执行ShellCode |
| RtlCreateUserThread      | 利用Windows RtlCreateUserThread函数（注入explorer.exe）  |
| UuidFromString           | 利用Windows UuidFromStringA函数                          |

编译参数说明，不包含`ldflag -s -w`及`-trimpath` ，默认自带

|                             参数                             |                           参数说明                           |
| :----------------------------------------------------------: | :----------------------------------------------------------: |
|                             race                             |       使用竞态检测器-race进行编译（可能提高免杀效果）        |
|                             Hide                             | ~~隐藏窗口ldflags -H windowsgui（可能降低免杀效果）~~<br>更换为调用https://github.com/lxn/win,免杀效果增强，但有一闪而过的黑框 |
|        [garble](https://github.com/burrowers/garble)         | 使用编译混淆器garble来编译，需事先安装好，编译速度会慢一些（推荐） |
| [literalobf](https://github.com/burrowers/garble#literal-obfuscation) |        garble特有的参数，混淆所有字符串等（建议勾选）        |
| [randomseed](https://github.com/burrowers/garble#determinism-and-seeds) | garble特有的参数，使编译变的更随机，更加难以逆向（建议勾选） |

## 安装

已经编译好的程序可以从[realeases](https://github.com/piiperxyz/AniYa/releases)下载

#### 从源码编译

~~构建源代码的需要依赖项是[keystone 引擎](https://github.com/keystone-engine/keystone)，请按照[这些](https://github.com/keystone-engine/keystone/blob/master/docs/COMPILE.md)说明安装库。然后按照以下步骤进行编译~~

因确定sgn被拉黑，取消相关功能，现在直接编译很方便
```
直接go build
或者安装fyne之后使用fyne的打包工具来打包fyne package -icon favicon.ico
```

~~keystone安装比较麻烦，可以自行将sgn的相关功能注释掉，人工对shelllcode进行sgn混淆。~~

## 环境准备

在生成免杀马之前请注意以下四件事

1. 确保安装`Golang`且环境变量中包含`go`否则无法编译
2. 请在当前目录先执行`go env -w GO111MODULE=on`然后`go mod download`命令下载依赖
3. 生成木马时需将杀软关闭，go产生的中间文件会被查杀
4. 如果下载依赖过慢配置镜像`go env -w GOPROXY=https://mirrors.aliyun.com/goproxy`。国内用户建议配置。

一切就绪后就可以开始生成了

## 免杀效果

很多大佬都根据同个优秀的loader写了一些框架，目前啥选项都不配置有越来越多的杀软可以查杀，基本不能使用。

建议使用分离shellcode，技术简单但效果好。

写了一个能过DF的增强功能暂不放出。

另heapalloc的效果好一点。

sgn加密疑似已被提取特征，被WD和360拉黑了。

自测下来开启一些选项还是能免360和WD，希望各位大佬测试的时候关闭360和WD的自动上传样本功能，测试环境测，不要直接拖到VT上，火绒断网测就OK。

## 参考

欢迎各位大佬提PR!

感谢[不羡](https://github.com/V1rtu0l)师傅提供的GUI建议及反沙箱模块

- https://github.com/safe6Sec/GolangBypassAV
- https://github.com/Ne0nd0g/go-shellcode
- https://github.com/afwu/GoBypass\(4ra1n大佬的好像删了)

## 更新

- 1.1.0

更新HTTP分离免杀、变更窗口隐藏功能实现方式。

优化UI，现loader会根据打包EXE还是shellcode进行变更。

修复BUG。

感谢[夜中空想](https://github.com/imkitsch)提供的HTTP分离免杀和隐藏窗口功能

- 1.0.1

修复HIDE参数的bug

## TODO

- 签名伪造

- 自定义icon和versioninfo
- 文件捆绑功能

- ~~更多增强功能，如脱钩技术（halo's gate hell's gate unhook），考虑到免杀实效性，暂不考虑公开~~

