# ginxapp

这是一个为使用ginx提供api服务而创建的一个模板项目，旨在方便使用者可以快速搭建一个ginx基础项目。因此，你需要先安装gonew工具。

## 安装gonew工具

```bash
❯ go install golang.org/x/tools/cmd/gonew@latest
go: downloading golang.org/x/tools v0.24.0
go: downloading golang.org/x/mod v0.20.0
```

检查下是否安装成功：
```bash
❯ gonew
usage: gonew srcmod[@version] [dstmod [dir]]
See https://pkg.go.dev/golang.org/x/tools/cmd/gonew. 
```

## Copy项目到本地

```bash
gonew github.com/whencome/ginxapp your.module/path
```

至此，你已经在本地搭建了一个使用ginx的项目框架。

## 其他说明

ginxapp本身只是用于项目框架搭建，主要用于共享代码结构，统一项目布局。其中的代码并没有经过多少测试，在使用的时候需根据需要修改。
