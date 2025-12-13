## 一、核心概念（必须理解）
- 模式（Mode）
  - 普通模式（Normal）——默认模式，用于移动、删除、复制、粘贴、命令等。
  - 插入模式（Insert）——输入文本（按 `i`、`a`、`o` 进入）。
  - 可视模式（Visual）——选择文本（按 `v`、`V`、`Ctrl-v`）。
  - 命令行模式（Command）——执行 `:` 开头的命令（保存、替换、退出等）。

---

## 二、启动与退出
- 打开文件：`vim filename`
- 保存：`:w`
- 退出：`:q`
- 保存并退出：`:wq` 或 `:x`
- 强制退出（放弃修改）：`:q!`

示例：
```bash
vim notes.txt
:w
:wq
```

---

## 三、从普通模式进入插入模式（常用）
- `i`：光标前插入
- `a`：光标后插入
- `o`：在下一行新开一行并进入插入
- `I`、`A`：行首/行尾插入

返回普通模式：`Esc` 或 `Ctrl-[`

---

## 四、基本移动（普通模式）
- h j k l：左 下 上 右（练习替代箭头）
- w / e / b：按词移动（到词首/词尾/词首向后）
- 0：行首，^：第一个非空字符，$：行尾
- G：跳到文件末尾，gg：跳到文件开始
- nG：跳到第 n 行（例如 `5G` 跳到第 5 行）

---

## 五、删除、复制、粘贴
- 删除（普通模式）
  - `x`：删除字符
  - `dd`：删除整行
  - `d0`：删除到行首，`d$`：删除到行尾
  - `dw`：删除到下一个词首
- 复制（Yank）
  - `yy`：复制整行
  - `yw`：复制一个词
  - `y$`：复制到行尾
- 粘贴
  - `p`：在光标后粘贴
  - `P`：在光标前粘贴

说明：可以与数字结合，如 `3dd` 删除三行，`2yw` 复制两个词。

---

## 六、撤销与重做
- 撤销：`u`
- 重做：`Ctrl-r`

---

## 七、可视选择（Visual）
- `v`：字符可视，`V`：整行可视，`Ctrl-v`：块选择
- 选择后可按 `d` 删除、`y` 复制、`>` `<` 缩进等

---

## 八、搜索与替换
- 搜索（普通模式）
  - `/pattern` 向下查找，按 `n` 下一个，`N` 上一个
  - `?pattern` 向上查找
- 替换（命令行）
  - 替换当前行：`:s/old/new/g`
  - 替换整个文件：`:%s/old/new/g`
  - 带确认：`:%s/old/new/gc`

---

## 九、缓冲区、窗口与标签
- 打开新文件：`:e otherfile`
- 缓冲区切换：`:bnext` / `:bn`（下一个缓冲区），`:bprev` / `:bp`
- 水平分屏：`:split filename` 或 `:sp filename`
- 垂直分屏：`:vsplit filename` 或 `:vs filename`
- 在分屏间移动：`Ctrl-w` 然后方向键（h/j/k/l）
- 新标签页：`:tabnew filename`，切换 `:tabnext` / `:tabn`，关闭 `:tabclose`

---

## 十、常用配置建议（放到 ~/.vimrc）
最小 .vimrc：
```vim
set number            " 显示行号
set relativenumber    " 相对行号（可选）
syntax on             " 语法高亮
set tabstop=4
set shiftwidth=4
set expandtab         " 用空格代替 tab（可选）
set clipboard=unnamedplus " 使用系统剪贴板（Linux/macOS）
```

---

## 十一、速查表（常用）
- 插入模式：`i a o I A`
- 保存/退出：`:w`, `:q`, `:wq`, `:q!`
- 删除/复制/粘贴：`x`, `dd`, `yy`, `p`, `P`
- 撤销/重做：`u`, `Ctrl-r`
- 搜索：`/pattern` → `n`/`N`
- 替换：`:%s/old/new/g`
- 分屏：`:sp` / `:vs`，切换 `Ctrl-w` + `h/j/k/l`