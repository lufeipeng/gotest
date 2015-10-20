# gotest
学习Go的Project

//2015-10-19

//我的go开发环境
运行测试用例的方法, 下载代码后，进入对应的目录，比如mydoc, 运行go test -v 即可
根目录的webserver.go 简单的webserver测试函数
main.go 入口函数, elipse下面一个main package只能有一个main func, 暂时修改为了main_back

采用的是windows下面eclipse + php plugin
并把测试的代码放入到github上面, 可以通过github就行源代码管理

//关于godoc

godoc可以实时的刷新, 修改源代码后doc 网页直接刷新就可以看到

运行的命令 godoc -http:6060 
可以看到goroot下面自己的package信息, 比如自己定义的工程
http://localhost:6060/pkg/github.com/lufeipeng/gotest/godoc/
的function中不能使用fmt做输出, go test 的时候会报错

//gotest
以下目录测试mystrings package
go test ./mystrings -v -x
注意godoc目录的example文件example_test.go, 注释的output必须与实际function的输出对应上, 否则go test 失败

//oracle 关于eclipse ctrl + 右键跳转引用
需要下载golang tools, 插件原来的下载地址失效
go get -u golang/x/tools/cmd/oracle
从github.com\golang\tools 下载tools 放入到gopath的对应路径, 在运行go install 即可


