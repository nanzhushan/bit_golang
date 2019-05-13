### 循环
golang 只有一种循环结构--for循环. 它有四种方式:

* 1）由init statement(初始语句)、condition statement(条件语句)和post statement(循环语句)组成
```go
for i := 0; i < 10; i++ {
    // statement 
}
```

* 2)由condition statement和post statement组成
```go
for ; i < 10; i++ {
    // statement
}
```

* 3)由condition statement组成
```go
// 带有分号的表示方式
for ; i < 10; {
    // statement
}

// 另一种表示方式，像while循环
for i < 10 {
    // statement
}
```

* 4）无限循环
```
for {
    // statement
}
```
