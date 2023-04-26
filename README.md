<img src="http://t-blog-images.aijs.top/img/202304241057941.webp" style="width:500px;max-width:100%"/>

## 引用类型和值类型

1. Go 中严格区分引用类型和值类型

## 零值和空值

在 Golang 的 Gin 框架中，可以使用以下方法来处理零值和空值：

1. 使用指针类型表示空值：在 Go 中，指针类型的零值为 `nil`，可以用来表示空值。因此，可以将某些变量定义为指针类型，并在需要表示空值的地方将其赋值为 `nil`。例如：

```go
var myString *string

// 将 myString 赋值为 nil，表示为空值
myString = nil

// 在路由处理函数中检查 myString 是否为空值
if myString == nil {
    // 处理空值的情况
}
```

2. 使用零值表示默认值：在 Go 中，基本数据类型的零值是其对应类型的默认值，可以用来表示零值。例如，整数类型的零值为 `0`，布尔类型的零值为 `false`。可以在定义变量时直接赋值为零值，表示默认值。例如：

```go
var myInt int

// 将 myInt 赋值为 0，表示为零值
myInt = 0

// 在路由处理函数中检查 myInt 是否为零值
if myInt == 0 {
    // 处理零值的情况
}
```

3. 使用空字符串表示空值：在 Go 中，空字符串 `""` 可以用来表示空值。可以将某些字符串类型的变量初始化为空字符串，表示为空值。例如：

```go
var myString string

// 将 myString 赋值为空字符串，表示为空值
myString = ""

// 在路由处理函数中检查 myString 是否为空值
if myString == "" {
    // 处理空值的情况
}
```

需要注意的是，处理零值和空值时，应根据具体的业务逻辑和需求进行判断和处理。以上只是一些示例方法，实际使用时应根据实际情况选择合适的方式。同时，在处理请求参数时，应注意对参数进行合法性校验，以确保代码的健壮性和安全性。

### gin 在接收 put 请求时候如何区分 空置和零值

在 Golang 的 Gin 框架中，接收 PUT 请求时，如果需要区分空置和零值，可以通过以下方法来实现：

1. 使用指针类型表示空值：将需要接收的字段定义为指针类型，例如 `*string`、`*int` 等。如果请求体中的字段值为空置，则对应的指针将为 `nil`，如果请求体中的字段值为零值，则对应的指针将为指向零值的指针。例如：

```go
type MyData struct {
    MyString *string `json:"my_string"`
    MyInt    *int    `json:"my_int"`
}
```

在路由处理函数中使用 `c.ShouldBindJSON()` 方法将请求体中的 JSON 数据绑定到 `MyData` 结构体中，然后可以通过判断指针是否为 `nil` 或指向零值来区分空置和零值的情况。例如：

```go
func UpdateData(c *gin.Context) {
    var data MyData
    if err := c.ShouldBindJSON(&data); err != nil {
        // 处理绑定错误
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 判断字段是否为空置
    if data.MyString == nil {
        // 处理空置的情况
    } else if *data.MyString == "" {
        // 处理空字符串的情况
    }

    // 判断字段是否为零值
    if data.MyInt == nil {
        // 处理空置的情况
    } else if *data.MyInt == 0 {
        // 处理零值的情况
    }

    // 处理其他情况
    // ...
}
```

2. 使用自定义类型并实现 `UnmarshalJSON()` 方法：可以自定义一个类型，并为其实现 `UnmarshalJSON()` 方法，在该方法中判断空置和零值的情况。例如：

```go
type MyString string

func (s *MyString) UnmarshalJSON(data []byte) error {
    if len(data) == 0 || string(data) == `""` {
        // 处理空置的情况
        *s = MyString("")
    } else {
        // 处理其他情况
        var str string
        if err := json.Unmarshal(data, &str); err != nil {
            return err
        }
        *s = MyString(str)
    }
    return nil
}
```

在路由处理函数中使用 `c.ShouldBindJSON()` 方法将请求体中的 JSON 数据绑定到结构体中，其中的自定义类型字段会自动调用 `UnmarshalJSON()` 方法进行解析，从而可以在其中处理空置和零值的情况。

需要注意的是，以上只是一些示例方法，实际使用时应根据实际情况选择合适的方式，并根据业务逻辑和需求进行判断和处理。同时，在处理 PUT 请求时，应注意对请求体数据进行合法性校验，以确保代码的健壮性和安全性。

### gin 中使用 UnmarshalJSON 处理 put 请求中对象某些字段零值和空置的完整代码例子

以下是一个使用 Gin 框架并结合 `UnmarshalJSON` 方法处理 PUT 请求中对象某些字段零值和空置的完整代码示例：

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyData struct {
	MyString *string `json:"my_string"`
	MyInt    *int    `json:"my_int"`
}

func (d *MyData) UnmarshalJSON(b []byte) error {
	var temp struct {
		MyString *string `json:"my_string"`
		MyInt    *int    `json:"my_int"`
	}

	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}

	if temp.MyString != nil && *temp.MyString == "" {
		d.MyString = nil // 处理空字符串的情况，设置为 nil
	} else {
		d.MyString = temp.MyString
	}

	if temp.MyInt != nil && *temp.MyInt == 0 {
		d.MyInt = nil // 处理零值的情况，设置为 nil
	} else {
		d.MyInt = temp.MyInt
	}

	return nil
}

func main() {
	r := gin.Default()

	r.PUT("/data", UpdateData)

	r.Run(":8080")
}

func UpdateData(c *gin.Context) {
	var data MyData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.MyString == nil {
		fmt.Println("MyString is empty or not provided")
	} else {
		fmt.Println("MyString:", *data.MyString)
	}

	if data.MyInt == nil {
		fmt.Println("MyInt is empty or not provided")
	} else {
		fmt.Println("MyInt:", *data.MyInt)
	}

	// 处理其他情况
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}
```

在这个示例中，通过自定义的 `UnmarshalJSON` 方法对 `MyString` 和 `MyInt` 进行了处理，可以正确处理 PUT 请求中的零值和空置的情况，并根据业务逻辑进行相应的处理。通过使用这种方式，可以在 Gin 框架中轻松地处理 PUT 请求中对象某些字段的零值和空置，并进行相应的处理。

### 更好的例子

抱歉，之前的例子可能并不是最佳实践。在处理 PUT 请求中对象某些字段零值和空置的情况时，通常可以直接使用 `omitempty` 标签来简化处理逻辑。以下是一个更为简洁的示例：

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyData struct {
	MyString string  `json:"my_string,omitempty"`
	MyInt    int     `json:"my_int,omitempty"`
}

func main() {
	r := gin.Default()

	r.PUT("/data", UpdateData)

	r.Run(":8080")
}

func UpdateData(c *gin.Context) {
	var data MyData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("MyString:", data.MyString)
	fmt.Println("MyInt:", data.MyInt)

	// 处理其他情况
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}
```

在这个示例中，使用了 `omitempty` 标签在结构体字段中，它的作用是在序列化时忽略零值和空置字段，从而避免了使用 `UnmarshalJSON` 方法来处理零值和空置的复杂逻辑。在处理 PUT 请求时，如果请求中未包含某些字段，那么这些字段会被置为其类型的零值，从而方便地处理了零值和空置的情况。这种方式更加简洁、直观，并且符合 Go 语言的最佳实践。

### 这个事例的请求输入输出的结果

假设使用以下请求体作为 PUT 请求的输入：

```
{
	"my_string": "",
	"my_int": 0
}
```

则上述示例代码的输出结果如下：

```
MyString:
MyInt: 0
```

可以看到，请求体中的空字符串 `""` 会被解析为空字符串 `""`，而请求体中的零值 `0` 会被解析为整型字段的零值 `0`。这是因为在结构体字段中使用了 `omitempty` 标签，使得零值和空置字段在序列化时被忽略，因此在处理 PUT 请求时，字段的零值和空置会被正确地反映到结构体中，从而可以根据业务逻辑进行相应的处理。

### 没传 my_string 字段的情况，输入输出

如果 PUT 请求的输入没有包含 `my_string` 字段，那么示例代码的输出结果如下：

```
MyString:
MyInt: 0
```

可以看到，`my_string` 字段被解析为空字符串 `""`，因为在请求中并没有提供该字段的值。而 `my_int` 字段被解析为整型字段的零值 `0`，因为在请求中未指定该字段的值。这是因为在结构体字段中使用了 `omitempty` 标签，使得零值和空置字段在序列化时被忽略，从而在处理 PUT 请求时，字段的零值和空置会被正确地反映到结构体中，方便进行后续的业务逻辑处理。

### 这样就区分不了 my_string 是不是真正的空字符串

您是对的，如果请求中传递了空字符串 `""`，而不是完全省略了 `my_string` 字段，那么在使用 `omitempty` 标签的情况下，无法区分字段的零值和空字符串。

如果需要明确区分字段的零值和空字符串，在处理 PUT 请求时，可以通过在结构体中使用指针类型来实现，示例如下：

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyData struct {
	MyString *string `json:"my_string,omitempty"`
	MyInt    int     `json:"my_int,omitempty"`
}

func main() {
	r := gin.Default()

	r.PUT("/data", UpdateData)

	r.Run(":8080")
}

func UpdateData(c *gin.Context) {
	var data MyData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.MyString != nil {
		fmt.Println("MyString:", *data.MyString)
	} else {
		fmt.Println("MyString: nil")
	}
	fmt.Println("MyInt:", data.MyInt)

	// 处理其他情况
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}
```

在这个示例中，`MyString` 字段的类型更改为 `*string`，即指针类型，这样当请求中包含空字符串 `""` 时，`MyString` 字段的值将被设置为指向一个空字符串的指针，而不是直接设置为空字符串 `""`。这样在处理 PUT 请求时，就可以通过检查指针是否为 `nil` 来区分字段的零值和空字符串。需要注意的是，在使用指针类型时，需要额外处理指针的空值和解引用操作，以避免空指针引发的错误。

## 结构体

1. 嵌套
2. 继承

<a href="https://juejin.cn/post/6969574358142418975#heading-1" target="_blank" >见</a>

### golang 结构体继承

在 Go 中，结构体（struct）是一种复合数据类型，用于存储一组相关的字段。但是，Go 不支持传统的面向对象编程中的类继承。Go 语言使用组合（Composition）而不是继承（Inheritance）来实现代码的复用。

组合（Composition）是指在一个结构体中嵌套其他结构体，从而实现结构体之间的组合关系。被嵌套的结构体可以获得嵌套结构体的字段和方法，从而实现代码的复用。

下面是一个简单的示例，展示了在 Go 中如何使用结构体组合来实现类似继承的效果：

```go
// 父结构体
type Person struct {
	Name string
	Age  int
}

// 子结构体，嵌套了父结构体
type Employee struct {
	Person   // 嵌套 Person 结构体
	JobTitle string
	Salary   float64
}

func main() {
	// 创建一个 Employee 实例
	emp := Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		JobTitle: "Software Engineer",
		Salary:   5000.0,
	}

	// 访问 Employee 的字段
	fmt.Println("Name:", emp.Name)
	fmt.Println("Age:", emp.Age)
	fmt.Println("Job Title:", emp.JobTitle)
	fmt.Println("Salary:", emp.Salary)
}
```

在上面的示例中，`Person` 结构体被嵌套到 `Employee` 结构体中，从而实现了类似继承的效果。`Employee` 结构体可以访问 `Person` 结构体的字段 `Name` 和 `Age`，并且可以定义自己的字段和方法。

需要注意的是，在使用结构体组合时，如果嵌套的结构体中存在字段或方法与外层结构体中的字段或方法同名，那么在访问这个字段或方法时，需要使用完整的字段或方法名称，包括结构体名称。例如，在上面的示例中，如果 `Employee` 结构体中也定义了 `Name` 字段，那么在访问 `Person` 结构体的 `Name` 字段时，应该使用 `emp.Person.Name`。

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
3. 你一定会发现，变量 Class 貌似没有解析为我们期待样子。
   1. 因为此时的 Class 是个 interface{}类型的变量，而 json 串中 key 为 CLASS 的 value 是个复合结构，不是可以直接解析的简单类型数据（如“张三”，18，true 等）。
   2. 所以解析时，由于没有指定变量 Class 的具体类型，json 自动将 value 为复合结构的数据解析为 `map[string]interface{}`类型的项。
   3. 也就是说，此时的 struct Class 对象与 StuRead 中的 Class 变量没有半毛钱关系，故与这次的 json 解析没有半毛钱关系。

### golang json API

在 Go 语言中，处理 JSON 数据有丰富的标准库（`encoding/json`）提供的 API。以下是一些常用的 JSON 相关的 API：

1. `json.Marshal(v interface{}) ([]byte, error)`: 用于将 Go 数据结构转换为 JSON 字节切片。`v` 参数是要转换的 Go 数据结构，返回值是 JSON 格式的字节切片和可能出现的错误。

2. `json.Unmarshal(data []byte, v interface{}) error`: 用于将 JSON 字节切片解析为 Go 数据结构。`data` 参数是要解析的 JSON 字节切片，`v` 参数是目标 Go 数据结构的指针，用于存储解析结果。返回值是可能出现的错误。

3. `json.NewEncoder(w io.Writer) *json.Encoder`: 提供一个用于将 Go 数据结构编码为 JSON 格式并写入 `io.Writer` 接口的编码器。`w` 参数是用于写入 JSON 数据的 `io.Writer` 实例。

4. `json.NewDecoder(r io.Reader) *json.Decoder`: 提供一个用于从 `io.Reader` 接口读取 JSON 数据并解码为 Go 数据结构的解码器。`r` 参数是用于读取 JSON 数据的 `io.Reader` 实例。

5. `json.UnmarshalJSON(data []byte, v interface{}) error`: 类似于 `json.Unmarshal`，但是可以用于解析包含嵌套 JSON 数据的字段值。

6. `json.MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)`: 类似于 `json.Marshal`，但是可以为生成的 JSON 数据添加前缀和缩进，以便更好地进行格式化和阅读。

7. `json.Valid(data []byte) bool`: 用于检查 JSON 字节切片是否是有效的 JSON 格式。

这些只是 `encoding/json` 包提供的一些常用 API，还有其他一些 API 用于处理 JSON 数据的细节操作，例如处理 JSON 数字精度、自定义 JSON 编码/解码方法等。在 Go 中，`encoding/json` 是处理 JSON 数据的标准库，通常情况下足够满足大多数 JSON 数据处理的需求。

### 每个 API 的例子

当然，下面是每个 API 的简单示例：

### 1 `json.Marshal(v interface{}) ([]byte, error)`:

```go
type Person struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Gender string `json:"gender"`
}

person := Person{
    Name:   "Alice",
    Age:    30,
    Gender: "Female",
}

jsonData, err := json.Marshal(person)
if err != nil {
    fmt.Println("JSON Marshal error:", err)
    return
}

fmt.Println(string(jsonData))
// 输出：{"name":"Alice","age":30,"gender":"Female"}
```

### 2 `json.Unmarshal(data []byte, v interface{}) error`:

```go
type Person struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Gender string `json:"gender"`
}

jsonStr := `{"name":"Alice","age":30,"gender":"Female"}`
jsonData := []byte(jsonStr)

var person Person
err := json.Unmarshal(jsonData, &person)
if err != nil {
    fmt.Println("JSON Unmarshal error:", err)
    return
}

fmt.Println(person.Name)
fmt.Println(person.Age)
fmt.Println(person.Gender)
// 输出：
// Alice
// 30
// Female
```

### 3 `json.NewEncoder(w io.Writer) *json.Encoder`:

```go
type Person struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Gender string `json:"gender"`
}

person := Person{
    Name:   "Alice",
    Age:    30,
    Gender: "Female",
}

file, err := os.Create("person.json")
if err != nil {
    fmt.Println("Failed to create file:", err)
    return
}
defer file.Close()

jsonEncoder := json.NewEncoder(file)
err = jsonEncoder.Encode(person)
if err != nil {
    fmt.Println("JSON Encode error:", err)
    return
}

// person.json 文件中将包含编码后的 JSON 数据：
// {"name":"Alice","age":30,"gender":"Female"}
```

### 4 `json.NewDecoder(r io.Reader) *json.Decoder`:

```go
type Person struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Gender string `json:"gender"`
}

file, err := os.Open("person.json")
if err != nil {
    fmt.Println("Failed to open file:", err)
    return
}
defer file.Close()

var person Person
jsonDecoder := json.NewDecoder(file)
err = jsonDecoder.Decode(&person)
if err != nil {
    fmt.Println("JSON Decode error:", err)
    return
}

fmt.Println(person.Name)
fmt.Println(person.Age)
fmt.Println(person.Gender)
// 输出：
// Alice
// 30
// Female
```

### 5 `json.UnmarshalJSON(data []byte, v interface{}) error` (续):

```go
type Person struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Gender string `json:"gender"`
}

type Data struct {
    RawJSON json.RawMessage `json:"raw_json"`
}

jsonStr := `{"raw_json": {"name":"Alice","age":30,"gender":"Female"}}`
jsonData := []byte(jsonStr)

var data Data
err := json.Unmarshal(jsonData, &data)
if err != nil {
    fmt.Println("JSON Unmarshal error:", err)
    return
}

var person Person
err = person.UnmarshalJSON(data.RawJSON)
if err != nil {
    fmt.Println("JSON Unmarshal error:", err)
    return
}

fmt.Println(person.Name)
fmt.Println(person.Age)
fmt.Println(person.Gender)
// 输出：
// Alice
// 30
// Female
```

### 6 `json.Compact(dst *bytes.Buffer, src []byte) error`:

```go
jsonStr := `{
    "name": "Alice",
    "age": 30,
    "gender": "Female"
}`

jsonData := []byte(jsonStr)

var pretty bytes.Buffer
err := json.Indent(&pretty, jsonData, "", "\t")
if err != nil {
    fmt.Println("JSON Indent error:", err)
    return
}

fmt.Println(pretty.String())
// 输出：
// {
// 	"name": "Alice",
// 	"age": 30,
// 	"gender": "Female"
// }

var compact bytes.Buffer
err = json.Compact(&compact, pretty.Bytes())
if err != nil {
    fmt.Println("JSON Compact error:", err)
    return
}

fmt.Println(compact.String())
// 输出：{"name":"Alice","age":30,"gender":"Female"}
```

### 7 `json.HTMLEscape(dst *bytes.Buffer, src []byte)`:

```go
jsonStr := `{"name":"<Alice>","age":30,"gender":"Female"}`
jsonData := []byte(jsonStr)

var escaped bytes.Buffer
json.HTMLEscape(&escaped, jsonData)

fmt.Println(escaped.String())
// 输出：{"name":"\u003cAlice\u003e","age":30,"gender":"Female"}
```

在 Go 语言中，`json`包中的`HTMLEscape`函数用于将特殊字符转义为 HTML 实体，以便在生成 JSON 时确保输出的字符串是安全的，不会引发 HTML 注入攻击。

下面是`json`包中`HTMLEscape`函数的基本使用示例：

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := map[string]string{
		"name":     "John",
		"age":      "25",
		"location": "<script>alert('XSS')</script>",
	}

	// 将数据编码为JSON
	encoded, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON encoding failed:", err)
		os.Exit(1)
	}

	// 输出原始的JSON字符串
	fmt.Println("Original JSON:", string(encoded))

	// 使用HTMLEscape函数转义特殊字符
	escaped := json.HTMLEscape(nil, encoded)

	// 输出转义后的JSON字符串
	fmt.Println("Escaped JSON:", string(escaped))
}
```

在以上示例中，我们定义了一个包含一些特殊字符的 JSON 数据对象，并使用`json.Marshal`函数将其编码为 JSON 格式的字节数组。然后，我们通过调用`json.HTMLEscape`函数将这些特殊字符转义为 HTML 实体，返回一个新的字节数组。最后，我们将转义后的 JSON 字符串输出到控制台。

需要注意的是，`json.HTMLEscape`函数在原地修改了传入的字节数组，并返回修改后的字节数组，因此在使用时应谨慎处理原始数据。此外，`json.HTMLEscape`函数只会转义特定的字符，例如`<`, `>`, `&`, `'`, `"`等，而不会对所有字符进行转义。因此，在处理用户输入或其他潜在的不受信任的数据时，应该使用其他方法来保护 JSON 数据的安全性，例如使用输入验证和输出编码等措施。

### 8 `json.Valid(data []byte) bool`:

```go
jsonStr := `{"name":"Alice","age":30,"gender":"Female"}`
jsonData := []byte(jsonStr)

isValid := json.Valid(jsonData)

fmt.Println(isValid)
// 输出：true

invalidJSON := []byte(`{"name":"Alice","age":30,"gender"}`)
isValid = json.Valid(invalidJSON)

fmt.Println(isValid)
// 输出：false
```

这些就是一些常见的 Go JSON API 的简单示例。

## 内存模型

<a href="./go程/go模型.md" target="_blank" >Go 内存模型</a>

###

Golang（或称为 Go）是一种编程语言，其内存模型描述了程序在运行时如何管理和访问内存。理解 Golang 的内存模型对于编写高效、并发和安全的 Go 程序非常重要。

Golang 的内存模型基于以下几个主要概念：

1. 堆（Heap）：堆是一块大的内存区域，用于分配和存储程序运行时动态分配的内存。在 Golang 中，通过 `new`、`make` 等关键字进行堆上的内存分配。

2. 栈（Stack）：栈是一块较小的内存区域，用于存储程序的局部变量和函数调用的上下文信息。在 Golang 中，每个 Goroutine 都有自己的栈空间。

3. Goroutine：Goroutine 是 Go 中的轻量级线程，用于实现并发执行。Goroutine 的创建和销毁由 Go 运行时管理，每个 Goroutine 都可以在一个独立的栈上运行。

4. 通道（Channel）：通道是 Goroutine 之间进行通信和同步的一种机制。通道可以用来发送和接收数据，并且在读写时会自动进行同步，从而避免了显式的锁操作。

5. 内存模型：Golang 的内存模型规定了 Goroutine 之间的内存访问和同步行为。Golang 使用了基于 happens-before 原则的内存模型，通过同步操作（例如使用通道、互斥锁、原子操作等）来确保 Goroutine 之间的内存一致性。

在 Golang 的内存模型中，Goroutine 之间的通信和同步是通过通道进行的，而不是显式的锁操作。这使得编写并发程序更加简单和安全，避免了传统多线程编程中的许多常见问题，例如竞态条件、死锁等。

总的来说，理解 Golang 的内存模型对于编写高效、并发和安全的 Go 程序非常重要，需要熟悉堆、栈、Goroutine、通道以及 Golang 的内存模型规定的同步行为。

## 并发与一致

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

## 加密（skip）

## token/cookie

<a href="https://learnku.com/articles/71845" target="_blank" >见</a>

jwt 库很多了 各有各的优势 有些库是不维护了

选择了 `github.com/golang-jwt/jwt` 库

获取命令：`go get -u github.com/golang-jwt/jwt/v4`

### Header

header 典型的由两部分组成：token 的类型（“JWT”）和算法名称（比如：HMAC SHA256 或者 RSA 等等

```go
{
  'typ': 'JWT',
  'alg': 'HS256'
}
```

### Payload

载荷就是存放有效信息的地方。这个名字像是特指飞机上承载的货品，这些有效信息包含三个部分

标准中注册的声明
公共的声明
私有的声明

jwt.StandardClaims 标准中注册的声明 (建议但不强制使用) ：

1. iss: jwt 签发者
2. sub: jwt 所面向的用户
3. aud: 接收 jwt 的一方
4. exp: jwt 的过期时间，这个过期时间必须要大于签发时间
5. nbf: 定义在什么时间之前，该 jwt 都是不可用的.
6. iat: jwt 的签发时间
7. jti: jwt 的唯一身份标识，主要用来作为一次性 token, 从而回避重放攻击。

私有的声明 ：
私有声明是提供者和消费者所共同定义的声明，一般不建议存放敏感信息，因为 base64 是对称解密的，意味着该部分信息可以归类为明文信息

私有定义的内容根据自己业务需要来，这里简单加了 UID

```go
type AuthClaim struct {
    UID int64 `json:"uid"`
    jwt.StandardClaims
}
```

### Signature 签名

secret 是保存在服务器端的，jwt 的签发生成也是在服务器端的，secret 就是用来进行 jwt 的签发和 jwt 的验证，所以，它就是你服务端的私钥，在任何场景都不应该流露出去。一旦客户端得知这个 secret, 那就意味着客户端是可以自我签发 jwt 了。

```go
var Secret = "私钥"
var hmacSampleSecret = []byte(Secret)
```

### 生成 token

生成了两个小时过期时间的 token

```go
const TokenExpireDuration = 2 * time.Hour //过期时间

func New(uid int64) (tokenStr string) {
    var authClaim AuthClaim
    authClaim.UID = uid
    authClaim.StandardClaims.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaim)
    tokenString, _ := token.SignedString(hmacSampleSecret) //私钥加密
    return tokenString
}
```

### 解析 token

```go
func Parse(tokenString string) (auth AuthClaim, Valid bool) {
    token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Don't forget to validate the alg is what you expect:
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
        return hmacSampleSecret, nil
    })
    Valid = token.Valid//token是否有效 true有效  false无效
    if claims, ok := token.Claims.(jwt.MapClaims); ok && Valid {
        auth.UID = int64(claims["uid"].(float64)) //自定义的UID
        auth.ExpiresAt = int64(claims["exp"].(float64)) //过期时间
    }
    return
}
```

## 跨域问题

<a href="https://juejin.cn/post/6871583587062415367" target="_blank" >见</a>

## gin

1. 中间件开发 https://juejin.cn/post/7034338727883177997#heading-1
2. 插件开发
3. 数据绑定
4. 返回数据
5. 封装返回
6. 异常统一处理

### gin 中间件

Gin 是一个用于构建 Web 应用程序的 Go 语言框架，它提供了一些用于处理 HTTP 请求和响应的功能。Gin 中间件是一种特殊类型的功能，它可以在请求到达路由处理器之前或之后执行一些操作。

Gin 中间件通常用于以下几个目的：

1. 路由处理前的预处理：可以在路由处理器执行之前进行一些操作，例如身份验证、权限验证、请求参数解析等。这样可以在进入路由处理器之前对请求进行预处理，从而减轻路由处理器的负担。

2. 路由处理后的后处理：可以在路由处理器执行完毕之后进行一些操作，例如日志记录、响应结果封装等。这样可以在路由处理器执行完毕后对响应进行后处理，从而实现一些共享的逻辑。

3. 错误处理：可以在发生错误时进行处理，例如错误日志记录、错误信息返回等。这样可以统一处理应用程序中可能发生的错误，避免在每个路由处理器中都进行错误处理逻辑的重复。

Gin 中间件通过注册到 Gin 的引擎（Engine）上，并按照注册的顺序依次执行。每个中间件可以通过 Gin 的上下文（Context）对象来访问请求和响应的信息，并在需要时对其进行修改。

以下是一个使用 Gin 中间件的示例：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 自定义中间件示例
	r.Use(func(c *gin.Context) {
		// 在请求处理前执行的操作
		fmt.Println("执行中间件操作前")

		// 执行下一个中间件或路由处理器
		c.Next()

		// 在请求处理后执行的操作
		fmt.Println("执行中间件操作后")
	})

	// 路由处理器示例
	r.GET("/", func(c *gin.Context) {
		// 在路由处理器中的操作
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	r.Run(":8080")
}
```

在上面的示例中，通过 `r.Use()` 注册了一个自定义的中间件，它会在每次请求到达时先执行预处理操作，然后调用 `c.Next()` 继续执行下一个中间件或路由处理器，最后在请求处理完毕后执行后处理操作。这样可以在请求处理前后执行一些共享的操作。注意，中间件的注册顺序很重要，它们会按照注册的顺序依次执行。

Gin 中间件是一种强大的功能，可以在处理 HTTP 请求和响应时灵活地进行预处理和后处理操作，从而实现一些共享的逻辑和功能。

### 返回数据 统一处理返回 事例 1

在 Gin 中，可以使用中间件来实现对响应数据的统一处理和封装，从而实现一些共享的逻辑，例如格式化响应数据、添加公共响应头、处理错误等。以下是一个简单的示例：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := gin.Default()

	// 自定义中间件示例
	r.Use(func(c *gin.Context) {
		// 在请求处理前执行的操作
		// ...

		// 执行下一个中间件或路由处理器
		c.Next()

		// 在请求处理后执行的操作
		// ...

		// 统一处理响应数据
		code := c.Writer.Status()
		message := ""
		data := c.Keys["response_data"]
		if err, ok := data.(error); ok {
			// 处理错误信息
			code = 500 // 错误状态码
			message = err.Error() // 错误消息
			data = nil // 清空数据
		}
		response := Response{
			Code:    code,
			Message: message,
			Data:    data,
		}
		c.JSON(response.Code, response)
	})

	// 路由处理器示例
	r.GET("/", func(c *gin.Context) {
		// 在路由处理器中的操作
		// ...

		// 设置响应数据到上下文中
		c.Set("response_data", gin.H{"message": "Hello, Gin!"})

		// 抛出一个错误
		// c.AbortWithError(500, fmt.Errorf("发生了一个错误"))

		c.Next()
	})

	r.Run(":8080")
}
```

在上面的示例中，自定义的中间件会在每次请求到达时先执行预处理操作，然后调用 `c.Next()` 继续执行下一个中间件或路由处理器，最后在请求处理完毕后执行后处理操作。在后处理操作中，根据需要可以统一处理响应数据，例如格式化响应数据、添加公共响应头等。在路由处理器中，可以通过 `c.Set()` 方法将响应数据存储到 Gin 的上下文中，供后续的中间件和路由处理器使用。

需要注意的是，如果发生错误，可以通过 `c.AbortWithError()` 方法终止请求处理，并传递错误信息到后续的中间件和路由处理器，从而实现统一的错误处理。在后处理操作中，可以根据错误信息进行相应的处理，例如设置合适的错误状态码和错误消息。

这样，通过使用 Gin 中间件来统一处理和封装响应数据，可以实现对响应的格式化、公共头部的添加、错误处理等逻辑的统一管理，从而提高代码的可维护性和可扩展性。

### 返回数据 统一处理返回 事例 2

好的，这里是一个更详细的示例，演示了如何使用 Gin 中间件来统一处理和封装响应数据。

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	r := gin.Default()

	// 自定义中间件示例
	r.Use(func(c *gin.Context) {
		// 在请求处理前执行的操作
		fmt.Println("请求处理前")

		// 执行下一个中间件或路由处理器
		c.Next()

		// 在请求处理后执行的操作
		fmt.Println("请求处理后")

		// 统一处理响应数据
		code := c.Writer.Status()
		message := ""
		data := c.Keys["response_data"]
		if err, ok := data.(error); ok {
			// 处理错误信息
			code = 500 // 错误状态码
			message = err.Error() // 错误消息
			data = nil // 清空数据
		}
		response := Response{
			Code:    code,
			Message: message,
			Data:    data,
		}
		c.JSON(response.Code, response)
	})

	// 路由处理器示例
	r.GET("/", func(c *gin.Context) {
		// 在路由处理器中的操作
		fmt.Println("处理请求")

		// 设置响应数据到上下文中
		c.Set("response_data", gin.H{"message": "Hello, Gin!"})

		// 抛出一个错误
		// c.AbortWithError(500, fmt.Errorf("发生了一个错误"))

		c.Next()
	})

	r.Run(":8080")
}
```

在这个示例中，我们定义了一个自定义的中间件，它会在每次请求到达时先执行预处理操作，然后调用 `c.Next()` 继续执行下一个中间件或路由处理器，最后在请求处理完毕后执行后处理操作。在后处理操作中，我们根据需要可以统一处理响应数据，例如格式化响应数据、添加公共响应头等。

在路由处理器中，我们通过 `c.Set()` 方法将响应数据存储到 Gin 的上下文中，供后续的中间件和路由处理器使用。在这里，我们简单地设置了一个包含 "message" 字段的 map 作为响应数据。

需要注意的是，如果发生错误，我们可以通过 `c.AbortWithError()` 方法终止请求处理，并传递错误信息到后续的中间件和路由处理器，从而实现统一的错误处理。在后处理操作中，我们可以根据错误信息进行相应的处理，例如设置合适的错误状态码和错误消息。

通过这样的方式，我们可以在 Gin 中使用中间件来统一处理和封装响应数据，实现对响应的格式化、公共头部的添加、错误处理等逻辑的统一管理，从而提高代码的可维护性和可扩展性。

### token 校验中间件

在 Gin 框架中，可以通过自定义中间件来实现 Token 校验功能，用于验证客户端请求中的 Token 是否合法。下面是一个简单的示例：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 自定义 Token 校验中间件示例
	r.Use(func(c *gin.Context) {
		// 获取请求中的 Token
		token := c.GetHeader("Authorization")

		// 进行 Token 校验逻辑
		if token == "" || token != "valid_token" {
			// 如果 Token 为空或者不合法，返回错误响应
			c.JSON(401, gin.H{"error": "Invalid Token"})
			c.Abort() // 终止请求处理
			return
		}

		// Token 校验通过，继续执行下一个中间件或路由处理器
		c.Next()
	})

	// 路由处理器示例
	r.GET("/", func(c *gin.Context) {
		// 在路由处理器中的操作
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	r.Run(":8080")
}
```

在这个示例中，我们定义了一个自定义的中间件来进行 Token 校验。在中间件中，我们通过 `c.GetHeader()` 方法获取请求头中的 "Authorization" 字段，该字段用于携带 Token。然后，我们可以进行自定义的 Token 校验逻辑，例如判断 Token 是否为空、是否合法等。

如果 Token 校验失败，我们通过 `c.JSON()` 方法返回错误的 JSON 响应，并通过 `c.Abort()` 方法终止请求处理，从而避免继续执行后续的中间件或路由处理器。

如果 Token 校验通过，我们可以调用 `c.Next()` 方法继续执行下一个中间件或路由处理器。

通过这种方式，我们可以在 Gin 中使用中间件来实现 Token 校验功能，从而实现对客户端请求中的 Token 进行统一的验证和处理，提高应用的安全性。需要根据实际需求进行相应的定制和扩展，例如可以将 Token 存储在数据库或缓存中，并在中间件中进行查询和验证。

### 异常统一处理中间件

在 Gin 框架中，可以通过自定义中间件来实现异常统一处理功能，用于捕获和处理路由处理器中的异常，从而对异常进行统一的处理和返回错误响应。下面是一个简单的示例：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 自定义异常统一处理中间件示例
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 捕获路由处理器中的异常

				// 可以根据实际需求进行异常处理，例如记录日志、返回错误响应等
				fmt.Println("Recovered error:", err)

				// 返回错误响应
				c.JSON(500, gin.H{"error": "Internal Server Error"})
				c.Abort() // 终止请求处理
			}
		}()

		// 继续执行下一个中间件或路由处理器
		c.Next()
	})

	// 路由处理器示例
	r.GET("/", func(c *gin.Context) {
		// 在路由处理器中可能出现的异常
		panic("Some error occurred")

		// 路由处理器中的其他操作
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	r.Run(":8080")
}
```

在这个示例中，我们定义了一个自定义的中间件来进行异常统一处理。在中间件中，我们使用 `recover()` 函数捕获了路由处理器中的异常，并在 `defer` 语句中进行了处理。在处理中，我们可以根据实际需求进行异常处理，例如记录日志、返回错误响应等。

如果发生异常，我们通过 `c.JSON()` 方法返回错误的 JSON 响应，并通过 `c.Abort()` 方法终止请求处理，从而避免继续执行后续的中间件或路由处理器。

通过这种方式，我们可以在 Gin 中使用中间件来实现异常统一处理功能，从而对路由处理器中的异常进行捕获和处理，提高应用的可靠性和稳定性。需要根据实际需求进行相应的定制和扩展，例如可以根据不同类型的异常返回不同的错误响应或执行不同的处理逻辑。

### 单个中间件管理

在 Gin 框架中，可以将中间件单独放置在项目的某个目录中，以便进行组织和管理。下面是一个简单的示例：

1. 创建一个名为 `middlewares` 的目录，用于存放中间件文件。可以将这个目录放置在项目的任意位置，例如项目的根目录下。

2. 在 `middlewares` 目录中创建一个中间件文件，例如 `error_handler.go`，用于实现异常统一处理中间件。

```go
package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// ErrorHandler 中间件用于捕获和处理路由处理器中的异常
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 捕获路由处理器中的异常

				// 可以根据实际需求进行异常处理，例如记录日志、返回错误响应等
				fmt.Println("Recovered error:", err)

				// 返回错误响应
				c.JSON(500, gin.H{"error": "Internal Server Error"})
				c.Abort() // 终止请求处理
			}
		}()

		// 继续执行下一个中间件或路由处理器
		c.Next()
	}
}
```

3. 在项目的主文件（例如 `main.go`）中引入 `middlewares` 包，并使用其中的中间件。

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yourproject/middlewares" // 引入自定义中间件包
)

func main() {
	r := gin.Default()

	// 使用自定义异常统一处理中间件
	r.Use(middlewares.ErrorHandler())

	// 路由处理器示例
	r.GET("/", func(c *gin.Context) {
		// 在路由处理器中可能出现的异常
		panic("Some error occurred")

		// 路由处理器中的其他操作
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	r.Run(":8080")
}
```

通过将中间件单独放置在项目的某个目录中，并在需要使用中间件的地方引入和使用，可以更好地组织和管理中间件，使代码更加清晰和可维护。需要根据实际项目结构和需求进行相应的调整和扩展。

### 多个中间件管理

在 Gin 框架中，可以通过以下步骤来管理三个中间件：

1. 创建中间件文件：在项目的某个目录下创建三个中间件文件，例如 `middleware1.go`、`middleware2.go` 和 `middleware3.go`，分别实现三个中间件的逻辑。

2. 编写中间件逻辑：在每个中间件文件中，实现对应的中间件逻辑。例如，可以使用 `gin.HandlerFunc` 函数签名来定义中间件函数，并在其中编写处理逻辑，例如校验 token、异常处理、日志记录等。

```go
// middleware1.go
package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 中间件1的逻辑处理
		// ...
		c.Next() // 继续执行下一个中间件或路由处理器
	}
}

// middleware2.go
package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 中间件2的逻辑处理
		// ...
		c.Next() // 继续执行下一个中间件或路由处理器
	}
}

// middleware3.go
package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Middleware3() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 中间件3的逻辑处理
		// ...
		c.Next() // 继续执行下一个中间件或路由处理器
	}
}
```

3. 在项目中引入和使用中间件：在项目的主文件（例如 `main.go`）或路由文件中，通过引入中间件文件，并调用其中的中间件函数，将中间件添加到 Gin 引擎中的中间件链中。

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourproject/middlewares" // 引入自定义中间件包
)

func main() {
	r := gin.Default()

	// 添加三个自定义中间件到 Gin 引擎中
	r.Use(middlewares.Middleware1())
	r.Use(middlewares.Middleware2())
	r.Use(middlewares.Middleware3())

	// 路由处理器示例
	r.GET("/", func(c *gin.Context) {
		// 路由处理器的逻辑
		// ...
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	r.Run(":8080")
}
```

通过这种方式，可以方便地管理多个中间件，将它们按照需要添加到 Gin 引擎中的中间件链中，实现各自的逻辑处理。需要根据实际项目需求进行中间件的调用顺序、添加位置等设置。

### 跨域问题处理中间件

在 Gin 框架中，解决跨域问题可以通过使用中间件来处理。以下是一个简单的跨域中间件示例：

```go
package middlewares

import (
	"github.com/gin-gonic/gin"
)

// CorsMiddleware 是处理跨域请求的中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的域名，* 表示允许任何域名访问，也可以设置具体的域名
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		// 设置允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// 允许携带跨域请求中的 cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 如果是 OPTIONS 请求，直接返回 200，表示允许该请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// 继续执行下一个中间件或路由处理器
		c.Next()
	}
}
```

在项目中使用这个跨域中间件：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourproject/middlewares" // 引入自定义中间件包
)

func main() {
	r := gin.Default()

	// 使用自定义的跨域中间件
	r.Use(middlewares.CorsMiddleware())

	// 路由处理器示例
	r.GET("/", func(c *gin.Context) {
		// 路由处理器的逻辑
		// ...
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	r.Run(":8080")
}
```

通过将跨域中间件添加到 Gin 引擎中的中间件链中，可以处理来自不同域名的跨域请求，并设置允许的请求方法、请求头等相关参数，解决跨域问题。需要根据实际项目需求进行相关设置，例如设置允许的域名、请求方法、请求头等。

### 分组路由跨域问题处理中间件

在 Gin 框架中，可以通过为不同的路由分组使用中间件来处理跨域问题。以下是一个使用分组路由处理跨域问题的示例：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourproject/middlewares" // 引入自定义中间件包
)

func main() {
	r := gin.Default()

	// 定义路由组
	api := r.Group("/api")
	{
		// 使用自定义的跨域中间件处理 /api 下的所有路由跨域问题
		api.Use(middlewares.CorsMiddleware())

		// 路由处理器示例
		api.GET("/users", func(c *gin.Context) {
			// 路由处理器的逻辑
			// ...
			c.JSON(200, gin.H{"message": "Hello, users!"})
		})

		api.GET("/posts", func(c *gin.Context) {
			// 路由处理器的逻辑
			// ...
			c.JSON(200, gin.H{"message": "Hello, posts!"})
		})

		// 可以在这里添加更多的路由处理器
	}

	r.Run(":8080")
}
```

在这个示例中，我们创建了一个名为 `/api` 的路由组，并在该路由组中使用了自定义的跨域中间件 `middlewares.CorsMiddleware()`。这样，该路由组下的所有路由都会经过跨域中间件的处理，从而解决了跨域问题。

需要注意的是，在使用分组路由处理跨域问题时，中间件的顺序也很重要。例如，如果在路由组中先添加了其他中间件，再添加跨域中间件，可能会导致跨域中间件无法生效。因此，建议将跨域中间件放在路由组的最前面，以确保跨域中间件先于其他中间件执行。

### 还能做什么

中间件是一种在处理 HTTP 请求和响应时插入逻辑处理的机制，可以在 Gin 框架中用于实现各种功能。除了处理跨域问题和异常统一处理之外，中间件还可以用于以下几种情况：

1. 认证和授权：中间件可以用于验证请求的身份信息，例如基于 token 的认证，JWT 鉴权等。可以在中间件中进行身份验证的逻辑，如果验证失败则可以返回错误响应，否则继续处理请求。

2. 日志记录：中间件可以用于记录请求和响应的日志，包括请求方法、路径、参数、响应状态码等信息，用于监控和排查问题。

3. 参数校验和请求处理：中间件可以用于对请求参数进行校验，例如检查请求参数是否符合要求、是否为空等，可以在中间件中进行参数校验，如果不符合要求则返回错误响应，从而减轻路由处理器的负担。

4. 数据处理和转换：中间件可以用于对请求和响应的数据进行处理和转换，例如对请求数据进行解析、对响应数据进行格式化、压缩等操作。

5. 缓存和性能优化：中间件可以用于实现缓存策略，例如在中间件中判断请求是否命中缓存，如果命中则直接返回缓存数据，从而提高响应速度和性能。

6. 错误处理和异常处理：中间件可以用于捕获和处理应用程序中的错误和异常，例如处理请求过程中发生的错误，并返回相应的错误响应，用于统一的错误处理。

需要注意的是，中间件应该按照业务逻辑的不同进行划分和组织，避免中间件过于复杂和耦合，保持中间件的单一职责和可维护性。在使用中间件时，也应该注意中间件的执行顺序，确保它们按照预期的顺序执行。

### 日志中间件

日志中间件是一种用于记录请求和响应的日志信息的中间件，在 Gin 框架中可以用于记录请求的方法、路径、参数、响应状态码等信息，用于监控和排查问题。以下是一个简单的日志中间件的示例：

```go
package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 记录请求结束时间和耗时
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		// 获取请求信息
		requestMethod := c.Request.Method
		requestURI := c.Request.RequestURI
		statusCode := c.Writer.Status()

		// 记录日志
		fmt.Printf("[GIN] %v | %3d | %12v | %s | %s\n",
			endTime.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latencyTime,
			requestMethod,
			requestURI,
		)
	}
}
```

在上面的示例中，`LoggerMiddleware` 是一个日志中间件，它通过在请求处理前记录请求开始时间，然后在请求处理后记录请求结束时间和耗时，再根据请求方法、请求路径、响应状态码等信息输出日志。在使用该中间件时，只需要在路由组或路由处理器中使用 `Use` 方法添加该中间件即可，例如：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/your-package/middleware"
)

func main() {
	r := gin.Default()

	// 添加日志中间件
	r.Use(middleware.LoggerMiddleware())

	// 添加路由
	// ...

	r.Run(":8080")
}
```

这样，每次请求处理时都会通过该日志中间件记录请求和响应的日志信息，便于监控和排查问题。当然，你还可以根据需要对日志中间件进行定制化，例如输出的日志格式、日志存储方式等，以满足项目的具体需求。

### .next()作用

在 Gin 框架中，`.Next()` 是一个用于在中间件中传递控制权到下一个中间件或处理器函数的方法。当一个请求经过多个中间件时，每个中间件都可以执行一些操作，然后通过调用 `.Next()` 方法将请求传递给下一个中间件或处理器函数处理。

使用 `.Next()` 方法的主要作用是将控制权传递给链中的下一个中间件或处理器函数，允许后续的中间件或处理器函数继续对请求进行处理。如果不调用 `.Next()` 方法，请求将被中间件截断，不会传递给后续的中间件或处理器函数，从而导致请求无法继续处理。

在 Gin 中，中间件是按照添加的顺序执行的，通过 `.Next()` 方法的调用，Gin 会将请求传递给链中的下一个中间件或处理器函数，直到最后一个中间件或处理器函数处理完请求或调用了 `.Next()` 方法。这样，每个中间件都可以根据需要进行处理，例如校验、记录日志、修改请求或响应等，然后将请求传递给下一个中间件或处理器函数继续处理。

### 认证和授权

在 Web 应用程序中，认证和授权是两个关键的安全概念，用于确保用户的身份验证和控制用户对资源的访问权限。

认证（Authentication）是指确认用户的身份是否合法，通常通过用户名和密码进行验证。认证主要用于验证用户是谁，以便在应用程序中标识用户的身份。

授权（Authorization）是指确定用户是否有权限访问特定资源或执行特定操作。授权主要用于控制用户对应用程序中的资源的访问权限，例如访问某个页面、执行某个操作或访问某个 API 等。

在 Gin 框架中，可以使用中间件来实现认证和授权的功能。以下是一个简单的示例，展示如何使用中间件进行基于 JWT（JSON Web Token）的认证和授权：

```go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your-package/auth"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 Authorization
		authorizationHeader := c.GetHeader("Authorization")

		// 验证 Authorization 是否为空
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 解析并验证 JWT
		token, err := auth.ParseJWT(authorizationHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 在上下文中保存解析后的用户信息
		c.Set("user", token.Claims.User)

		c.Next()
	}
}

// AuthRequired 授权中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查上下文中是否存在用户信息
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		// 进行授权判断，例如检查用户权限、角色等
		// ...

		c.Next()
	}
}
```

在上面的示例中，`AuthMiddleware` 是一个认证中间件，它用于验证请求头中的 Authorization 字段中的 JWT，并将解析后的用户信息保存在上下文中，供后续的中间件或处理器函数使用。`AuthRequired` 是一个授权中间件，它用于检查上下文中是否存在用户信息，并根据具体的授权逻辑判断用户是否有权限访问资源。在使用这两个中间件时，可以根据需要在路由组或路由处理器中使用 `Use` 方法添加这两个中间件，以实现认证和授权的功能。

### 认证授权使用

在使用 Gin 框架中的中间件时，一般需要按照以下步骤进行：

1. 定义中间件函数：根据需求，编写一个或多个中间件函数，每个中间件函数都应该符合 `gin.HandlerFunc` 类型的定义，接收一个 `*gin.Context` 参数，用于处理请求和响应。

2. 注册中间件函数：在项目中的合适位置，将中间件函数注册到 Gin 框架中。例如，在应用程序的主函数或初始化函数中，使用 `router.Use()` 方法将中间件函数注册到 Gin 的路由器（`*gin.Engine`）中。

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/your-package/middleware"
)

func main() {
	router := gin.Default()

	// 注册认证中间件
	router.Use(middleware.AuthMiddleware())

	// 注册授权中间件
	router.Use(middleware.AuthRequired())

	// 添加路由处理器函数
	router.GET("/protected", func(c *gin.Context) {
		// 这里可以访问到已经经过认证和授权的用户信息
		user := c.MustGet("user").(string)
		c.JSON(200, gin.H{"user": user})
	})

	router.Run(":8080")
}
```

在上面的示例中，`middleware.AuthMiddleware()` 和 `middleware.AuthRequired()` 分别是认证和授权中间件函数，通过 `router.Use()` 方法将它们注册到 Gin 的路由器中。随后，在路由处理器函数中可以通过 `c.MustGet()` 方法获取中间件函数中设置的上下文值，例如获取已经经过认证和授权的用户信息。

需要注意的是，中间件的注册顺序会影响它们的执行顺序。在使用多个中间件时，可以通过调整注册的顺序来控制它们的执行顺序，从而满足业务需求。

### 发送请求

使用 curl 进行请求时，可以通过 `-H` 参数设置请求头部，将认证令牌放入请求头部进行传递。

以下是一个使用 curl 发送带认证令牌的请求的示例命令：

```bash
curl -H "Authorization: my_token" http://localhost:8080/protected
```

上面的命令中，`-H` 参数用于设置请求头部，将认证令牌 "my_token" 放入请求头部的 "Authorization" 字段中。然后，通过指定请求的 URL "http://localhost:8080/protected" 发送请求。如果认证令牌有效，将会收到正常的响应；如果认证令牌无效，将会收到 HTTP 401 错误响应。

## curl

### 是什么

curl 是一个用于发送 HTTP 请求的命令行工具，支持多种协议，包括 HTTP、HTTPS、FTP、SMTP 等。其基本语法如下：

```
curl [options] [URL]
```

其中，`options` 是 curl 命令的选项，用于设置请求的各种参数，`URL` 是目标 URL，指定请求的目标地址。

以下是一些常用的 curl 选项：

- `-H`：设置请求头部，可以指定多个头部，用逗号分隔。
- `-X`：指定请求方法，如 GET、POST、PUT、DELETE 等。
- `-d`：设置请求体数据，用于 POST 请求。
- `-u`：设置认证用户名和密码。
- `-s`：静默模式，不显示进度和错误信息。
- `-o`：将响应保存到文件。
- `-L`：自动跟随重定向。
- `-k`：忽略 SSL 证书验证。
- `-i`：显示响应头部。

以下是一个示例命令，用于发送一个简单的 GET 请求：

```
curl http://example.com
```

以下是一个示例命令，用于发送一个带自定义请求头部和请求体的 POST 请求：

```
curl -X POST -H "Content-Type: application/json" -d '{"name": "John", "age": 30}' http://example.com
```

这只是 curl 的基本语法和常用选项，curl 还有许多其他选项和用法，可以通过 `curl --help` 命令查看完整的文档。

### 以下是 curl 常见的 HTTP 请求方法的示例代码：

1. GET 请求：

```bash
curl http://example.com
```

2. POST 请求：

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "John", "age": 30}' http://example.com
```

3. PUT 请求：

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name": "John", "age": 30}' http://example.com/1
```

4. DELETE 请求：

```bash
curl -X DELETE http://example.com/1
```

5. PATCH 请求：

```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"age": 31}' http://example.com/1
```

6. HEAD 请求：

```bash
curl -I http://example.com
```

7. OPTIONS 请求：

```bash
curl -X OPTIONS http://example.com
```

以上示例中，`-X` 参数用于指定请求方法，`-H` 参数用于设置请求头部，`-d` 参数用于设置请求体数据，`-I` 参数用于只返回响应头部，`-d` 参数用于只发送请求头部并获取服务器端支持的 HTTP 方法。请注意，实际的请求 URL 需要替换为您的目标地址。

### 动机

curl 是一个命令行工具，用于在终端中发送 HTTP 请求并获取响应，常用于测试和调试 Web 服务。它的产生主要是为了方便地在命令行中进行 HTTP 请求，以便快速测试和调试 API 接口、验证服务器的响应、模拟客户端请求等。

以下是一些 curl 的主要动机：

1. 简单易用：curl 提供了简单的命令行接口，可以通过简单的命令和选项来发送 HTTP 请求，无需编写复杂的代码，非常方便和易用。

2. 多协议支持：curl 支持多种协议，包括 HTTP、HTTPS、FTP、SMTP 等，可以用于测试和调试各种类型的 Web 服务。

3. 强大的选项：curl 提供了丰富的选项，可以设置请求头部、请求方法、请求体数据、认证信息、代理、SSL 证书验证等，可以满足不同的测试和调试需求。

4. 跨平台支持：curl 可以在多个平台上运行，包括 Windows、macOS、Linux 等，具有很好的跨平台支持，方便在不同的操作系统上使用。

5. 广泛应用：curl 在开发者社区中非常流行，广泛应用于 Web 开发、API 测试、网络调试等场景，被广大开发者广泛使用和推崇。

综上所述，curl 的产生主要是为了提供一种简单、方便、跨平台的方式来发送 HTTP 请求，用于测试和调试 Web 服务，满足开发者的需求。

## gorm

<a href="https://blog.csdn.net/u010525694/article/details/94294890" target="_blank" >gorm 中文文档</a>

### golang gorm 有哪些 API

Golang 的 ORM 库 GORM 是一个流行的数据库操作库，提供了丰富的 API 用于进行数据库操作。以下是 GORM 库的一些常用 API：

1. 创建连接：

```go
db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local")
```

2. 定义模型（定义数据库表结构）：

```go
type User struct {
    ID   uint
    Name string
    Age  int
}
```

3. 创建表：

```go
db.CreateTable(&User{})
```

4. 插入记录：

```go
user := User{Name: "John", Age: 30}
db.Create(&user)
```

5. 查询记录：

```go
var user User
db.First(&user, 1) // 查询ID为1的记录并将结果存入user变量
```

6. 更新记录：

```go
db.Model(&user).Update("Age", 40) // 将user的Age字段更新为40
```

7. 删除记录：

```go
db.Delete(&user) // 删除user记录
```

8. 查询多条记录：

```go
var users []User
db.Find(&users)
```

9. 使用条件查询：

```go
db.Where("age > ?", 30).Find(&users) // 查询年龄大于30的记录
```

10. 使用事务：

```go
tx := db.Begin()
// 在事务中执行数据库操作
tx.Create(&user)
tx.Delete(&user)
tx.Commit() // 提交事务
```

11. 使用预加载（Eager Loading）优化查询：

```go
var users []User
db.Preload("Profile").Find(&users) // 预加载关联的Profile模型，减少查询次数
```

以上只是 GORM 库提供的一部分 API，GORM 还提供了许多其他功能，如关联查询、分页查询、原生 SQL 执行、钩子函数等。你可以查看 GORM 官方文档（https://gorm.io）以获取更详细的API参考和示例代码。

### 一对一

在 GORM 中，一对一（One-to-One）关联是指两个模型之间的关联关系，其中一个模型的实例对应着另一个模型的实例。以下是 GORM 中一对一关联的常用 API：

1. 定义模型：

```go
type User struct {
    ID       uint
    Name     string
    Age      int
    Profile  Profile // 与Profile模型建立一对一关联
}

type Profile struct {
    ID       uint
    UserID   uint // 通过UserID字段与User模型建立外键关联
    Address  string
}
```

2. 创建表：

```go
db.CreateTable(&User{})
db.CreateTable(&Profile{})
```

3. 插入记录：

```go
profile := Profile{Address: "1234 Main St."}
user := User{Name: "John", Age: 30, Profile: profile} // 将Profile关联到User
db.Create(&user)
```

4. 查询记录（包含关联记录）：

```go
var user User
db.Preload("Profile").First(&user, 1) // 查询ID为1的User记录，并关联查询其对应的Profile记录
```

5. 更新关联记录：

```go
db.Model(&user).Update("Profile.Address", "5678 Elm St.") // 更新关联的Profile记录的Address字段
```

6. 解除关联关系：

```go
db.Model(&user).Association("Profile").Delete(&user.Profile) // 解除User与Profile的关联关系
```

需要注意的是，GORM 中的一对一关联需要通过外键字段进行关联，如上面的示例中的`UserID`字段。GORM 还提供了其他丰富的一对一关联操作，如设置外键、设置级联关系、查询关联记录、创建关联记录等。详细的一对一关联操作请参考 GORM 官方文档（https://gorm.io/docs/associations.html#one-to-one）以获取更详细的API参考和示例代码。

### 理解

1.  belongto，我属于你，你给我打标记
2.  hasone 我拥有你，我给你打标记
3.  belongto 是 hasone 的反向，从我的角度：我是你的，从你的角度：你是拥有我的

### belongto 与 hasone 的区别

1. belongto: 在包含外键的模型中定义关联字段来实现
2. hasone: 在包含外键的模型中定义关联字段，并在关联的模型中添加外键字段来实现

在 GORM 中，`BelongsTo`和`HasOne`是两种不同的关联类型，用于建立模型之间的关联关系，分别表示一对一的反向关联和正向关联。

1. `BelongsTo`关联：`BelongsTo`表示一对一的反向关联，即一个模型关联到另一个模型，但另一个模型没有关联回来。在数据库中，这通常通过在包含外键的模型中定义关联字段来实现。例如：

```go
type User struct {
    ID       uint
    Profile  Profile // BelongsTo关联，一个User对应一个Profile
}

type Profile struct {
    ID       uint
    Address  string
    UserID   uint // 外键字段，用于关联到User模型
}
```

在上面的示例中，`Profile`模型通过在其定义中添加一个`UserID`字段与`User`模型建立了`BelongsTo`关联，表示一个`Profile`对应一个`User`。

2. `HasOne`关联：`HasOne`表示一对一的正向关联，即一个模型关联到另一个模型，并且另一个模型也关联回来。在数据库中，这通常通过在包含外键的模型中定义关联字段，并在关联的模型中添加外键字段来实现。例如：

```go
type User struct {
    ID       uint
    Profile  Profile // HasOne关联，一个User对应一个Profile
}

type Profile struct {
    ID       uint
    Address  string
    User     User // 关联回User模型
    UserID   uint // 外键字段，用于关联到User模型
}
```

在上面的示例中，`User`模型通过在其定义中添加一个`Profile`字段与`Profile`模型建立了`HasOne`关联，表示一个`User`对应一个`Profile`，并且`Profile`模型通过在其定义中添加一个`UserID`字段与`User`模型建立了外键关联，并关联回了`User`模型。

总结而言，
`BelongsTo`关联表示一对一的反向关联，只在包含外键的模型中定义关联字段；
而`HasOne`关联表示一对一的正向关联，既在包含外键的模型中定义关联字段，也在关联的模型中添加外键字段并关联回来。

两者在使用上略有差异，具体的选择取决于你的数据模型设计和业务需求。

### gorm BelongsTo 和 HasOne 的适用场景

在 GORM 中，`BelongsTo`和`HasOne`是两种不同的关联类型，适用于不同的场景。

1. `BelongsTo`适用场景：
   `BelongsTo`表示一对一的反向关联，适用于以下场景：

- 当一个模型（称为子模型）属于另一个模型（称为父模型），并且子模型中包含了父模型的外键字段时。
- 当你需要在查询子模型时，自动加载关联的父模型。

例如，一个用户（User）模型关联一个配置文件（Profile）模型，其中配置文件（Profile）模型包含了用户（User）模型的外键字段，那么可以使用`BelongsTo`关联来表示这种关系。

```go
// 父模型
type User struct {
    ID       uint
    Profile  Profile // BelongsTo关联，一个User对应一个Profile
}
// 子模型 子模型中包含了父模型的外键字段时
type Profile struct {
    ID       uint
    Address  string
    UserID   uint // 外键字段，用于关联到User模型
}
```

2. `HasOne`适用场景：
   `HasOne`表示一对一的正向关联，适用于以下场景：

- 当一个模型（称为父模型）关联一个模型（称为子模型），并且子模型中包含了父模型的外键字段时。
- 当你需要在查询父模型时，自动加载关联的子模型。

例如，一个用户（User）模型关联一个配置文件（Profile）模型，其中用户（User）模型包含了配置文件（Profile）模型的外键字段，那么可以使用`HasOne`关联来表示这种关系。

```go
// 父模型
type User struct {
    ID       uint
    Profile  Profile // HasOne关联，一个User对应一个Profile
}
// 子模型， 子模型中包含了父模型的外键字段时
type Profile struct {
    ID       uint
    Address  string
    User     User // 关联回User模型
    UserID   uint // 外键字段，用于关联到User模型
}
```

需要根据具体的数据模型设计和业务需求来选择使用`BelongsTo`还是`HasOne`关联。如果子模型包含了父模型的外键字段，则使用`BelongsTo`关联；如果父模型包含了子模型的外键字段，并且需要在查询父模型时自动加载关联的子模型，则使用`HasOne`关联。

### 两者在数据库数据查询返回有什么不同

在 GORM 中，`BelongsTo`和`HasOne`在数据库数据查询返回时有一些不同。

1. `BelongsTo`关联：

- 当使用`BelongsTo`关联时，通过查询子模型（包含外键字段）可以自动加载关联的父模型。例如，当查询用户（User）模型时，可以自动加载关联的配置文件（Profile）模型，从而在结果中包含了关联的父模型的数据。

```go
var user User
db.Preload("Profile").First(&user)
// 上述代码中，使用了Preload方法加载了关联的Profile模型，从而在查询用户模型时自动加载了关联的父模型数据。
```

2. `HasOne`关联：

- 当使用`HasOne`关联时，通过查询父模型（包含子模型的外键字段）可以自动加载关联的子模型。例如，当查询配置文件（Profile）模型时，可以自动加载关联的用户（User）模型，从而在结果中包含了关联的子模型的数据。

```go
var profile Profile
db.Preload("User").First(&profile)
// 上述代码中，使用了Preload方法加载了关联的User模型，从而在查询配置文件模型时自动加载了关联的子模型数据。
```

需要注意的是，在使用`BelongsTo`和`HasOne`关联时，都可以使用 GORM 提供的`Preload`方法来进行预加载，从而在查询时自动加载关联的模型数据，避免了 N+1 查询问题。具体的查询返回结果会根据你的数据模型设计和数据库查询操作而有所不同，但通过预加载关联模型，可以方便地获取关联的数据并避免额外的数据库查询。

### User BelongsTo Profile

对不起，我之前的回答中确实遗漏了 GORM 中关联关系的 tag 设置。下面是正确的示例代码：

使用 `BelongsTo` 关联时的结构体定义：

```go
type User struct {
    ID       uint
    Profile  Profile   `gorm:"foreignkey:ProfileID"` // BelongsTo 关联，一个 User 属于一个 Profile
}

type Profile struct {
    ID       uint
    Address  string
    UserID   uint // 外键字段，用于关联到 User 模型
}
```

### User HasOne Profile

使用 `HasOne` 关联时的结构体定义：

```go
type User struct {
    ID       uint
    Profile  Profile   `gorm:"foreignkey:UserID"` // HasOne 关联，一个 User 拥有一个 Profile
    ProfileID uint      // 外键字段，用于关联到 Profile 模型
}

type Profile struct {
    ID       uint
    Address  string
}
```

在以上示例代码中，使用了 `gorm` 的 tag 来指定外键字段的名称，通过 `foreignkey` tag 来显式指定外键字段的名称，从而与关联的模型建立正确的关联关系。

`BelongsTo` 关联中，`User` 模型中的 `Profile` 字段的 tag 设置为 `foreignkey:ProfileID`，表示 `ProfileID` 字段在 `Profile` 模型中用作外键字段，关联到 `User` 模型的 `ID` 字段。

`HasOne` 关联中，`User` 模型中的 `Profile` 字段的 tag 设置为 `foreignkey:UserID`，表示 `UserID` 字段在 `User` 模型中用作外键字段，关联到 `Profile` 模型的 `ID` 字段。

根据具体的业务需求和数据库设计，可以根据 `gorm` 的 tag 设置来正确定义结构体，以建立合适的关联关系。

### BelongsTo 建表的结果

在使用 GORM 中的 `BelongsTo` 方法建立一对一关联关系时，对应的数据库表结构通常会包含一个外键字段，用于关联到另一个表的主键。以下是一个示例：

假设我们有两个模型 `User` 和 `Profile`，其中 `User` 模型通过 `BelongsTo` 方法与 `Profile` 模型建立一对一关联关系，其中 `User` 模型包含一个外键字段 `ProfileID`，用于关联到 `Profile` 模型的主键字段 `ID`。

定义 `User` 和 `Profile` 模型如下：

```go
type User struct {
    ID        uint
    ProfileID uint     // 外键字段，用于关联到 Profile 模型
    Name      string
    Profile   Profile  `gorm:"foreignkey:ProfileID"`  // BelongsTo 关联，一个 User 属于一个 Profile
}

type Profile struct {
    ID       uint
    Address  string
}
```

在以上示例中，`User` 模型包含了一个名为 `ProfileID` 的外键字段，用于关联到 `Profile` 模型的主键字段 `ID`。通过在 `User` 模型的 `Profile` 字段上使用 `gorm` 的 `foreignkey` tag，指定了外键字段的名称为 `ProfileID`。

对应的数据库表结构可能如下所示：

```
users 表：
+----+-------------+------+
| ID | ProfileID   | Name |
+----+-------------+------+
| 1  | 1           | John |
+----+-------------+------+
| 2  | 2           | Jane |
+----+-------------+------+

profiles 表：
+----+---------+
| ID | Address |
+----+---------+
| 1  | Addr1   |
+----+---------+
| 2  | Addr2   |
+----+---------+
```

在以上示例中，`users` 表包含了一个外键字段 `ProfileID`，用于关联到 `profiles` 表的主键字段 `ID`，从而建立了一对一关联关系。每个 `User` 对应一个 `Profile`，通过 `ProfileID` 字段进行关联。通过 GORM 的 `BelongsTo` 方法和外键字段的定义，可以在数据库中建立一对一关联关系，并进行相应的查询和操作。

### 查询的结果

通过 GORM 进行查询时，使用 `BelongsTo` 方法建立的一对一关联关系会在查询结果中包含关联的数据。例如，使用以下的查询代码：

```go
var user User
db.Preload("Profile").First(&user, 1)
```

其中 `db` 是 GORM 的数据库连接实例，`User` 是模型结构体，`Profile` 是关联的模型结构体。

以上查询会查询 ID 为 1 的 `User` 记录，并将关联的 `Profile` 数据一并查询出来，将其填充到 `User` 模型的 `Profile` 字段中。

查询结果会包含关联的 `Profile` 数据，例如：

```go
{
  "ID": 1,
  "ProfileID": 1,
  "Name": "John",
  "Profile": {
    "ID": 1,
    "Address": "Addr1"
  }
}
```

可以看到，查询结果中包含了 `Profile` 数据，它作为 `User` 模型的嵌套字段被填充到了 `Profile` 字段中。这样可以方便地获取一对一关联关系的数据，并进行操作和处理。

### sqlite 完整事例代码

以下是一个完整的示例代码，演示如何使用 GORM 进行一对一关联关系的查询：

```go
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    ID        uint
    ProfileID uint     // 外键字段，用于关联到 Profile 模型
    Name      string
    Profile   Profile  `gorm:"foreignkey:ProfileID"`  // BelongsTo 关联，一个 User 属于一个 Profile
}

type Profile struct {
    ID       uint
    Address  string
}

func main() {
    // 连接数据库
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect database")
    }
    db.AutoMigrate(&User{}, &Profile{})  // 自动迁移表结构

    // 创建一条 User 记录和对应的 Profile 记录
    user := User{Name: "John"}
    profile := Profile{Address: "Addr1"}
    db.Create(&user)
    db.Create(&profile)
    user.ProfileID = profile.ID
    db.Save(&user)

    // 查询 User 记录并包含关联的 Profile 数据
    var result User
    db.Preload("Profile").First(&result, 1)

    // 输出查询结果
    fmt.Printf("User: %+v\n", result)
    fmt.Printf("Profile: %+v\n", result.Profile)
}
```

在以上示例中，首先定义了 `User` 和 `Profile` 两个模型，并使用 `BelongsTo` 方法建立了一对一关联关系。然后通过 GORM 进行数据库连接，并创建了一条 `User` 记录和对应的 `Profile` 记录。接着通过 `Preload` 方法查询 `User` 记录时，包含了关联的 `Profile` 数据，最终将查询结果输出到控制台。

请注意，示例中使用了 SQLite 作为数据库，需要提前安装 SQLite 并创建一个名为 `test.db` 的数据库文件。实际使用中，数据库连接参数和表结构定义可能需要根据具体情况进行调整。

### 使用 mysql 作为数据库的代码怎么写

以下是一个使用 MySQL 数据库作为存储的完整示例代码，演示如何使用 GORM 进行一对一关联关系的查询：

```go
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	ProfileID uint     // 外键字段，用于关联到 Profile 模型
	Name      string
	Profile   Profile  `gorm:"foreignkey:ProfileID"`  // BelongsTo 关联，一个 User 属于一个 Profile
}

type Profile struct {
	ID       uint
	Address  string
}

func main() {
	// 连接数据库
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&User{}, &Profile{}) // 自动迁移表结构

	// 创建一条 User 记录和对应的 Profile 记录
	user := User{Name: "John"}
	profile := Profile{Address: "Addr1"}
	db.Create(&user)
	db.Create(&profile)
	user.ProfileID = profile.ID
	db.Save(&user)

	// 查询 User 记录并包含关联的 Profile 数据
	var result User
	db.Preload("Profile").First(&result, 1)

	// 输出查询结果
	fmt.Printf("User: %+v\n", result)
	fmt.Printf("Profile: %+v\n", result.Profile)
}
```

在以上示例中，我们使用了 GORM 的 MySQL 数据库驱动 `gorm.io/driver/mysql`，并通过 `dsn` 参数设置了 MySQL 数据库的连接信息，包括用户名、密码、地址、端口、数据库名称以及连接参数。接着创建了 `User` 和 `Profile` 两个模型，并使用 `BelongsTo` 方法建立了一对一关联关系。其余的创建记录、查询记录和输出查询结果的代码与之前的示例类似。

请注意，在实际使用中，需要根据自己的 MySQL 数据库配置和表结构定义进行相应的调整。

### 上述代码的表结构，及建表语句

根据以上示例代码，对应的 MySQL 数据库表结构可以如下定义：

```sql
-- User 表
CREATE TABLE `users` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `profile_id` INT(10) UNSIGNED NOT NULL,
  `name` VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_profile` (`profile_id`),
  CONSTRAINT `fk_users_profile` FOREIGN KEY (`profile_id`) REFERENCES `profiles` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Profile 表
CREATE TABLE `profiles` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `address` VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

以上 SQL 语句创建了两张表，分别是 `users` 表和 `profiles` 表。`users` 表包含了 `id`、`profile_id` 和 `name` 三个字段，其中 `id` 是主键，`profile_id` 是外键，与 `profiles` 表的 `id` 字段关联。`profiles` 表包含了 `id` 和 `address` 两个字段，其中 `id` 是主键。

在 `users` 表中，`profile_id` 字段与 `profiles` 表的 `id` 字段建立了外键关系，并通过 `CONSTRAINT` 定义了外键约束，指定了在删除 `profiles` 表中的记录时要执行的操作为级联删除 (`ON DELETE CASCADE`)，即删除 `profiles` 表中的记录时，会同时删除关联的 `users` 表中的记录。这样就实现了 User 和 Profile 之间的一对一关联关系。

### gorm 注意事项

### map

1. 注意从 map 创建时，不会调用钩子，不会保存关联，也不会重新填充主键值

### gorm 有哪些 tag

`gorm` 是 Go 中一款流行的 ORM（对象关系映射）库，用于操作数据库。`gorm` 支持使用标签（tag）来对模型（Model）的字段进行注解，以指定其在数据库中的映射行为。以下是 `gorm` 支持的一些常用的标签（tag）：

1. `gorm:"column:column_name"`：指定字段在数据库表中对应的列名。
2. `gorm:"primary_key"`：指定字段作为数据库表的主键。
3. `gorm:"type:data_type"`：指定字段在数据库中的数据类型，如 `int`, `varchar`, `text` 等。
4. `gorm:"size:size"`：指定字段在数据库中的大小限制，如 `size:255` 表示该字段最大长度为 255。
5. `gorm:"not null"`：指定字段不能为空。
6. `gorm:"unique"`：指定字段在数据库表中的值必须唯一。
7. `gorm:"index"`：为字段创建数据库索引。
8. `gorm:"default:default_value"`：指定字段在数据库中的默认值。
9. `gorm:"auto_increment"`：指定字段在数据库中自增。
10. `gorm:"precision:precision; scale:scale"`：指定字段的精度和小数位数，仅适用于浮点数字段。
11. `gorm:"uniqueIndex:unique_index_name"`：为字段创建指定名称的唯一索引。
12. `gorm:"index:idx_name"`：为字段创建指定名称的索引。
13. `gorm:"comment:comment"`：为字段添加注释。

以上只是 `gorm` 支持的一些常用的标签，实际上 `gorm` 还支持更多的标签，用于配置模型在数据库中的映射行为。通过使用这些标签，我们可以灵活地定义模型的字段属性，并与数据库表的列进行对应，从而实现方便的数据库操作。

当然，还有更多的 `gorm` 标签可以用于进一步配置模型在数据库中的映射行为，包括但不限于以下几个：

14. `gorm:"embedded"`：指定字段为嵌入式结构体。
15. `gorm:"embeddedPrefix:prefix"`：指定嵌入式结构体在数据库表中的列名前缀。
16. `gorm:"polymorphic:polymorphic_field;polymorphic_value"`：指定字段为多态关联，用于多表关联查询。
17. `gorm:"foreignkey:foreign_key_field"`：指定关联的外键字段。
18. `gorm:"association_foreignkey:association_foreign_key_field"`：指定关联的外键字段。
19. `gorm:"many2many:join_table"`：指定多对多关联的中间表名称。
20. `gorm:"many2many:join_table:foreignkey:foreign_key;association_foreignkey:association_foreign_key"`：指定多对多关联的中间表及外键字段。
21. `gorm:"polymorphic:polymorphic_field;polymorphic_identity:polymorphic_identity_value"`：指定多态关联的字段及标识值。

这些标签可以用于进一步配置模型在数据库中的映射行为，实现更复杂的数据库操作需求。详细的 `gorm` 标签用法可以参考 `gorm` 的官方文档：https://gorm.io/docs/models.html#field-attributes

当然，这还不是全部，`gorm` 还提供了一些其他的标签，如下：

22. `gorm:"unique"`：指定字段在数据库中为唯一值。
23. `gorm:"not null"`：指定字段在数据库中不能为空。
24. `gorm:"default:value"`：指定字段在数据库中的默认值。
25. `gorm:"index"`：为字段创建数据库索引。
26. `gorm:"size:value"`：指定字段的数据库大小。
27. `gorm:"precision:value"`：指定字段的数据库精度。
28. `gorm:"scale:value"`：指定字段的数据库刻度。
29. `gorm:"comment:comment"`：为字段添加数据库注释。

这些标签可以用于进一步定制模型在数据库中的表现形式，满足不同的业务需求。具体使用方法和示例可以参考 `gorm` 的官方文档：https://gorm.io/docs/models.html#field-attributes

是的，还有一些其他的 `gorm` 标签，包括：

30. `gorm:"createIndex"`：在创建表时为字段创建索引。
31. `gorm:"updateIndex"`：在更新表时为字段创建索引。
32. `gorm:"deleteIndex"`：在删除表时删除字段的索引。
33. `gorm:"uniqueIndex"`：为字段创建唯一索引。
34. `gorm:"fulltextIndex"`：为字段创建全文索引。
35. `gorm:"compositeIndex:composite_index_name"`：为字段创建复合索引，并指定索引名称。
36. `gorm:"type:column_type"`：指定字段在数据库中的列类型。
37. `gorm:"column:column_name"`：指定字段在数据库中的列名称。
38. `gorm:"autoCreateTime"`：在创建记录时自动设置字段的创建时间。
39. `gorm:"autoUpdateTime"`：在更新记录时自动设置字段的更新时间。
40. `gorm:"primaryKey"`：将字段设置为主键。
41. `gorm:"uniqueIndex:unique_index_name"`：为字段创建唯一索引，并指定索引名称。
42. `gorm:"check:constraint_name:constraint_expression"`：为字段创建检查约束，并指定约束名称和约束表达式。

这些标签可以用于进一步定制模型在数据库中的映射行为和表现形式，满足更复杂的数据库操作需求。具体使用方法和示例可以参考 `gorm` 的官方文档：https://gorm.io/docs/models.html#field-attributes

抱歉，以上列出的是 gorm 中比较常用的标签，目前没有其他的 gorm 标签了。如果有新的版本发布，可能会有新的标签被引入。您可以随时查阅 gorm 的官方文档，以获取最新的信息和更新。

### 区别

1. db.Table
2. db.Rows
3. db.Scan

### db.Table

`db.Table` 是 `gorm` 库中用于指定查询表名的方法。通过 `db.Table` 方法，可以在查询时显式指定要查询的数据库表名，而不是使用 `gorm` 的默认表名推断规则。

例如，下面是使用 `db.Table` 方法查询名为 "users" 的数据库表的示例：

```go
// 定义 User 结构体
type User struct {
	ID   uint
	Name string
	Age  int
}

// 使用 db.Table 指定表名为 "users" 进行查询
var users []User
db.Table("users").Find(&users)
```

在这个示例中，`db.Table("users")` 方法指定了查询的数据库表名为 "users"，然后使用 `Find` 方法执行查询操作，并将查询结果存储到 `users` 变量中。

需要注意的是，使用 `db.Table` 方法后，`gorm` 将会完全按照指定的表名进行查询，而不再使用默认的表名推断规则。这可以在需要查询非默认表名的情况下使用，例如当查询多个不同表结构相似的表时，或者当需要查询不在 `gorm` 模型定义中的表时。

### FirstOrInit/FirstOrCreate

1. db.FirstOrInit 有点像多态
2. Attrs 不会被用于生成查询 SQL, 找不到仅用于创建，找到了无效
3. 不管是否找到记录，Assign 都会将属性赋值给 struct，但这些属性不会被用于生成查询 SQL，也不会被保存到数据库

**区别：**

FirstOrCreate 和 FirstOrInit 的区别： FirstOrCreate 的创建或更新会触发 callback ，而 FirstOrInit 不会

### 零值处理

1. Save 会保存所有的字段，即使字段是零值（如果保存值不包含主键，它将执行 Create，否则将执行 Update（包含所有字段））
2. Update
   1. 更新支持使用 struct 或 map[string]face{}更新，使用 struct 更新时默认只会更新非零字段
   2. 注意使用 struct 更新时，GORM 只会更新非零字段。您可能希望使用 map 来更新属性或使用 Select 来指定要更新的字段

### gorm dryrun

在 Gorm 中，DryRun 模式是一种用于调试 SQL 语句的特殊模式。当你在开发阶段使用 Gorm 时，你可能会需要查看 Gorm 生成的 SQL 语句，以便确定 Gorm 是否正确地生成了预期的 SQL 语句。DryRun 模式允许你在不实际执行 SQL 语句的情况下查看生成的 SQL 语句。

在 Gorm 中启用 DryRun 模式非常简单。只需将 `db.Session(&gorm.Session{DryRun: true})` 作为查询的第一个参数传递给 Gorm，即可启用 DryRun 模式。例如：

```
db.Session(&gorm.Session{DryRun: true}).Find(&users)
```

在上面的示例中，我们将 DryRun 模式启用，并通过 `Find` 方法查询了用户数据。由于我们启用了 DryRun 模式，Gorm 只会生成 SQL 语句，并将其打印到控制台上，而不会实际执行 SQL 语句。你可以通过查看控制台输出来检查生成的 SQL 语句是否正确。

### 安全

在使用 Gorm 或任何其他 ORM（对象关系映射）库时，确保 SQL 安全是非常重要的，以防止 SQL 注入攻击。下面是一些在使用 Gorm 时保持 SQL 安全的建议：

1. 使用参数化查询（Prepared Statements）：Gorm 默认使用参数化查询，这可以防止 SQL 注入攻击。参数化查询将查询参数与 SQL 查询语句分开，确保参数不会被解释为 SQL 代码。在 Gorm 中，你可以使用 `db.Preload()`、`db.Where()`、`db.First()` 等方法进行查询，并将查询参数作为参数传递，而不是将其直接嵌入到 SQL 查询语句中。

例如：

```go
name := "John"
age := 25
db.Where("name = ?", name).Where("age > ?", age).Find(&users)
```

2. 不要直接拼接 SQL 查询语句：避免将用户输入或其他未经验证的数据直接拼接到 SQL 查询语句中，这可能会导致 SQL 注入攻击。相反，使用 Gorm 提供的方法来构建查询，如 `db.Where()`、`db.Or()`、`db.Not()` 等，这些方法会自动处理参数化查询。

例如：

```go
name := "John"
age := 25
db.Where("name = ?", name).Where("age > ?", age).Find(&users)
```

3. 验证用户输入：在将用户输入数据传递给 SQL 查询之前，务必对其进行验证和过滤，以防止恶意 SQL 注入攻击。可以使用合适的验证库或编写自己的验证逻辑来确保用户输入数据的合法性。

4. 限制数据库权限：在使用数据库时，确保数据库用户只具有最小必需的权限，例如只允许执行特定的查询和操作，以限制潜在的攻击面。

5. 更新 Gorm 版本：定期更新 Gorm 到最新版本，以确保你使用的是最新的安全补丁和修复了可能存在的安全漏洞的版本。

总之，确保 SQL 安全对于使用 Gorm 或任何其他 ORM 库都是至关重要的。遵循最佳实践，使用参数化查询、验证用户输入、限制数据库权限等措施，以保护你的应用程序免受 SQL 注入攻击。

### 验证用户输入

验证用户输入是确保应用程序安全的重要步骤，包括在使用 Gorm 或任何其他 ORM 库时。下面是一些实现验证用户输入的方法：

1. 使用 Go 的内置验证库：Go 标准库中提供了多个用于验证用户输入的包，例如 `regexp`、`strconv`、`strings` 等。你可以使用这些包来检查用户输入是否符合期望的格式、类型、长度等。例如，你可以使用正则表达式验证电子邮件地址的格式，使用 `strconv` 包将字符串转换为数字等。

```go
import (
    "regexp"
)

func ValidateEmail(email string) bool {
    // 使用正则表达式验证电子邮件地址格式
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    return regexp.MustCompile(pattern).MatchString(email)
}

func main() {
    email := "user@example.com"
    if ValidateEmail(email) {
        fmt.Println("Email is valid")
    } else {
        fmt.Println("Email is invalid")
    }
}
```

2. 自定义验证逻辑：根据你的应用程序需求，你可以编写自己的验证逻辑来验证用户输入。例如，你可以检查密码的长度、检查输入是否为合法的日期等。可以通过编写自定义的验证函数或方法来实现。

```go
func ValidatePassword(password string) bool {
    // 检查密码长度是否大于等于8
    return len(password) >= 8
}

func main() {
    password := "mypassword"
    if ValidatePassword(password) {
        fmt.Println("Password is valid")
    } else {
        fmt.Println("Password is invalid")
    }
}
```

3. 使用第三方验证库：除了 Go 标准库外，还有许多第三方验证库可供选择，这些库提供了更丰富的验证功能，如 `validator`、`go-playground/validator` 等。这些库提供了各种验证规则和自定义验证功能，可以方便地在你的应用程序中使用。

```go
import (
    "github.com/go-playground/validator/v10"
)

type User struct {
    Name     string `validate:"required"`
    Age      int    `validate:"gte=18"`
    Email    string `validate:"email"`
}

func main() {
    user := User{Name: "John", Age: 20, Email: "user@example.com"}
    validate := validator.New()
    err := validate.Struct(user)
    if err != nil {
        fmt.Println("Validation failed:", err)
    } else {
        fmt.Println("Validation passed")
    }
}
```

以上只是验证用户输入的一些简单示例，实际上，验证用户输入可能会涉及更复杂的业务逻辑和多个字段的验证。根据你的应用程序需求，选择合适的验证方法，并根据最佳实践来保护应用程序免受恶意用户输入造成的安全威胁。

### gin 防 sql 注入校验

在使用 Gin 进行 SQL 注入校验时，可以考虑以下几种方法：

1. 使用预处理语句：使用预处理语句来执行 SQL 查询和更新操作，而不是直接将用户输入的数据插入 SQL 查询中。预处理语句可以在 SQL 查询执行之前将用户输入的数据作为参数传递给数据库，从而防止 SQL 注入攻击。在使用 Gin 进行数据库操作时，可以使用 `?` 占位符来代替直接插入用户输入的数据，然后将用户输入的数据作为参数传递给数据库库的方法，数据库库会自动使用预处理语句来执行 SQL 操作。

例如，在使用 Go 自带的 `database/sql` 进行数据库操作时，可以使用 `?` 占位符来代替直接插入用户输入的数据，并将用户输入的数据作为参数传递给 `Exec` 或 `Query` 方法，`database/sql` 会自动使用预处理语句来执行 SQL 操作。

```go
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func getUserByName(c *gin.Context) {
    // 获取用户输入的用户名
    username := c.Query("username")

    // 使用预处理语句，将用户输入作为参数传递
    stmt, err := db.Prepare("SELECT * FROM users WHERE username = ?")
    if err != nil {
        // 处理错误
        // ...
    }
    defer stmt.Close()

    rows, err := stmt.Query(username)
    if err != nil {
        // 处理错误
        // ...
    }

    // 处理查询结果
    // ...
}
```

2. 使用 ORM 的安全特性：一些数据库 ORM 库（如 Gorm、Xorm 等）提供了安全特性，可以防止 SQL 注入攻击。这些特性通常包括自动转义用户输入的数据，避免将用户输入的数据作为原始 SQL 查询的一部分。在使用 ORM 进行数据库操作时，尽量使用 ORM 提供的安全特性，以确保用户输入的数据不会导致 SQL 注入。

例如，在使用 Gorm 进行查询时，可以使用 `Find` 方法，并将用户输入的数据作为参数传递给方法，Gorm 会自动转义用户输入的数据，从而防止 SQL 注入攻击。

```go
import "github.com/gin-gonic/gin"
import "gorm.io/gorm"

func getUserByName(c *gin.Context) {
    // 获取用户输入的用户名
    username := c.Query("username")

    // 使用 Gorm 进行查询，将用户输入作为参数传递，Gorm 会自动转义用户输入的数据
    var user User
    if err := db.Find(&user, "username = ?", username).Error; err != nil {
        // 处理错误
        // ...
    }

    // 处理查询结果
    // ...
}
```

3. 自定义输入验证：根据你的应用程序需求，可以编写自定义的输入验证逻辑，对用户输入的数据进行校验。例如，可以使用正则表达式、字符串处理函数等方法对用户输入的数据进行验证，确保其符合预期的格式，从而防止 SQL 注入攻击。在使用 Gin 进行输入验证时，可以使用 `regexp` 包来进行正则表达式验证，或者使用字符串处理函数来对用户输入的数据进行格式化、过滤或转义。

例如，可以编写一个自定义的验证函数 `isValidUsername` 来验证用户名是否符合预期的格式，例如只包含字母和数字，长度在 6 到 20 个字符之间：

```go
import (
    "github.com/gin-gonic/gin"
    "regexp"
)

func getUserByName(c *gin.Context) {
    // 获取用户输入的用户名
    username := c.Query("username")

    // 自定义输入验证逻辑，例如使用正则表达式验证用户名格式
    if !isValidUsername(username) {
        // 处理用户名格式错误
        // ...
    }

    // 使用 Gorm 进行查询，将用户输入作为参数传递，Gorm 会自动转义用户输入的数据
    // ...
}

func isValidUsername(username string) bool {
    // 定义用户名格式的正则表达式
    pattern := "^[a-zA-Z0-9]{6,20}$"

    // 使用正则表达式验证用户名格式
    match, _ := regexp.MatchString(pattern, username)

    return match
}
```

以上是一些常见的防止 SQL 注入攻击的方法，在实际应用中，可以根据具体的业务场景和安全需求，结合使用这些方法或其他安全措施来确保应用程序的数据安全。同时，还应定期更新依赖库，并保持应用程序的安全最佳实践，以保护应用程序免受 SQL 注入等安全威胁。

### gorm 多态【我觉得：就算不晓得所谓的多态，我也不可能设计出后面那种表】

在 Gorm 中，多态（Polymorphism）是指在一个数据库表中存储多种类型的数据，并通过一个公共的字段来标识不同的类型。这样可以在一个表中存储多种实体的数据，从而避免了创建多张表的复杂性。

Gorm 提供了多态关联（Polymorphic Associations）的功能，可以通过使用 `gorm:"polymorphic"` 标签来定义多态关联。下面是一个示例：

```go
type Comment struct {
    ID           uint
    Comment      string
    CommentableID uint
    CommentableType string
}

type Post struct {
    ID       uint
    Title    string
    Comments []Comment
}

type Video struct {
    ID       uint
    Title    string
    Comments []Comment
}
```

在上面的示例中，`Comment` 模型有一个 `CommentableID` 字段用来存储关联的实体的 ID，以及一个 `CommentableType` 字段用来存储关联实体的类型。`CommentableID` 字段存储了实体的 ID，`CommentableType` 字段存储了实体的类型（例如 "Post" 或 "Video"），用于标识关联的实体类型。

在定义多态关联时，可以使用 `gorm:"polymorphic"` 标签来声明 `CommentableID` 和 `CommentableType` 字段的多态关联，如下所示：

```go
type Comment struct {
    ID           uint
    Comment      string
    CommentableID uint
    CommentableType string `gorm:"polymorphic:Commentable;"`
}

type Post struct {
    ID       uint
    Title    string
    Comments []Comment `gorm:"polymorphic:Commentable;"`
}

type Video struct {
    ID       uint
    Title    string
    Comments []Comment `gorm:"polymorphic:Commentable;"`
}
```

在以上示例中，`Comment` 模型中的 `CommentableType` 字段使用了 `gorm:"polymorphic:Commentable;"` 标签，其中 "Commentable" 是关联的多态类型名，用于在数据库中存储关联实体的类型。

使用多态关联后，可以通过 Gorm 进行查询和操作，例如：

```go
// 查询评论的所有子评论
var comments []Comment
db.Where("CommentableID = ? AND CommentableType = ?", parentCommentID, "Comment").Find(&comments)

// 查询文章的评论
var post Post
db.Preload("Comments").Find(&post)

// 查询视频的评论
var video Video
db.Preload("Comments").Find(&video)
```

通过使用多态关联，可以轻松地在一个表中存储多种类型的数据，并通过公共的字段进行关联和查询，从而实现多态的功能。

### gorm 多态与不使用多态看上去没有什么区别

在某些情况下，使用多态和不使用多态的确会看上去没有什么区别，但是多态的优势在于更好的代码可读性和可维护性。

使用多态关联的一个重要好处是，可以避免重复代码和逻辑，特别是当有多个模型需要关联到同一模型时，可以避免在每个模型中重复定义相同的关联字段和关联方法。另外，使用多态关联还可以使代码更加简洁和易于理解。

例如，假设有一个 `Comment` 模型需要关联到 `Post` 和 `Video` 两个模型，如果不使用多态关联，需要在 `Comment` 模型中定义两个不同的外键，如下所示：

```go
type Comment struct {
    ID        uint
    Comment   string
    PostID    uint
    VideoID   uint
}

type Post struct {
    ID       uint
    Title    string
    Comments []Comment
}

type Video struct {
    ID       uint
    Title    string
    Comments []Comment
}
```

这样的代码会让数据模型变得混乱，而且在查询和操作数据时也会变得更加麻烦。

相反，使用多态关联，可以将外键字段和关联方法定义在 `Comment` 模型中，从而使代码更加简洁和易于理解，如前面所示的示例。使用多态关联还可以在代码中使用多态的概念，从而使代码更加可读性和可维护性。

因此，虽然使用多态关联可能看起来并没有显著的区别，但它确实可以提高代码的可读性和可维护性。

### polymorphicValue

`polymorphicValue` 是 GORM 中用于处理多态关联的字段，它允许在关联关系中使用不同的数据类型，从而实现一对多或多对多关联的灵活性。

**在 GORM 中，`polymorphicValue` 是一个用于表示多态关联的字段名，它通常与 `polymorphicType` 字段一起使用，用于指定多态关联的实际类型。 `polymorphicValue` 字段的值通常是一个字符串，用于表示关联的实际类型的名称或标识。**

以下是一个简单的示例，演示了如何在 GORM 中使用 `polymorphicValue` 字段来实现多态关联：

```go
type Comment struct {
    ID              uint
    Body            string
    CommentableID   uint
    CommentableType string
    Commentable     interface{} `gorm:"polymorphic:Commentable;polymorphicValue:commentable_type"` // 定义多态关联
}

type Post struct {
    ID       uint
    Title    string
    Comments []Comment `gorm:"polymorphic:Commentable;polymorphicValue:commentable_type"` // 定义多态关联
}

type Video struct {
    ID       uint
    Title    string
    Comments []Comment `gorm:"polymorphic:Commentable;polymorphicValue:commentable_type"` // 定义多态关联
}
```

在上面的例子中，`Comment` 模型通过 `Commentable` 字段与 `Post` 和 `Video` 模型建立了多态关联，表示一条评论可以关联到多个不同类型的模型（即 `Post` 和 `Video`），同时通过 `CommentableType` 字段和 `polymorphicValue` 标签来标识实际关联的类型。

在数据库中，可以通过添加 `CommentableID` 和 `CommentableType` 字段来表示多态关联的关系。同时，GORM 会自动解析 `Commentable` 字段和 `polymorphicValue` 标签，从而在查询和操作数据时可以直接通过关联字段进行操作，例如：

```go
// 查询帖子及其评论
var post Post
db.Preload("Comments").Find(&post, 1) // 查询 ID 为 1 的帖子及其评论

// 查询视频及其评论
var video Video
db.Preload("Comments").Find(&video, 1) // 查询 ID 为 1 的视频及其评论

// 创建评论关联到帖子
post := Post{Title: "Post 1"}
comment := Comment{Body: "Comment 1", Commentable: &post}
db.Create(&post)
db.Create(&comment)

// 创建评论关联到视频
video := Video{Title: "Video 1"}
comment := Comment{Body: "Comment 2", Commentable: &video}
db.Create(&video)
db.Create(&comment)
```

通过以上示例，可以看到如何在 GORM 中使用 `polymorphicValue` 字段来实现多态关联，从而实现模型之间的灵活关联查询和操作。

### 得到的表结构

上面的示例中，Comment 模型通过 Commentable 字段与 Post 和 Video 模型建立了多态关联，同时通过 CommentableType 字段和 polymorphicValue 标签指定了实际关联类型的名称为 commentable_type。

在生成的 MySQL 表结构中，会包含 commentable_type 字段来表示实际关联类型，例如：

```sql
CREATE TABLE `comments` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `body` VARCHAR(255) NOT NULL,
  `commentable_id` INT(10) UNSIGNED NOT NULL,
  `commentable_type` VARCHAR(255) NOT NULL, -- 多态关联类型字段
  PRIMARY KEY (`id`)
);

CREATE TABLE `posts` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `videos` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);
```

从上面的表结构可以看到，comments 表中包含了 commentable_type 字段来表示实际关联类型，这样在查询和操作多态关联数据时，GORM 可以根据该字段来判断实际关联类型，从而实现多态关联的查询和操作。

### 插入数据

好的，以下是使用 GORM 插入几条数据的示例，假设我们已经定义了如下的数据模型：

```go
type Comment struct {
    ID              uint
    Body            string
    CommentableID   uint
    CommentableType string `gorm:"polymorphic:Commentable;polymorphicValue:commentable_type"`
}

type Post struct {
    ID       uint
    Title    string
    Comments []Comment `gorm:"polymorphic:Commentable;polymorphicValue:commentable_type"`
}

type Video struct {
    ID       uint
    Title    string
    Comments []Comment `gorm:"polymorphic:Commentable;polymorphicValue:commentable_type"`
}
```

现在我们可以通过创建实例并设置相关字段的值，然后使用 GORM 的 `Create` 方法插入数据到对应的表中。

```go
// 创建 Post
post := Post{
    Title: "Post 1",
    Comments: []Comment{
        {Body: "Comment 1 for Post 1"},
        {Body: "Comment 2 for Post 1"},
    },
}
gormDB.Create(&post)

// 创建 Video
video := Video{
    Title: "Video 1",
    Comments: []Comment{
        {Body: "Comment 1 for Video 1"},
        {Body: "Comment 2 for Video 1"},
    },
}
gormDB.Create(&video)
```

上面的示例中，我们分别创建了一个带有多个评论的 `Post` 和 `Video` 实例，并使用 `Create` 方法将它们插入到对应的表中。由于评论模型 `Comment` 中定义了 `CommentableID` 和 `CommentableType` 字段作为多态关联的键值对，GORM 会根据 `polymorphicValue` 标签自动设置 `CommentableType` 字段的值为 `commentable_type`，从而插入多态关联数据到 `comments` 表中。

插入数据后，可以通过查询相关的表来验证数据是否插入成功。例如，可以使用以下代码查询 `comments` 表中的数据：

```go
var comments []Comment
gormDB.Find(&comments)
fmt.Println(comments)
```

如果一切设置正确，你应该能够在输出中看到插入的评论数据，并且 `CommentableType` 字段的值应该与 `polymorphicValue` 标签中指定的值相同。

### 至此数据库数据

假设我们使用的是 MySQL 数据库，并且已经成功执行了上面的示例代码插入了一些数据，那么在数据库中存储的数据将类似于以下的结构：

comments 表：

```
+----+------------------------+----------------+---------------------+------------------+
| ID | Body                   | CommentableID  | CommentableType     | created_at       |
+----+------------------------+----------------+---------------------+------------------+
| 1  | Comment 1 for Post 1  | 1              | Post                | 2023-04-26 12:34 |
| 2  | Comment 2 for Post 1  | 1              | Post                | 2023-04-26 12:34 |
| 3  | Comment 1 for Video 1 | 1              | Video               | 2023-04-26 12:34 |
| 4  | Comment 2 for Video 1 | 1              | Video               | 2023-04-26 12:34 |
+----+------------------------+----------------+---------------------+------------------+
```

posts 表：

```
+----+--------+------------------+
| ID | Title  | created_at       |
+----+--------+------------------+
| 1  | Post 1 | 2023-04-26 12:34 |
+----+--------+------------------+
```

videos 表：

```
+----+--------+------------------+
| ID | Title  | created_at       |
+----+--------+------------------+
| 1  | Video 1| 2023-04-26 12:34 |
+----+--------+------------------+
```

从上面的数据表结构可以看出，comments 表中的 `CommentableID` 字段表示关联的评论所属的记录的 ID，`CommentableType` 字段表示关联的评论所属的记录类型，即 `Post` 或者 `Video`。这种方式实现了多态关联，允许一个表中的记录关联到多个不同表中的记录。
### CommentableType是从哪里获取到的

在使用 GORM 中的多态关联时，`CommentableType` 字段的值是由 GORM 根据模型关联的结构自动设置的。在示例中，`CommentableType` 是由 `Comment` 模型中定义的多态字段 `PolymorphicType` 决定的，其值是关联到的模型的名称。

在示例中，`Comment` 模型中定义了 `PolymorphicType` 字段，其值为 `gorm:"polymorphic:Commentable;"`。这里的 `Commentable` 是一个字符串，表示多态关联的关联名，即关联到的模型的名称。当创建一个关联到 `Comment` 模型的记录时，GORM 会自动将 `CommentableType` 字段设置为 `Post` 或者 `Video`，具体取决于关联的模型是 `Post` 还是 `Video`。这样在数据库中就可以记录下每条评论关联到的记录类型，从而实现多态关联。
### gorm has many 自引用

在 GORM 中，可以通过使用 `ForeignKey` 和 `AssociationForeignKey` 标签来实现模型的自引用关系，从而建立一个模型与自身的 "Has Many" 关联。

下面是一个简单的示例，演示了如何在 GORM 中定义一个自引用的模型关联：

```go
type User struct {
    ID        uint
    Name      string
    ParentID  uint
    Children  []User `gorm:"foreignkey:ParentID"` // 定义子用户的关联
    Parent    *User  `gorm:"foreignkey:ID"`       // 定义父用户的关联
}
```

在上面的例子中，`User` 模型通过 `ParentID` 字段与自身建立了 "Has Many" 关联，表示一个用户可以有多个子用户，同时通过 `Parent` 字段与自身建立了 "Belongs To" 关联，表示一个用户只能有一个父用户。

在数据库中，可以通过添加 `ParentID` 字段来表示用户之间的关系。同时，GORM 会自动解析模型中定义的关联字段，从而在查询和操作数据时可以直接通过关联字段进行操作，例如：

```go
// 查询用户及其子用户
var user User
db.Preload("Children").Find(&user, 1) // 查询 ID 为 1 的用户及其子用户

// 创建父用户及其子用户
parent := User{Name: "Parent"}
child1 := User{Name: "Child 1", Parent: &parent}
child2 := User{Name: "Child 2", Parent: &parent}
db.Create(&parent)
db.Create(&child1)
db.Create(&child2)

// 查询子用户的父用户
var child User
db.Preload("Parent").Find(&child, 1) // 查询 ID 为 1 的子用户及其父用户
```

通过以上示例，可以看到如何在 GORM 中使用 `Has Many` 和 `Belongs To` 关联来实现模型的自引用关系，从而实现模型之间的关联查询和操作。

### 以上示例查询结果

以上示例中的查询结果可能会根据具体的数据库记录和数据模型定义而有所不同。以下是可能的查询结果的示例：

1. 查询用户及其子用户：

```go
var user User
db.Preload("Children").Find(&user, 1)
```

查询 ID 为 1 的用户及其子用户，结果可能如下：

```go
user := User{
    ID:        1,
    Name:      "Parent",
    ParentID:  0, // 或者 null，表示该用户为顶级用户，没有父用户
    Children: []User{
        {ID: 2, Name: "Child 1", ParentID: 1},
        {ID: 3, Name: "Child 2", ParentID: 1},
    },
}
```

2. 查询子用户的父用户：

```go
var child User
db.Preload("Parent").Find(&child, 1)
```

查询 ID 为 1 的子用户及其父用户，结果可能如下：

```go
child := User{
    ID:        1,
    Name:      "Child 1",
    ParentID:  1,
    Parent: &User{
        ID:       1,
        Name:     "Parent",
        ParentID: 0,
    },
}
```

注意：以上示例中的查询结果仅为示意，并可能因实际数据和数据模型的定义而有所不同。实际的查询结果会根据数据库中的数据以及数据模型的字段定义和标签设置而定。

## redis

## 如何组织项目
