<img src="http://t-blog-images.aijs.top/img/202304241057941.webp" style="width:500px;max-width:100%"/>

## 引用类型和值类型

1. Go 中严格区分引用类型和值类型

## 区分大小写

1.  首字母大写的方法可以被引用
2.  首字母大写的结构体可以被引用
3.  首字母大写的键可以被转化

## 打印日志

https://www.liwenzhou.com/posts/Go/fmt/

## 方法 中 值 vs 指针

区别在于：

1. 值方法：可通过指针和值调用，
2. 指针方法：只能通过指针来调用。

**为什么会有这条规则？**

1. 因为指针方法可以修改接收者；通过值调用它们会导致方法接收到该值的副本， 因此任何修改都将被丢弃，因此该语言不允许这种错误。
2. 若该值是可寻址的， 那么该语言就会自动插入取址操作符来对付一般的通过值调用的指针方法

如：变量 b 是可寻址的，因此我们只需通过 b.Write 来调用它的 Write 方法，编译器会将它重写为 (&b).Write

## JSON

**结构体数据-> string**

1. 大写被转化，小写不被转化
2. 有标签的用标签，无标签的，用 key 名字
3. Channel， complex 以及函数不能被编码 json 字符串, 循环的数据结构也不行，它会导致 marshal 陷入死循环
4. 有时为了通用性，或使代码简洁，我们希望有一种类型**可以接受各种类型**的数据，并进行 json 编码。这就用到了 interface{}类型

**json 字符串解析**

1. json 字符串解析时，需要一个“接收体”接受解析后的数据，且 Unmarshal 时接收体必须传递指针。否则解析虽不报错，但数据无法赋值到接受体中。如这里用的是 StuRead{}接收。
2. 解析时，接收体可自行定义。json 串中的 key 自动在接收体中寻找匹配的项进行赋值。匹配规则是：
   1. 先查找与 key 一样的 json 标签，找到则赋值给该标签对应的变量(如 Name)。
   2. 没有 json 标签的，就从上往下依次查找变量名与 key 一样的变量，如 Age。或者变量名忽略大小写后与 key 一样的变量。如 HIgh，Class。第一个匹配的就赋值，后面就算有匹配的也忽略。
   3. (前提是该变量必需是可导出的，即首字母大写)。
   4. 不可导出的变量无法被解析（如 sex 变量，虽然 json 串中有 key 为 sex 的 k-v，解析后其值仍为 nil,即空值）
   5. 当接收体中存在 json 串中匹配不了的项时，解析会自动忽略该项，该项仍保留原值。如变量 Test，保留空值 nil。
3. 你一定会发现，变量 Class 貌似没有解析为我们期待样子。因为此时的 Class 是个 interface{}类型的变量，而 json 串中 key 为 CLASS 的 value 是个复合结构，不是可以直接解析的简单类型数据（如“张三”，18，true 等）。所以解析时，由于没有指定变量 Class 的具体类型，json 自动将 value 为复合结构的数据解析为 `map[string]interface{}`类型的项。也就是说，此时的 struct Class 对象与 StuRead 中的 Class 变量没有半毛钱关系，故与这次的 json 解析没有半毛钱关系。

## 内存模型

<a href="./go程/go模型.md" target="_blank" >Go 内存模型</a>

## 并发

<a href="https://www.modb.pro/db/65265" target="_blank" >见</a>

## Go Select 详解

**GO 为什么引入 select?**

select 是一种 go 可以处理多个通道之间的机制，看起来和 switch 语句很相似，但是 select 其实和 IO 机制中的 select 一样，多路复用通道，随机选取一个进行执行。

如果说通道(channel)实现了多个 goroutine 之前的同步或者通信，那么 select 则实现了多个通道(channel)的同步或者通信

select 具有阻塞的特性。

**有怎样的使用场景？**

<a href="https://www.jianshu.com/p/66edceabd5f6" target="_blank" >Go Select 详解</a>

## defer

设计动机：两点好处：

1. 第一，它能确保你不会忘记关闭文件。如果你以后又为该函数添加了新的返回路径时， 这种情况往往就会发生。
2. 第二，它意味着“关闭”离“打开”很近， 这总比将它放在函数结尾处要清晰明了

使用场景：

1. 场景：解锁互斥和关闭文件

特点：

1. 被推迟的函数按照`后进先出（LIFO）`的顺序执行

## 追加

1. append 会在切片末尾追加元素并返回结果
2. 将一个切片追加到另一个切片，**必须**在调用的地方使用 ...

```Go
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...) // 这里...不可缺少
fmt.Println(x)
```

## 区分某项是不存在还是其值为零值

**提及**
<a href="https://go-zh.org/doc/effective_go.html#%E8%BF%BD%E5%8A%A0:~:text=%E6%9C%89%E6%97%B6%E4%BD%A0%E9%9C%80%E8%A6%81-,%E5%8C%BA%E5%88%86%E6%9F%90%E9%A1%B9%E6%98%AF%E4%B8%8D%E5%AD%98%E5%9C%A8%E8%BF%98%E6%98%AF%E5%85%B6%E5%80%BC%E4%B8%BA%E9%9B%B6%E5%80%BC,-%E3%80%82%E5%A6%82%E5%AF%B9%E4%BA%8E%E4%B8%80%E4%B8%AA" target="_blank" >提及：区分某项是不存在还是其值为零值</a>

**应对**

<a href="https://learnku.com/go/t/49332" target="_blank" >Golang 中使用 JSON 时如何区分空字段和未设置字段？</a>

<a href="https://www.cnblogs.com/joyswings/p/9864568.html" target="_blank" >go 语言的 null 值问题</a>

## 仓库

<!-- git init
git add .
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:841660202/go-examples.git
git push -u origin main -->

## 改变背景色

```js
var st = document.createElement('style');
st.innerHTML = `
   *{
      color: #a1a1a1!important;
      background: #000!important;
      text-shadow: none!important;
   }
`;
document.head.append(st);
```
