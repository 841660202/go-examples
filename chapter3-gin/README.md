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
4. DefaultQuery/Query/c.Request.URL.Query().Get https://gin-gonic.com/zh-cn/docs/examples/query-and-post-form/
5. ShouldBind/ShouldBindJSON/ShouldBindXML
6. 模型绑定和验证
7. Param
8. Multipart/Urlencoded 绑定 https://gin-gonic.com/zh-cn/docs/examples/multipart-urlencoded-binding/
9. Multipart/Urlencoded 表单 https://gin-gonic.com/zh-cn/docs/examples/multipart-urlencoded-form/
10. 想绑定多次怎么办 https://gin-gonic.com/zh-cn/docs/examples/bind-body-into-dirrerent-structs/

<a href="https://cloud.tencent.com/developer/article/1955340" target="_blank" >golang 的 gin 框架，各种接收参数的方式和各种绑定的区别</a>

1. 使用接收单个参数各种方法：
2. 使用各种绑定方法

TODO:
<a href="https://blog.csdn.net/wohu1104/article/details/121928193" target="_blank" >Gin 框架学习笔记（02）— 参数自动绑定到结构体</a>

## gin 日志切割

# tag / validator

gin 使用 go validator, git@github.com:go-playground/validator.git
validator

1. 校验 struct
2. 校验 var

为了搞清楚使用，有必要学习:

1. validator 的 example
2. gin 的 tag 有哪些
3. 如何自定义 tag

--

1. gin 的 tag 有哪些，各有什么用

<a href="https://www.cnblogs.com/jiujuan/p/13823864.html" target="_blank" >见</a>
<a href="https://juejin.cn/post/6863765115456454664" target="_blank" >见</a>

2. 如何自定义 tag

# validator 总结

1. 校验单个字段用：validate.Var
2. 校验结构体用： validate.Struct
3. 对于 slice/map 类型字段，需要结合 dive,keys,规则,endkeys
4. 嵌套结构
   1. 规则和结构体可以拆分开：validate.ValidateMap
   2. 嵌套结构，可以使用嵌套规则：validate.ValidateMap
   3. validate.ValidateMap，错误信息需要使用 len 进行判断
5. 自定义
   1. 自定义 tag：validate.RegisterValidation
   2. 自定义校验：validate.RegisterValidation
6. 对 tag 进行重命名：validate.RegisterAlias
7. 给结构体注册校验规则：validate.RegisterStructValidationMapRules
8. 给某些类型注册自定义的处理函数：validate.RegisterCustomTypeFunc
