# go-curl

[![MIT licensed][3]][4]

[3]: https://img.shields.io/badge/license-MIT-blue.svg
[4]: LICENSE

#### 介绍
基于 Golang 内置 net/http 包中 http.Client 结构方法，实现 HTTP 客户端请求方式，可以直接使用 Get\Post\Put\Patch\Delete
方式发起 HTTP 请求，实现远程数据请求，使得开发过程调用更简单、便捷。

#### 安装
```shell
// github
go get github.com/lihao1988/go-curl

// gitee
go get gitee.com/lihao1988/go-curl
```

## 版本要求
Go 1.15 or above.

##HTTP 函数
### 1. Get Request
```go
// example - 1
client := NewClient("http://www.example.com")
dataBytes, err := client.Get("/api?param=1", nil) // param url
fmt.Println("Get: ", string(dataBytes), err)

// example - 2
client := NewClient("http://www.example.com")
data := map[string]string{"param":"1"} 
dataBytes, err := client.Get("/api", data) // param rawQuery
fmt.Println("Get: ", string(dataBytes), err)

// example - 3
client := NewClient("http://www.example.com")
data := map[string]string{"param":"1"}
dataBytes, err := client.Curl("/api", Get, data, JsonType)
fmt.Println("Get: ", string(dataBytes), err)
```

### 2. Post Request
```go
// example - 1
client := NewClient("http://www.example.com")
data := map[string]string{"param":"1"}
dataBytes, err := client.Post("/api", data) // json
fmt.Println("Post: ", string(dataBytes), err)

// example - 2
client := NewClient("http://www.example.com")
data := map[string]string{"param":"1"}
dataBytes, err := client.PostByForm("/api", data) // form
fmt.Println("Post: ", string(dataBytes), err)

// example - 3
client := NewClient("http://www.example.com")
data := map[string]string{"param":"1"}
dataBytes, err := client.Curl("/api", Post, data, FormType)
fmt.Println("Post: ", string(dataBytes), err)
```

### 3. Put Request
```go
// example - or PutByForm\Curl
client := NewClient("http://www.example.com")
data := map[string]string{"param":"1"}
dataBytes, err := client.Put("/api", data) // json
fmt.Println("Put: ", string(dataBytes), err)
```

### 4. Patch Request
```go
// example - or PatchByForm\Curl
client := NewClient("http://www.example.com")
data := map[string]string{"param":"1"}
dataBytes, err := client.Patch("/api", data) // json
fmt.Println("Patch: ", string(dataBytes), err)
```

### 5. Delete Request
```go
// example - or Curl
client := NewClient("http://www.example.com")
data := map[string]string{"param":"1"}
dataBytes, err := client.Delete("/api", data)
fmt.Println("Delete: ", string(dataBytes), err)
```
