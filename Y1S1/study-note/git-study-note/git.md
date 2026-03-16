
# Git 常用命令速查笔记

> 这是一个面向日常使用的 Git 命令速查笔记，覆盖基础配置、工作流、分支管理、历史查看、远程协作、恢复与高级用法。每条命令附简短说明与常见示例。

---

## 目录
- 基础配置
- 仓库初始化与克隆
- 工作区 / 暂存区 / 提交
- 分支与合并
- 查看历史与差异
- 变基 / 重写历史
- 暂存与储藏（stash）
- 远程操作
- 标签（tags）
- 恢复 / 回滚 / 参考日志
- 子模块 / 工作树
- 高级与诊断命令
- 常用别名 & 小技巧

---

## 基础配置
配置 Git 行为与用户信息。
```bash
git config --global user.name "Your Name"
git config --global user.email "you@example.com"
git config --global core.editor "code --wait"      # 设置默认编辑器
git config --list                                   # 列出所有配置
```

---

## 仓库初始化与克隆
```bash
git init                  # 在当前目录初始化新仓库
git init --bare repo.git   # 初始化裸仓库（服务器端）
git clone <url>            # 克隆远程仓库
git clone -b <branch> <url>  # 克隆并切换到指定分支
```

---

## 工作区 / 暂存区 / 提交
基本工作流：编辑 → 暂存 → 提交
```bash
git status                # 查看工作区/暂存区状态
git add <file>            # 将文件加入暂存区
git add .                 # 暂存所有改动（注意：也会暂存删除）
git restore <file>        # 恢复工作区文件（丢弃未暂存更改）
git restore --staged <file> # 从暂存区移除（保留工作区文件）
git commit -m "msg"       # 提交暂存区到本地仓库
git commit --amend        # 修改最新提交（改消息或合并更改）
git commit --allow-empty -m "empty commit" # 允许空提交
```

常见示例：
```bash
git add -p                # 交互式分块暂存（精细控制）
git diff                  # 查看未暂存改动
git diff --staged         # 查看已暂存但未提交的改动
```

---

## 分支与合并
```bash
git branch                # 列出本地分支
git branch -a             # 列出所有（本地+远程）分支
git branch <name>         # 新建分支（不切换）
git switch <name>         # 切换分支（推荐新命令）
git switch -c <name>      # 新建并切换分支
git checkout <branch>     # 老命令：切换分支
git merge <branch>        # 将 <branch> 合并到当前分支
git merge --no-ff <branch># 强制创建合并提交
git merge --abort         # 中止正在进行的合并
git branch -d <branch>    # 删除分支（安全删除，若未合并会报错）
git branch -D <branch>    # 强制删除分支
```

常见协作策略：
- feature 分支+merge 或 rebase 到主分支
- 使用 pull request / merge request 做代码审查

---

## 查看历史与差异
```bash
git log                               # 查看提交历史
git log --oneline --graph --decorate  # 精简图形化日志
git show <commit>                     # 查看某个提交的详情
git show <commit>:<path>              # 查看某个提交中某文件内容
git blame <file>                      # 查看每行最后一次修改信息
git diff <a> <b>                      # 比较两个提交/分支/文件
git shortlog -s -n                    # 按提交数统计贡献者
```

过滤与格式化：
```bash
git log --author="Name" --since="2 weeks ago"
git log -p -S"函数名"                 # 搜索改动包含指定文本的提交并显示差异
```

---

## 变基 / 重写历史
重写提交历史以保持线性历史（谨慎：不要变基已推到共享远程的提交）。
```bash
git rebase <branch>            # 将当前分支变基到 <branch>
git rebase -i <base>           # 交互式变基（squash、reword、reorder、drop）
git rebase --abort             # 取消变基并回到变基前状态
git rebase --continue          # 解决冲突后继续变基
```

常见用法：
- git rebase -i origin/main  用于整理本地提交

---

## 暂存（stash）
临时保存未提交的改动以便切换分支
```bash
git stash                      # 暂存修改并恢复干净工作区
git stash push -m "wip"        # 带说明的 stash
git stash list                 # 列出 stash
git stash apply <stash@{N}>    # 应用 stash（不删除）
git stash pop                  # 应用并删除最新 stash
git stash drop <stash@{N}>     # 删除某个 stash
git stash clear                # 清空所有 stash
```

---

## 远程操作
```bash
git remote -v                  # 列出远程仓库
git remote add origin <url>    # 添加远程
git fetch                      # 获取远程更新（不合并）
git fetch --all
git pull                       # fetch + merge（默认）
git pull --rebase              # fetch + rebase（偏好线性历史）
git push                       # 推送当前分支到远程
git push --set-upstream origin <branch>  # 推送并设置上游分支
git push origin --delete <branch>       # 删除远程分支
git push --force-with-lease    # 强制推送（相对安全）
git push --tags                # 推送标签
```

---

## 标签（Tags）
```bash
git tag                        # 列出标签
git tag v1.0.0                 # 轻量标签
git tag -a v1.0.0 -m "release" # 注释标签（推荐）
git show v1.0.0
git push origin v1.0.0
git push origin --tags
```

---

## 恢复 / 回滚 / 参考日志（reflog）
```bash
git reset --soft <commit>      # 回退到某提交，保留暂存区和工作区
git reset --mixed <commit>     # 回退到某提交，保留工作区（默认）
git reset --hard <commit>      # 回退到某提交，丢弃暂存/工作区改动（危险）
git revert <commit>            # 产生新提交，撤销指定提交（安全用于公共历史）
git reflog                     # 查看 HEAD 的移动历史（找回丢失提交）
git checkout <commit> -- <file># 恢复某个文件到指定提交的状态（工作区）
```

提示：使用 reflog 可以找回因 reset 或误删除而“丢失”的提交。

---

## 子模块 & 工作树（submodules / worktree）
```bash
git submodule add <url> <path>
git submodule update --init --recursive
git submodule foreach git pull origin main

git worktree add ../other-branch <branch>  # 在不同目录检出另一个分支
```

---

## 高级与诊断命令
```bash
git fsck                      # 检查仓库完整性
git gc                        # 垃圾回收，优化仓库
git prune                     # 删除未引用的对象
git bisect start               # 二分查找引入 bug 的提交
git bisect good <rev>
git bisect bad <rev>
git archive -o release.zip HEAD
git bundle create repo.bundle --all
git rev-parse --abbrev-ref HEAD   # 获取当前分支名
git show-ref                      # 列出引用
git cat-file -p <hash>            # 显示对象内容
```

---

## 常用别名（建议加入 ~/.gitconfig）
```ini
[alias]
  st = status
  co = checkout
  br = branch
  ci = commit
  lg = log --oneline --graph --decorate --all
  last = log -1 HEAD
```

---

## 常见小技巧
- 若需临时实验，使用 git worktree 创建并行工作目录。
- 提交前使用 git add -p 精确选择变更。
- 协作时 prefer pull --rebase 可以减少无意义合并提交（与团队约定）。
- 永远在推送前确认 git status、git log --oneline -n 5。
- 对已公开的提交不要随意 git rebase 或 git reset --hard；若必须，使用强制推送并与团队沟通。

---

如果你需要，我可以：
- 把这份笔记按主题生成 PDF 或更短的一页速查卡；
- 根据你的工作流（例如 GitFlow、Github Flow）生成对应的命令模板；
- 或者把常用命令做成 Git aliases 配置文件，方便直接拷贝使用。
