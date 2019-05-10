当for 和 select结合使用时，break语言是无法跳出for之外的，因此若要break出来，这里需要加一个标签，使用goto， 或者break 到具体的位置

code 如下:
```go
func test(){
    i := 0
    for {
        select {
        case <-time.After(time.Second * time.Duration(2)):
            i++
            if i == 5{
                fmt.Println("break now")
                break 
            }
            fmt.Println("inside the select: ")
        }
        fmt.Println("inside the for: ")
    }
}
```
解决方法一：使用golang中break的特性，在外层for加一个标签:

```go
func test(){
    i := 0
    ForEnd:
    for {
        select {
        case <-time.After(time.Second * time.Duration(2)):
            i++
            if i == 5{
                fmt.Println("break now")
                break ForEnd
            }
            fmt.Println("inside the select: ")
        }
        fmt.Println("inside the for: ")
    }
}
```

解决方法二: 使用goto直接跳出循环
```go
func test(){
    i := 0
 
    for {
        select {
        case <-time.After(time.Second * time.Duration(2)):
            i++
            if i == 5{
                fmt.Println("break now")
                goto ForEnd
            }
            fmt.Println("inside the select: ")
        }
        fmt.Println("inside the for: ")
    }
    ForEnd：
}
```

