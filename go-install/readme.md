下载链接：https://studygolang.com/dl
* 源码：一般不用

## 在 window 上安装Go

1. 下载后，双击安装，除了修改安装目录外，一路下一步
    * 安装目录可设置为：`D:/program files/go`
2. 在安装目录中，创建目录
    * `D:/program files/go/third_go`
    * 再创建两个目录
        * `D:/program files/go/third_go/bin` 安装一些第三方的可执行
        * `D:/program files/go/third_go/pkg` 安装第三方的一些Go语言包
3. 打开window设置-环境变量-系统的环境变量-新建
    * GOROOT：D:/program files/go
    * GOPATH：D:/program files/go/third_go
    * GOPROXT（GO语言代理）：https://goproxy.cn,direct
        * 设置原因：方便之后安装一些第三方库的话可能网速很慢，或者根本无法安装，这时候，如果通过代理的话情况会好一些。
        * 这不是必须的：如果在国外的话，根本不需要设置。
4. 往 PATH 中-新增
    * %GOROOT%\bin
    * %GOPATH%\bin
5. 至此，window下的GO语言开发环境就搭建好了，使用如下：
    * `go version`、`go env`

## 在 Linux 上搭建GO语言的开发环境（Mac上和Linux上一致）

1. 右键复制 Linux 安装包的链接，打开Linux系统，随便在一个目录下，使用 wget 下载
    * wget https://studygolang.com/dl/golang/go1.21.3.linux-amd64.tar.gz
3. 解压 (默认解压到当前目录，可以通过 `-c` 指定解压目录)
    * `tar zxvf goxxx`
4. 在当前目录下创建 third_go
    ```shell
    mkdir third_go
    cd third_go
    mkdir pkg
    mkdir bin
    ```
5. 设置环境变量 将当前目录 `pwd` 放到 GOROOT 环境变量中
    * 设置当前用户的：`vim ~/.bashrc`
    * 设置整个系统的：`sudo vim /etc/profile`, 到文件末尾
        * `export GOROOT=/home/ubuntu/soft/go`  刚刚下载的压缩包解压的路径
        * `export GOPATH=/home/ubuntu/soft/third_go`
        * `export GOPROXY=https://goproxy.cn,direct`
        * `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`
        * `wq` 保存退出，使用 `source /etc/profile` 立即生效
6. 验证 `go version` 、 `go env`