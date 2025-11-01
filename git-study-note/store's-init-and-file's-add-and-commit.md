# git中的仓库初始化和文件提交
## 仓库初始化
1. `git init`可以初始化一个仓库，在其后加名称可以初始化该仓库的名称。且加上名称就会在这个名称目录之下生成.git目录
2. `git clone`可以从github中复制一个仓库，格式`git clone url`其中URL是对应仓库的URL
## 将文件添加仓库中
* `git status`查看仓库状态
* `git add`可以将文件放入暂存区（小货车）
    1. 注意文件名中**不要有空格**，会导致添加失败，可以用**-**来替代
    2. `git add *.txt`可以将txt格式文件批量提交
    3. `git rm --cached <file>`可用于取消暂存文件
    4. `git add .`可将当前文件夹下的所有文件提交暂存区中
* `git commit`可用于提交暂存区的文件，且仅会提交暂存区文件
   * 在命令后加-m 可以添加提交的信息，如fix等，详见standardized-commits-in-git
* `git log`可查看提交记录，其后加上`--oneline`可查看简洁的提交记录