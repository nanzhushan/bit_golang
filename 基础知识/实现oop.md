## go不是面向对象的语言，那如何实现面向对象的多态封装和继承了?
方法: 函数+接收者

* 1.方法其实是接收者(或者说receiver)上的一个函数。特殊类型的函数罢了
* 2.接受者必须是内置类型或者结构体的一个值或者指针

##### Go的方法 vs PHP的类：
php:
```
    class person {
        public function PrintName($name) {
            print_r($name);
        }
    }

    $person = new person();
    $person->PrintName('咖啡色的羊驼');
```
go:
```
    package main

    import (
        "fmt"
    )
    // 定义一个结构体，当作方法接收者
    type Person struct {}      

    // 方法声明其实就是在函数的基础上在func后面多增加定义接收者
    // func (r ReceiverType) funcName(parameters) (results)
    func (p Person) PrintName(name string) {
        fmt.Println(name)
    }

    func main() {
        cbs := Person {}            // 字面量的方式定义一个person类型
        cbs.PrintName("咖啡色的羊驼") // 直接调用方法
    }

```
##### Go的继承 vs PHP的继承
php例子就不写了，直接上go:
```
    package main

    import (
        "fmt"
    )

    type Person struct {
    }

    type SuperMan struct {
        Person      // 通过结构体内嵌包含的方式，来实现继承一样的效果
        Age int 
    }

    func (p Person) PrintName(name string) {
        fmt.Println(name)
    }

    func (s SuperMan) PrintAge() {
        fmt.Println(s.Age)
    }

    func main() {
        cbs := SuperMan{Age:18}
        cbs.PrintName("咖啡色的羊驼")       // 输出"咖啡色的羊驼"
        cbs.Person.PrintName("咖啡色的羊驼")// 输出"咖啡色的羊驼"
        cbs.PrintAge()                    // 输出18
    }

```

##### go的重载 vs php的重载
go里头没有重载

##### go的封装vs php的封装
首字母大写就是可以从包中导出(public),go语言中封装的单元是包,而不是类型.无论在函数内的代码，还是方法内，或者结构体内的字段对于同一个包中的所有代码都是可见的
```
    type Person struct {
        age float64
    }

    func (s Person) PrintAge() {
        fmt.Println(s.age)
    }

    func main() {
        cbs := Person{age:18}
        cbs.PrintAge()      
        // 输出18
        // 说明"方法当中可以访问结构当中的私有字段"
        // go语言中封装的单元是包，而不是类型,小写字母开头就表示私有属性
    }
```