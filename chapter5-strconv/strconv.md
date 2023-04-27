## strconv

strconv 是 Go 语言标准库中的一个包，提供了一些常用的字符串和基本数据类型之间的转换函数。它的名称源自“string conversion”，也就是字符串转换的意思。

在 Go 语言中，字符串和基本数据类型之间的转换是一种常见的操作。strconv 包提供了将字符串转换为整数、浮点数、布尔值等基本类型的函数，也提供了将基本类型转换为字符串的函数。

例如，strconv.Atoi 函数可以将字符串转换为 int 类型，strconv.ParseFloat 函数可以将字符串转换为 float64 类型，strconv.FormatBool 函数可以将 bool 类型转换为字符串类型等等。此外，strconv 还提供了一些处理 Unicode 字符的函数，如将 Unicode 字符编码为 UTF-8 字节序列和将 UTF-8 字节序列解码为 Unicode 字符等。

以下是一些 strconv 包中常用的函数：

1. strconv.Atoi(s string) (int, error)：将字符串转换为 int 类型。
2. strconv.ParseFloat(s string, bitSize int) (float64, error)：将字符串转换为 float64 类型。
3. strconv.FormatBool(b bool) string：将 bool 类型转换为字符串类型。
4. strconv.Quote(s string) string：将字符串 s 转义并添加引号。
5. strconv.Unquote(s string) (string, error)：将经过引号引用和转义的字符串还原为原始字符串。
6. strconv.QuoteToASCII(s string) string：将字符串转义为 ASCII 编码的字符串。
7. strconv.QuoteRune(r rune) string：将单个 Unicode 字符转义并添加引号。
8. strconv.AppendInt(dst []byte, i int64, base int) []byte：将 int64 类型转换为指定进制的字节切片，添加到 dst 中。

需要注意的是，在使用这些函数时要注意错误处理，因为输入的字符串可能不是合法的数字或格式不正确。

除了上述提到的常用函数，strconv 还提供了其他一些函数，下面列举一些常用的函数：

1. strconv.Itoa(i int) string：将 int 类型转换为字符串类型。
2. strconv.ParseInt(s string, base int, bitSize int) (int64, error)：将字符串转换为 int64 类型，base 指定进制（如 2、8、10 或 16），bitSize 指定结果的位数（如 32 或 64）。
3. strconv.ParseUint(s string, base int, bitSize int) (uint64, error)：将字符串转换为无符号整数（uint64 类型），其他参数和 ParseInt 函数类似。
4. strconv.FormatInt(i int64, base int) string：将 int64 类型转换为字符串类型，base 指定进制（如 2、8、10 或 16）。
5. strconv.FormatUint(i uint64, base int) string：将无符号整数（uint64 类型）转换为字符串类型，其他参数和 FormatInt 函数类似。
6. strconv.ParseBool(str string) (bool, error)：将字符串转换为 bool 类型，字符串只能是"true"或"false"。
7. strconv.ParseFloat(s string, bitSize int) (float64, error)：将字符串转换为 float64 类型，bitSize 指定结果的位数（如 32 或 64）。

总之，strconv 是 Go 语言中非常常用的包之一，它提供了很多方便的函数来进行字符串和基本数据类型之间的转换。在使用这些函数时，需要注意输入参数的正确性和错误处理，以确保程序能够正确地运行。

## 例子

好的，下面举几个 strconv 包中的函数，并给出一个简单的示例：

### 1. strconv.Atoi(s string) (int, error)

这个函数将字符串 s 转换为 int 类型，如果转换失败会返回一个错误。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "123"
    i, err := strconv.Atoi(s)
    if err != nil {
        fmt.Println("Conversion failed:", err)
        return
    }
    fmt.Printf("String '%s' converted to int %d\n", s, i)
}
```

输出：

```go
String '123' converted to int 123
```

### 2. strconv.FormatBool(b bool) string

这个函数将 bool 类型转换为字符串类型。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    b := true
    s := strconv.FormatBool(b)
    fmt.Printf("Bool %t converted to string '%s'\n", b, s)
}
```

输出：

```go
Bool true converted to string 'true'
```

### 3. strconv.Quote(s string) string

这个函数将字符串 s 转义并添加引号。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "Hello, world!"
    q := strconv.Quote(s)
    fmt.Printf("String '%s' quoted as '%s'\n", s, q)
}
```

输出：

```go
String 'Hello, world!' quoted as '"Hello, world!"'
```

### 4. strconv.ParseInt(s string, base int, bitSize int) (int64, error)

这个函数将字符串 s 转换为 int64 类型，base 指定进制（如 2、8、10 或 16），bitSize 指定结果的位数（如 32 或 64）。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "100"
    i, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
        fmt.Println("Conversion failed:", err)
        return
    }
    fmt.Printf("String '%s' converted to int64 %d\n", s, i)
}
```

输出：

```go
String '100' converted to int64 100
```

这里将字符串"100"转换为 int64 类型，并使用 10 进制表示。

### 5. strconv.AppendInt(dst []byte, i int64, base int) []byte

这个函数将 int64 类型的整数 i 转换为字符串类型，并将结果追加到 byte 类型的切片 dst 中，base 指定进制（如 2、8、10 或 16）。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var dst []byte
    i := int64(100)
    dst = strconv.AppendInt(dst, i, 10)
    fmt.Printf("Int64 %d appended to byte slice as string '%s'\n", i, dst)
}
```

输出：

```go
Int64 100 appended to byte slice as string '100'
```

这里将 int64 类型的整数 100 转换为字符串类型，并将结果追加到空的 byte 类型切片中。

### 6. strconv.IsPrint(r rune) bool

这个函数判断 rune 类型的字符是否是可打印字符。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    r := rune('\t')
    if strconv.IsPrint(r) {
        fmt.Printf("Rune '%c' is printable\n", r)
    } else {
        fmt.Printf("Rune '%c' is not printable\n", r)
    }
}
```

输出：

```
Rune '    ' is not printable
```

这里判断 tab 字符是否是可打印字符。因为 tab 字符不是可打印字符，所以输出结果是"Rune ' ' is not printable"。

### 7. strconv.ParseFloat(s string, bitSize int) (float64, error)

这个函数将字符串 s 转换为 float64 类型，bitSize 指定结果的位数（如 32 或 64）。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "3.14"
    f, err := strconv.ParseFloat(s, 64)
    if err != nil {
        fmt.Println("Conversion failed:", err)
        return
    }
    fmt.Printf("String '%s' converted to float64 %f\n", s, f)
}
```

输出：

```
String '3.14' converted to float64 3.140000
```

这里将字符串"3.14"转换为 float64 类型。

### 8. strconv.QuoteToASCII(s string) string

这个函数将字符串 s 转义并添加 ASCII 码的引号。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "こんにちは"
    q := strconv.QuoteToASCII(s)
    fmt.Printf("String '%s' quoted as ASCII '%s'\n", s, q)
}
```

输出：

```
String 'こんにちは' quoted as ASCII '"\u3053\u3093\u306B\u3061\u306F"'
```

这里将字符串"こんにちは"转义并添加 ASCII 码的引号。

### 9. strconv.ParseBool(str string) (bool, error)

这个函数将字符串 str 解析为 bool 类型的值。例如：

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "true"
	s2 := "false"
	s3 := "True"
	s4 := "False"
	s5 := "abc"

	b1, err1 := strconv.ParseBool(s1)
	fmt.Printf("%v %v\n", b1, err1)

	b2, err2 := strconv.ParseBool(s2)
	fmt.Printf("%v %v\n", b2, err2)

	b3, err3 := strconv.ParseBool(s3)
	fmt.Printf("%v %v\n", b3, err3)

	b4, err4 := strconv.ParseBool(s4)
	fmt.Printf("%v %v\n", b4, err4)

	b5, err5 := strconv.ParseBool(s5)
	fmt.Printf("%v %v\n", b5, err5)
}
```

输出：

```
true <nil>
false <nil>
true <nil>
false <nil>
false strconv.ParseBool: parsing "abc": invalid syntax
```

这里分别将字符串"true"、"false"、"True"、"False"、"abc"解析为 bool 类型的值，前四个可以成功转换，最后一个无法转换。

### 10. strconv.Unquote(s string) (string, error)

这个函数将一个被引号包围的字符串 s 去掉引号并解码转义字符。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "\"hello\\nworld\\t\""
    u, err := strconv.Unquote(s)
    if err != nil {
        fmt.Println("Unquoting failed:", err)
        return
    }
    fmt.Printf("String '%s' unquoted as '%s'\n", s, u)
}
```

输出：

```
String '"hello\nworld\t"' unquoted as 'hello
world	'
```

这里将一个被引号包围的字符串"\"hello\\nworld\\t\""去掉引号并解码转义字符。

### 11. strconv.AppendInt(dst []byte, i int64, base int) []byte

这个函数将 int64 类型的整数 i 转换为指定进制 base 的字符串，并将其附加到 dst 字节片中。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    i := int64(123456)
    dst := make([]byte, 0)
    dst = strconv.AppendInt(dst, i, 10)
    fmt.Printf("Int64 %d converted to string %s\n", i, string(dst))
}
```

输出：

```go
Int64 123456 converted to string 123456
```

这里将 int64 类型的整数 123456 转换为十进制的字符串，并将其附加到空的字节片中。

### 12. strconv.AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte

这个函数将 float64 类型的浮点数 f 转换为指定格式 fmt 的字符串，并将其附加到 dst 字节片中。prec 指定精度，bitSize 指定结果的位数（如 32 或 64）。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    f := 3.14159
    dst := make([]byte, 0)
    dst = strconv.AppendFloat(dst, f, 'f', 2, 64)
    fmt.Printf("Float64 %f converted to string %s\n", f, string(dst))
}
```

输出：

```go
Float64 3.141590 converted to string 3.14
```

这里将 float64 类型的浮点数 3.14159 转换为小数点后两位的字符串，并将其附加到空的字节片中。

### 13. strconv.Quote(s string) string

这个函数将字符串 s 包围在引号中并返回其转义版本。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "hello\nworld\t"
    q := strconv.Quote(s)
    fmt.Printf("String '%s' quoted as '%s'\n", s, q)
}
```

输出：

```go
String 'hello
world	' quoted as '"hello\nworld\t"'
```

这里将字符串"hello\nworld\t"包围在引号中并返回其转义版本。

### 14. strconv.QuoteToASCII(s string) string

这个函数将字符串 s 包围在引号中并返回其只包含 ASCII 字符的转义版本。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "你好，世界"
    q := strconv.QuoteToASCII(s)
    fmt.Printf("String '%s' quoted to ASCII as '%s'\n", s, q)
}
```

输出：

```go
String '你好，世界' quoted to ASCII as '"\u4f60\u597d\uff0c\u4e16\u754c"'
```

这里将字符串"你好，世界"包围在引号中并返回其只包含 ASCII 字符的转义版本。

### 15. strconv.CanBackquote(s string) bool

这个函数报告字符串 s 是否可以用反引号包围，而不需要转义任何字符。例如：

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s1 := "hello\nworld"
    s2 := "`hello\nworld`"
    fmt.Printf("String '%s' can be backquoted: %t\n", s1, strconv.CanBackquote(s1))
    fmt.Printf("String '%s' can be backquoted: %t\n", s2, strconv.CanBackquote(s2))
}
```

输出：

```go
String 'hello
world' can be backquoted: false
String '`hello\nworld`' can be backquoted: true
```

这里报告字符串"hello\nworld"不能用反引号包围而不需要转义任何字符，而字符串"`hello\nworld`"可以。

## Atoi、 Itoa

<a href="https://www.zhihu.com/question/266044920" target="_blank" >golang 的 API 设计的这么烂？</a>

1. atoi：ASCII to integer,将字符串转换成整形，从数字或正负号开始转换，一直到非数字为止
2. itoa：integer to ASCII--将整形转换成字符串
3. atof：ascii to float--字符串转换成浮点型
4. atol：ascii to long---字符串转换成长整形 gcvt：浮点型转换成字符串
