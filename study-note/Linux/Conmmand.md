- **绝对路径** 从根目录开始写，例如 **/home/root/workspace/sample。**
- **相对路径** 以当前目录为基础， **./** 表示当前目录（./ 等价于不写）， **../** 表示父级目录（当前路径所在的上一级路径）， **/** 表示当前目录的根目录。

### 1. ls：**列出该目录下的文件（list）**

**ls 常用参数：**

- **-l：** 列出文件的详细信息
- **-a：** all，列出所有文件，包括隐藏文件
#### 颜色区分文件
1. 白色：普通文件，如txt等
2. 青色：跳转链接，类似快捷方式
3. 绿色：可执行文件
4. 红色：压缩文件，如.gz
5. 黄色：设备文件

#### -l显示的参数意义
格式示例：`-rwxr-xr-x`
* 首字符：文件类型
	-：普通文件  
	d：目录  
	l：符号链接  
	b：块设备文件  
	c：字符设备文件  
	p：管道文件  
	s：套接字文件
* 剩余每三个一组，依次表示所有者user，所属组group，其他用户other的权限
	`rwx`依次表示，读、写、执行
### 2. pwd：显示当前目录的绝对路径（Print Working Directory）

### 3. cd：切换目录（Change Directory）

```python
## cd用法
cd /home    # 切换/进入home目录
cd ..       # 到上一目录（父目录）
cd ../..	# 到父目录的父目录
cd .		# 进入当前目录
```

### 4. cp：复制（Copy）

**cp 常用参数：**

- **-i** interactive mode，若有同名文件，会询问是否覆盖（如果没这个参数，会不提示，直接覆盖）
- **-r** recursive copy，复制文件夹时连同子文件(夹)一起复制，如果是对文件夹进行操作，一定要带这个参数

```python
## cp用法
cp -ir sourceDir/ home/targetDir/	
# 把当前路径下的sourceDir文件夹复制到home目录下，取名为targetDir，且带参数-i和-r
```

### 5. mv：移动（Move）

**mv 参数：**

- **-i** interactive mode ，同 cp 的 -i 参数，若覆盖会询问

```python
## mv用法
mv -i sourceFile /home/targetFile	# 把当前目录下的sourceFile剪切到/home目录下并命名为targetFile
```
 
### 6. rm：删除给定的文件（Remove）

**rm 参数：**

- **-i** interactive，同上，若覆盖，先询问
- **-r** recursive mode，删除所有子文件(夹)

```python
## rm用法
rm Dir/	# 删除Dir文件夹（错误示例，会报错）
rm -r Dir/	# 删除Dir文件夹（正确，对文件夹操作一定要带-r）
```

### 7. mkdir：创建一个新目录（Make Directory）

```python
## mkdir用法
mkdir newDir/	# 在当前路径创建一个空文件夹newDir/
```

### 8. rmdir：删除文件夹（Remove Directory）

```python
## rmdir用法
rmdir oldDir/	# 在当前路径删除oldDir文件夹及其子文件（夹）
```

### 9. cat：查看文件内容（concatenate and print files）

```python
## cat用法
cat myFile	# 显示myFile
```

### 10. tar：打包压缩、解压

**tar 常用参数：**

- **- jcv** 压缩
- **- jxv** 解压

```python
## tar用法
tar -jcv myDir/		# 压缩myDir文件夹
tar -jxv DownloadDir.tar.gz myDir/	# 解压DownloadDir.tar.gz到当前文件夹下，并命令为myDir
```

### 11. zip、unzip：打包压缩、解压

- **-r** 递归处理，将指定目录下的所有文件和子目录一并处理
- **-d** 解压用，用来指定解压目录

### 12. ps：查看进程（Process Select）

**ps 常用参数：**

- **-A** 显示所有进程
- **-a** 不与 terminal 有关的所有进程
- **-u** 有效用户的相关进程
- **-x** 一般与 -a 一起用，列出完整的进程信息
- **-l** long，详细列出 PID 的信息

```python
## ps用法
ps Aux 	# 查看系统所有的进程数据
ps ax	
```

### 13. kill：杀死进程

**kill 常用参数 :**

- **SIGHUP** 启动被终止的进程
- **SIGINT** 相当于 **Ctrl + C**，中断进程
- **SIGKILL** 强制中断进程
- **SIGTERM** 以正常的结束进程方式来终止进程
- **SIGSTOP** 相当于 **Ctrl + Z**，暂停进程

```python
## kill用法
kill -SIGKILL 10876	# 强制中断PID=10876的进程（PID可以通过ps查到，有时可以加上| grep进行筛选）
```

### 14. passwd：修改密码（Password）

```python
## passwd用法
passwd	# 修改当前用户的密码
```

### 15. tee：显示并保存

显示内容并将内容保存在文件中：

```python
python3.6 test.py | tee result.log	# 运行test.py文件，显示编译与运行结果并保存成result.log文件
```

### 16. reboot：重启

```python
## reboot用法reboot	# 输完立马重启（记得保存文件）
```

### 17. date：时间相关指令

- 用来显示当前时间
- 手动指定显示时间的格式

date 指定格式显示时间：**date +%Y:%m:%d**  
date 用法：**date [OPTION]... [+FORMAT]**

### 18. find：查找

语法： **find pathname -options**

功能： 用于在文件树种查找文件，并作出相应的处理（可能访问磁盘）

常用选项：

- **find 单独使用**时，必须指定目录查找或查找当前目录的文件
- **find -name 文件名：**遍历指定位置查找（范围较大时，较费时间）
- “*.txt”表示后缀名为txt的文件
- / 表示整个系统

### 19. grep：按行查找并匹配

**grep 参数：**

- **-r** recursive，查找所有子文件(夹)
- **-n** number，显示行号
- **-w** word，完整匹配整个单词
- **-i** insensitive search，忽略大小写
- **-l** 显示文件名称，而非匹配到的行的内容
- **-v** 反向选择，显示出没匹配到的行的内容

语法： **grep [选项]** 搜寻字符串的文件

功能： 在文件中搜索字符串，将找到的行打印出来，默认区分大小写

|**选项**|**功能** |
| ------ | ------------------ |
| -i     | **取消区分大小写**        |
| -n     | **输出行号**           |
| -v     | **反向选择，选择不带关键字的行** |

# 20.lsof
用于列出正在被进程打开的文件

参数：
* -c 进程名：显示该进程打开的文件
* -p 进程id：~
* -u 用户名：特定用户
* 文件路径：显示该文件被哪些进程使用
* -i：所有网络连接
	* :xx：特定端口（1-1024表示这个范围）/连接（TCP/UDP）/主机
```
# 查看目录被哪些进程使用
lsof +D /var/log

# 查看特定文件系统
lsof /dev/sda1

# 查看已删除但仍被进程占用的文件
lsof | grep deleted# AND 条件（同时满足）
lsof -a -u apache -c apache2

# 排除特定用户
lsof -u ^root

# 查看进程打开的文件数量
lsof -p 1234 | wc -l

#输出传递给另一个程序
lsof -i :8080 | ==Qa4VXb
```

# 21.cat
查看文件
`cat [选项] [文件]`
- `-n`：显示行号，会在输出的每一行前加上行号。
- `-b`：显示行号，但只对非空行进行编号。
- `-s`：压缩连续的空行，只显示一个空行。
- `-E`：在每一行的末尾显示 `$` 符号。
- `-T`：将 Tab 字符显示为 `^I`。
- `-v`：显示一些非打印字符