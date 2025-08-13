## 不同类型声明方式对比
### slice 和 map
切以片为例，不同声明方式对比：

| 声明方式 | 底层数组 | 优缺点 | 使用场景 | 备注 |
| -------- | ---- | ---- | ---- | -------- |
|`var s []T`| nil | 全局变量声明，且值为nil | 全局变量、函数签名 | 1. 具名返回值 `[] T` 和 `[]int(nil)` 等同于该方式<br />2. 不分配底层数组内存，仅切片结构体本身占用少量内存<br />3. `nil` 切片可以直接append |
|`s := []T{}`| 空数组 | 使用方便，但切片扩容时性能下降 | 局部变量声明 | 1. 最常用方式<br />2. 序列化后为`[]`，而不是`nil` |
|`s := make([]T, N)`| 预分配N个零值 | 声明长度后可有效避免切片扩容和数组越界访问，但如果长度>0且进行append操作，容易踩坑 | 已知容量大小，如切片值的拷贝 | 如果确定在原值上修改，再声明长度，否则使用默认长度为0 |

### Channel 

| 声明方式 | 特性 | 使用场景 | 备注 |
| -------- | ---- | -------- | ---- |
| `var ch chan T` | nil channel | 永远阻塞 | 常用于select中禁用某个case |
| `ch := make(chan T)` | 无缓冲channel | 同步通信 | 发送方会阻塞直到接收方就绪 |
| `ch := make(chan T, N)` | 有缓冲channel | 异步通信 | 缓冲区未满时发送不阻塞 |

## 类型转换
### []byte 使用指南
在Go中，string 底层是 `[]byte`，可以转换为 `[]rune` 来处理 Unicode 字符。其中`[]byte` 适用于 ASCII 字符，`[]rune` 适用于 Unicode 字符。`string(b)` 和 `[]byte(s)` 是常用的 string 与 byte 的转换方式。

#### []byte 常用操作对比

| 操作方式 | 性能 | 内存分配 | 使用场景 | 备注 |
| -------- | ---- | -------- | -------- | ---- |
| `[]byte(str)` | 较慢 | 分配新内存 | string转[]byte | 会复制数据，安全但有开销 |
| `bytes.NewBuffer()` | 中等 | 按需分配 | 频繁读写操作 | 适合构建字符串 |
| `unsafe.Slice()` | 最快 | 无分配 | 性能敏感场景 | 危险，需要确保string不被GC |

#### []byte vs string 转换

| 转换方式 | 开销 | 安全性 | 何时使用 |
| -------- | ---- | ------ | -------- |
| `string(b)` | O(n) 复制 | 安全 | 一般场景 |
| `unsafe.String()` | O(1) 无复制 | 不安全 | 性能关键路径 |

#### []byte 常用库函数

| 函数包 | 主要功能 | 典型用法 | 性能特点 |
| ------ | -------- | -------- | -------- |
| `bytes` | 字节操作 | `bytes.Contains()`, `bytes.Split()` | 针对[]byte优化 |
| `strings` | 字符串操作 | `strings.Contains()` | 需要转换，有开销 |

#### 实际使用建议

| 场景 | 推荐方案 | 原因 |
| ---- | -------- | ---- |
| JSON解析 | `[]byte` | 避免string转换开销 |
| 网络I/O | `[]byte` | 直接读写，无需转换 |
| 文本处理 | `string` | API更丰富，使用方便 |
| 性能敏感 | `[]byte` + `bytes` 包 | 避免不必要的内存分配 |

#### 常见陷阱

```go
// ❌ 频繁转换，性能差
for _, str := range strs {
    data := []byte(str)  // 每次都分配内存
    // process data...
}

// ✅ 复用buffer，性能好
var buf []byte
for _, str := range strs {
    buf = buf[:0]        // 重置长度，保留容量
    buf = append(buf, str...)
    // process buf...
}
```

#### 内存优化技巧

| 技巧 | 说明 | 示例 |
| ---- | ---- | ---- |
| 预分配容量 | 避免频繁扩容 | `make([]byte, 0, 1024)` |
| 复用slice | 重置长度保留容量 | `buf = buf[:0]` |
| 使用sync.Pool | 对象池复用 | 高并发场景下复用[]byte |

### 数字与字符串间的转换

#### strconv 包 - 标准转换方式

| 函数 | 功能 | 示例 | 性能 | 备注 |
| ---- | ---- | ---- | ---- | ---- |
| `strconv.Itoa(i)` | int转string | `strconv.Itoa(123)` → `"123"` | 快 | 等价于`Formatint(int64(i), 10)` |
| `strconv.Atoi(s)` | string转int | `strconv.Atoi("123")` → `123, nil` | 快 | 返回`(int, error)` |
| `strconv.FormatInt(i, base)` | int64转string(指定进制) | `FormatInt(255, 16)` → `"ff"` | 快 | 支持2-36进制 |
| `strconv.ParseInt(s, base, bitSize)` | string转int64 | `ParseInt("ff", 16, 64)` → `255` | 快 | 最灵活的解析方式 |
| `strconv.FormatFloat(f, fmt, prec, bitSize)` | float转string | `FormatFloat(3.14, 'f', 2, 64)` → `"3.14"` | 快 | fmt: 'f','e','g'等 |
| `strconv.ParseFloat(s, bitSize)` | string转float | `ParseFloat("3.14", 64)` → `3.14` | 快 | bitSize: 32或64 |

#### fmt 包 - 通用但较慢的方式

| 函数 | 功能 | 示例 | 性能 | 适用场景 |
| ---- | ---- | ---- | ---- | ---- |
| `fmt.Sprintf()` | 任意类型转string | `fmt.Sprintf("%d", 123)` | 慢 | 复杂格式化 |
| `fmt.Sprint()` | 默认格式转string | `fmt.Sprint(123)` | 慢 | 简单调试 |

#### 性能对比

| 转换方式 | 性能排序 | 内存分配 | 推荐度 |
| -------- | -------- | -------- | ------ |
| `strconv.Itoa()` | 最快 | 最少 | ⭐⭐⭐⭐⭐ |
| `strconv.FormatInt()` | 快 | 少 | ⭐⭐⭐⭐ |
| `fmt.Sprint()` | 慢 | 多 | ⭐⭐ |
| `fmt.Sprintf()` | 最慢 | 最多 | ⭐ |

#### 常见使用场景

| 场景 | 推荐方法 | 示例 |
| ---- | -------- | ---- |
| 简单int转string | `strconv.Itoa()` | `strconv.Itoa(42)` |
| 简单string转int | `strconv.Atoi()` | `strconv.Atoi("42")` |
| 十六进制转换 | `strconv.FormatInt()/ParseInt()` | `strconv.FormatInt(255, 16)` |
| 浮点数转换 | `strconv.FormatFloat()/ParseFloat()` | `strconv.FormatFloat(3.14, 'f', 2, 64)` |
| 复杂格式化 | `fmt.Sprintf()` | `fmt.Sprintf("value: %d", 42)` |

#### 错误处理最佳实践

```go
// ✅ 推荐：检查错误
if num, err := strconv.Atoi(str); err != nil {
    return fmt.Errorf("invalid number: %w", err)
} else {
    // 使用 num
}

// ❌ 不推荐：忽略错误
num, _ := strconv.Atoi(str)  // 可能panic或得到意外值
```

#### 性能优化技巧
```go
// ❌ 性能差：使用fmt
result := fmt.Sprintf("%d", number)

// ✅ 性能好：使用strconv
result := strconv.Itoa(number)

// ✅ 批量转换时复用buffer
var buf []byte
for _, num := range numbers {
    buf = strconv.AppendInt(buf[:0], int64(num), 10)
    // 使用 buf
}
```