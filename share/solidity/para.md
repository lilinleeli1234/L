# 平行链环境搭建

<!-- TOC -->
- [官方版本（仅支持类linux环境）](#官方版本仅支持类linux环境)
    - [1 获取文件](#1-获取文件)
    - [2 解压](#2-解压)
    - [3 配置文件修改](#3-配置文件修改)
- [本地编译](#本地编译)
    - [1 切换路径](#1-切换路径)
    - [2 获取最新Chain33代码分支进行编译](#2-获取最新chain33代码分支进行编译)
        - [2.1 linux下编译](#21-linux下编译)
        - [2.2 windows下编译](#22-windows下编译)
<!-- /TOC -->

## 官方版本（仅支持类linux环境）
### 1 获取文件
```bash 
wget https://bty.oss-ap-southeast-1.aliyuncs.com/chain33/paraChain.tar.gz
```

### 2 解压
```bash
tar zxvf paraChain.tar.gz
```

### 3 配置文件修改

ParaRemoteGrpcClient项取值为："101.37.227.226:8802,39.97.20.242:8802,47.107.15.126:8802,jiedian2.33.cn"

mainnetJrpcAddr项取值为："http://jiedian1.33.cn:8801"


修改后如下所示：

![paraRemoteGrpcClient](https://github.com/lynAzrael/L/blob/master/share/img/paraRemoteGrpcClient.png)

![mainnetJrpcAddr](https://github.com/lynAzrael/L/blob/master/share/img/mainnetJrpcAddr.png)

启动chain33进程
```bash
cd paraChain && ./chain33 -f chain33.para.toml
```

由于官方提供的bin文件仅支持类linux环境，因此windows环境下需要[本地编译](#本地编译)二进制文件。

## 本地编译
### 1 切换路径

```bash
cd ${GoPath}/src/github.com/33cn
```

>GoPath为安装Go编译器之后的环境变量，如果未安装Go，请参考环境准备中的[Go环境安装](https://chain.33.cn/document/81#1.1%20Go%20%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85)

>除GoPath之外的其他子目录如果不存在，请手工创建。

### 2 获取最新Chain33代码分支进行编译
```bash
git clone git@github.com:33cn/plugin.git
```
#### 2.1 linux下编译
```bash
cd plugin && make
```
#### 2.2 windows下编译
cmd模式使用如下指令：

```bash
cd plugin && build.bat
```

powershell使用如下指令：

```bash
(cd .\plugin) -or (.\build.bat)
```

