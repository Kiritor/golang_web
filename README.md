# golang_web
golang简单web服务器(登录小例子),效果预览:
<center>![golang_web](http://kiritor.github.io/img/golang_web.gif)</center>
##目录结构
bin--gobuild生成的执行程序目录
    /template --静态文件(和执行程序同级)
src--源代码
    /golang_web --包
template--静态文件(编译之后,将其置于执行程序所在的目录下)
sql--sql建表语句

##运行
1、搭建go环境,代码复制到gopath中(可以多个gopath)
2、安装go语言mysql驱动,mymysql
3、创建mysql数据库
4、修改ajaxController.go代码中数据库设置的参数
5、编译go build golang_web
6、将生成的执行文件置于template的父级目录运行服务(或者将template置于可执行文件目录)
7、输入localhost:8888/ 访问