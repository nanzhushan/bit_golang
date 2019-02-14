## go不是面向对象的语言，那如何实现面向对象的多态封装和继承了?
### 实现继承:
```
package main

type Person struct {
}
func (p *Person) Say() {
    println("I'm a person.")
}

//通过结构体和方法的组合实现继承
type Student struct {
    Person
}

func main() {
    var s Student
    s.Say()
}
```
