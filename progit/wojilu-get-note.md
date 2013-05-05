http://www.wojilu.com/Forum1/Topic/2236


下面是我的部分笔记，仅供参考。

--------------------------------------------------------

常用命令：
 
The most commonly used git commands are:
 
   add       
Add file contents to the index
   bisect    
Find by binary search the change that introduced a bug
   branch    
List, create, or delete branches
   checkout  
Checkout a branch or paths to the working tree
   clone     
Clone a repository into a new directory
   commit    
Record changes to the repository
   diff      
Show changes between commits, commit and working tree, etc
   fetch     
Download objects and refs from another repository
   grep      
Print lines matching a pattern
   init      
Create an empty git repository or reinitialize an existing one
   log       
Show commit logs
   merge     
Join two or more development histories together
   mv        
Move or rename a file, a directory, or a symlink
   pull      
Fetch from and merge with another repository or a local branch
   push      
Update remote refs along with associated objects
   rebase    
Forward-port local commits to the updated upstream head
   reset     
Reset current HEAD to the specified state
   rm        
Remove files from the working tree and from the index
   show      
Show various types of objects
   status    
Show the working tree status
   tag       
Create, list, delete or verify a tag object signed with GPG
 
基本概念
 
1）工作区
2）暂存区(stage)
3）版本库
 
M的意义：
第二列的M表示“工作区”和“暂存区”有改动；
第一列的M表示“暂存区”和“版本库”有改动。
 
对象ID：是一个40位的SHA1哈希值
 
[问]为什么需要一个暂存区？
[答]多了一个缓存，避免了不必要的提交错误，更加方便commit反悔撤销（反悔可以不影响工作目录）。
（通过命令 git reset HEAD^ 反悔，或者使用 git reset --soft HEAD^ 反悔）
 
常用命令
 
git命令分成两类操作：
一是将数据提交到版本库的操作，这是主要操作；
二是将版本库的数据恢复到工作区，这是逆操作，会改变当前工作区目录下的文件和状态。这类操作主要是两个命令：
 
1）git reset
 
在书本p98
 
 
--soft
(a)将某个commit替换master/HEAD
默认
--mixed
(a)将某个commit替换master/HEAD   (b)同时替换暂存区
 
--hard
(a)将某个commit替换master/HEAD   (b)同时替换暂存区   (c)同时替换工作区
 
git reset
使用默认的--mixed，即用HEAD替换HEAD；同时使用HEAD替换暂存区。
git reset HEAD
同上，使用HEAD替换暂存区
git reset --filename
使用HEAD中的file替换暂存区的file，也就是撤销暂存区的某文件的最新改动，相当于对add file做反向操作
git reset HEAD filename
同上
git reset --soft HEAD^
工作区和暂存区不变，但版本库向前倒退一步。
git commit --amend (修补最新提交的说明)相当于——
git reset --soft HEAD^
git commit -e -F .git/COMMIT_EDITMSG
git reset HEAD^
除了工作区，暂存区和版本库都后退一步
git reset --hard HEAD^
工作区、暂存区、版本库都后退，所有修改丢失
 
2）git checkout
 
checkout 是 git 最常用的命令，也是最危险的命令之一，它会改变工作区。
 
简介
git checkout 4902dc^
checkout 的本质是修改 HEAD 的指向，虽然默认是指向master的，但可以改成指向其他 commit ，比如 git checkout 4902dc^  会导致提醒：You ar in 'detached HEAD' state。如果查看此时的 HEAD 指向：cat .git/HEAD ，结果就是一个具体的 commit 的ID（即一个 40位的ID号）
提交
commit
在 'detached HEAD' 状态下commit，会提醒 'commit in detached HEAD mode'，提交之后，HEAD 会指向新的commit的40位ID号。
检出
git checkout master
即将 HEAD 指向 master ，同时会导致工作区和暂存区也全部倒退到master的文件状态。
合并
git merge acc2f69
将 acc2f69 的提交合并到当前 master
 
从暂存区检出：上面都是从版本库检出(修改HEAD并改写工作区)，但如果不填写commit，则默认从暂存区检出，也就是用暂存区的内容改写工作区。它和 reset 不同，reset 有三种参数，默认的mixed是替换暂存区，除非用hard 否则不会改写工作区。而checkout默认是用暂存区替换工作区。如果不省略commit，则除了替换工作区，也替换暂存区。
 
git checkout branch
更新HEAD指向branch，并用branch指向的树更新暂存区和工作区
git checkout branch -- filename
HEAD指向不变，用branch分支的文件覆盖暂存区和工作区的文件
git checkout
汇总显示工作区、暂存区、HEAD的差异
git checkout -- filename
用暂存区中的文件覆盖工作区的文件，这个命令很危险，会覆盖本地的修改
--和filename之间有空格
git checkout .
非常危险，将暂存区所有文件替换工作区所有文件
 
【reset 和 checkout 区别】
 
checkout 和 reset 不同，reset 有三种参数，默认的mixed是替换暂存区，除非用hard 否则不会改写工作区。
而checkout默认是用暂存区替换工作区。同时它的本质是修改 HEAD 指向。
 
【设置命令别名】
 
$ git config --global alias.unstage 'reset HEAD --'
这样一来，下面的两条命令完全等同：
$ git unstage fileA
$ git reset HEAD fileA
 
 
一、工作区和暂存区的互动
 
1）未add：在工作区增加了文件之后，如果查看 git status -s，则会看到带问号的文件列表
??  filename
它表示这个文件没有add到暂存区。因为没有纳入版本管理(连暂存区都没有纳入)，所以你可以随意修改、删除，版本库不会有任何记录。只有add到暂存区后，文件的修改才会被监控。
 
2）add之后：如果add之后，会看到文件列表前面的问号编程了A
A  filename
此时文件已经被纳入了版本管理，当然仅仅是暂存区。
注意：只有commit到版本库，此处的A才会变成M
 
3）修改之后：如果修改了工作区的文件，则不会自动和暂存区同步。通过 git diff，会看到工作区的文件内容和暂存区的内容是不一样的。如果需要同步，请继续使用 git add 命令，将工作区变动同步到暂存区。
 
4）撤销add(撤销对暂存区的改动)：
git rm --cached filename
或者
git reset -- filename 或 git reset HEAD <file>
本质上是使用 HEAD 中的内容覆盖暂存区内容
 
5）用暂存区内容替换工作区内容(撤销工作区的改动)：
git checkout -- filename 或者 git checkout .
 
6）查看add日志：
查看暂存区目录树：git ls-files -s
 
7）删除之后：暂存区不会变化，如果要同步删除，必须 git add.
 
 
二、暂存区和版本库的互动
 
1）commit：
git commit -m "msg"
 
2）查看commit日志
git log --prettty=oneline
 
3）修改刚才的commit消息
git commit --amend -m "msg"
 
修改刚才的提交者：
git commit --amend --reset-author
 
4）撤销刚才的提交
git reset --soft HEAD^
 
5）撤销提交，并撤销暂存区修改
 
 
三、综合实例（p106）
 
如果文件状态如下 git status -s
A    a/b/c/hello.txt
  M  welcome.txt
 
那么——
git commit
将暂存区内容提交
git reset --soft HEAD^
返回了，取消刚刚的提交
git add welcome.txt
想将welcome.txt提交
git reset HEAD a/b/c
将hello.txt撤出暂存区
git reset
将剩下的文件(welcome.txt)从暂存区撤出
 
 
git add somefile
加入到缓存去(使用 git add .添加所有文件)
git commit -m "note"
提交，将暂存区中文件保存到版本库。
如果只填写 git commit 则会出现vim窗口，提示你添加提交信息。你必须会使用vim进行编辑输入，最后:wq保存信息退出。
git log
git log --prettty=oneline
查看日志历史
在命令行窗口中，如果内容过多，按j往下移动(vim的命令)；按q表示退出
 
git log --pretty=oneline --max-count=2
git log --pretty=oneline --since='5 minutes ago'
git log --pretty=oneline --until='5 minutes ago'
git log --pretty=oneline --author=<your name>
git log --pretty=oneline --all
git log --all --pretty=format:"%h %cd %s (%an)" --since='7 days ago'
推荐格式
git log --pretty=format:"%h %ad | %s%d [%an]" --graph --date=short
git status 
git status -s
最常用的命令。查看哪些文件修改了，哪些文件被删除了，哪些需要commit，当然，是在git add之后，也就是缓存之后，才能看到结果。
 
 
--pretty="..." 定义了输出格式
%h
代表提交的精简hash码
%d
这个提交的一些修饰信息（比如，分支顶或者tag）
%ad
作者提交的日期
%s
提交时的注释
%an
提交人的名字
--graph
让git用字符图的方式显示提交树
--date=short
显示简短的日期格式
 
 
配置命令——
 
.GITCONFIG
[alias]
  co = checkout
  ci = commit
  st = status
  br = branch
  hist = log --pretty=format:\"%h %ad | %s%d [%an]\" --graph --date=short
  type = cat-file -t
  dump = cat-file -p
 
 
【分支】
 
1）创建分支并切换到此分支
 
git checkout -b <branchName>
 
等效于
 
git branch <branchname>
git checkout <branchname>
 
2）切换到某个分支
 
git checkout <branchName>
 
 
3）分支合并
 
git merge master ——将master分支的内容和当前分支合并
 
合并之后可以通过 git hist --all 查看分支图
 
【使用rebase】
更改基线的结果和合并很相似。greet分支现在含有所有的修改，包括master分支的修改。 不过，提交树却完全不一样了。greet分支的提交树整个重写了，包含了master分支的提交历史 这让提交链表看上去更连续，可读性更好。
 
4）删除分支
 
只能删除已经被合并的分支
git branch -d branchName
如果要强行删除某分支
git branch -D branchName
 
4）撤销合并
 
git reset --hard HEAD
 
 
【tag】
 
一、tag操作
 
1）增加 git tag -m "description msg" tagName
2）删除 git tag -d tagName
 
二、tag的作用
 
1）在不同的tag之间切换，查看历史文件
 
命令：git checkout tag1
会切换到tag1的状态，但命令行会提醒你：
You are in 'detached HEAD' state. You can look around, make experimental changes
在这种状态下，你可以查看历史文件，做一些实验性的修改测试或提交，不会影响任何分支。
 
2）基于历史tag创建分支
如果要基于当前tag创建分支，可以在当前tag中使用命令 git checkout -b newBranchName 创建分支，创建之后即切换到此新分支，然后可以提交文件。在此新分支所作的任何更改，不会影响到相关的tag
 
 
【文件处理】
 
0）文件查找
 
git grep someText
查找版本库中某段文字
git grep -n someText
结果显示行号
git grep --name-only someText
只显示文件名
 
1）移动文件：git mv fileName dirName，它等效于
 
mv fileName dirName
 
git add dirName/fileName
git rm fileName
 
2）删除文件(此命令和 git add 正好相反)
 
git rm fileName
 
或者先在工作区手动删除，然后使用 git add -u 命令将本地文件的所有变更(包括修改和删除)同步到暂存区
 
删除目录和所有子目录
 
git rm dirName -r
 
3）恢复被删除的文件
虽然工作区没有文件了，但版本库的历史中仍然保存着文件。使用命令：
git cat-file -p HEAD~1:fileName.txt > fileName.txt
它等效于使用
git show HEAD~1:fileName.txt > fileName.txt
 
注意，其中 HEAD 后面的数字是HEAD的历史版本，0表示当前的提交，1表示上一次提交，2表示1的上一次提交……
 
更直接的方法是
git checkout HEAD~1 -- fileName.txt
(直接恢复，但不能在恢复的时候重命名)
 
4）文件忽略
 
在任意目录下，放置一个 .gitignore 文件即可，其内容格式一行一个，比如：
 
# 以'#' 开始的行，被视为注释.
# 忽略掉所有文件名是 foo.txt 的文件.
foo.txt
# 忽略所有生成的 html 文件,
*.html
# foo.html是手工维护的，所以例外.
!foo.html
#  忽略所有.o 和 .a文件.
*.[oa]
# 忽略obj文件夹
obj/
 
 
此 .gitignore 文件会对当前目录以及所有子目录生效
 
5）文件打包(归档)
 
如果你手工打包，比如使用 7zip/winzip/winrar 等工具打包，往往会将需要忽略的文件、或.git文件夹等打包进去。
而使用git自带的打包归档命令则不会：
 
git archive -o latest.zip HEAD
基于最新提交打包归档文件
git archive -o partial.tar HEAD src doc
只将src和doc目录打包归档
 
6）文件比较
 
git diff
工作区和暂存区文件的比较
git diff HEAD
工作区和版本库HEAD的比较
git diff --cached
暂存区和版本库的比较(或者使用 git diff --staged也一样)
 
 
 
【远程库管理】
 
1）分步管理
 
git fetch
将远程origin库的修改拉抓到本地
git merge origin/master
 
 
2）一步管理
 
git pull
等价于 git fetch 再 git merge
 
3）跟踪远程分支
 
git branch --track greet origin/greet
将远程的origin/greet分支拉到本地greet分支
 
4）裸库
 
裸库创建
git clone --bare hello hello.git
根据hello这个目录下的.git内容，创建一个没有文件的hello.git裸库
裸库加入
git remote add shared ../hello.git
将裸库以远程库的形式加入到原始库里
加入一个hello.git的远程库，取名为shared
将改动加入共享库
git push shared master
将改动更新到shared库中
 
如何从某个库中拉取内容？比如从共享库中更新内容到本地？
 
git remote add shared ../hello.git
将hello.git裸库加入本地的远程库，并取名为shared
git branch --track shared master
让裸库和本地库同步
git pull shared master
将裸库内容同步到本地
 
【查看历史commit的内容】
 
一、逐步查看版本库中的文件
 
1）先获得commit的ID， 通过log，得到最近的commit的ID
git hist --max-count=2
 
2）根据此commit的ID，查看其内容
git cat-file -p commitID
 
输出结果：
 
$ git cat-file -p 0cc067f
tree 096b74c56bfc6b40e754fc0725b8c70b2038b91e    ——这里就是目录树
parent e6790bb380b06feaee08426b99a1afc3bab17281
author Jim Weirich <jim (at) edgecase.com> 1316495180 +0800
committer Jim Weirich <jim (at) edgecase.com> 1316495180 +0800
 
3）根据tree的ID，获取目录树的内容，是一个文件名列表，每个文件都也有一个hash的ID
 
$ git cat-file -p 096b74c
100644 blob 28e0e9d6ea7e25f35ec64a43f569b550e8386f90        Rakefile
040000 tree e46f374f5b36c6f02fb3e9e922b79044f754d795        lib
 
4）根据文件的ID，查看文件的内容
 
git cat-file -p commitID
 
此处命令可以用 git show commitID ，效果相同
 
二、直接查看文件内容
 
不通过如上步骤，而是直接查看文件内容：
git show HEAD~0:fileName.txt
git show HEAD~1:fileName.txt
git show HEAD~2:fileName.txt
 
【git服务器】
 
建立本地git服务器
git daemon --verbose --export-all --base-path=.
