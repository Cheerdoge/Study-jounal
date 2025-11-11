结构体学过，方法光看名居然没啥印象（）
# 结构体
1. new创建结构体变量返回的是指针
2. 链表
	1. 通过在结构体中添加一个指向下一个结构体的指针成员实现指向临近节点，如
		```go
		type person struct{
		pr *struct1
		id int
		su *struct2
		}
		```
	2. 
		其中指针指向的结构体也可是和该结构体一个类型的![[Linked-LIst.png]]
3. **结构体工厂**，就是定义一个函数，实现快速产生一个结构体变量，注意返回值最好是指针类型，通常将函数名命名为new+结构体名![[Struct-Factary.png]]
4. 查看结构体T的实例所占内存`size := unsafe.Sizeof(T{})` 
5. 调用其他包里的结构体就可以使用工厂方法，也可以通过`new(packagename.typename)` 
6. 结构体中的成员可以不给名字，就单给个类型，此时类型名就是这个成员的名字（每个类型只能有一个）
	1. 如果内嵌了一个匿名的结构体，那么这个结构体中的成员的名字就是相应成员的名字
	2. 如果内嵌结构体指针，又想一次打印所有内容，要么手动格式化输出，要么String方法
	3. 多说无益，来段代码
		```go
		type s1 struct{
		a int
		b int
		}
		type s2 struct{
		c int
		d int
		int
		s1
		}
		func main(){
			text := new(s2)
			s2.c = xxx
			s2.int = xxx//这个就是s2中第三个成员
			s2.a = xxx//这个就是内嵌匿名结构体中的成员
			}
		```
7. 当内嵌结构体和外层结构体使用了相同的成员名字
	* 当二者类型相同->手动修改
	* 二者类型不同->外层结构体.内层结构体.名字和外层结构体.名字
# 方法
1. 定义：函数定义的基础上，在函数名（func之后）添加（接收者名和接收者类型），如果方法中**未使用接收者，接收者可以不命名** 
2. 区别：方法只能为接收者类型的变量使用
3. 上例子
	* 函数
		```go
		package main
		
		import "fmt"
		
		// 定义一个“人”的结构体类型
		type Person struct {
		    Name string
		    Age  int
		}
		
		// 这是一个【函数】。它需要一个Person类型的参数。
		func SayHelloFunction(p Person) {
		    fmt.Printf("函数：你好，我是 %s，我 %d 岁了。\n", p.Name, p.A.Age)
		}
		
		func main() {
		    // 创建一个Person实例
		    bob := Person{Name: "Bob", Age: 25}
		
		    // 调用函数，需要把bob作为参数传进去
		    SayHelloFunction(bob)
		}```
	* 方法
	  ```go
		package main
	
		import "fmt"
		
		type Person struct {
		    Name string
		    Age  int
		}
		
		// 这是一个【方法】。
		// 注意看 (p Person) 这部分，它叫“接收者”（Receiver）。
		// 它意味着这个函数是“属于”Person这个类型的。
		// 现在，SayHelloMethod 就是 Person 的一个专属行为。
		func (p Person) SayHelloMethod() {
		    fmt.Printf("方法：你好，我是 %s，我 %d 岁了。\n", p.Name, p.Age)
		}
		
		func main() {
		    bob := Person{Name: "Bob", Age: 25}
		
		    // 调用方法！语法是：变量.方法名()
		    // 看，这里不需要再把bob作为参数传入了。
		    // bob 就是这个方法的主人，方法内部天然就能访问bob的数据。
		    bob.SayHelloMethod()
		}
			```
	4. 同函数传入参数一样，传入值参数不影响原本的值，传入指针参数影响原本的值
	5. 类型和其方法必须在一个包内，所以不支持在常规类型上定义它的结构体，但是非要定义，可以给常规类型起个别名
	6. 结构体中内嵌一个匿名结构体，那么这个结构体的实例可以直接调用内嵌结构体的方法
	7. 如果内嵌结构体不是匿名的，就需要先访问内嵌的结构体的名称，再使用方法
	8. 可以通过方法来访问内嵌结构体的方法，可以做到
		1. 不同包的方法的访问
		2. 添加额外的逻辑
	9. 外层结构体的同名方法可以通过在方法中使用内层结构体方法达到重写的效果，就是**运行外层结构体的方法**
	10. 当内嵌结构体的类型是指向这个类型的指针（用的时候确定它不是nill）时，需要先给这个**指针初始化内存，再进一步处理** 
	11. 定义的String（）方法会在Printf（“%v”）和Println（）时被采用，相当于定制怎么输出。注意**不要在定义String方法中使用这些** ，会导致无限递归
# 垃圾回收
1. GC会搜索不再使用的变量并释放内存，runtime包可以访问其进程
2. `runtime.GC()` 可以主动触发内存释放