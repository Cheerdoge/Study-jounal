用ai检查出错误后尝试修改后的版本，主要错误如下
1. 变量初始化问题：`UserInfomation` 和 `Sessions` 是map，但没有初始化，会导致nil panic。
	
2. 并发锁使用问题：在`CreateNewUser`中，我们使用了读锁，但实际上我们需要写锁，因为我们要修改map。
	
3. 在`GetSession`中，我们检查session过期后调用了`DeleteSession`，但`DeleteSession`没有加锁，而且我们在读锁内执行删除操作，这会导致并发问题。
	
4. 在`DeleteSession`中，我们直接删除，没有加锁，需要加写锁。
	
5. 中间件`MiddleAthusation`中，我们使用了`Sessions` map来获取session，但这里没有使用`GetSession`函数，导致过期检查逻辑没有统一使用。
	
6. 在`Logout`函数中，我们解析了请求体，但实际上登出不需要请求体，而且我们也没有使用解析出来的数据。登出应该是根据cookie来删除session的。
	
7. 在`ChangeProfile`函数中，我们修改了用户信息，但没有加锁，因为`GetUser`返回的是指针，我们直接修改了指针指向的内容，但修改操作不是原子的，而且如果多个goroutine同时修改同一个用户，会有数据竞争。
	
8. 在`ChangePassword`函数中，同样存在并发修改用户密码的问题，没有加锁。
	
9. 在`ChangeInformation`函数中，我们使用了一个map来传递修改，但是map的key是"Age"，而函数内使用的是"age"（小写），这会导致无法修改年龄。
	
10. 在`ChangeInformation`函数中，我们调用`ChangeProfile`，但`ChangeProfile`返回的bool值我们只判断了false，没有处理true，实际上我们总是返回修改成功，即使修改失败也会返回成功（因为false时我们写了错误，但true时没有返回，会继续执行后面的WriteSuccess）。
	
11. 在`ChangeProfile`函数中，我们通过`GetUser`获取用户，然后直接修改，但`GetUser`返回的是指针，我们修改的是原始数据，所以实际上不需要再通过map更新，但需要注意的是，我们修改的时候需要加锁。
	
12. 在`main`函数中，我们先调用了`http.ListenAndServe`，然后才注册路由，这会导致监听时路由还没有注册。