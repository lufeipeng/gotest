# gotest
学习Go的Project

//2015-10-19

//我的go开发环境

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

