# 一级标题

这是一级标题下的内容

## 二级标题

这是二级标题1下的内容

### 哇哇

这是二级标题1下的三级标题的内容

## 二级标题2
111111111111111

### 三级标题2

2222222222222

这是一个**粗体文字**

# 正式练习

现在开始依次练习代码块、数字/圆点列表、加粗、删除线

## 代码

目前学业上要求的是c语言，它的一个程序如下

```c
#include<stdio.h>

int main(){
    int y, m;
    
    scanf("%d%d",&y, &m);
   
    switch (m)
    {
    case 1:
    case 3:
    case 5:
    case 7:
    case 8:
    case 10:
    case 12:
             printf("31");
             break;
    case 4:
    case 6:
    case 9:
    case 11:
            printf("30");
            break;
    case 2:
           if((y % 4 == 0 && y % 100 != 0)||(y % 400 == 0))
               printf("29");
            else 
            printf("28");
            break;
    }

    return 0;
}

```

社团要求的是go，它的一个程序如下

```go
package main

import (
    "fmt"
)

func main() {
    var n, E, r int
    var input string
    var ld int
    fmt.Scanf("%d%d%d\n", &n, &E, &r)
    fmt.Scanf("%s", &input)
    a := []byte(input)

    switch a[0] {
    case 45:
        fmt.Println("-1")
        return
    case 43:
        E = E + r
    }

    var list []int
    for num := 1; num <= n; num++ {
        for k := 1; k <= E; k++ {

            if ld == n {
                list = append(list, num)
                break
            }
            switch a[ld-1+k] {
            case 45:
                continue
            case 43:
                E = E + r
                ld = ld + k
            case 48:
                ld = ld + k
            }

        }
    }

    if list == nil {
        fmt.Println("-1")
    } else {
        m := len(list)
        for i := 0; i < m-1; i++ {
            for j := 0; j < m-1-i; j++ {
                if list[j] > list[j+1] {
                    list[j], list[j+1] = list[j+1], list[j]
                }
            }

        }
        fmt.Println(list[0])
    }
}

```

## 数字/圆点列表

1. 按下数字加小数点加空格就可以快速调用

2. 先“/”，再找有序列表

3. 后续会自动添加

   1. 先回车再按下tab键可以嵌套一个有序列表

   2. 如果想去除，shift+tab

   3. 再输入新的上一级有序列表，先输入内容再取消缩进

4. 如果想加入无序列表

   1. 先创建有序的

   2. 再按无序的快捷键就可以了

      * 只会影响从这一行开始之后的

      * 测试

## 加粗

加粗很简单，**旺柴** 用两个\*框起来再空格

## 删除线

我是~~大帅哥~~
1. 先用~~把要划的文字框起来
2. ~~像这样~~

## 引用链接

不知道快捷键，斜杠然后找吧

* 如果在空格中，要先空一格再斜杠才能打开搜索
* [百度](www.baidu.com)

## 斜体
用单个星号括起来的就是斜体，如：*斜体*
