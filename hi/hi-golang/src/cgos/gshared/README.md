# gshard
> go 输出dll文件


## 运行

> 指定文件编译
```shell
# 编译目录
go build -buildmode=c-shared -o main.dll main.go
```



> 指定目录编译

```shell
# 编译目录
go build -buildmode=c-shared -o shared.dll .
```