# 错误处理
go中已经预先定义了一个名为error的接口类型
## 定义错误
使用errors包中的errors.New("错误信息")来创建一个error接口类型变量
`err := errors.New(“math - square root of negative number”)` 
用来直接输出一段错误信息
## 判定错误类型
对错误使用类型判断可以拿来补充一些补救代码
```go
//  err != nil
if e, ok := err.(*os.PathError); ok {
    // remedy situation
}
```
或者用switch
```go
switch err := err.(type) {
    case ParseError:
        PrintParseError(err)
    case PathError:
        PrintPathError(err)
    ... 
    default:
        fmt.Printf(“Not a special error, just %s\n”, err)
}
```
## 运行时异常或panic
### panic函数
1. panic通常接受一个字符串，用于中止程序同时打印该语句并给出调试信息，当然也可以通过字符串+error变量打印出error记录的错误信息
2. 有些函数包含Must前缀，这些函数出错时会panic
## Go panicking
在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者。这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情况：这个终止过程就是 _panicking_。
意思就是，多层嵌套panic，会由内向外依次执行defer，直到
## recover函数
用于从panic中恢复
1. 只能在defer修饰的函数中使用
2. 使用了recover但是崩溃点后的代码不会执行
3. 会捕获panic的信息，返回的值类型与panic传入参数一致
4. 举个使用例子
   ```go
   defer func(){
   if r := recover(); r != nill{
	   fmt.Printf("....%v", r)
		//var ok bool
		//err, ok = r.(error)
		//if !ok {
		//err = fmt.Errorf("pkg: %v", r)
		}
   }()
   }
   ```
当需要把错误返回给上层函数时，使用注释中的代码，记得在该层函数返回值添加error类型的err返回值
## 利用闭包处理错误
可以减少重复的错误检查代码

1. **为什么不直接定义一个函数？**
   * 定义函数处理错误需要每一个函数中手动调用这个处理函数
   * 这个处理函数必须针对每一个函数写一次，而且可能因为外层函数参数的问题导致参数名不同，需要单独写一次，不能通用（比如在函数1中第一个变量是a，函数2中第一个变量是b，这样就必须在处理函数中的参数中手动修改） 
   * 可以简化传入参数，比如一个bool值，为true怎么处理，为false怎么处理，用闭包可以在创建这个函数变量时提前传入，后续只用传入其他参数
 2. 要求必须说同一签名的函数（指参数、返回值的个数和类型一一对应）
 3.  格式如下
```go
    func createSafeFileReader() func(string) {
    return func(filename string) {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("文件操作失败: %v", r)
            }
        }()
        
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            panic(err) // 抛出错误，由闭包统一处理
        }
        
        fmt.Printf("文件内容: %s\n", string(data))
    }
}
```
