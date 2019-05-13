##### 执行结果
```
E:\go_study>go run 4.go
1
123

E:\go_study>go run 4.go
1
123

E:\go_study>go run 4.go
1
123

E:\go_study>go run 4.go
1
123

E:\go_study>go run 4.go
panic: hello

goroutine 1 [running]:
main.main()
        E:/go_study/4.go:18 +0x2f4
exit status 2
```

结论：
```
如果两个case都满足条件，是伪随机选择一个执行的，而不是之前想着的从上到下依次判断哪个case能执行。

还有一点，当某个case得到执行后，就会退出select，因为打印出了 123 。

最后一点，如果没有case可以执行，则立即执行default，然后退出select
```
