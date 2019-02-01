### 看了go语言，python和go在devops中应用还是很大，比如go的项目有k8s,docker,grafana等，python的开源软件有salt，ansible，openstack等

## beego
[beego官方文档](https://beego.me/docs/intro/)

## go 和java都是强类型语言，他们都很像
[go英文官方文档](https://golang.org/doc/)

[go中文文档](http://docscn.studygolang.com/doc/)

[Golang标准库文档](https://studygolang.com/pkgdoc)

[入门指南](https://www.kancloud.cn/kancloud/the-way-to-go/72457)

[正则表达式](https://www.cnblogs.com/benlightning/articles/4440940.html)

[go入门文档git_book](https://go.fdos.me)

##对比java/python/php  确实go还有很多不足的地方
* 导入自定义的要放到gopath目录
* .....

### go使用场景
* 高并发处理，类似于c,支持协程，为并发而生
* 高性能的http服务器
* 代码加密

### py为什么有全局锁
* 全局锁，同一时刻只准许一个线程在跑，这样的目的是为了防止多个线程访问同一资源排斥的问题，所有就有互斥锁的概念
* php 天生不支持线程，官方也不建议
* 全局锁是解释性语言的通病，所以解释性语言执行效率比编译型语言效率低

解释性语言有perl,python,php ,编译型语言有 java,golang

## go 引入自定义的包(只能放到gopath中)
```
(1)要想引入自定义的包，需要自定义包的路径为GOPATH路径中。
(2)golang和C或php不一样，不会自动查找当前路径下的文件，必须先在$GOAPTH里添加自己工程的路径
(3)自定义包里面对外提供的API函数，首字母必须大写
(4)通常情况下，import的包都是相对$GOPATH/src目录引入的，比如从github上面clone下来的项目，直接放到$GOPATH/src目录下，就可以直接import,例如：
如果项目的import路径是这样写的：
import "github.com/yourname/projectname"
需要将项目代码放置在：
$GOAPTH/src/github.com/yourname/projectname/下
```

