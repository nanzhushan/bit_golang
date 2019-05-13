### 循环
golang 只有一种循环结构--for循环. 它有五种方式:

* 1）由init statement(初始语句)、condition statement(条件语句)和post statement(循环语句)组成
```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

}
```

* 2)由condition statement和post statement组成
```go
package main

import "fmt"

func main() {
	var i = 0
	for ; i < 10; i++ {   // 条件语句和修饰语句组成
		fmt.Printf("%d ", i)
	}

}
```

* 3)由condition statement组成
```go
package main

import "fmt"

func main() {
	var i int
	for i <= 10 {  // 只有条件语句
		fmt.Println(i)
		i++
	}
}

```

* 4）用range关键字遍历集合
```go
package main

import "fmt"

func main() {
	str := "123ABCabc学习"
	for i, value := range str {
		fmt.Printf("第%d位的字符是 %c \n", i, value)
	}
}
```
* 5）无限循环
```
for {
    // statement
}

```
