##### 每个项目都在开发工具自己的gopath中开发

#### 导入的每个包需要有init函数，就像py一样,init 函数可以为空
import _ “fmt” 表示不使用该包，而是只是使用该包的init函数，并不显示的使用该包的其他内容