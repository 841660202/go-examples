<img src="http://t-blog-images.aijs.top/img/202304241057941.webp" style="width:500px;max-width:100%"/>

## 引用类型和值类型

1. Go 中严格区分引用类型和值类型

## golang append

`append` 是 Go 语言中的内置函数，用于向切片中追加元素。其语法为：

```
append(slice []Type, elems ...Type) []Type
```

其中，`slice` 是要追加元素的切片，`elems` 是要追加的元素，可以是单个元素或一个切片。`append` 函数会返回一个新的切片，其中包含了原有的元素以及追加的元素。

如果追加的元素的数量超过了切片的容量，那么 `append` 函数会重新分配内存，并将原有的元素复制到新的内存中。在重新分配内存后，原有的切片仍然指向旧的内存，而新的切片指向新的内存。

`append` 函数的返回值是一个新的切片，因此，你需要使用新的变量来存储它。例如：

```
// 定义一个切片
slice := []int{1, 2, 3}

// 追加元素
slice = append(slice, 4, 5)

// 输出新的切片
fmt.Println(slice) // [1 2 3 4 5]
```

在上面的例子中，我们首先定义了一个包含三个元素的切片，然后使用 `append` 函数向其中追加两个元素，并将新的切片赋值给原来的变量。最后，我们输出了新的切片，其中包含了原有的元素以及追加的元素。

### 上述内容中原 slice 会被怎么处理，会不会被垃圾回收

在 Go 中，如果 `append` 函数追加元素后，超出了原有切片的容量，那么 `append` 函数会重新分配内存，并将原有的元素复制到新的内存中。此时，原有的切片仍然指向旧的内存，而新的切片指向新的内存。

如果原有的切片不再被使用，那么它会被垃圾回收。垃圾回收器会自动识别不再使用的对象，并将它们从内存中删除。

需要注意的是，如果原有的切片仍然被其它变量所引用，那么它不会被垃圾回收。只有在所有引用它的变量都被销毁后，垃圾回收器才会将其回收。

因此，在使用 `append` 函数时，我们需要注意内存的使用情况，避免产生不必要的内存分配和泄漏。在需要追加元素时，最好提前分配足够的容量，避免频繁的内存分配。同时，当不再需要使用一个切片时，最好将其置为 `nil`，以便垃圾回收器能够及时回收它所占用的内存。

### 分析下 append 的源码

`append` 是 Go 语言内置函数之一，用于向切片中追加元素。下面是 `append` 函数的源码：

```go
func append(slice []Type, elems ...Type) []Type {
    // 获取切片的容量和长度
    var (
        s = len(slice)
        c = cap(slice)
    )
    // 如果追加的元素个数超过了剩余的容量，需要重新分配内存
    if s+len(elems) > c {
        // 分配新的切片，容量为原切片长度的两倍或者追加元素后的长度，取最大值
        newCap := max(2*c, s+len(elems))
        newSlice := makeSlice(len(slice), newCap)
        // 复制原有的元素到新的切片中
        copy(newSlice, slice)
        // 将新的切片赋值给原变量，以实现扩容
        slice = newSlice
    }
    // 将新的元素追加到切片中
    slice = slice[:s+len(elems)]
    for i, x := range elems {
        slice[s+i] = x
    }
    return slice
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func makeSlice(oldLen, newCap int) []Type {
    // 分配新的内存
    newPtr := mallocgc(newCap * sizeof(Type), nil, false)
    // 将原有的元素复制到新的内存中
    if raceenabled || msanenabled {
        // 开启竞态检测和内存检测
        typ := sliceType(Type{})
        oldPtr := mallocgc(oldLen*int(typ.size), typ, false)
        memmove(newPtr, oldPtr, uintptr(oldLen)*typ.size)
        if raceenabled {
            racewriterangepc(unsafe.Pointer(newPtr), uintptr(oldLen)*typ.size, getcallerpc(), funcPC(append))
        }
        if msanenabled {
            msanwrite(unsafe.Pointer(newPtr), uintptr(oldLen)*typ.size)
        }
        typedmemmove(typ, newPtr, unsafe.Pointer(&oldLen), unsafe.Pointer(&newCap))
        return *(*[]Type)(unsafe.Pointer(&SliceHeader{
            Data: newPtr,
            Len:  oldLen,
            Cap:  newCap,
        }))
    }
    // 普通情况下，直接复制元素到新的内存中
    typ := sliceType(Type{})
    newSlice := *(*[]Type)(unsafe.Pointer(&SliceHeader{
        Data: newPtr,
        Len:  oldLen,
        Cap:  newCap,
    }))
    if needmem(typ) {
        // 需要进行类型转换
        for i := range newSlice {
            typedmemmove(typ, unsafe.Pointer(&newSlice[i]), unsafe.Pointer(&slice[i]), nil)
        }
    } else {
        // 直接复制元素
        memmove(newPtr, data(slice), uintptr(len(slice))*typ.size)
    }
    return newSlice
}

func data(slice []Type) unsafe.Pointer {
    return unsafe.Pointer(&slice[0])
}

```

`append` 函数的实现比较复杂，主要分为以下几个步骤：

1. 获取切片的长度和容量。

   在函数一开始，通过 `len` 和 `cap` 内置函数获取原始切片的长度和容量。

2. 判断是否需要扩容。

   如果新的元素个数超过了原切片的剩余容量，则需要扩容。这时会重新分配一个新的切片，容量为原切片长度的两倍或者追加元素后的长度，取两者中的最大值。

   ```go
   if s+len(elems) > c {
       // 分配新的切片，容量为原切片长度的两倍或者追加元素后的长度，取最大值
       newCap := max(2*c, s+len(elems))
       newSlice := makeSlice(len(slice), newCap)
       // 复制原有的元素到新的切片中
       copy(newSlice, slice)
       // 将新的切片赋值给原变量，以实现扩容
       slice = newSlice
   }
   ```

   这里调用了 `max` 函数，用于计算新的容量。然后调用 `makeSlice` 函数分配新的内存，并将原有元素复制到新的切片中，最后将新的切片赋值给原变量。

3. 追加新的元素。

   在新的容量可以满足追加元素的情况下，就可以将新的元素直接追加到切片的尾部。这里使用了 `slice` 切片的切片操作，将 `slice` 的长度扩展到 `s+len(elems)`，然后使用 `for` 循环将新的元素一个一个追加到切片的尾部。

   ```go
   // 将新的元素追加到切片中
   slice = slice[:s+len(elems)]
   for i, x := range elems {
       slice[s+i] = x
   }
   ```

4. 返回新的切片。

   最后返回的是新的切片，其中包含了追加的新元素。

关于内存分配和释放的问题，Go 语言使用了垃圾回收器来自动管理内存，因此在 `append` 函数中不需要手动释放旧的切片内存空间。当旧的切片被扩容时，Go 语言的垃圾回收器会自动回收旧的内存空间。

### 上述源码中为什么要 开启竞态检测和内存检测

在 `append` 函数中，开启了竞态检测和内存检测，是因为这两种检测可以帮助我们在编写程序时发现一些潜在的错误和问题，从而提高程序的健壮性和可靠性。

竞态检测（Race Detector）是 Go 语言内置的一个工具，它可以用于检测多个 goroutine 之间的竞态条件，比如数据竞态、锁竞争等问题。在 `append` 函数中，可能存在多个 goroutine 对同一个切片进行并发的追加操作，这时如果不进行竞态检测，就有可能出现数据竞态等问题，导致程序出错或者崩溃。

内存检测（Memory Sanitizer）是 GCC、Clang 等编译器中的一个功能，它可以用于检测程序中的内存访问越界、未初始化内存、使用已释放内存等问题。在 `append` 函数中，如果没有开启内存检测，就可能出现一些内存访问越界等问题，导致程序出错或者崩溃。

因此，在开发过程中，开启竞态检测和内存检测可以帮助我们及时发现这些问题，从而及时修复，提高程序的稳定性和可靠性。当然，开启这些检测也会带来一些性能上的损失，因此在生产环境中一般不会开启这些检测。

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

## gorm

## redis

## 如何组织项目
