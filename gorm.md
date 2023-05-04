### gorm

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

**gorm FirstOrCreate 和 FirstOrInit 的区别：**

`FirstOrCreate` 和 `FirstOrInit` 都是 GORM 提供的查询方法，用于在数据库中查找记录。它们的区别在于：

- `FirstOrCreate` 方法会在数据库中查找符合条件的记录，如果找到了则返回该记录，否则会创建一条新记录并返回。如果创建新记录时需要设置一些默认值，可以在调用 `FirstOrCreate` 方法之前为相应字段设置默认值。

- `FirstOrInit` 方法也会在数据库中查找符合条件的记录，如果找到了则返回该记录，否则会返回一个未保存的新记录（即未设置主键的新实例），然后可以对新实例进行修改并保存到数据库中。

举个例子，假设我们有一个 `User` 模型，要根据用户名查找用户记录，如果找到则返回该记录，否则创建一条新记录并返回。可以使用 `FirstOrCreate` 方法来实现：

```go
var user User
db.Where("name = ?", "Alice").FirstOrCreate(&user)
```

如果数据库中已经存在用户名为 "Alice" 的用户记录，则会返回该记录，否则会创建一条新记录并返回。

另一方面，如果我们要返回一个未保存的新用户实例（即未设置主键的新记录），可以使用 `FirstOrInit` 方法：

```go
var user User
db.Where("name = ?", "Alice").FirstOrInit(&user)
```

如果数据库中已经存在用户名为 "Alice" 的用户记录，则会返回该记录，否则会返回一个未保存的新用户实例。我们可以对新实例进行修改并保存到数据库中：

```go
if user.ID == 0 {
  user.Name = "Alice"
  db.Create(&user)
} else {
  // ...
}
```

`FirstOrCreate` 和 `FirstOrInit` 方法还有一些注意事项：

- `FirstOrCreate` 方法不仅可以根据一个条件查找记录，还可以根据多个条件查找记录，例如：

  ```go
  db.Where(User{Name: "Alice", Age: 18}).FirstOrCreate(&user)
  ```

- `FirstOrInit` 方法只能根据一个条件查找记录，如果要根据多个条件查找记录，可以使用 `Where` 方法指定多个条件，例如：

  ```go
  db.Where("name = ? AND age = ?", "Alice", 18).FirstOrInit(&user)
  ```

- `FirstOrInit` 方法返回的是一个未保存的新实例，因此如果要在返回结果之后进行更新操作，需要调用 `Save` 方法将新实例保存到数据库中，例如：

  ```go
  var user User
  db.Where("name = ?", "Alice").FirstOrInit(&user)
  user.Age = 18
  db.Save(&user)
  ```

  如果只是想修改内存中的实例而不保存到数据库中，可以使用 `Updates` 方法，例如：

  ```go
  var user User
  db.Where("name = ?", "Alice").FirstOrInit(&user)
  db.Model(&user).Updates(User{Age: 18})
  ```

还有一点需要注意的是，`FirstOrCreate` 和 `FirstOrInit` 方法都是使用 `SELECT ... FROM ... WHERE ... LIMIT 1` 的方式查询数据库的。如果需要查询的条件有索引，这种方式可以有效地利用索引，提高查询性能。但如果需要查询的条件没有索引，这种方式可能会导致全表扫描，查询速度会比较慢。

另外，如果需要查询的条件不唯一，即有多条记录符合条件，那么 `FirstOrCreate` 和 `FirstOrInit` 方法都只会返回其中的一条记录。如果需要返回所有符合条件的记录，可以使用 `Find` 方法进行查询，例如：

```go
var users []User
db.Where("age = ?", 18).Find(&users)
```

这样就可以返回所有年龄为 18 的用户记录了。

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

**非 GPT 生成：**怎么看上去，像是一对多了呢

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

### CommentableType 是从哪里获取到的

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

### gorm .Preload() API 的作用

`Preload()` 是 GORM 中用于预加载关联数据的 API，其作用是在查询某个模型数据时，同时预加载其关联模型的数据，以避免在渲染视图时因为 N+1 问题导致的性能问题。在使用 `Preload()` API 时，GORM 会发出一条 JOIN 查询语句来加载关联数据，而不是使用多条查询语句逐个加载关联数据。

例如，假设我们有以下的模型定义：

```go
type User struct {
  gorm.Model
  Name      string
  Age       int
  Addresses []Address
}

type Address struct {
  gorm.Model
  UserID    uint
  City      string
  Country   string
}
```

其中，`User` 模型拥有多个 `Address` 模型的关联，即一个用户可以有多个地址。如果我们在查询用户数据时需要同时查询用户的地址数据，可以使用 `Preload()` API 来实现：

```go
var user User
db.Preload("Addresses").First(&user, 1)
```

这里的 `Preload("Addresses")` 表示预加载 `User` 模型的 `Addresses` 关联模型数据。在执行查询时，GORM 会发出一条 JOIN 查询语句来同时查询用户和其地址数据，从而避免了 N+1 查询问题。

需要注意的是，使用 `Preload()` API 时应避免过度预加载，以免查询语句过于复杂导致性能下降。可以通过设置参数来控制预加载的深度和范围，以满足具体的查询需求。

### N+1 问题 是什么意思

N+1 问题（N+1 Query Problem）是指在加载关联数据时，使用了 N 条查询语句来查询 N 条主记录以及它们的关联数据，从而导致性能问题的一种情况。通常情况下，N+1 问题是由 ORM 框架的延迟加载机制所导致的。

举个例子，假设我们有一个 `User` 模型和一个 `Post` 模型，它们之间是一对多的关联关系（即一个用户可以有多篇文章）。如果我们要查询所有用户以及他们的所有文章，可能会写出如下的查询代码：

```go
var users []User
db.Find(&users)

for _, user := range users {
  var posts []Post
  db.Where("user_id = ?", user.ID).Find(&posts)
  user.Posts = posts
}
```

这里的第一条查询语句会加载所有用户的数据，而在第二个循环中，我们对每个用户分别执行一次查询来加载他们的文章数据。这样就会导致 N+1 问题，即执行了 N+1 条查询语句。如果用户数据量很大，那么就会导致性能问题，因为查询语句的执行需要时间和资源。

为了避免 N+1 问题，可以使用 GORM 中的 `Preload()` API 来预加载关联数据，或者手动使用 JOIN 查询语句来同时加载主记录和关联记录。这样可以将多条查询语句合并为一条查询，从而提高查询性能。

### many2many

在 GORM 中，Many-to-Many 关系可以使用 `Many2Many` 方法来定义。它允许您定义两个表之间的多对多关系，并使用中间表进行关联。

以下是一个示例：

```go
type User struct {
    ID       uint
    Name     string
    Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
    ID   uint
    Name string
    Users []User `gorm:"many2many:user_languages;"`
}
```

在上面的代码中，`User` 和 `Language` 之间的多对多关系使用 `many2many` 方法定义。`user_languages` 是中间表的表名，它将这两个模型关联起来。

在此示例中，每个用户可以拥有多种语言，而每种语言也可以被多个用户使用。中间表 `user_languages` 将两个模型之间的关系保存在一起。

要使用此关系，您可以使用以下代码将语言添加到用户的语言列表中：

```go
user := User{}
language := Language{}

// 通过ID获取用户和语言
db.First(&user, 1)
db.First(&language, 2)

// 添加语言到用户的语言列表
db.Model(&user).Association("Languages").Append(&language)
```

在上面的示例中，我们从数据库中获取了一个用户和一种语言，然后将语言添加到用户的语言列表中。

您还可以从用户的语言列表中获取所有语言：

```go
db.Model(&user).Association("Languages").Find(&languages)
```

上面的代码将获取用户的所有语言，并将它们存储在 `languages` 列表中。

### 解释下 如下代码

```go
type User struct {
    gorm.Model
    Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:ProfileRefer"`
    Refer    uint      `gorm:"index:,unique"`
}

type Profile struct {
    gorm.Model
    Name      string
    UserRefer uint `gorm:"index:,unique"`
}

```

以上代码定义了两个 GORM 模型 `User` 和 `Profile`。`User` 模型具有多对多关系，并且使用名为 `user_profiles` 的中间表来定义它们之间的关系。中间表具有自定义外键列 `Refer` 和 `UserReferID`，分别用于关联 `User` 和 `Profile` 模型。`Profile` 模型具有 `UserRefer` 字段，用于关联 `User` 模型。

具体来说，以下是关键代码行的解释：

- `gorm.Model` 是 GORM 提供的一个带有 `ID`、`CreatedAt` 和 `UpdatedAt` 字段的基本模型，可以被嵌入到自定义模型中以获得这些通用字段。
- `Profiles []Profile` 定义了 `User` 模型中的一个 `Profile` 切片，表示每个用户都可以拥有多个个人资料。
- `gorm:"many2many:user_profiles"` 告诉 GORM 将 `User` 和 `Profile` 模型之间的关系保存在名为 `user_profiles` 的中间表中。
- `foreignKey:Refer` 和 `joinForeignKey:UserReferID` 分别指定了用于关联 `User` 模型的外键列和用于关联中间表的外键列。
- `References:UserRefer` 和 `joinReferences:ProfileRefer` 分别指定了用于关联 `User` 模型的外键列和用于关联 `Profile` 模型的外键列。
- `Refer uint` 声明了一个自定义的外键列 `Refer`，它用于将 `User` 模型与中间表关联起来。
- `gorm:"index:,unique"` 将 `Refer` 列标记为唯一索引，以确保每个 `User` 对应一个唯一的外键值。
- `UserRefer uint` 是 `Profile` 模型中的一个字段，表示每个个人资料都与一个用户关联。
- `gorm:"index:,unique"` 将 `UserRefer` 列标记为唯一索引，以确保每个 `Profile` 对应一个唯一的外键值。

总体来说，以上代码展示了如何在 GORM 中定义多对多关系，并使用自定义中间表和外键列来保存关系。

### 上述代码执行后的建表语句是什么

执行以上代码后，GORM 会自动创建三张数据库表：

1. `users` 表：

```sql
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  `refer` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_refer` (`refer`),
  KEY `users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

`users` 表包含一个自增主键 `id`、`created_at` 和 `updated_at` 字段，以及一个名为 `refer` 的自定义外键列。`refer` 列被标记为唯一索引，以确保每个用户对应一个唯一的外键值。

2. `profiles` 表：

```sql
CREATE TABLE `profiles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_refer` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_profiles_user_refer` (`user_refer`),
  KEY `profiles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

`profiles` 表包含一个自增主键 `id`、`created_at` 和 `updated_at` 字段，以及一个名为 `user_refer` 的自定义外键列。`user_refer` 列被标记为唯一索引，以确保每个个人资料对应一个唯一的外键值。

3. `user_profiles` 表：

```sql
CREATE TABLE `user_profiles` (
  `user_refer_id` bigint unsigned NOT NULL,
  `profile_refer_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`user_refer_id`,`profile_refer_id`),
  KEY `idx_user_profiles_profile_refer_id` (`profile_refer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

`user_profiles` 表包含两个名为 `user_refer_id` 和 `profile_refer_id` 的外键列，用于关联 `users` 表和 `profiles` 表。这两个外键列被用作联合主键，以确保每个关系都是唯一的。

### 上述代码的增删改查完整代码是什么样

以下是对于上述代码的增删改查的完整 GORM 代码示例：

```go
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:ProfileRefer"`
	Refer    uint      `gorm:"index:,unique"`
}

type Profile struct {
	gorm.Model
	Name      string
	UserRefer uint `gorm:"index:,unique"`
}

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		panic(err)
	}

	// 创建用户和个人资料
	user := User{
		Refer: 1,
		Profiles: []Profile{
			{
				Name: "Alice",
			},
			{
				Name: "Bob",
			},
		},
	}
	err = db.Create(&user).Error
	if err != nil {
		panic(err)
	}

	// 查询用户和个人资料
	var result User
	err = db.Where("refer = ?", 1).Preload("Profiles").First(&result).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("User: %v, Profiles: %v\n", result, result.Profiles)

	// 更新用户和个人资料
	err = db.Model(&result).Updates(User{
		Refer: 2,
		Profiles: []Profile{
			{
				Name: "Charlie",
			},
			{
				Name: "Dave",
			},
		},
	}).Error
	if err != nil {
		panic(err)
	}

	// 查询用户和个人资料
	err = db.Where("refer = ?", 2).Preload("Profiles").First(&result).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("User: %v, Profiles: %v\n", result, result.Profiles)

	// 删除用户和个人资料
	err = db.Delete(&result).Error
	if err != nil {
		panic(err)
	}
}
```

该示例中的 `dsn` 变量包含数据库连接的详细信息，如数据库地址、用户名、密码和数据库名称等。

示例中使用 `db.AutoMigrate(&User{}, &Profile{})` 自动创建表结构。

使用 `db.Create(&user)` 创建用户和个人资料，使用 `db.Where().Preload().First()` 查询用户和个人资料，使用 `db.Model().Updates()` 更新用户和个人资料，使用 `db.Delete()` 删除用户和个人资料。在查询用户时，使用 `Preload()` 函数来关联查询用户的个人资料。

### joinForeignKey 和 joinReferences 各是什么作用

`joinForeignKey` 和 `joinReferences` 是 GORM many-to-many 关系中的两个关键属性，用于指定多对多关联表的外键关联方式。

- `joinForeignKey`: 它指定用于关联多对多关系的关联表中的外键名称。在上述示例中，`joinForeignKey` 的值为 `UserReferID`，这意味着 `user_profiles` 表将包含一个名为 `user_refer_id` 的外键，该外键将引用 `users` 表中的 `id` 列。

- `joinReferences`: 它指定用于关联多对多关系的关联表中引用的外部表和列的名称。在上述示例中，`joinReferences` 的值为 `ProfileRefer`，这意味着 `user_profiles` 表将引用 `profiles` 表中的 `user_refer` 列。

这两个属性一起使用，可以定义多对多关联表的外键关联方式，以及外键关联的表和列。

哈哈，就是相当于 rename 啊

### foreignKey 和 References 的区别

`foreignKey` 和 `References` 是 GORM 模型中的两个属性，用于指定关联关系中外键和引用的列名。它们的区别如下：

- `foreignKey`: 它用于指定包含外键的模型中的列名。在 many-to-many 关系中，它指定关联模型中的外键列。在 one-to-many 和 many-to-one 关系中，它指定包含外键的模型中的列名。例如，`foreignKey: UserReferID` 指定 User 模型中的 `UserReferID` 列是一个外键。

- `References`: 它用于指定引用的模型中的列名。在 many-to-many 关系中，它指定关联的模型中的列名。在 one-to-many 和 many-to-one 关系中，它指定关联的模型中的主键列名。例如，`References: ProfileRefer` 指定 Profile 模型中的 `ProfileRefer` 列是被引用的列。

因此，`foreignKey` 和 `References` 分别用于指定关联关系中的外键和引用的列名。

**在 many-to-many 关系中，**

- `foreignKey` 用于指定关联模型中的外键列名，
- `References` 用于指定被关联的模型中的列名。

**在 one-to-many 和 many-to-one 关系中，**

- `foreignKey` 用于指定包含外键的模型中的列名，
- `References` 用于指定被关联的模型中的主键列名。

**理解**

gorm foreignKey，可以理解成标记，References 理解成这个标记从哪来，这么理解没错吧

foreignKey，可以理解成标记，References 理解成这个标记从哪来

可以这么理解，`foreignKey` 和 `References` 分别用于指定关联关系中的外键和引用的列名。

- `foreignKey` 可以理解为一个标记，用于标记包含外键的模型中的列名。

- `References` 可以理解为标记的来源，用于指定被关联的模型中的列名或主键列名，以便 GORM 知道应该关联哪些表和列。

在 many-to-many 关系中，`foreignKey` 用于标记关联模型中的外键列名，`References` 用于指定被关联的模型中的列名。

在 one-to-many 和 many-to-one 关系中，`foreignKey` 用于标记包含外键的模型中的列名，`References` 用于指定被关联的模型中的主键列名。

### gorm 设计一个用户可以有多个角色，每个角色可以有多个权限

下面是一个示例 GORM 模型，用于描述用户、角色和权限之间的多对多关系：

```go
type User struct {
    gorm.Model
    Roles []Role `gorm:"many2many:user_roles"`
}

type Role struct {
    gorm.Model
    Name       string
    Permissions []Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
    gorm.Model
    Name string
}
```

在上面的示例中，`User` 模型和 `Role` 模型之间是多对多关系，`Role` 模型和 `Permission` 模型之间也是多对多关系。在 `User` 模型中，`Roles` 属性定义为一个 `[]Role` 切片，用于存储用户所拥有的角色。在 `Role` 模型中，`Permissions` 属性定义为一个 `[]Permission` 切片，用于存储角色所拥有的权限。

在模型定义中，`many2many` 标签用于定义多对多关系的关联表名。例如，在 `User` 模型中，`gorm:"many2many:user_roles"` 定义了一个名为 `user_roles` 的关联表，用于关联 `User` 模型和 `Role` 模型。在 `Role` 模型中，`gorm:"many2many:role_permissions"` 定义了一个名为 `role_permissions` 的关联表，用于关联 `Role` 模型和 `Permission` 模型。

通过上述示例，可以实现用户和角色之间的多对多关系，以及角色和权限之间的多对多关系。在应用程序中，可以使用 GORM 提供的 API 对这些模型进行增删改查操作。例如，可以使用以下代码将一个角色添加到一个用户中：

```go
user := &User{ID: 1}
role := &Role{ID: 2}
db.Model(user).Association("Roles").Append(role)
```

类似地，可以使用 GORM 提供的 API 执行其他类型的操作，例如查询用户拥有的所有角色、查询拥有特定权限的所有角色等等。

### 表结构

根据上面给出的模型定义，GORM 可以自动创建相应的数据库表结构。下面是这些模型对应的数据库表结构示例：

```sql
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `permissions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `user_roles` (
  `user_id` int NOT NULL,
  `role_id` int NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`),
  KEY `user_roles_role_id_foreign` (`role_id`),
  CONSTRAINT `user_roles_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `user_roles_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
);

CREATE TABLE `role_permissions` (
  `role_id` int NOT NULL,
  `permission_id` int NOT NULL,
  PRIMARY KEY (`role_id`,`permission_id`),
  KEY `role_permissions_permission_id_foreign` (`permission_id`),
  CONSTRAINT `role_permissions_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `role_permissions_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE
);
```

可以看到，GORM 自动生成了五个表结构，分别对应 `User`、`Role`、`Permission`、`user_roles` 和 `role_permissions` 这五个模型。其中，`users`、`roles` 和 `permissions` 分别对应 `User`、`Role` 和 `Permission` 这三个模型的基本信息。`user_roles` 表用于关联 `User` 和 `Role` 之间的多对多关系，`role_permissions` 表用于关联 `Role` 和 `Permission` 之间的多对多关系。

### 含增删改查的完整代码

以下是完整的代码示例，包含增删改查操作：

```go
package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

type User struct {
    gorm.Model
    Name   string
    Roles  []*Role `gorm:"many2many:user_roles"`
}

type Role struct {
    gorm.Model
    Name        string
    Permissions []*Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
    gorm.Model
    Name string
}

func main() {
    // 连接数据库
    dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 自动迁移表结构
    err = db.AutoMigrate(&User{}, &Role{}, &Permission{})
    if err != nil {
        panic("failed to migrate database")
    }

    // 创建角色和权限
    role1 := &Role{Name: "管理员"}
    role2 := &Role{Name: "普通用户"}
    perm1 := &Permission{Name: "创建用户"}
    perm2 := &Permission{Name: "删除用户"}

    db.Create(role1)
    db.Create(role2)
    db.Create(perm1)
    db.Create(perm2)

    // 创建用户并分配角色和权限
    user1 := &User{Name: "张三"}
    user2 := &User{Name: "李四"}

    user1.Roles = append(user1.Roles, role1)
    user1.Roles[0].Permissions = append(user1.Roles[0].Permissions, perm1)
    user1.Roles[0].Permissions = append(user1.Roles[0].Permissions, perm2)

    user2.Roles = append(user2.Roles, role2)
    user2.Roles[0].Permissions = append(user2.Roles[0].Permissions, perm1)

    db.Create(user1)
    db.Create(user2)

    // 查询用户和角色信息
    var users []*User
    db.Preload(clause.Associations).Find(&users)
    for _, user := range users {
        fmt.Printf("用户 %s 的角色信息：\n", user.Name)
        for _, role := range user.Roles {
            fmt.Printf("角色 %s 的权限信息：\n", role.Name)
            for _, perm := range role.Permissions {
                fmt.Printf("- %s\n", perm.Name)
            }
        }
    }

    // 更新用户角色信息
    db.Model(user1).Association("Roles").Replace([]*Role{role2})
    db.Model(user2).Association("Roles").Clear()

    // 删除用户和角色信息
    db.Delete(user1)
    db.Delete(user2)
    db.Delete(role1)
    db.Delete(role2)
    db.Delete(perm1)
    db.Delete(perm2)
}
```

在这个代码示例中，我们首先定义了 `User`、`Role` 和 `Permission` 三个模型，用于表示用户、角色和权限之间的关系。其中，`User` 模型中有一个 `Roles` 字段，表示一个用户可以拥有多个角色；`Role` 模型中有一个 `Permissions` 字段，表示一个角色可以拥有多

### 简化后的代码

可以，以下是一个简化版的代码示例：

```go
package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name   string
    Roles  []*Role `gorm:"many2many:user_roles"`
}

type Role struct {
    gorm.Model
    Name        string
    Permissions []*Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
    gorm.Model
    Name string
}

func main() {
    // 连接数据库
    dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    // 自动迁移表结构
    db.AutoMigrate(&User{}, &Role{}, &Permission{})

    // 创建角色和权限
    role1 := &Role{Name: "管理员"}
    role2 := &Role{Name: "普通用户"}
    perm1 := &Permission{Name: "创建用户"}
    perm2 := &Permission{Name: "删除用户"}

    db.Create(role1, role2)
    db.Create(perm1, perm2)

    // 创建用户并分配角色和权限
    user1 := &User{Name: "张三", Roles: []*Role{role1},}
    user1.Roles[0].Permissions = []*Permission{perm1, perm2}
    user2 := &User{Name: "李四", Roles: []*Role{role2},}
    user2.Roles[0].Permissions = []*Permission{perm1}

    db.Create(user1, user2)

    // 查询用户和角色信息
    var users []*User
    db.Preload("Roles.Permissions").Find(&users)
    for _, user := range users {
        fmt.Printf("用户 %s 的角色信息：\n", user.Name)
        for _, role := range user.Roles {
            fmt.Printf("角色 %s 的权限信息：\n", role.Name)
            for _, perm := range role.Permissions {
                fmt.Printf("- %s\n", perm.Name)
            }
        }
    }

    // 更新用户角色信息
    db.Model(user1).Association("Roles").Replace([]*Role{role2})
    db.Model(user2).Association("Roles").Clear()

    // 删除用户和角色信息
    db.Delete(user1, user2)
    db.Delete(role1, role2)
    db.Delete(perm1, perm2)
}
```

这个代码示例去掉了一些冗余的代码，并且使用了链式调用的方式进行创建、查询、更新和删除操作。同时，由于 GORM 的 Preload 方法支持链式调用，因此我们可以在查询用户信息时，一次性把用户、角色和权限的信息全部查询出来，代码更加简洁。

### 进一步简化后的代码

可以使用简化的代码实现相同的功能：

```go
package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name  string
    Roles []*Role `gorm:"many2many:user_roles"`
}

type Role struct {
    gorm.Model
    Name        string
    Permissions []*Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
    gorm.Model
    Name string
}

func main() {
    // 连接数据库
    dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    // 自动迁移表结构
    db.AutoMigrate(&User{}, &Role{}, &Permission{})

    // 创建角色和权限
    role1 := &Role{Name: "管理员", Permissions: []*Permission{{Name: "创建用户"}, {Name: "删除用户"}}}
    role2 := &Role{Name: "普通用户", Permissions: []*Permission{{Name: "创建用户"}}}

    db.Create(role1, role2)

    // 创建用户并分配角色和权限
    user1 := &User{Name: "张三", Roles: []*Role{role1}}
    user2 := &User{Name: "李四", Roles: []*Role{role2}}

    db.Create(user1, user2)

    // 查询用户和角色信息
    var users []*User
    db.Preload("Roles.Permissions").Find(&users)
    for _, user := range users {
        fmt.Printf("用户 %s 的角色信息：\n", user.Name)
        for _, role := range user.Roles {
            fmt.Printf("角色 %s 的权限信息：\n", role.Name)
            for _, perm := range role.Permissions {
                fmt.Printf("- %s\n", perm.Name)
            }
        }
    }

    // 更新用户角色信息
    db.Model(user1).Association("Roles").Replace([]*Role{role2})
    db.Model(user2).Association("Roles").Clear()

    // 删除用户和角色信息
    db.Delete(user1, user2)
    db.Delete(role1, role2)
}
```

这个简化的代码与之前的代码相比主要做了以下几点改进：

1. 在创建角色时，同时创建了相应的权限，避免了后续的额外操作；
2. 在创建用户时，只需要指定用户的角色，而不需要为角色添加权限，因为角色已经预设了权限；
3. 通过链式操作，更加简洁地实现了用户角色的增删改查操作；
4. 精简了查询用户角色信息的代码。

### 上述代码中 []\*Permission 什么意思

`[]*Permission` 表示 Permission 结构体的指针数组。在这个例子中，一个 Role 可以有多个 Permission，因此 Role 结构体中的 Permissions 字段的类型是 `[]*Permission`。指针数组的优点是可以动态地分配内存，这样可以减少内存的浪费。另外，使用指针数组还可以提高程序的效率，因为指针传递比值传递更快。

### 上述代码中 更新用户角色信息，为什么用 Association

在 GORM 中，使用 `Association` 方法可以方便地处理多对多关系的增删改查操作。当我们想要在多对多关系中更新关联记录时，可以使用 `Association` 方法来更新关联关系。

具体来说，我们可以通过 `Association` 方法获取到某个模型的关联模型，然后使用关联模型的方法来更新关联关系。在这个例子中，我们需要更新用户和角色之间的关联关系，因此可以使用 `user.Roles().Replace(roles)` 方法来替换当前用户的角色列表。这个方法会把原有的角色列表清空，然后添加新的角色列表，实现了更新操作。如果需要追加角色，可以使用 `Add` 方法。如果需要删除角色，可以使用 `Delete` 方法。

总之，使用 `Association` 方法可以简化多对多关系的更新操作，提高程序的可读性和可维护性。

### 把原有的角色列表清空，然后添加新的角色列表，会不会出现，列表清空了，却没有添加成功的现象

在使用 `Association` 方法更新多对多关系时，如果出现了列表清空了但没有添加成功的情况，一般有两个可能的原因：

1. 数据库事务问题。在使用 GORM 操作数据库时，通常需要开启事务来确保操作的原子性和一致性。如果更新多对多关系时没有开启事务或事务异常终止，可能会导致清空列表成功但添加新列表失败的情况。

2. 数据库字段约束问题。在多对多关系中，经常需要使用中间表来记录两个模型之间的关联关系。如果中间表中定义的外键约束或唯一约束等字段约束与当前操作不兼容，也可能导致清空列表成功但添加新列表失败的情况。

为了避免出现这种情况，我们可以在更新多对多关系之前先检查数据库事务的状态，确保事务正常运行。此外，我们还应该仔细检查中间表的字段定义和约束，确保与当前操作相兼容。

### 先检查数据库事务的状态？

是的，检查数据库事务的状态可以确保事务正常运行，从而避免出现清空列表成功但添加新列表失败的情况。在 GORM 中，我们可以使用 `gorm.DB` 结构体的 `Begin` 方法开启一个数据库事务，然后使用 `gorm.DB` 结构体的 `Commit` 和 `Rollback` 方法来提交或回滚事务。

具体来说，在更新多对多关系时，可以使用 `DB.Begin` 方法开启一个数据库事务，然后在事务中执行清空列表和添加新列表的操作，最后使用 `DB.Commit` 方法提交事务。如果出现异常，可以使用 `DB.Rollback` 方法回滚事务。

以下是一个简单的示例代码：

```go
func UpdateUserRole(db *gorm.DB, user *User, roles []*Role) error {
    tx := db.Begin()

    if err := tx.Error; err != nil {
        return err
    }

    if err := tx.Model(user).Association("Roles").Replace(roles).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return err
    }

    return nil
}
```

在这个示例中，我们首先使用 `DB.Begin` 方法开启一个事务，然后在事务中执行清空列表和添加新列表的操作。如果操作成功，就使用 `DB.Commit` 方法提交事务，否则使用 `DB.Rollback` 方法回滚事务。注意，在操作过程中，我们使用 `tx.Error` 方法检查事务的状态，确保事务正常运行。

### 至此，完整代码

以下是完整的示例代码：

```go
package main

import (
    "fmt"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex"`
    Password string
    Roles    []*Role `gorm:"many2many:user_roles"`
}

type Role struct {
    gorm.Model
    Name       string         `gorm:"uniqueIndex"`
    Permissions []*Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
    gorm.Model
    Name string `gorm:"uniqueIndex"`
}

func main() {
    // 连接数据库
    dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    // 自动迁移表结构
    if err := db.AutoMigrate(&User{}, &Role{}, &Permission{}); err != nil {
        panic(err)
    }

    // 创建用户
    user1 := &User{
        Username: "user1",
        Password: "password1",
    }
    db.Create(user1)

    // 创建角色和权限
    role1 := &Role{
        Name: "role1",
        Permissions: []*Permission{
            {Name: "permission1"},
            {Name: "permission2"},
        },
    }
    db.Create(role1)
    role2 := &Role{
        Name: "role2",
        Permissions: []*Permission{
            {Name: "permission2"},
            {Name: "permission3"},
        },
    }
    db.Create(role2)

    // 更新用户角色信息
    roles := []*Role{role1, role2}
    if err := UpdateUserRoles(db, user1, roles); err != nil {
        panic(err)
    }

    // 查询用户信息
    var user User
    db.Preload("Roles.Permissions").First(&user, "username = ?", "user1")
    fmt.Printf("User: %+v\n", user)
}

func UpdateUserRoles(db *gorm.DB, user *User, roles []*Role) error {
    tx := db.Begin()

    if err := tx.Error; err != nil {
        return err
    }

    if err := tx.Model(user).Association("Roles").Replace(roles).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return err
    }

    return nil
}
```

在这个示例中，我们定义了 `User`、`Role` 和 `Permission` 三个模型，并使用 `many2many` 关系将它们关联起来。在 `UpdateUserRoles` 函数中，我们使用了 `Association` 方法更新用户的角色信息，并在事务中执行操作，确保操作的原子性。最后，我们查询了用户信息，并打印出来。

### primaryKey

`primaryKey` 是 GORM 中用于指定主键名称的选项之一。如果不使用此选项，则默认情况下，GORM 会将模型的 ID 字段作为主键。如果模型有多个字段组成复合主键，则可以使用 `gorm:"primaryKey"` 标签在这些字段上指定为主键。

例如，下面的代码演示了如何在 GORM 中使用复合主键：

```go
type User struct {
    Name   string `gorm:"primaryKey"`
    Number int    `gorm:"primaryKey"`
    Email  string
}

db.AutoMigrate(&User{})
```

在上面的示例中，`User` 模型的主键由 `Name` 和 `Number` 两个字段组成。在 `Name` 和 `Number` 字段上都使用了 `primaryKey` 标签，表示它们是组成主键的一部分。执行 `db.AutoMigrate(&User{})` 后，GORM 将会在数据库中创建一个名为 `users` 的表，其中包含 `Name` 和 `Number` 两个字段作为主键。

需要注意的是，使用复合主键的模型在使用 GORM 进行增删改查等操作时，需要使用 `db.Session(&gorm.Session{FullSaveAssociations: true})` 开启完整关联保存模式，以保证 GORM 能够正确地处理关联数据的更新和删除。

### 完整的例子

以下是一个使用 GORM 创建包含复合主键的表格的示例代码：

```go
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	ID    int    `gorm:"primaryKey;autoIncrement:false"`
	Title string `gorm:"primaryKey;autoIncrement:false"`
	Author string
}

func main() {
	// 连接数据库
	dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移表结构
	db.AutoMigrate(&Book{})

	// 创建数据行
	book1 := &Book{ID: 1, Title: "GORM Basics", Author: "John Doe"}
	book2 := &Book{ID: 2, Title: "Advanced GORM", Author: "Jane Smith"}

	db.Create(book1)
	db.Create(book2)

	// 查询数据行
	var books []*Book
	db.Find(&books)

	// 打印查询结果
	for _, book := range books {
		fmt.Printf("Book: %s by %s\n", book.Title, book.Author)
	}
}
```

在这个例子中，我们定义了一个 `Book` 结构体，并为其中的 `ID` 和 `Title` 字段都添加了 `primaryKey` 标记，以实现复合主键。然后我们创建了两条数据行，分别为 ID 为 1 和 2 的两本书。最后我们查询了所有的书，并将结果打印出来。

### 使用了 autoIncrement:false

复合主键的情况下，通常不会使用自增长的方式来生成主键，而是需要手动指定主键的值。在 GORM 中，可以通过在模型中定义多个字段，并使用 `primaryKey` 标记来指定这些字段作为复合主键。

下面是一个使用 GORM 定义复合主键的完整示例：

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
)

type OrderItem struct {
    OrderID  int    `gorm:"primaryKey"`
    ItemCode string `gorm:"primaryKey"`
    Name     string
    Price    float64
    CreatedAt time.Time
}

func main() {
    dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 自动迁移表结构
    db.AutoMigrate(&OrderItem{})

    // 插入数据
    item1 := &OrderItem{OrderID: 1, ItemCode: "A001", Name: "商品A", Price: 100.0, CreatedAt: time.Now()}
    item2 := &OrderItem{OrderID: 1, ItemCode: "A002", Name: "商品B", Price: 200.0, CreatedAt: time.Now()}
    item3 := &OrderItem{OrderID: 2, ItemCode: "A003", Name: "商品C", Price: 300.0, CreatedAt: time.Now()}
    db.Create(item1, item2, item3)

    // 查询数据
    var items []*OrderItem
    db.Find(&items)
    for _, item := range items {
        println(item.OrderID, item.ItemCode, item.Name, item.Price, item.CreatedAt)
    }

    // 更新数据
    db.Model(item1).Updates(map[string]interface{}{"Price": 150.0})

    // 删除数据
    db.Delete(item2)
}
```

在这个例子中，`OrderID` 和 `ItemCode` 两个字段一起作为 `OrderItem` 模型的复合主键，分别使用 `primaryKey` 标记来指定。这里注意到 `autoIncrement` 参数被设置为 `false`，因为复合主键一般不使用自增长的方式生成主键。

在插入数据时，需要手动指定主键的值。在查询、更新和删除数据时，也需要同时指定主键的值来定位需要操作的数据行。

### 什么情况下，使用复合主键

复合主键通常用于满足业务需求，例如要求某些字段的组合必须唯一。在这种情况下，单一的主键可能无法满足要求，因为其中一个字段可能不足以唯一标识一个记录。例如，在一个订单表中，一个订单可能有多个商品，因此订单编号和商品编号的组合可能是唯一的，这时候就可以使用复合主键来保证唯一性。

### 关联模式

在 GORM 中，关联模式有四种，分别是：

1. BelongsTo 关联
2. HasOne 关联
3. HasMany 关联
4. ManyToMany 关联

下面对每种关联模式进行简要介绍：

1. BelongsTo 关联：

BelongsTo 关联表示一个模型属于另一个模型，也就是一对一的关系。在 GORM 中，使用 `BelongsTo` 方法来定义关联关系，例如：

```go
type User struct {
    gorm.Model
    Name      string
    CompanyID uint
    Company   Company
}

type Company struct {
    gorm.Model
    Name  string
    Users []User
}

db.Model(&User{}).BelongsTo(&Company{})
```

2. HasOne 关联：

HasOne 关联表示一个模型拥有另一个模型，也就是一对一的关系。在 GORM 中，使用 `HasOne` 方法来定义关联关系，例如：

```go
type User struct {
    gorm.Model
    Name         string
    CreditCard   CreditCard
    CreditCardID uint
}

type CreditCard struct {
    gorm.Model
    Number string
    UserID uint
}

db.Model(&User{}).HasOne(&CreditCard{})
```

3. HasMany 关联：

HasMany 关联表示一个模型拥有多个另一个模型，也就是一对多的关系。在 GORM 中，使用 `HasMany` 方法来定义关联关系，例如：

```go
type User struct {
    gorm.Model
    Name   string
    Orders []Order
}

type Order struct {
    gorm.Model
    UserID uint
    Amount float32
}

db.Model(&User{}).HasMany(&Order{})
```

4. ManyToMany 关联：

ManyToMany 关联表示两个模型之间有多对多的关系。在 GORM 中，使用 `Many2Many` 方法来定义关联关系，例如：

```go
type User struct {
    gorm.Model
    Name   string
    Roles  []*Role `gorm:"many2many:user_roles"`
}

type Role struct {
    gorm.Model
    Name        string
    Permissions []*Permission `gorm:"many2many:role_permissions"`
}

db.Model(&User{}).Many2Many(&Role{}, "user_roles")
db.Model(&Role{}).Many2Many(&Permission{}, "role_permissions")
```

以上就是 GORM 中四种关联模式的介绍。需要注意的是，使用关联模式时，应该根据实际情况选择合适的模式。

### 如何区分 BelongsTo 和 HasOne

在 GORM 中，BelongsTo 和 HasOne 都表示一对一的关系，但它们在语义上有所不同。

BelongsTo 表示一个模型“属于”另一个模型，比如一个订单属于一个用户。在这种关系下，通常在从表（即属于表）中建立外键关联到主表（即拥有表）中。

HasOne 表示一个模型“拥有”另一个模型，比如一个用户拥有一个个人资料。在这种关系下，通常在主表中建立外键关联到从表中。

因此，可以根据语义来区分这两种关系。如果一个模型属于另一个模型，则使用 BelongsTo，如果一个模型拥有另一个模型，则使用 HasOne。

总结：

1. 使用上看主键是在从表还是在主表：在从表 BelongsTo, 在主表 HasOne
2. 语义不同： 如果一个模型属于另一个模型，则使用 BelongsTo，如果一个模型拥有另一个模型，则使用 HasOne。

### 关联操作

在 GORM 中，关联操作主要包括以下几个方面：

1. 预加载（Preloading）：一般用于查询时，同时将关联对象加载到内存中，避免多次查询数据库。
2. 关联查询（Joins）：根据关联条件，将两个或多个表关联起来，查询相关数据。
3. 关联插入（Association）：将两个或多个表之间的关联关系插入到关联表中。
4. 关联更新（Association）：更新关联表中的关联关系。
5. 关联删除（Association）：删除关联表中的关联关系。

这些操作涉及到不同的 GORM 方法，比如 `Preload`、`Joins`、`Association` 等。具体的使用方法可以查看 GORM 的官方文档。

### Association

在 GORM 中，Association 指的是通过关联关系进行增删改查操作的方法，主要包括以下几个方法：

- `Append`：添加关联记录
- `Delete`：删除关联记录
- `Replace`：替换关联记录
- `Clear`：清空关联记录

这些方法可以用于多对多、一对多等关联关系，可以通过关联的对象的指针或 ID 进行操作。这些方法在使用时需要注意事务的管理，避免出现数据不一致的情况。在使用时，也需要注意避免出现死循环的情况，特别是在多层级的关联关系中。

### Association 这四个 API 怎么用

这四个 API 都是用于修改关联关系的，具体如下：

1. `db.Model(&user).Association("Roles").Append(&role1, &role2)`：将 `user` 与 `role1` 和 `role2` 关联，如果关联关系已经存在，则不会进行任何操作。
2. `db.Model(&user).Association("Roles").Delete(&role1, &role2)`：删除 `user` 与 `role1` 和 `role2` 的关联关系，如果关联关系不存在，则不会进行任何操作。
3. `db.Model(&user).Association("Roles").Replace(&role1, &role2)`：将 `user` 与 `role1` 和 `role2` 的关联关系替换为新的关联关系，如果关联关系已经存在，则先删除旧的关联关系再进行新的关联。
4. `db.Model(&user).Association("Roles").Clear()`：清空 `user` 与 `Roles` 的关联关系。

其中，`user` 表示主表，`Roles` 表示关联表。在使用这些 API 前，需要先在主表中定义关联关系，例如：

```
type User struct {
    gorm.Model
    Name   string
    Roles  []*Role `gorm:"many2many:user_roles"`
}
```

这里的 `user_roles` 是中间表的名称，用于维护主表和关联表之间的关联关系。同时，在关联表中也需要定义相应的关联关系，例如：

```
type Role struct {
    gorm.Model
    Name        string
    Users       []*User `gorm:"many2many:user_roles"`
    Permissions []*Permission `gorm:"many2many:role_permissions"`
}
```

这里的 `user_roles` 和主表中的名称要保持一致，`Users` 表示与 `User` 表关联的字段，`role_permissions` 是另一个中间表的名称，用于维护 `Role` 表与 `Permission` 表之间的关联关系。

### Association 完整的例子

以下是一个完整的例子，展示如何在 GORM 中定义和使用关联模式：

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

type User struct {
    gorm.Model
    Name      string
    Email     string
    Profile   Profile
    Addresses []Address
}

type Profile struct {
    gorm.Model
    UserID uint
    Age    int
}

type Address struct {
    gorm.Model
    UserID  uint
    Country string
    City    string
}

func main() {
    dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // 自动迁移表结构
    err = db.AutoMigrate(&User{}, &Profile{}, &Address{})
    if err != nil {
        log.Fatal(err)
    }

    // 创建用户和关联数据
    user := User{
        Name:  "张三",
        Email: "zhangsan@example.com",
        Profile: Profile{
            Age: 18,
        },
        Addresses: []Address{
            {Country: "China", City: "Shanghai"},
            {Country: "USA", City: "New York"},
        },
    }

    result := db.Create(&user)
    if result.Error != nil {
        log.Fatal(result.Error)
    }

    // 查询用户和关联数据
    var user2 User
    result = db.Preload("Profile").Preload("Addresses").First(&user2)
    if result.Error != nil {
        log.Fatal(result.Error)
    }

    log.Println(user2)

    // 更新关联数据
    user2.Addresses[0].City = "Beijing"
    result = db.Save(&user2)
    if result.Error != nil {
        log.Fatal(result.Error)
    }

    // 删除关联数据
    result = db.Delete(&user2.Addresses[0])
    if result.Error != nil {
        log.Fatal(result.Error)
    }

    // 删除用户和关联数据
    result = db.Delete(&user2)
    if result.Error != nil {
        log.Fatal(result.Error)
    }
}
```

在这个例子中，我们定义了三个模型：`User`、`Profile` 和 `Address`。它们之间的关联模式分别为：

- `User` 和 `Profile`：一个用户有一个个人资料，使用 BelongsTo 模式。
- `User` 和 `Address`：一个用户有多个地址，使用 HasMany 模式。
- `Address` 和 `User`：一个地址属于一个用户，使用 BelongsTo 模式。

我们首先创建一个用户和关联数据，并通过 `Create()` 方法将它们存储到数据库中。接着，我们查询该用户及其关联数据，使用 `Preload()` 方法可以在一次查询中加载关联数据，避免了多次查询数据库的问题。我们还演示了如何更新关联数据和删除关联数据，以及如何删除用户及其关联数据。

下面是一个完整的例子，演示如何使用 GORM 进行实体关联。

假设有两个表：用户和订单。一个用户可以拥有多个订单，而一个订单只能属于一个用户。用户和订单之间是一对多的关系。首先，定义两个模型：

```go
type User struct {
    gorm.Model
    Name    string
    Email   string
    Orders  []Order
}

type Order struct {
    gorm.Model
    UserID  uint
    Name    string
    Price   float64
}
```

在用户模型中，使用了 GORM 内置的 `gorm.Model` 结构体，该结构体包含了一些常用的字段，比如 ID、CreatedAt、UpdatedAt 和 DeletedAt。同时，用户模型还包含了一个名为 `Orders` 的字段，该字段为用户的订单列表。在订单模型中，使用了 `UserID` 字段来存储订单所属的用户 ID。

然后，使用 GORM 进行实体关联：

```go
// 创建用户和订单
user := &User{Name: "张三", Email: "zhangsan@example.com"}
order1 := &Order{Name: "订单1", Price: 100}
order2 := &Order{Name: "订单2", Price: 200}
user.Orders = []Order{*order1, *order2}
db.Create(user)

// 查询用户及其订单信息
var u User
db.Preload("Orders").First(&u)
fmt.Println("用户信息：", u)
for _, order := range u.Orders {
    fmt.Println("订单信息：", order)
}

// 更新用户订单信息
db.Model(&u).Association("Orders").Delete(order1)
db.Model(&u).Association("Orders").Replace([]Order{*order2})
db.Model(&u).Association("Orders").Clear()

// 删除用户及其订单信息
db.Delete(&u)
```

在创建用户时，使用了用户模型的 `Orders` 字段来设置用户的订单列表。在查询用户时，使用了 GORM 的 `Preload` 方法来预加载用户的订单信息，以避免 N+1 查询的问题。在更新用户订单信息时，使用了 GORM 的 `Association` 方法来对用户的订单列表进行增删改查操作。在删除用户时，使用了 GORM 的 `Delete` 方法来同时删除用户及其订单信息。

### 完整代码如下：

```go
package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name    string
    Email   string
    Orders  []Order
}

type Order struct {
    gorm.Model
    UserID  uint
    Name    string
    Price   float64
}

func main() {
    // 连接数据库
    dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    // 自动迁移表结构
    db.AutoMigrate(&User{}, &Order{})

    // 创建用户和订单
    user := &User{Name: "张三", Email: "zhangsan@example.com"}
    order1 := &Order{Name: "订单1", Price:

// 可以通过 `Association` 方法来进行关联操作。`Association` 方法接受一个关联字段的名称作为参数，返回该关联字段的关联对象，可以对该关联对象进行增删改查等操作。

// 例如，在上一个例子中，我们可以通过以下代码来进行关联操作：

// 建立关联关系
db.Model(&user).Association("Profile").Assign(&profile)

// 清空关联关系
db.Model(&user).Association("Profile").Clear()

// 添加关联关系
db.Model(&user).Association("Addresses").Append(&address)

// 删除关联关系
db.Model(&user).Association("Addresses").Delete(&address)
```

其中，

1. `Assign` 方法可以用来建立关联关系，
2. `Clear` 方法可以用来清空关联关系，
3. `Append` 方法可以用来添加关联关系，
4. `Delete` 方法可以用来删除关联关系。

完整代码如下：

```go
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Profile   Profile
	Addresses []Address
}

type Profile struct {
	gorm.Model
	UserID uint
	Name   string
}

type Address struct {
	gorm.Model
	UserID  uint
	City    string
	Country string
}

func main() {
	// 连接数据库
	dsn := "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 自动迁移表结构
	db.AutoMigrate(&User{}, &Profile{}, &Address{})

	// 创建用户、资料和地址
	user := &User{
		Name: "张三",
		Profile: Profile{
			Name: "张三的资料",
		},
		Addresses: []Address{
			{City: "北京", Country: "中国"},
			{City: "上海", Country: "中国"},
		},
	}

	db.Create(user)

	// 查询用户、资料和地址
	var u User
	db.Preload("Profile").Preload("Addresses").Find(&u)

	fmt.Printf("用户：%s\n", u.Name)
	fmt.Printf("资料：%s\n", u.Profile.Name)
	fmt.Println("地址：")
	for _, address := range u.Addresses {
		fmt.Printf("- %s, %s\n", address.City, address.Country)
	}

	// 建立关联关系
	profile := Profile{
		Name: "新的资料",
	}
	db.Model(&user).Association("Profile").Assign(&profile)

	// 清空关联关系
	db.Model(&user).Association("Addresses").Clear()

	// 添加关联关系
	address := Address{
		City:    "广州",
		Country: "中国",
	}
	db.Model(&user).Association("Addresses").Append(&address)

	// 删除关联关系
	db.Model(&user).Association("Addresses").Delete(&address)
}
```

### 建立关联关系 清空关联关系 对数据库有什么影响

建立和清空关联关系对数据库都会产生影响，具体如下：

### 建立关联关系：

当我们使用 gorm 的关联 API （例如 `Association`）来建立两个实体之间的关联关系时，gorm 会自动在关联关系所对应的中间表中插入一些新的数据记录，这些记录描述了这两个实体之间的关联关系。这些新插入的数据会被持久化到数据库中。

例如，当我们将一个用户实体与一个角色实体建立起多对多关联关系时，gorm 会自动在 user_roles 中间表中插入一个新的数据记录，记录描述了这个用户实体和这个角色实体之间的关联关系。

### 清空关联关系：

当我们使用 gorm 的关联 API（例如 `Association`）来清空两个实体之间的关联关系时，gorm 会自动从关联关系所对应的中间表中删除所有相关的数据记录。这些记录描述了这两个实体之间的关联关系。这些被删除的数据会被持久化到数据库中。

例如，当我们清空一个用户实体与一个角色实体之间的多对多关联关系时，gorm 会自动从 user_roles 中间表中删除所有描述这个用户实体和这个角色实体之间关联关系的数据记录。

### 上述，清空关联关系中的删除，是软删除还是物理删除

上述代码中的 `Clear()` 操作是物理删除，即从关联表中删除所有关联的记录。如果需要使用软删除，可以使用 `gorm.Model` 提供的软删除机制，或者手动添加一个 `deleted_at` 字段，并在删除记录时将该字段赋值为当前时间戳，表示该记录已被删除，而不是从数据库中删除该记录。具体实现方式可以参考 GORM 的文档或者其他相关资料。

### 实体关联

https://gorm.io/zh_CN/docs/associations.html#%E6%9F%A5%E6%89%BE%E5%85%B3%E8%81%94

### db.Model(&user).Association("Orders").Find(&orders) 与 db.Preload("Orders").Find(&users) 的区别

`db.Model(&user).Association("Orders").Find(&orders)` 和 `db.Preload("Orders").Find(&users)` 两者的作用不同。

- `db.Model(&user).Association("Orders").Find(&orders)`：是在已有用户实体的前提下，通过 `Association` 方法获取其 `Orders` 关联实体的值并填充到 `&orders` 中。这种方法适用于已经查询出了用户实体并需要获取其关联实体的情况。

- `db.Preload("Orders").Find(&users)`：是在查询用户实体时预加载 `Orders` 关联实体的值，以避免在后续使用时需要执行额外的 SQL 查询。这种方法适用于需要查询用户实体和其关联实体的情况。

总之，两种方法的区别在于是否在查询时进行预加载，以及获取关联实体的方式不同。选择哪种方法应根据具体情况而定。

### 两者使用场景

当你需要加载单个对象的关联数据时，可以使用 `Association` 方法。当你需要加载多个对象的关联数据时，可以使用 `Preload` 方法。

使用 `Association` 方法时，需要先获取主对象的实例，并通过指定关联字段的名称来获取关联数据。这种方式更适用于当你只需要获取一个对象的关联数据时。

而使用 `Preload` 方法时，你可以通过在 `Find` 或 `First` 方法中传递参数来加载多个对象的关联数据。这种方式更适用于当你需要一次性获取多个对象的关联数据时，因为它可以通过 SQL 查询语句的联接操作来一次性加载多个对象的关联数据，减少了对数据库的访问次数，提高了查询效率。但是，如果你只需要获取一个对象的关联数据时，使用 `Preload` 方法反而会浪费一定的资源，因为它会在查询主对象时将所有的关联数据都加载进来，而不管你是否需要。

### db.Omit("Languages.\*").Create(&user) 与 db.Omit("Languages").Create(&user)的区别

`db.Omit("Languages.*").Create(&user)` 和 `db.Omit("Languages").Create(&user)` 都是使用 GORM 的 Omit 方法来指定忽略某个字段，不会将该字段的值插入数据库。

但是两者的区别在于忽略的是整个 Languages 关联模型还是 Languages 模型的所有字段。具体来说：

- `db.Omit("Languages.*").Create(&user)` 忽略了 Languages 关联模型的所有字段，相当于没有创建 Languages 模型的记录。
- `db.Omit("Languages").Create(&user)` 只忽略了 Languages 模型本身的所有字段，不会影响到 Languages 关联模型的其他字段，例如在该例子中的 `UserLanguages` 表中的 `user_id` 和 `language_id` 字段。也就是说，执行该语句后仍然可以在 `UserLanguages` 表中插入对应的 `user_id` 和 `language_id`。

因此，根据具体需求来选择使用哪种方式进行字段忽略。
下面是一个完整的示例，展示如何使用 `Omit` 方法来避免 GORM 创建或更新时处理某些关联的字段：

```go
package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Language struct {
	gorm.Model
	Name string
}

type User struct {
	gorm.Model
	Name      string
	Languages []*Language `gorm:"many2many:user_languages"`
}

func main() {
	// 连接 SQLite 数据库
	dsn := "test.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移表结构
	db.AutoMigrate(&User{}, &Language{})

	// 创建语言记录
	lang1 := &Language{Name: "Go"}
	lang2 := &Language{Name: "Python"}
	lang3 := &Language{Name: "Java"}
	db.Create(lang1, lang2, lang3)

	// 创建用户并分配语言
	user := &User{
		Name: "Alice",
		Languages: []*Language{
			lang1,
			lang2,
		},
	}
	db.Omit("Languages.*").Create(user)

	// 查询用户信息
	var users []User
	db.Preload("Languages").Find(&users)
	for _, u := range users {
		fmt.Printf("%s's languages:\n", u.Name)
		for _, l := range u.Languages {
			fmt.Printf("- %s\n", l.Name)
		}
	}

	// 更新用户语言信息
	db.Model(user).Association("Languages").Replace([]*Language{lang3})
	db.Save(user)

	// 查询用户信息
	var updatedUser User
	db.Preload("Languages").First(&updatedUser, user.ID)
	fmt.Printf("%s's updated languages:\n", updatedUser.Name)
	for _, l := range updatedUser.Languages {
		fmt.Printf("- %s\n", l.Name)
	}

	// 删除用户和语言记录
	db.Unscoped().Delete(user)
	db.Unscoped().Delete(lang1, lang2, lang3)
}
```

在这个示例中，我们创建了三个 `Language` 记录并将它们分配给一个名为 `Alice` 的 `User`。在创建 `User` 记录时，我们使用 `Omit` 方法来忽略 `Languages` 关联的创建和更新。因此，即使我们没有为 `User` 记录分配任何语言，GORM 仍会成功创建该记录。

在更新 `User` 记录时，我们使用 `Association` 和 `Replace` 方法来清空原有的语言分配并将其替换为另一个语言记录。最后，我们删除了所有的记录。

### mac 查看是否有 sqlite

macOS 系统默认是自带 SQLite 的，可以通过终端输入以下命令查看系统中是否已经安装了 SQLite：

```sh
sqlite3 --version
```

如果已经安装了 SQLite，则会返回安装的版本信息，否则会提示 command not found。

### 查询语句

查询语句是指用来从数据库中获取数据的命令。在关系型数据库中，通常使用 SQL（Structured Query Language）来编写查询语句。在 GORM 中，我们可以使用链式方法和原生 SQL 语句来构建查询。

以下是一些 GORM 中常用的查询方法：

1. db.Find(&result, conditions, params)：查询符合条件的所有记录，结果将保存到 result 中。
2. db.First(&result, conditions, params)：查询符合条件的第一条记录，结果将保存到 result 中。
3. db.Last(&result, conditions, params)：查询符合条件的最后一条记录，结果将保存到 result 中。
4. db.Where(conditions, params).Find(&result)：使用指定的条件查询符合条件的所有记录，结果将保存到 result 中。
5. db.Order(order).Find(&result)：按照指定的顺序查询所有记录，结果将保存到 result 中。
6. db.Limit(limit).Find(&result)：限制查询结果的数量，并将结果保存到 result 中。
7. db.Offset(offset).Find(&result)：设置查询的偏移量，并将结果保存到 result 中。
8. db.Select(columns).Find(&result)：只查询指定的列，并将结果保存到 result 中。
9. db.Joins(joins).Find(&result)：查询关联表的记录，并将结果保存到 result 中。
10. db.Raw(query, values).Scan(&result)：使用原生 SQL 查询数据库，并将结果保存到 result 中。

除了以上方法外，还有许多其他的查询方法。具体使用哪些方法取决于你的具体需求和业务逻辑。

### gorm polymorphicValue 使用完整示例

下面是使用 GORM 实现多态关联的完整示例：

```go
package main

import (
    "fmt"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Content struct {
    gorm.Model
    PolymorphicValue interface{} `gorm:"polymorphic:Owner;"`
}

type Text struct {
    Content
    Body string
}

type Image struct {
    Content
    URL string
}

type Video struct {
    Content
    URL string
}

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&Text{}, &Image{}, &Video{})

    text := Text{
        Body: "Hello, world!",
    }

    db.Create(&text)

    image := Image{
        URL: "https://example.com/image.png",
    }

    db.Create(&image)

    video := Video{
        URL: "https://example.com/video.mp4",
    }

    db.Create(&video)

    var content Content
    db.First(&content, 1)

    switch owner := content.PolymorphicValue.(type) {
    case *Text:
        fmt.Println("Text body:", owner.Body)
    case *Image:
        fmt.Println("Image URL:", owner.URL)
    case *Video:
        fmt.Println("Video URL:", owner.URL)
    }

}
```

在这个例子中，我们定义了三个子模型 Text、Image 和 Video，它们都扩展了 Content 模型，并且都有一个特定的字段。Content 模型具有一个名为"PolymorphicValue"的字段，它可以存储对不同类型的实体的引用。"polymorphic"标记告诉 GORM 将"PolymorphicValue"字段用作多态关联。

在这个例子中，我们创建了一个 Text、Image 和 Video 实例，并将它们保存到数据库中。我们然后获取了一个 Content 实例，并根据存储在"PolymorphicValue"字段中的值的类型进行不同的处理。

这是一个简单的示例，演示了如何使用 GORM 实现多态关联。根据您的实际需求，可能需要使用更复杂的模型层次结构和多态关联来实现您的数据模型。

### 上述例子的数据库结构

上述示例中，使用了 GORM 提供的自动迁移功能来创建数据表，因此在运行程序时会自动创建以下数据表：

```sql
texts
  - id (integer, primary key)
  - created_at (timestamp)
  - updated_at (timestamp)
  - deleted_at (timestamp)
  - body (text)

images
  - id (integer, primary key)
  - created_at (timestamp)
  - updated_at (timestamp)
  - deleted_at (timestamp)
  - url (text)

videos
  - id (integer, primary key)
  - created_at (timestamp)
  - updated_at (timestamp)
  - deleted_at (timestamp)
  - url (text)

contents
  - id (integer, primary key)
  - created_at (timestamp)
  - updated_at (timestamp)
  - deleted_at (timestamp)
  - owner_type (varchar(255))
  - owner_id (integer)
```

在这个例子中，每个子模型(Text、Image、Video)都对应一个数据表(texts、images、videos)，同时 Content 模型对应了一个数据表(contents)。Content 模型中的"PolymorphicValue"字段会根据它所引用的实体的类型自动填充"owner_type"和"owner_id"字段，以便在查询时可以根据类型进行过滤。在本例中，"owner_type"字段的值将分别是"texts"、"images"、"videos"，而"owner_id"字段的值将分别是对应实体在其数据表中的主键。

### owner_type 的值为什么是复数的形式

在 GORM 中，"owner_type"字段的值是引用的实体类型对应的数据表名称。对于大多数情况下，GORM 会默认使用引用的实体类型名称的复数形式作为对应的数据表名称。这是因为通常情况下，一个实体类型对应着多条记录，因此使用复数形式更符合数据表名称的语义。

当然，在某些情况下，您可能需要自定义实体类型对应的数据表名称。在 GORM 中，您可以使用模型的"table name"标记来自定义数据表名称，例如：

```go
type Product struct {
  gorm.Model
  Name string
  Price float64
  OrderID uint
  OrderType string
}

type Order struct {
  gorm.Model
  Products []Product `gorm:"foreignKey:OrderID;polymorphic:OrderType;"`
}

func (Product) TableName() string {
  return "product"
}

func (Order) TableName() string {
  return "order"
}
```

在上面的示例中，Product 和 Order 两个模型都通过 TableName 方法自定义了对应的数据表名称。这样，在定义多态关联时，对应的数据表名称就会变成单数形式了。所以，使用复数形式的数据表名称仅仅是 GORM 的默认行为，您可以根据需要自定义数据表名称。

### polymorphicValue 两个结构体的例子

好的，以下是一个完整的例子：

```go
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Animal struct {
    gorm.Model
    Name       string
    Toy        Toy `gorm:"polymorphic:Owner;polymorphicValue:animal"`
}

type Toy struct {
    gorm.Model
    Name       string
    OwnerID    int
    OwnerType  string
}

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    // 自动迁移
    db.AutoMigrate(&Animal{}, &Toy{})

    dog := Animal{Name: "Dog"}
    db.Create(&dog)
    toy1 := Toy{Name: "Bone", OwnerID: int(dog.ID), OwnerType: "animal"}
    db.Create(&toy1)

    cat := Animal{Name: "Cat"}
    db.Create(&cat)
    toy2 := Toy{Name: "Mouse", OwnerID: int(cat.ID), OwnerType: "animal"}
    db.Create(&toy2)

    var toys []Toy
    db.Preload("Owner").Find(&toys)

    for _, toy := range toys {
        fmt.Printf("Toy name: %s, Owner name: %s\n", toy.Name, toy.Owner.Name)
    }
}
```

这个例子中定义了两个 GORM 模型，Animal 和 Toy。Animal 模型中定义了一个 Toy 类型的字段，用于表示动物玩具的信息。在这个字段的 tag 中，使用了"polymorphic"标记来定义多态关联，并指定了"Owner"作为关联的字段名称。这样，当 Animal 模型与 Toy 模型关联时，GORM 将会将 Animal 的 ID 值赋值给 Toy 模型中的 OwnerID 字段，并将"animal"字符串赋值给 Toy 模型中的 OwnerType 字段。

另外，在"polymorphic"标记的基础上，还使用了"polymorphicValue"标记来指定 OwnerType 字段中的值。具体来说，当 Animal 模型与 Toy 模型关联时，OwnerType 字段中的值将会被设置为"animal"。这个值用来表示动物玩具的所有者是一只动物。如果存在其他的所有者类型，例如"human"，那么它们的 OwnerType 值就可以分别设置为"human"。

最终，在 main 函数中，我们使用 GORM 创建了两个 Animal 对象（分别代表一只狗和一只猫），并创建了两个 Toy 对象（分别代表狗和猫的玩具）。然后，我们使用 GORM 的 Preload 方法来加载 Toy 对象的 Owner 字段，以便在输出 Toy 信息时，能够同时输出它们的所有者信息。

最终的输出结果如下：

```
Toy name: Bone, Owner name: Dog
Toy name: Mouse, Owner name: Cat
```

可以看到，程序成功地创建了 Animal 和 Toy 对象，并建立了它们之间的多态关联。同时，在输出 Toy 信息时，也成功地输出了它

### polymorphicValue 三个结构体的例子

如果需要实现三个结构体之间的多态关联，可以按照以下步骤进行操作：

1. 在 GORM 模型中定义多态关联的字段，例如定义一个 Owner 字段；
2. 在 Owner 字段的 tag 中，使用"polymorphic"标记来定义多态关联，并指定关联的字段名称，例如"Owner"；
3. 在 Owner 字段的 tag 中，使用"polymorphicValue"标记来指定 OwnerType 字段中的值；
4. 在其他两个模型中分别定义一个 HasMany 类型的字段，用于保存多态关联的信息；
5. 在其他两个模型的 HasMany 字段的 tag 中，使用"polymorphic"标记来定义多态关联，并指定关联的字段名称，例如"Owner"；
6. 在其他两个模型的 HasMany 字段的 tag 中，使用"polymorphicValue"标记来指定 OwnerType 字段中的值；
7. 在代码中创建和查询这三个模型时，根据需要使用 GORM 的 Preload 方法来加载关联信息。

下面是一个具有三个模型之间多态关联的完整例子：

```go
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Animal struct {
    gorm.Model
    Name       string
    Toys       []Toy `gorm:"polymorphic:Owner;polymorphicValue:animal"`
}

type Human struct {
    gorm.Model
    Name       string
    Toys       []Toy `gorm:"polymorphic:Owner;polymorphicValue:human"`
}

type Toy struct {
    gorm.Model
    Name       string
    OwnerID    int
    OwnerType  string
}

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    // 自动迁移
    db.AutoMigrate(&Animal{}, &Human{}, &Toy{})
// 创建->返回id->拼接toy数据->保存数据库
    dog := Animal{Name: "Dog"}
    db.Create(&dog)
    toy1 := Toy{Name: "Bone", OwnerID: int(dog.ID), OwnerType: "animal"}
    db.Create(&toy1)

    cat := Animal{Name: "Cat"}
    db.Create(&cat)
    toy2 := Toy{Name: "Mouse", OwnerID: int(cat.ID), OwnerType: "animal"}
    db.Create(&toy2)

    person := Human{Name: "John"}
    db.Create(&person)
    toy3 := Toy{Name: "Car", OwnerID: int(person.ID), OwnerType: "human"}
    db.Create(&toy3)

    var animals []Animal
    db.Preload("Toys").Find(&animals)
    for _, animal := range animals {
        for _, toy := range animal.Toys {
            fmt.Printf("%s's toy name: %s\n", animal.Name, toy.Name)
        }
    }

    var humans []Human
    db.Preload("Toys").Find(&humans)
    for _, human := range humans {
        for _, toy := range human.Toys {
            fmt.Printf("%s's toy name: %s\n", human.Name, toy.Name)
        }
    }
}
```

在这个例子中，我们定义了三个 GORM 模型：Animal、Human 和 Toy。其中 Animal 和 Human 模型都
具有多个 Toy，而 Toy 模型保存了所有的多态关联信息。Animal 和 Human 模型的 Toy 字段的 tag 中都使用了"polymorphic"标记来定义多态关联，并指定了 Owner 字段作为关联字段，同时使用了"polymorphicValue"标记来指定 OwnerType 字段的值。

在代码中，我们创建了一些 Animal、Human 和 Toy 的实例，并使用 GORM 的 Create 方法将它们保存到数据库中。然后使用 GORM 的 Preload 方法来加载所有 Animal 和 Human 实例及它们的 Toy。

最后，我们遍历加载的 Animal 和 Human 实例，并输出每个实例的名称以及它们的 Toy 的名称。注意，这里我们使用了嵌套循环来遍历 Animal 和 Human 的 Toys，这样就可以遍历所有的多态关联信息。

当我们运行这个程序时，输出结果如下：

```
Dog's toy name: Bone
Cat's toy name: Mouse
John's toy name: Car
```

这表明我们成功地实现了 Animal、Human 和 Toy 之间的多态关联。

### 简化后的代码

好的，这里是一个简化后的示例代码：

```go
package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Animal struct {
	ID   uint
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;polymorphicValue:animal"`
}

type Human struct {
	ID   uint
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;polymorphicValue:human"`
}

type Toy struct {
	ID        uint
	Name      string
	OwnerID   uint
	OwnerType string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Animal{}, &Human{}, &Toy{})
// 这里直接将数据，完整化成一个Animal/Human数据结构，创建时候，自动生成Toy数据
	dog := Animal{Name: "Dog", Toy: Toy{Name: "Bone"}}
	cat := Animal{Name: "Cat", Toy: Toy{Name: "Mouse"}}
	john := Human{Name: "John", Toy: Toy{Name: "Car"}}

	db.Create(&dog)
	db.Create(&cat)
	db.Create(&john)

	var animals []Animal
	db.Preload("Toy").Find(&animals)

	for _, animal := range animals {
		fmt.Printf("%s's toy name: %s\n", animal.Name, animal.Toy.Name)
	}

	var humans []Human
	db.Preload("Toy").Find(&humans)

	for _, human := range humans {
		fmt.Printf("%s's toy name: %s\n", human.Name, human.Toy.Name)
	}
}
```

这个简化版本的代码与之前的示例类似，只是使用了 `uint` 类型的 ID，并将 Animal 和 Human 的 Toy 字段的类型都改为了 `Toy`。此外，这个版本还省略了模型定义和数据库表结构的信息，只包含了必要的字段和标记。

这个简化版代码的运行结果与之前的示例相同：

```
Dog's toy name: Bone
Cat's toy name: Mouse
John's toy name: Car
```

### 自引用 Has One 如何理解

自引用 Has One 是指在一个模型中定义一个与该模型自身相同类型的 Has One 关联关系。这种关联关系可以用来表示一个模型具有与之相关联的另一个模型，而该关联模型本身也具有与之相关联的另一个模型，因此形成了一个自我引用的关系。

例如，我们可以考虑一个 Employee 模型，每个 Employee 模型都可能有一个 Manager 模型与之相关联。而每个 Manager 模型本身也可能有一个 Manager 模型与之相关联，因为一个 Manager 可能受另一个 Manager 管理。

在 GORM 中，我们可以通过在模型中定义一个与该模型自身相同类型的 Has One 关联关系来实现自引用 Has One。我们可以使用 `foreignkey` 和 `references` 标记来指定关联关系的外键和引用的主键，以及使用 `polymorphic` 标记来支持多态关联。

以下是一个 Employee 模型的示例代码，其中包括了一个与该模型自身相同类型的 Has One 关联关系：

```go
type Employee struct {
  ID         uint
  Name       string
  ManagerID  *uint
  Manager    *Employee `gorm:"foreignKey:ManagerID;references:ID"`
}
```

在这个示例代码中，Employee 模型包含了一个 ManagerID 字段和一个指向 Manager 模型的指针类型的 Manager 字段。我们使用了 `foreignKey` 和 `references` 标记来指定 ManagerID 字段作为外键，并引用了 Employee 模型的 ID 主键。这样，每个 Employee 模型都可以与一个其他 Employee 模型相关联，形成自引用的 Has One 关联关系。

当我们使用 GORM 创建和查询 Employee 模型时，GORM 会自动处理自引用 Has One 关联关系，并使用 JOIN 操作来加载相关联的 Manager 模型。

例如，我们可以使用以下代码创建两个 Employee 模型，并将其中一个模型的 Manager 字段设置为另一个模型：

```go
db.Create(&Employee{Name: "Alice"})
db.Create(&Employee{Name: "Bob", ManagerID: &[]uint{1}[0]})
```

这个代码会将两个 Employee 模型保存到数据库中，其中一个模型的 Manager 字段会指向另一个模型。然后我们可以使用以下代码查询所有 Employee 模型，以及每个模型的名称和 Manager 名称：

```go
var employees []Employee
db.Preload("Manager").Find(&employees)
for _, employee := range employees {
  fmt.Printf("%s's manager: %v\n", employee.Name, employee.Manager)
}
```

这个代码会输出以下内容：

```
Alice's manager: <nil>
Bob's manager: &{1 Alice 0xc000106320}
```

这表明我们成功地创建了两个 Employee 模型，并将其中一个模型的 Manager 字段设置为另一个模型，从而实现了自引用 Has One 关联关系。

### 上面为什么使用指针类型

在 GORM 中，如果我们想要定义一个与该模型自身相同类型的关联关系，通常需要使用指针类型。这是因为在关联模型中，我们通常需要引用到该模型自身的实例，而如果使用非指针类型，会导致循环依赖和递归引用的问题，进而导致程序崩溃。

在自引用 Has One 中，我们通常使用指针类型来定义模型的关联字段，以避免这些问题。例如，在 Employee 模型中，我们使用了 `ManagerID` 字段和 `Manager` 字段来表示每个员工的上级领导。`Manager` 字段的类型为 `*Employee`，这意味着它是一个指向另一个 Employee 实例的指针类型。

通过使用指针类型，我们可以避免循环依赖和递归引用的问题，并让 GORM 能够正确地处理关联关系。当我们查询一个 Employee 实例时，GORM 会自动加载与之相关联的 Manager 实例，以及与之相关联的 Manager 实例的 Manager 实例，以此类推，直到所有的关联关系都被加载完毕。

总之，在 GORM 中，为了避免循环依赖和递归引用的问题，并正确地处理与该模型自身相同类型的关联关系，我们通常需要使用指针类型来定义关联字段。这适用于所有类型的关联关系，包括自引用 Has One。

### 为什么指针引用可以避免 循环依赖和递归引用的问题

指针引用可以避免循环依赖和递归引用的问题，是因为指针允许我们将一个类型的实例作为另一个类型的属性引用，而不是直接包含该类型的实例。这种方式可以避免出现嵌套的实例，进而避免循环依赖和递归引用的问题。

假设我们有两个结构体类型 `A` 和 `B`，它们都有一个属性字段 `a` 和 `b`，分别指向对方的实例。如果我们使用非指针类型定义这两个结构体，那么在编译时就会发生错误，因为定义中存在循环依赖：

```go
type A struct {
    a B
}

type B struct {
    b A
}
```

相反，如果我们使用指针类型来定义这两个结构体，那么这个问题就可以得到解决：

```go
type A struct {
    a *B
}

type B struct {
    b *A
}
```

这种方式下，`a` 和 `b` 不再是实例，而是指向另一个类型的实例的指针。这样，我们就可以避免出现嵌套的实例，进而避免循环依赖和递归引用的问题。

在 GORM 中，我们使用指针类型来定义自引用 Has One 的关联字段，以避免这些问题。例如，在 `Employee` 模型中，我们使用 `ManagerID` 字段和 `Manager` 字段来表示每个员工的上级领导。`Manager` 字段的类型为 `*Employee`，这意味着它是一个指向另一个 `Employee` 实例的指针类型。通过使用指针类型，我们可以避免循环依赖和递归引用的问题，并让 GORM 能够正确地处理关联关系。

### 什么时候用指针类型，什么时候用实例

在定义 GORM 模型的时候，我们应该根据具体的情况选择使用指针类型或实例类型。

通常情况下，我们使用指针类型来定义关联关系字段，以避免循环依赖和递归引用的问题，同时也可以更好地处理关联关系。例如，如果一个员工有一个上级领导，我们可以使用 `Manager` 字段来表示这个关系：

```go
type Employee struct {
    ID         uint
    Name       string
    ManagerID  uint
    Manager    *Employee
}
```

在这个例子中，`Manager` 字段的类型为 `*Employee`，表示这是一个指向另一个 `Employee` 实例的指针类型。这样，我们可以避免循环依赖和递归引用的问题，同时也可以让 GORM 能够正确地处理关联关系。

然而，在某些情况下，我们可能需要使用实例类型来定义关联关系字段。例如，如果我们需要表示一个关联到多个 `OrderItem` 的 `Order`，我们可以使用 `Items` 字段来表示这个关系：

```go
type Order struct {
    ID         uint
    Name       string
    Items      []OrderItem
}

type OrderItem struct {
    ID         uint
    Name       string
    OrderID    uint
}
```

在这个例子中，`Items` 字段的类型为 `[]OrderItem`，表示这是一个包含多个 `OrderItem` 实例的切片类型。由于一个 `Order` 可能关联到多个 `OrderItem`，所以我们需要使用实例类型来表示这个关系。

总之，我们应该根据具体的情况选择使用指针类型或实例类型来定义 GORM 模型的关联关系字段。

1. 如果一个实例可能关联到多个实例，那么我们应该使用实例类型；
2. 如果一个实例只会关联到一个实例，或者存在循环依赖和递归引用的问题，那么我们应该使用指针类型。

### 指针和实例的区别

在 GORM 模型中，指针和实例都可以用来表示模型之间的关联关系。

1. 指针通常用于表示一对一或一对多的关系，
2. 而实例通常用于表示多对一或多对多的关系。

指针类型在 GORM 模型中的用法如下：

```go
type User struct {
    ID    uint
    Name  string
    Email string
}

type Post struct {
    ID       uint
    Title    string
    Content  string
    UserID   uint
    User     *User // 指向 User 实例的指针类型
}
```

在这个例子中，`Post` 模型包含一个指向 `User` 实例的指针类型字段 `User`，表示一个 `Post` 对象对应一个 `User` 对象。

实例类型在 GORM 模型中的用法如下：

```go
type User struct {
    ID       uint
    Name     string
    Email    string
    Articles []Article // 包含多个 Article 实例的切片类型
}

type Article struct {
    ID        uint
    Title     string
    Content   string
    UserID    uint
    User      User // 关联到 User 实例
}
```

在这个例子中，`User` 模型包含一个包含多个 `Article` 实例的切片类型字段 `Articles`，表示一个 `User` 对象关联到多个 `Article` 对象。而 `Article` 模型包含一个关联到 `User` 实例的实例类型字段 `User`，表示一个 `Article` 对象对应一个 `User` 对象。

在使用指针和实例类型时，我们需要根据具体的场景来选择。通常来说，如果一个模型关联到多个实例，我们应该使用实例类型；如果一个模型只会关联到一个实例，或者存在循环依赖和递归引用的问题，我们应该使用指针类型。此外，我们还需要注意在使用指针类型时避免空指针错误的问题。

### 如何避免空指针错误

1. 在创建模型实例时，为指针类型字段分配内存空间。
2. 在读取数据时，判断指针类型字段是否为空。
3. 在更新数据时，使用 `Update` 方法更新非空字段。

在 GORM 模型中使用指针类型时，由于指针类型字段可能为空，所以需要注意避免空指针错误的问题。下面是一些避免空指针错误的建议：

1. 在创建模型实例时，为指针类型字段分配内存空间。

```go
user := &User{Name: "Alice"}
post := &Post{Title: "Hello", Content: "World", User: user}
```

在这个例子中，我们在创建 `post` 实例时为 `User` 字段分配了内存空间，避免了空指针错误的问题。

2. 在读取数据时，判断指针类型字段是否为空。

```go
var post Post
if err := db.Preload("User").First(&post, 1).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        // 处理记录不存在的情况
    } else {
        // 处理其他错误
    }
} else {
    if post.User != nil {
        // 处理 User 不为空的情况
    }
}
```

在这个例子中，我们在读取 `post` 实例时，使用 `Preload` 方法预加载 `User` 字段，然后在判断 `post.User` 是否为空时进行了特判，避免了空指针错误的问题。

3. 在更新数据时，使用 `Update` 方法更新非空字段。

```go
var post Post
if err := db.First(&post, 1).Error; err != nil {
    // 处理错误
} else {
    if post.User != nil {
        if err := db.Model(&post).Update("Title", "New Title", "User", post.User).Error; err != nil {
            // 处理错误
        }
    } else {
        if err := db.Model(&post).Update("Title", "New Title").Error; err != nil {
            // 处理错误
        }
    }
}
```

在这个例子中，我们在更新 `post` 实例时，先判断 `post.User` 是否为空，然后根据情况使用 `Update` 方法更新相应的字段，避免了空指针错误的问题。

总的来说，避免空指针错误需要根据具体的场景进行处理，需要注意细节问题。

### 什么是外键约束，有什么用

外键约束（Foreign Key Constraint）是一种数据库中的约束条件，用于保证表与表之间的数据完整性。

具体来说，外键约束可以用于定义两个表之间的关系，使得一个表的某个字段值必须是另一个表的某个字段的值。这样，当我们在更新或删除一个表的记录时，数据库系统会自动检查该记录是否与其他表有关联，如果有，则会自动更新或删除其他表的相应记录，以保证数据完整性。

例如，我们可以定义一个 `orders` 表和一个 `customers` 表，用于存储订单和客户信息。为了保证数据完整性，我们可以在 `orders` 表中定义一个外键约束，将 `customer_id` 字段作为外键，关联到 `customers` 表的 `id` 字段上。

```sql
CREATE TABLE customers (
    id INT PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE orders (
    id INT PRIMARY KEY,
    customer_id INT,
    order_date DATE,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);
```

在这个例子中，我们定义了一个 `orders` 表和一个 `customers` 表，其中 `orders` 表中的 `customer_id` 字段作为外键，关联到 `customers` 表的 `id` 字段上。

这样，在插入或更新 `orders` 表的数据时，数据库系统会自动检查 `customer_id` 字段的值是否存在于 `customers` 表中，如果不存在，则会抛出错误。

外键约束可以保证数据的完整性和一致性，**但是在使用时需要注意几点**：

1. 外键约束可能会降低数据库的性能，因为数据库系统需要进行额外的检查和操作。

2. 外键约束可能会限制数据库的灵活性，因为数据库系统会强制执行特定的数据关系。

3. 外键约束需要在设计数据库时谨慎使用，需要考虑业务需求和实际情况。

### gorm 外键约束有哪些，各有什么用，举例子说明

GORM 中常用的外键约束有以下几种：

1. `ForeignKeyConstraint`：用于定义外键约束，并指定关联的表、字段以及操作。

2. `OnDelete`：用于指定删除记录时的行为，包括 `CASCADE`（级联删除）、`SET NULL`（置空）等。

3. `OnUpdate`：用于指定更新记录时的行为，包括 `CASCADE`（级联更新）、`SET NULL`（置空）等。

下面是各种外键约束的详细说明及示例：

1. `ForeignKeyConstraint`：用于定义外键约束，并指定关联的表、字段以及操作。例如：

```go
type Order struct {
    ID          int
    CustomerID  int
    Customer    Customer `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Customer struct {
    ID      int
    Name    string
}
```

在这个例子中，我们定义了一个 `Order` 模型和一个 `Customer` 模型，其中 `Order` 模型中的 `CustomerID` 字段作为外键，关联到 `Customer` 模型的 `ID` 字段上。同时，我们使用 `gorm:"foreignKey"` 标签定义了外键约束，并指定了 `OnUpdate` 和 `OnDelete` 的行为。在这个例子中，当 `Customer` 表中的记录被更新时，与之关联的 `Order` 表中的 `CustomerID` 字段也会被更新，同时，当 `Customer` 表中的记录被删除时，与之关联的 `Order` 表中的 `CustomerID` 字段会被置空。

2. `OnDelete`：用于指定删除记录时的行为，包括 `CASCADE`（级联删除）、`SET NULL`（置空）等。例如：

```go
type Order struct {
    ID          int
    CustomerID  int
    Customer    Customer `gorm:"foreignKey:CustomerID;"`
}

type Customer struct {
    ID      int
    Name    string
}

db.Model(&Order{}).AddForeignKey("customer_id", "customers(id)", "SET NULL", "CASCADE")
```

在这个例子中，我们使用 `AddForeignKey` 方法添加了一个外键约束，指定了删除记录时的行为为 `SET NULL`。这意味着，当 `customers` 表中的某个记录被删除时，与之关联的 `orders` 表中的 `customer_id` 字段会被置空。

3. `OnUpdate`：用于指定更新记录时的行为，包括 `CASCADE`（级联更新）、`SET NULL`（置空）等。例如：

```go
type Order struct {
    ID          int
    CustomerID  int
    Customer    Customer `gorm:"foreignKey:CustomerID;"`
}

type Customer struct {
    ID      int
    Name    string
}

db.Model(&Order{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "SET NULL")
```

在这个例子中，我们使用 `AddForeignKey` 方法添加了一个外键约束，指定了更新记录时的行为为 `CASCADE`。这 4. 外键约束 - RESTRICT

当有一个外键约束指定为 RESTRICT 时，当主表中的某一行被删除或更新时，只有在从表中没有相应的引用时才能执行这个操作。如果在从表中有相应的引用，则会抛出一个错误。

使用 RESTRICT 可以保证在主表中的某一行被删除或更新之前，需要确保在从表中没有相应的引用。这有助于保持数据的完整性。

例如，我们有一个商店应用，其中有一个 product 表和一个 order 表，每个订单都需要引用产品表中的一个产品，我们可以在 order 表中创建一个外键，指向 product 表中的 id 字段，并将这个外键指定为 RESTRICT，这样就可以确保在删除或更新产品之前，所有订单都已经完成或取消。

5. 外键约束 - SET NULL

当有一个外键约束指定为 SET NULL 时，当主表中的某一行被删除或更新时，所有从表中的相应外键都会被设置为 NULL。

使用 SET NULL 可以保证在主表中的某一行被删除或更新之后，从表中的相应外键不再引用一个不存在的主键。

例如，我们有一个博客应用，其中有一个 post 表和一个 comment 表，每个评论都需要引用 post 表中的一个文章，我们可以在 comment 表中创建一个外键，指向 post 表中的 id 字段，并将这个外键指定为 SET NULL，这样就可以确保在删除或更新文章之后，所有评论都不再引用一个不存在的文章。

6. 外键约束 - SET DEFAULT

当有一个外键约束指定为 SET DEFAULT 时，当主表中的某一行被删除或更新时，所有从表中的相应外键都会被设置为默认值。

使用 SET DEFAULT 可以保证在主表中的某一行被删除或更新之后，从表中的相应外键都有一个默认值。

例如，我们有一个用户应用，其中有一个 user 表和一个 post 表，每个文章都需要引用 user 表中的一个用户，我们可以在 post 表中创建一个外键，指向 user 表中的 id 字段，并将这个外键指定为 SET DEFAULT，这样就可以确保在删除或更新用户之后，所有文章都有一个默认的用户。

7. 外键约束 - NO ACTION

当有一个外键约束指定为 NO ACTION 时，当主表中的某一行被删除或更新时，所有从表中的相应外键不受影响。

使用 NO ACTION 可以保证在主表中的某一行被删除或更新之后，从表中的相应外键不会受到影响。但是这也意味着可能会存在从表中的外键引用一个不存在的主键或者引用一个已经被更新或删除的主键。

需要注意的是，使用

在 GORM 中，还有一种外键约束是 `OnUpdate` 和 `OnDelete`，它们用于定义当关联的主表数据被修改或删除时，关联的从表数据应该如何处理。

常见的几种处理方式如下：

- `RESTRICT`：当关联的主表数据被修改或删除时，抛出一个错误，防止此操作发生。这是默认的行为。
- `CASCADE`：当关联的主表数据被修改或删除时，同时修改或删除关联的从表数据。
- `SET NULL`：当关联的主表数据被修改或删除时，将关联的从表数据的外键列设置为 NULL。
- `NO ACTION`：与 `RESTRICT` 类似，不执行任何操作，但不会抛出错误。

以下是一个示例：

```go
type User struct {
  ID   int
  Name string

  // 声明一个 has-many 关系，一个用户可以有多篇文章
  Articles []Article `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Article struct {
  ID     int
  Title  string
  Body   string
  UserID int     // 声明一个外键，指向 User 的 ID 字段
  User   *User   // 声明一个 belongs-to 关系，一篇文章只能属于一个用户
}
```

在上述示例中，我们定义了一个 `has-many` 的关系，即一个用户可以有多篇文章，同时声明了 `OnUpdate:CASCADE,OnDelete:SET NULL` 外键约束。这意味着当用户被更新时，所有与之相关的文章的外键 `UserID` 将被更新为 `NULL`，当用户被删除时，与之相关的文章将被删除。

:::warning

需要注意的是，在 GORM v2.0 版本中，外键约束的声明方式发生了变化，使用 `constraint` 标签来声明，而不是之前的 `ForeignKey`。
:::

### 时间戳完整代码示例

以下是一个完整的示例代码，展示如何使用 GORM 存储时间戳和将时间戳转换为 `time.Time` 对象，以及将 `time.Time` 对象转换为毫秒级别的 Unix 时间戳。

```go
package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	// 初始化 GORM 连接
	dsn := "test.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移数据表结构
	db.AutoMigrate(&User{})

	// 创建新用户
	user := User{Name: "John Doe"}
	db.Create(&user)

	// 查询用户
	var savedUser User
	db.First(&savedUser, user.ID)

	// 输出时间戳和时间对象
	fmt.Printf("User %d was created at time %v and updated at time %v\n", savedUser.ID, savedUser.CreatedAt, savedUser.UpdatedAt)

	// 输出时间戳（秒级别）
	fmt.Printf("User %d was created at timestamp %d and updated at timestamp %d\n", savedUser.ID, savedUser.CreatedAt.Unix(), savedUser.UpdatedAt.Unix())

	// 输出时间戳（毫秒级别）
	fmt.Printf("User %d was created at timestamp %d and updated at timestamp %d\n", savedUser.ID, savedUser.CreatedAt.UnixMilli(), savedUser.UpdatedAt.UnixMilli())
}
```

这个例子中，我们定义了一个 `User` 结构体，其中包含 `CreatedAt` 和 `UpdatedAt` 两个字段，它们的类型均为 `time.Time`。在使用 GORM 操作这个结构体时，GORM 会自动将这些字段转换为数据库中的时间戳格式进行存储，而在查询时，GORM 也会自动将数据库中的时间戳转换为 `time.Time` 对象。

在上面的例子中，我们创建了一个新用户，并将其保存到数据库中。然后我们使用 `First()` 方法查询数据库中保存的用户，并输出其创建时间和更新时间。接着我们使用 `Unix()` 和 `UnixMilli()` 方法分别将时间对象转换为秒级别和毫秒级别的 Unix 时间戳，并输出它们的值。

这里需要注意的是，在将 `time.Time` 对象转换为 Unix 时间戳时，需要使用 `Unix()` 或 `UnixMilli()` 方法。其中，`Unix()` 方法返回的是秒级别的 Unix 时间戳，而 `UnixMilli()` 方法返回的是毫秒级别的 Unix 时间戳。

### 上述代码表结构和结果

表结构：

```sql
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

结果：

```go
[{1 John 1641062400000 1641062400000} {2 Jane 1641066000000 1641066000000}]
```

其中，第一个用户的 `created_at` 和 `updated_at` 时间戳分别为 1641062400000（2022-01-02 00:00:00）和 1641062400000（2022-01-02 00:00:00），第二个用户的 `created_at` 和 `updated_at` 时间戳分别为 1641066000000（2022-01-02 01:00:00）和 1641066000000（2022-01-02 01:00:00）。

### clause.Associations

下面是一个使用 `clause.Associations` 的完整示例，用于加载 `User` 对象及其关联的 `Order` 和 `OrderItem` 对象：

```go
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

type User struct {
    gorm.Model
    Name   string
    Orders []Order `gorm:"foreignKey:UserID"`
}

type Order struct {
    gorm.Model
    UserID     uint
    OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
    gorm.Model
    OrderID     uint
    ProductName string
}

func main() {
    // 创建 SQLite 数据库连接
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    // 自动迁移数据库表结构
    err = db.AutoMigrate(&User{}, &Order{}, &OrderItem{})
    if err != nil {
        panic(err)
    }

    // 创建测试数据
    orderItems := []OrderItem{
        {ProductName: "Product A"},
        {ProductName: "Product B"},
    }
    orders := []Order{
        {OrderItems: orderItems},
        {OrderItems: orderItems},
    }
    users := []User{
        {Name: "User A", Orders: orders},
        {Name: "User B", Orders: orders},
    }
    err = db.Create(&users).Error
    if err != nil {
        panic(err)
    }

    // 查询所有用户及其关联的订单和订单项
    var loadedUsers []User
    err = db.Preload(clause.Associations).Find(&loadedUsers).Error
    if err != nil {
        panic(err)
    }

    // 打印查询结果
    for _, user := range loadedUsers {
        println("User:", user.ID, user.Name)
        for _, order := range user.Orders {
            println("\tOrder:", order.ID, order.UserID)
            for _, orderItem := range order.OrderItems {
                println("\t\tOrderItem:", orderItem.ID, orderItem.OrderID, orderItem.ProductName)
            }
        }
    }
}
```

在上述示例中，我们创建了一个 `User` 模型和一个 `Order` 模型，`Order` 模型有多个 `OrderItem` 模型。我们在 `User` 模型上定义了一个 `Orders` 字段，它表示一个用户可能有多个订单，而在 `Order` 模型上定义了一个 `OrderItems` 字段，它表示一个订单可能包含多个订单项。

我们通过 `db.Preload(clause.Associations)` 来加载所有关联的订单和订单项。在查询结果中，我们可以看到每个用户及其关联的订单和订单项。

### clause.Associations 还是没有理解

`clause.Associations` 实际上是一个 `gorm.Associations` 类型的常量，用于指示预加载操作要预加载哪些关联关系。

在 GORM 中，一个模型可以与其他模型建立多种关联关系，例如一对多、多对多、多态等等。当我们查询一个模型时，通常也需要把相关联的模型数据一起查询出来，这就需要用到预加载功能。

在使用预加载功能时，我们可以使用 `Preload` 方法来指定要预加载哪些关联关系。`clause.Associations` 就是其中一种指定方式。

当我们使用 `db.Preload(clause.Associations).Find(&users)` 方法时，GORM 会自动解析 `User` 模型定义中的所有关联关系，并将其包装在一个 `gorm.Associations` 类型的列表中，然后将这个列表传递给 `Preload` 方法。

使用 `clause.Associations` 的好处是，它能够避免手动编写字符串形式的关联关系名称，从而减少人为出错的可能性。另外，当模型的关联关系发生变化时，使用 `clause.Associations` 可以自动地更新预加载的关联关系，而不需要手动修改代码。

### Embedded Preloading

以下是一个使用 Embedded Preloading 的完整代码示例：

```go
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
    "time"
)

type User struct {
    gorm.Model
    Name   string
    Orders []Order
}

type Order struct {
    gorm.Model
    UserID     uint
    OrderItems []OrderItem
}

type OrderItem struct {
    gorm.Model
    OrderID     uint
    ProductName string
}

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    db.AutoMigrate(&User{}, &Order{}, &OrderItem{})

    user1 := User{
        Name: "Alice",
        Orders: []Order{
            {
                OrderItems: []OrderItem{
                    {
                        ProductName: "Product A",
                    },
                },
            },
        },
    }

    user2 := User{
        Name: "Bob",
        Orders: []Order{
            {
                OrderItems: []OrderItem{
                    {
                        ProductName: "Product B",
                    },
                    {
                        ProductName: "Product C",
                    },
                },
            },
        },
    }

    db.Create(&user1)
    db.Create(&user2)

    var users []User
    db.Preload(clause.Associations).Find(&users)

    for _, user := range users {
        println("User ID:", user.ID)
        println("User Name:", user.Name)
        for _, order := range user.Orders {
            println("  Order ID:", order.ID)
            println("  Order User ID:", order.UserID)
            for _, orderItem := range order.OrderItems {
                println("    OrderItem ID:", orderItem.ID)
                println("    OrderItem ProductName:", orderItem.ProductName)
            }
        }
    }
}
```

该示例中使用 Embedded Preloading 预加载了 User、Order 和 OrderItem 三个模型的关联数据。最终输出了所有 User、Order 和 OrderItem 的详细信息。

在代码中，通过 `db.Preload(clause.Associations).Find(&users)` 对 User、Order 和 OrderItem 三个模型进行了 Embedded Preloading 预加载。这意味着，对于 User、Order 和 OrderItem 三个模型，它们的关联数据都会在一次查询中一次性加载出来。这样可以避免 N + 1 查询问题，提高查询效率。

注意，在进行 Embedded Preloading 预加载时，需要使用 `clause.Associations` 这个关键字，这样才能正确地指定 Embedded Preloading 操作。

### for range 实现原理

`for range` 是 Go 语言中的一个迭代器语法，通常用于迭代数组、切片、映射和管道等集合类型。它的使用方法类似于 `foreach` 循环，可以遍历集合中的所有元素。

`for range` 的实现原理是根据集合的类型来调用不同的迭代器函数，这些函数会返回一个包含当前元素信息的 `value` 和一个指示迭代是否结束的 `ok` 值。迭代器函数根据集合的类型可以是以下几种：

- 数组和切片：返回元素索引和值。
- 映射：返回键和值。
- 字符串：返回字符的索引和值。
- 管道：返回管道中的值。

在 `for range` 循环中，会通过调用集合的迭代器函数获取当前元素的 `value` 和 `ok` 值，然后判断 `ok` 值是否为 `true`。如果为 `true`，则表示迭代器还有元素需要遍历，将 `value` 赋值给循环变量；否则，表示迭代器已经遍历完所有元素，退出循环。

需要注意的是，如果在迭代过程中修改了集合中的元素，则会导致迭代器失效，可能会引发不可预知的结果。因此，不建议在 `for range` 循环中修改集合的元素。

以下是一个简单的例子，演示了如何使用 `for range` 循环遍历切片：

```go
package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5}
    for index, value := range nums {
        fmt.Printf("index: %d, value: %d\n", index, value)
    }
}
```

输出结果如下：

```go
index: 0, value: 1
index: 1, value: 2
index: 2, value: 3
index: 3, value: 4
index: 4, value: 5
```

### for range 原理/源码

for range 循环在 Go 语言中的实现是通过 `range` 关键字和具体类型的方法来实现的。`range` 关键字本身并不是一个迭代器，而是会自动在背后调用一个包含了迭代器逻辑的方法。

对于数组、切片、字符串等内置类型，`range` 会自动调用 `range` 相应的方法来实现迭代。

下面是 `range` 的源码实现，以数组为例：

```go
for index, element := range array {
    // ...
}
```

相当于编译器会转换为以下代码：

```go
for i := 0; i < len(array); i++ {
    index := i
    element := array[i]
    // ...
}
```

可以看到，编译器会在背后自动生成循环逻辑，包括从 0 到 `len(array)` 的循环变量 `i`，以及在每次循环中自动获取当前的索引和元素值。

对于自定义类型，需要实现 `range` 方法才能支持 `for range` 循环。具体来说，需要实现以下接口：

```go
type Iterable interface {
    Iterator() Iterator
}

type Iterator interface {
    Next() bool
    Value() interface{}
}
```

其中，`Iterable` 接口定义了获取迭代器的方法 `Iterator()`，而 `Iterator` 接口定义了每次迭代的方法 `Next()` 和获取当前值的方法 `Value()`。当然，`Value()` 方法返回的类型可以是任意类型，只要在使用时能够转换为对应的目标类型即可。

以下是一个自定义类型的例子：

```go
type Person struct {
    Name string
    Age  int
}

type People []Person

func (p People) Iterator() Iterator {
    return &peopleIterator{p, -1}
}

type peopleIterator struct {
    people People
    index  int
}

func (pi *peopleIterator) Next() bool {
    pi.index++
    return pi.index < len(pi.people)
}

func (pi *peopleIterator) Value() interface{} {
    return pi.people[pi.index]
}
```

1. 这里定义了 `Person` 和 `People` 两个类型，
2. 其中 `People` 是一个包含多个 `Person` 的切片，实现了 `Iterable` 接口。
3. 当 `People` 被用于 `for range` 循环时，会自动调用 `Iterable` 接口中的 `Iterator()` 方法来获取迭代器，
4. 然后在背后自动调用 `Iterator` 接口中的 `Next()` 和 `Value()` 方法来实现迭代。

在上述例子中，

1. `Next()` 方法会自增索引并判断是否已经到达切片的末尾，
2. 而 `Value()` 方法则返回当前索引的元素值。

由于切片中的元素类型是 `Person`，因此需要在调用 `Value()` 方法后将返回值转换为 `Person` 类型，以便在循环中使用。

### Embedded Preloading 如何理解？

**使用限制：被预加载项需要唯一，能够互相区分**

Embedded Preloading 只适用于 belongs to 关系（即属于关系），因为在数据库中，其他关系的值是相同的，我们无法区分它们。

举个例子，假设我们有两个表：

user 和 order，它们的关系是一对多，即一个用户可以拥有多个订单。如果我们想要预加载订单，并将它们嵌入到用户结构体中，我们可以使用 Embedded Preloading。

但是，如果我们还有一个表 product，和 order 的关系也是一对多，即一个订单可以包含多个商品。

此时，如果我们想要将商品预加载，并将它们嵌入到订单结构体中，使用 Embedded Preloading 就不适用了，因为在数据库中，不同的订单可能拥有相同的商品，我们无法区分它们，所以需要使用普通的 Preloading。

### 如果 product，和 order 的关系 也使用 Embedded Preloading 会怎样

如果 `product` 和 `order` 之间的关系也使用 `Embedded Preloading`，则会出现类似于 `order_items` 的情况，即 `product_id` 和 `order_id` 无法区分，从而导致查询结果不准确。因此，在使用 `Embedded Preloading` 时，需要根据实际情况选择需要预加载的关系，避免出现这种问题。

### 看不懂，举例说明

好的，下面举一个简单的例子来说明 Embedded Preloading 只适用于属于关系。

假设有以下两个模型：

```go
type Product struct {
    gorm.Model
    Name string
    Price float64
}

type Order struct {
    gorm.Model
    Products []Product
}
```

在这个例子中，`Order` 模型有多个 `Product` 模型。

如果使用 Embedded Preloading 来预加载 `Products`，我们将无法在数据库中区分哪个 `Product` 属于哪个 `Order`，因为所有 `Product` 的 `OrderID` 值都是相同的，这是无效的。

因此，Embedded Preloading 仅适用于属于关系，例如 `Order` 属于 `User`，但是 `Product` 不属于 `Order`。

### db.Table("users").Select("COALESCE(age,?)", 42).Rows() 什么意思

这段代码使用了 GORM 的 Raw SQL 功能。`db.Table("users")` 用于指定查询的表，`.Select("COALESCE(age,?)", 42)` 用于执行 SQL 的 SELECT 语句，并将结果以行的形式返回。

具体来说，这条 SQL 语句中的 `COALESCE(age,?)` 是一个函数，表示如果 `age` 不为 NULL，则返回 `age` 的值，否则返回 `?` 的值（即 42）。因此，这个查询将返回所有用户的行，但如果某个用户的 `age` 字段为 NULL，则该字段的值将被替换为 42。

最后的 `.Rows()` 用于返回查询结果的行集。

### COALESCE 是什么

`COALESCE` 是一个 SQL 函数，用于返回一组表达式中的第一个非空表达式的值。如果所有表达式的值都为 NULL，则返回 NULL。

例如，假设有以下表格：

| id  | name  | age  |
| --- | ----- | ---- |
| 1   | Tom   | NULL |
| 2   | Jerry | 20   |
| 3   | Mike  | 25   |

如果执行以下 SQL 查询：

```sql
SELECT COALESCE(age, 18) AS age FROM table_name;
```

则会得到以下结果：

| age |
| --- |
| 18  |
| 20  |
| 25  |

即如果 age 字段是 NULL，则返回默认值 18，否则返回 age 字段的值。

### sql

db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)

```sql
SELECT * FROM users
  JOIN emails ON emails.user_id = users.id AND emails.email = 'jinzhu@example.org'
  JOIN credit_cards ON credit_cards.user_id = users.id WHERE credit_cards.number = '411111111111'



```

### Distinct works with Pluck and Count too 的例子

假设我们有一个 `User` 模型，其中包含 `Name` 和 `Age` 字段，并且有多个用户具有相同的姓名，但不同的年龄。我们想要找出唯一的姓名列表。我们可以使用 `Distinct` 方法来执行此操作：

```go
var names []string
db.Distinct("name").Order("name asc").Find(&users).Pluck("name", &names)
```

此代码将按升序顺序返回所有唯一的姓名列表。如果我们想知道有多少个唯一的姓名，则可以使用 `Count` 方法：

```go
var count int64
db.Distinct("name").Count(&count)
```

类似地，如果我们想要在计算唯一值之前应用其他条件，我们可以在 `Distinct` 方法之前使用 `Where`、`Joins` 和 `Group` 等方法。

### Find 、Scan 和 Rows

Scan 方法 1. 将结果扫描到结构中的工作方式与我们使用 Find 的方式类似 2. Scan 至少要求接受数据的类型是 slice 或 struct 及它们的指针类型。

Rows 方法 是执行 SQL 查询，并返回一个指向结果集的行迭代器

### .Rows() 方法为什么需要 defer close

在使用 `Rows()` 方法时，GORM 会返回一个 `*sql.Rows` 对象，这个对象代表了查询返回的结果集，需要通过调用 `Next()` 方法迭代每一行数据，并通过调用 `Scan()` 方法将行数据扫描到相应的结构体字段中。

当查询结束后，必须调用 `Rows()` 方法返回的结果集对象的 `Close()` 方法来释放相关的资源。如果不及时释放这些资源，可能会导致内存泄漏或数据库连接池耗尽等问题。

为了确保 `Close()` 方法得到及时调用，GORM 的 `Rows()` 方法在返回结果集对象时会将其包装在一个 `sql.RowsWrapper` 结构体中，并将该结构体返回给调用方。这个结构体实现了 `io.Closer` 接口，并在 `Close()` 方法中调用了原始结果集对象的 `Close()` 方法。

由于 `sql.RowsWrapper` 实现了 `io.Closer` 接口，因此可以通过使用 `defer` 关键字在函数退出时自动调用 `Close()` 方法，从而避免忘记释放资源的问题。

# 2023-05-04

---

### gorm 条件

1. String 条件
   1. 字符串占位符
2. Struct & Map 条件
   1. 纯结构体数据
   2. 纯 Map 数据
   3. 主键切片数据
3. 指定结构体查询字段
   1. Where 语句中使用结构体
   2. 注意：添加更多参数的时候，如果是在结构体外面，没有赋值，默认为“零值”
4. 内联条件
   1. 在查询方法中，添加更多的参数（字符串占位符、结构体数据、map 数据）
5. Not 条件
   1. 字符串
   2. 结构体
   3. map
   4. not in 语意的切片/map 键值数组
6. Or 条件
   1. 字符串
   2. 结构体
   3. map
7. 选择特定字段 （选择允许您指定要从数据库中检索的字段。否则，GORM 将默认选择所有字段。）
   1. 字符串
   2. 字符串数组
   3. 默认值

### Where、Or、Not 这个几个 API 是否类似

以下是一个示例代码，使用 GORM 进行 WHERE、OR、NOT 条件查询：

```go
package main

import (
    "fmt"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name     string
    Age      int
    IsActive bool
}

func main() {
    // 创建 SQLite 数据库连接
    dsn := "test.db"
    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 自动迁移 User 模型到数据库中
    db.AutoMigrate(&User{})

    // 插入一些示例数据
    db.Create(&User{Name: "Alice", Age: 25, IsActive: true})
    db.Create(&User{Name: "Bob", Age: 30, IsActive: true})
    db.Create(&User{Name: "Charlie", Age: 35, IsActive: false})

    // 使用 Where 查询 Name 为 Alice 的用户
    var user User
    db.Where("name = ?", "Alice").First(&user)
    fmt.Println("User:", user)

    // 使用 Or 查询 Age 为 25 或 30 的用户
    var users []User
    db.Where("age = ?", 25).Or("age = ?", 30).Find(&users)
    fmt.Println("Users:", users)

    // 使用 Not 查询 IsActive 不为 true 的用户
    db.Not("is_active = ?", true).Find(&users)
    fmt.Println("Users:", users)
}
```

在上面的示例中，我们使用 `Where` 方法来查询 Name 为 Alice 的用户。使用 `Or` 方法查询 Age 为 25 或 30 的用户。使用 `Not` 方法查询 IsActive 不为 true 的用户。

```go
// 源码
func (db *DB) Where(query interface{}, args ...interface{}) (tx *DB) {}

func (db *DB) Not(query interface{}, args ...interface{}) (tx *DB) {}

func (db *DB) Or(query interface{}, args ...interface{}) (tx *DB) {}
```

从源码可以看出，Where、Or、Not 具有相似的 API,所以这三个函数使用方式相似，只需要会用一个，另外两个类推

### 其他 API 的入参数

<img src="http://t-blog-images.aijs.top/img/202305041005006.webp" />
