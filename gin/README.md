## 重定向，在控制台返回的是 a 链接

```go
//  $ curl http://localhost:8080/test
//  <a href="http://www.google.com/">Moved Permanently</a>.
```

## HandleContext 内部会执行重置

```go
// HandleContext 内部会执行重置
// https://matthung0807.blogspot.com/2021/11/gin-engine-handlercontext-func.html
```

## 文件上传

<img src="http://t-blog-images.aijs.top/img/202304241621407.webp" />

## 时间格式化

<a href="https://cloud.tencent.com/developer/article/1467743" target="_blank" >时间格式化 </a>

## golang

1. 空置、零值过滤
2. 空置过滤，零值保留

## gin 数据绑定

1. ShouldBindQuery
2. ShouldBind
3. QueryMap/PostFormMap
4. DefaultQuery/Query/c.Request.URL.Query().Get
5. ShouldBind/ShouldBindJSON/ShouldBindXML
6. 模型绑定和验证
7. Param
8. 想绑定多次怎么办

<a href="https://cloud.tencent.com/developer/article/1955340" target="_blank" >golang 的 gin 框架，各种接收参数的方式和各种绑定的区别</a>

1. 使用接收单个参数各种方法：
2. 使用各种绑定方法

<a href="https://blog.csdn.net/wohu1104/article/details/121928193" target="_blank" >Gin 框架学习笔记（02）— 参数自动绑定到结构体</a>
