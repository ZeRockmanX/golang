# golang
study

Go 语言教程

      http://www.runoob.com/go/go-tutorial.html
      http://c.biancheng.net/golang/

Go 语言环境安装

Go 语言支持以下系统：

      Linux
      FreeBSD
      Mac OS X（也称为 Darwin）
      Windows

安装包下载地址为：https://golang.org/dl/。

如果打不开可以使用这个地址：https://golang.google.cn/dl/。

各个系统对应的包名：

      操作系统  包名
      Windows  go1.12.5.windows-amd64.msi
      Linux    go1.12.5.linux-amd64.tar.gz
      Mac      go1.12.5.darwin-amd64-osx10.8.pkg
      FreeBSD  go1.12.5.freebsd-amd64.tar.gz

UNIX/Linux/Mac OS X, 和 FreeBSD 安装
      
      1、直接安装: sudo apt install golang-go
      2、获取最新的软件包源，并添加至当前的apt库
      sudo add-apt-repository ppa:longsleep/golang-backports
      sudo apt-get update
      sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 123456789(Key)
      或
      sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 123456789(Key)
      sudo apt-get update
      sudo apt-cache policy golang-go #查看，如果有多个版本
      sudo apt-get install golang-go=<version>

以下介绍了在UNIX/Linux/Mac OS X, 和 FreeBSD系统下使用源码安装方法：

      1、下载二进制包：
            sudo apt install curl
            curl -O https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
            // 下载验证： sha256sum go1.12.5.linux-amd64.tar.gz
      2、将下载的二进制包解压
            tar xvf go1.12.5.linux-amd64.tar.gz
            sudo chown -R root:root ./go
            sudo mv go /usr/local
            虽然/usr/local/go是官方推荐的位置，但有些用户可能更喜欢或需要不同的路径。
            此时，使用Go将需要在命令行中指定其安装位置的完整路径。 
            为了使与Go的交互更加用户友好，我们将设置一些路径。
      3、将 /usr/local/go/bin 目录添加至PATH环境变量：
            vi ~/.profile
            export GOROOT="/usr/local/go"
            export GOPATH="$HOME/go"
            export PATH="$PATH:$GOROOT/bin:$GOPATH/bin"
            source ~/.profile
            添加同样内容到~/.bashrc

注意：MAC 系统下你可以使用 .pkg 结尾的安装包直接双击来完成安装，安装目录在 /usr/local/go/ 下。

Windows 系统下安装

Windows 下可以使用 .msi 后缀(在下载列表中可以找到该文件，如go1.4.2.windows-amd64.msi)的安装包来安装。默认情况下 .msi 文件会安装在 c:\Go 目录下。你可以将 c:\Go\bin 目录添加到 Path 环境变量中。添加后你需要重启命令窗口才能生效。

安装测试

创建工作目录 C:\>Go_WorkSpace。

test.go 文件代码：

      package main
      import "fmt"
      func main() {
         fmt.Println("Hello, World!")
      }

使用 go 命令执行以上代码输出结果如下：

C:\Go_WorkSpace>go run test.go

      Hello, World!