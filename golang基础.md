
## ���
```
uname -r
env GOOS=linux GOARCH=386 go build main.go
```
����ϵͳ����Ȼ����д��

##�����﷨
### 
* go ��ǿ��������(����dockerһ������Դʹ����������ô��)

### ����
* Go���������и� main() ����
* �����ڶ���ı�����Ϊ�ֲ�����
* �����ⶨ��ı�����Ϊȫ�ֱ���
* ���������еı�����Ϊ��ʽ����

1)`:=` �ṹ���var���б�����ֵ����ʹ���ں����⡣ 


## ����

```
var aa int32 = 10
fmt.Print("type:",reflect.TypeOf(aa))   // �ж�����, java ����instanceof

var aa int32 = 10
_=aa                    // �����������ʹ�ã���Ҫ���� �ձ�ʶ��
```

## ����
```
// һά����
var team [3]string
team[0] = "hammer"
team[1]="fdsaf"
print(team[0])
```
### �ṹ��
�ṹ�����һ���ֶεļ���
```
type aa struct {
	x int
	y int
}

func main()  {
	v := aa{1,4}
	v.x = 4    
	println(v.x)  // ͨ����ŷ���
	
}
```
## ָ��
```
 Go ����ָ�롣 ָ�뱣���˱������ڴ��ַ��
���� *T ��ָ������ T ��ֵ��ָ�롣����ֵ�� `nil`��

var p *int

& ���Ż�����һ��ָ�������ö����ָ�롣

i := 42
p = &i

* ���ű�ʾָ��ָ��ĵײ��ֵ��
fmt.Println(*p) // ͨ��ָ�� p ��ȡ i
*p = 21         // ͨ��ָ�� p ���� i

��Ҳ����ͨ����˵�ġ�������á��򡰷�ֱ�����á���
```
## �����slice
```
package main
import "fmt"
func main()  {
	var a[2] string   // ��javaһ��Ҫ�������鳤��
	a[0]="hello"
	a[1]="dd"
	fmt.Print(a)

	// slice �Ǹ������ֵ,���Բ����峤��.append �������Ԫ�أ�����py��list
	p:=[]int{2,3,5,7}
	fmt.Print(p)
	fmt.Print(len(p))
}
```
������������ַ�ʽ
```
func main()  {
	// ������������
	var aa = [2]string{"fdaf","fdaf"}   // �����ʼ��,����ó���
	var bb =[...]string{"cc","rr"}    // ����Ϊ�ɱ䳤��
	var cc = [3]string{0:"tom",1:"18"}    //  �����ַ�ʽ����
	print(aa[0])
	print(len(bb))
	print("----",cc[0])

}
```
## map��ֵ��
����: map[KeyType]ValueType
```
m := make(map[string]string)   // �������ֵ������
m["name"] =  "knight"
fmt.Print(m["name"])
```


## ���밲װ
* go install �ǽ����� GOPATH �ϵģ��޷��ڶ�����Ŀ¼��ʹ�� go install��
* GOPATH �µ� bin Ŀ¼���õ���ʹ�� go install ���ɵĿ�ִ���ļ�����ִ���ļ������������ڱ���ʱ�İ�����
* go install ���Ŀ¼ʼ��Ϊ GOPATH �µ� bin Ŀ¼���޷�ʹ��-o���Ӳ��������Զ��塣
* GOPATH �µ� pkg Ŀ¼���õ��Ǳ����ڼ���м��ļ���

## ѭ�����
```
func main() {
    for i := 0; i < 5; i++ {
        fmt.Printf("This is the %d iteration\n", i)
    }
}
```
## ����
1) ����(��������˷���������Ҫʹ��return�����û�ж��巵�����;Ͳ���Ҫ)
```
func main()  {
	Greting("uu","77")
}

func Greting(x string,y string) string{   // ���崫��Ĳ��������Լ���������
	print("��ף ",x,y)
	return x+y
}
```
����
```
func main()  {
	Greting("uu","77")
}

func Greting(x string,y string){   // ���崫��Ĳ��������Լ���������
	print("��ף ",x,y)
}
```
2)����ɱ����
```
func main()  {
	Greting("uu","77")
}
func Greting(x ...string) string{   // ����ɱ����
	print(x[0],x[1])         // Ĭ����������д���
	//print(x)
	return "haha"
}
```
3)defer(�Ƴ�)�ؼ���
```

func main()  {
	Greting("uu","77")
}
func Greting(x ...string) string{   // ����ɱ����
	print(x[0],x[1],len(x))    // len ���س���
	defer lai()     // defer ����return֮����ִ��һЩ�������java�е�finally,����д��retrun֮ǰ
	return "haha"

}
func lai()  {
	print("��������....")
}
```
## ����ʱ��
```
func main()  {
	start := time.Now()
	time.Sleep(3 * time.Second)   // sleep  3��
	end:= time.Now()
	time_all:=end.Sub(start)
	Greting("uu","77")
	print("�ܺ�ʱ",time_all.Seconds())
}
```
