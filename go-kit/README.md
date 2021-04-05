# Go Kit

Go kit 微服务示例

## 介绍

相关链接
- 官网： https://gokit.io/
- GitHub：https://github.com/go-kit/kit
- 官方例子：https://gokit.io/examples/
- Prometheus：https://prometheus.io/
- Prometheus Go Client：https://github.com/prometheus/client_golang

目录：

    / stringsvc 基础示例，字符串转换服务
    / stringsvc-middlewares 字符串转换服务 + 中间件（logger、prometheus 指标）

## Usage
如何运行这些 demo

### stringsvc
运行服务
```bash
cd stringsvc
go run main.go
```

POST 调用：
```bash
$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
{"v":"HELLO, WORLD"}

$ curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
{"v":12}
```

### stringsvc-middlewares
运行服务
```bash
$ go run main.go
msg=HTTP addr=:8080
method=uppercase input="hello, world" output="HELLO, WORLD" err=null took=823ns
method=uppercase input="hello, world" output="HELLO, WORLD" err=null took=747ns
```
类似上面那样调用接口，会输出日志。

同时，如果打开 `http://127.0.0.1:8080/metrics` ，可以看到相关指标的变动。
