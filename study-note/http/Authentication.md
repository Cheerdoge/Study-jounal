# Cookie
1. 键值对，要求必须是字符串
2. domain指的是所属域名
3. path表示在哪个路径生效
4. maxAge表示最大失效时间
	1. 单位秒
	2. 负数表示为临时cookie，关闭浏览器失效
	3. 为0表示1删除该cookie
5. expires表示过期时间
6. secure默认为false，为true在http无效，https有效
7. httpOnly设置一定程度防止XSS攻击
	XSS攻击指的是将恶意脚本代码注入网页中，其他用户浏览时被执行
# Session
1. 基于cookie实现，session在服务器，sessionid在客户端
2. 比cookie安全，键值对可以时任意数据类型
3. 持续时间短
4. 占用空间较大
# Token（令牌）
## Acess Token
1. 访问api（资源接口）时需要