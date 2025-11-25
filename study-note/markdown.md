# Markdown 语法速查表（可打印）

---

## 目录
- 基础语法
- 标题（Headings）
- 强调（Emphasis）
- 列表（Lists）
- 链接与图片
- 代码（Inline & Block）
- 表格（GFM）
- 引用（Blockquote）
- 水平线（Horizontal rule）
- 换行 / 段落
- 原始 HTML 与转义
- 任务列表（Task list）
- 行内表情与自动链接（GFM）
- 参考式链接与重用
- 扩展与常用元信息（YAML front matter）
- 常见注意事项与最佳实践

---

## 基础示例
- 文本即 Markdown。
- 段落通过空行分隔。
- 行尾末尾加两个空格表示强制换行（soft break 之外的换行）。

---

## 标题（Headings）
使用 # 表示，1~6 级：
```md
# H1
## H2
### H3
#### H4
##### H5
###### H6
```

---

## 强调（Emphasis）
```md
*斜体* 或 _斜体_
**加粗** 或 __加粗__
***加粗斜体*** 或 ___加粗斜体___
~~删除线~~        // GFM 支持
```

---

## 列表（Lists）
无序列表使用 `-`、`*` 或 `+`：
```md
- 项目 1
- 项目 2
  - 嵌套项
```
有序列表使用数字加点：
```md
1. 第一项
2. 第二项
   3. 嵌套 1
```

---

## 链接与图片
行内链接：
```md
这是 [Google](https://www.google.com) 的链接。
```
参考式链接（便于重用）：
```md
这是 [示例][1]。

[1]: https://example.com "可选标题"
```
图片（与链接类似，前面加 `!`）：
```md
![替代文字](https://example.com/image.png "可选标题")
```

---

## 代码（Inline & Block）
行内代码：
```md
使用 `fmt.Println()` 打印。
```
代码块（缩进或 fenced，推荐 fenced）：
<pre>
```go
package main
import "fmt"
func main() {
  fmt.Println("hello")
}
```
</pre>
在 fenced code block 指定语言可启用语法高亮（如 ` ```go `）。

---

## 表格（GFM 支持）
```md
| 列名1 | 列名2 | 列名3 |
|---|:---:|---:|   // 对齐: 左（默认）/ 中 / 右
| 左对齐 | 居中 | 右对齐 |
```

---

## 引用（Blockquote）
使用 `>`：
```md
> 这是引用
> 多行引用
```

---

## 水平线（Horizontal rule）
三个或更多的 `-`、`*` 或 `_` 单独成行：
```md
---
***
___
```

---

## 换行 / 段落
- 空行分隔段落。
- 强制换行：行尾加两个或更多空格然后回车。
- 不同渲染器对单行换行可能不同（GFM 将单个换行视为 <br>）。

---

## 原始 HTML 与转义
- 可以嵌入 HTML（如表格、细粒度样式），但某些平台会过滤危险标签（script）。
- 转义字符：`\* \_ \# \[ \] \( \) \` \>` 等，例如 `\*` 显示星号。

---

## 任务列表（GFM）
```md
- [x] 已完成项
- [ ] 未完成项
```

---

## 行内表情与自动链接（GFM）
- Emoji：`:smile:`（取决于渲染器是否支持）。
- 自动链接：`<https://example.com>` 会被渲染为可点击链接。
- 提及（@username）与 issue/PR 引用（#123）在 GitHub 上生效。

---

## 参考式链接（重用链接）
```md
请访问 [官网][site]。

[site]: https://example.com "站点标题"
```

---

## 扩展与元信息（YAML front matter）
很多静态站点生成器使用 YAML 前置块：
```md
---
title: "页面标题"
date: 2025-11-24
tags: [go, markdown]
---
```

---

## 常见技巧与注意事项
- 在代码块里使用三个反引号（```），若代码中含有三个反引号，可用更多反引号包裹外层（```` ``` ````）。
- 表格列数要一致，分隔行必须存在。
- 在表格或列表前后添加空行以保证兼容不同渲染器。
- 使用参考式链接能避免重复 URL 并简化长文档管理。
- 多使用 fenced code block + 指定语言，便于渲染器高亮。
- 若依赖平台特性（GFM 的 task list、表格、删除线），在文档顶部注明目标渲染器。
- 注意隐私：不要在公开 Markdown 中写明敏感 token 或密码。

---

## 常用快捷键（编辑器相关）
- VS Code：Markdown 预览（Ctrl+Shift+V / Cmd+Shift+V），预览侧边（Ctrl+K V）。
- Typora / Markdown 编辑器通常支持即时所见即所得。