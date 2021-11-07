
[Git基本原理介绍](https://space.bilibili.com/364122352/channel/detail?cid=150242&ctype=0)
[pro git 2nd Edition (2014)](https://git-scm.com/book/en/v2)

实时看文件变化
```shell
watch -n 1 -d find .
```

提交信息显示
```shell
# 图形化显示当前分支的提交日志
git log --graph --oneline

# 图形化显示当前分支的提交日志及每次提交的变更内容
git log --graph --patch

# 图形化显示所有分支的提交日志
git log --graph --oneline --all

# 图形化显示所有分支的提交日志及每次提交的变更内容
git log --graph --patch --all
```
#### **git init**
./.git/config 当前仓库配置文件

```shell
mkdir git-demo
cd git-demo
git init
git config --global -l
git config -l
```

#### **git add**

会在`.git`文件夹下添加`objects/xx`文件夹、`index`文件、`objects/xx/1223343435434abcd..`文件

git blob 对象存储文件内容，但是不存储文件名字

```shell
── info
│   └── exclude
├── objects
│   ├── 8d
│   │   └── 0e41234f24b6da002d962a26c2495ea16a425f
│   ├── info
│   └── pack

// 查看类型
// blob 存储文件内容
(base) lv@lv:git-demo$ git cat-file -t 8d0e41
blob

// 查看内容
(base) lv@lv:git-demo$ git cat-file -p 8d0e41
hello git


```
 

git的哈希是blob 文件内容的字节长度\0 文件内容
```shell
(base) lv@lv:git-demo$ cat hello.txt | git hash-object --stdin
8d0e41234f24b6da002d962a26c2495ea16a425f
```

#### **Working Directory, Staging Area(index)、git repo**

git add : add file info from Working Directory to Staging Area

git ls-files : list the files in `.git/index`

```shell
(base) lv@lv:git-demo$ git ls-files
hello.txt
tmp.txt
(base) lv@lv:git-demo$ git ls-files -s
100644 8d0e41234f24b6da002d962a26c2495ea16a425f 0	hello.txt
100644 8d0e41234f24b6da002d962a26c2495ea16a425f 0	tmp.txt
```
Staging Area save index of files.


git commit : sync file info from Staging Area to git repo.

generate `commit object`, save `tree object` hash and committer info.

```shell
(base) lv@lv:git-demo$ git commit -m "1st commit"
[master （根提交） 0c48c1a] 1st commit
 2 files changed, 2 insertions(+)
 create mode 100644 hello.txt
 create mode 100644 tmp.txt

(base) lv@lv:git-demo$ git cat-file -t 0c48
commit

(base) lv@lv:git-demo$ git cat-file -p 0c48
tree 570c2367196dad259e3f4e93707509ee94caeb2a
author lvsoso <wuqize5109@qq.com> 1636212657 +0800
committer lvsoso <wuqize5109@qq.com> 1636212657 +0800

1st committer

#local newest commit about each branch.
(base) lv@lv:git-demo$ cat .git/refs/heads/master 
0c48c1a380b7d98e4ef7633da17586543e6860e8

# current working branch.
(base) lv@lv:git-demo$ cat .git/HEAD 
ref: refs/heads/master
```

`tree object` save info about committed files.

```shell
(base) lv@lv:git-demo$ git cat-file -t 570c
tree
(base) lv@lv:git-demo$ git cat-file -p 570c
100644 blob 8d0e41234f24b6da002d962a26c2495ea16a425f	hello.txt
100644 blob 8d0e41234f24b6da002d962a26c2495ea16a425f	tmp.txt
```

modified file and add it, will create new blob object.

```shell
(base) lv@lv:git-demo$ git cat-file -t a201
blob
(base) lv@lv:git-demo$ git cat-file pt a201
fatal: invalid object type "pt"
(base) lv@lv:git-demo$ git cat-file -p a201
hello git
123

(base) lv@lv:git-demo$ git add tmp.txt 
(base) lv@lv:git-demo$ git commit -m "2nd commit"
[master 12d8445] 2nd commit
 1 file changed, 1 insertion(+)

# create new commit object.
(base) lv@lv:git-demo$ git ls-files -s
100644 8d0e41234f24b6da002d962a26c2495ea16a425f 0	hello.txt
100644 a201af5e44282fabf8c6cf57396ba71b106a4dbb 0	tmp.txt

# parent hash was last one commit before this time.
(base) lv@lv:git-demo$ git cat-file -p 12d8
tree a14914de69b490e7621439991557d0259bcb9944
parent 0c48c1a380b7d98e4ef7633da17586543e6860e8
author lvsoso <wuqize5109@qq.com> 1636215895 +0800
committer lvsoso <wuqize5109@qq.com> 1636215895 +0800

2nd commit

```

add new fold and new file, `commit` will create new `commit object`

```shell
(base) lv@lv:git-demo$ git cat-file -p a3519
tree eca43b8e136c4ca4e209c07c82a20f433e43544c
parent 12d8445ec1d43ba9e73435b2149d201f540c01e1
author lvsoso <wuqize5109@qq.com> 1636216767 +0800
committer lvsoso <wuqize5109@qq.com> 1636216767 +0800

3rd commit

(base) lv@lv:git-demo$ git cat-file -p eca4
040000 tree b4540ce0bad63a0f40de1619b97a4589a9259496	dir1
100644 blob 8d0e41234f24b6da002d962a26c2495ea16a425f	hello.txt
100644 blob a201af5e44282fabf8c6cf57396ba71b106a4dbb	tmp.txt
(base) lv@lv:git-demo$ git cat-file -p 12d8
tree a14914de69b490e7621439991557d0259bcb9944
parent 0c48c1a380b7d98e4ef7633da17586543e6860e8
author lvsoso <wuqize5109@qq.com> 1636215895 +0800
committer lvsoso <wuqize5109@qq.com> 1636215895 +0800

2nd commit

(base) lv@lv:git-demo$ git ls-files -s
100644 7c8ac2f8d82a1eb5f6aaece6629ff11015f91eb4 0	dir1/file3.txt
100644 8d0e41234f24b6da002d962a26c2495ea16a425f 0	hello.txt
100644 a201af5e44282fabf8c6cf57396ba71b106a4dbb 0	tmp.txt
```

#### **file status**
untracked: `git add` -> staged.
modified: staged file modified, `git add` -> staged.
staged: `git commit` -> unmodified.
unmodified: untrack -> untracked.


#### **Branches**
```shell
branch->commit->tree->blob
```

Branches: named pointers to commit.
Master: a branch, the canonical mainline branch by default.
HEAD: special pointer, current active branch, always point to the latest commit.

```shell
(base) lv@lv:git-demo$ cat .git/HEAD 
ref: refs/heads/master
(base) lv@lv:git-demo$ cat .git/refs/heads/master 
a3519dbbfea1b36440d11f11f0add366dea086d9
(base) lv@lv:git-demo$ git cat-file -p a351
tree eca43b8e136c4ca4e209c07c82a20f433e43544c
parent 12d8445ec1d43ba9e73435b2149d201f540c01e1
author lvsoso <wuqize5109@qq.com> 1636216767 +0800
committer lvsoso <wuqize5109@qq.com> 1636216767 +0800

3rd commit
(base) lv@lv:git-demo$ git log
commit a3519dbbfea1b36440d11f11f0add366dea086d9 (HEAD -> master)
Author: lvsoso <wuqize5109@qq.com>
Date:   Sun Nov 7 00:39:27 2021 +0800

    3rd commit
```

**git branch**

```shell
git branch
git branch <branch_name>
git branch -D <branch_name>
git branch --delete <branch_name>

# change the current active branch or commit
git checkout 
```
`git branch` create new alias to the pointer.

```shell
(base) lv@lv:git-demo$ git branch dev
(base) lv@lv:git-demo$ git branch
  dev
* master
(base) lv@lv:git-demo$ git log
commit a3519dbbfea1b36440d11f11f0add366dea086d9 (HEAD -> master, dev)
Author: lvsoso <wuqize5109@qq.com>
Date:   Sun Nov 7 00:39:27 2021 +0800

(base) lv@lv:git-demo$ ls .git/refs/heads/
dev  master

(base) lv@lv:git-demo$ git checkout dev 
切换到分支 'dev'
(base) lv@lv:git-demo$ cat .git/HEAD 
ref: refs/heads/dev
```
> delete branch not remove the object about this branch.
> not real delete the commit about this branch, just delete the pointer to the commit.

`git checkout` can create new branch from a commit.

`git checkout -b <branch_name>` create named branch from current pointer.

```shell
(base) lv@lv:git-demo$ git checkout 12d8445ec
注意：正在检出 '12d8445ec'。

您正处于分离头指针状态。您可以查看、做试验性的修改及提交，并且您可以通过另外
的检出分支操作丢弃在这个状态下所做的任何提交。

如果您想要通过创建分支来保留在此状态下所做的提交，您可以通过在检出命令添加
参数 -b 来实现（现在或稍后）。例如：

  git checkout -b <新分支名>

HEAD 目前位于 12d8445 2nd commit
```

`git reflog` log can help us find the branch which been deleted.

read info from `.git/logs/refs/`.

```shell
(base) lv@lv:git-demo$ git reflog 
12d8445 (HEAD) HEAD@{0}: checkout: moving from dev to 12d8445ec
a3519db (master, dev) HEAD@{1}: checkout: moving from master to dev
a3519db (master, dev) HEAD@{2}: commit: 3rd commit
12d8445 (HEAD) HEAD@{3}: commit: 2nd commit
0c48c1a HEAD@{4}: commit (initial): 1st commit

```

#### **git diff**
`git diff` compare the working directory and staged area.
`git diff --cached` compare the the staged area and git repository.
```shell
(base) lv@lv:git-demo$ git diff tmp.txt
diff --git a/tmp.txt b/tmp.txt
index a201af5..a50d0f9 100644
--- a/tmp.txt
+++ b/tmp.txt
@@ -1,2 +1,3 @@
 hello git
 123
+yy
(base) lv@lv:git-demo$ cat tmp.txt 
hello git
123
yy
(base) lv@lv:git-demo$ git cat-file -p a201
hello git
123
```

#### **remote repo**
```shell
(base) lv@lv:git-demo$ git remote add origin git@github.com:lvsoso/git-demo.git
(base) lv@lv:git-demo$ cat .git/config 
[core]
	repositoryformatversion = 0
	filemode = true
	bare = false
	logallrefupdates = true
[remote "origin"]
	url = git@github.com:lvsoso/git-demo.git
	fetch = +refs/heads/*:refs/remotes/origin/*

(base) lv@lv:git-demo$ git branch -M main
(base) lv@lv:git-demo$ cat .git/refs/heads/
dev   main

(base) lv@lv:git-demo$ git push -u origin main
对象计数中: 10, 完成.
Delta compression using up to 6 threads.
压缩对象中: 100% (6/6), 完成.
写入对象中: 100% (10/10), 786 bytes | 786.00 KiB/s, 完成.
Total 10 (delta 0), reused 0 (delta 0)
To github.com:lvsoso/git-demo.git
 * [new branch]      main -> main
分支 'main' 设置为跟踪来自 'origin' 的远程分支 'main'。

(base) lv@lv:git-demo$ cat .git/refs/remotes/origin/main 
a3519dbbfea1b36440d11f11f0add366dea086d9

(base) lv@lv:git-demo$ cat .git/refs/heads/main 
a3519dbbfea1b36440d11f11f0add366dea086d9

(base) lv@lv:git-demo$ git cat-file -t a351
commit

(base) lv@lv:git-demo$ git log
commit a3519dbbfea1b36440d11f11f0add366dea086d9 (HEAD -> main, origin/main, dev)
Author: lvsoso <wuqize5109@qq.com>
Date:   Sun Nov 7 00:39:27 2021 +0800
```

The `remote repo` in the server just like our `local repo`.

#### **object compress**

```shell
(base) lv@lv:files$ git init 
已初始化空的 Git 仓库于 /home/lv/tmp/files/.git/

# add  file
(base) lv@lv:files$ git add file1.txt 
(base) lv@lv:files$ tree .git/objects/
.git/objects/
├── 15
│   └── 762ba5ab98efb881983b0fb893d2d1d4196bd0
├── info
└── pack

3 directories, 1 file
(base) lv@lv:files$ ls -lh .git/objects/15/
总用量 3.4M
-r--r--r-- 1 lv lv 3.4M 11月  7 12:40 762ba5ab98efb881983b0fb893d2d1d4196bd0

# add random file
(base) lv@lv:files$ git add file2.txt 
(base) lv@lv:files$ ls -lh .git/objects/34/
总用量 5.9M
-r--r--r-- 1 lv lv 5.9M 11月  7 12:41 46ce38d78bd5afd6b4a3f88bba1ba4ab118b40

# add bin file
(base) lv@lv:files$ git add file.bin
(base) lv@lv:files$ ls -lh .git/objects/b4/
总用量 11M
-r--r--r-- 1 lv lv 11M 11月  7 12:42 4a77214709868f7a7850511ecf015e472d82d0

(base) lv@lv:files$ ls -l .git/objects/b4/
总用量 10244
-r--r--r-- 1 lv lv 10488970 11月  7 12:42 4a77214709868f7a7850511ecf015e472d82d0

(base) lv@lv:files$ ls -l file.bin 
-rw-rw-r-- 1 lv lv 10485760 11月  7 12:08 file.bin

```

Every time we add the file, will create a new blob object. 

If the file have random contents, it generates the blob object bigger than the file have unrandom contents.

Every time we commit will consume the space in the repo.

`git gc` 

```shell
(base) lv@lv:files$ du -h .git
4.0K	.git/objects/pack
8.0K	.git/objects/fa
8.0K	.git/objects/81
4.0K	.git/objects/info
5.9M	.git/objects/0c
8.0K	.git/objects/3d
8.0K	.git/objects/e8
5.9M	.git/objects/22
5.9M	.git/objects/34
18M	.git/objects
8.0K	.git/info
8.0K	.git/refs/heads
4.0K	.git/refs/tags
16K	.git/refs
52K	.git/hooks
4.0K	.git/branches
8.0K	.git/logs/refs/heads
12K	.git/logs/refs
20K	.git/logs
18M	.git
(base) lv@lv:files$ git gc
对象计数中: 6, 完成.
Delta compression using up to 6 threads.
压缩对象中: 100% (4/4), 完成.
写入对象中: 100% (6/6), 完成.
Total 6 (delta 1), reused 0 (delta 0)
(base) lv@lv:files$ du -h .git
5.6M	.git/objects/pack
8.0K	.git/objects/info
5.9M	.git/objects/34
12M	.git/objects
12K	.git/info
4.0K	.git/refs/heads
4.0K	.git/refs/tags
12K	.git/refs
52K	.git/hooks
4.0K	.git/branches
8.0K	.git/logs/refs/heads
12K	.git/logs/refs
20K	.git/logs
12M	.git
```

#### **pack and unpack**

```shell
# pack 
(base) lv@lv:files$ tree .git/objects/pack/
.git/objects/pack/
├── pack-a7140efc921894dd3a2b218606f05ea70dba7176.idx
└── pack-a7140efc921894dd3a2b218606f05ea70dba7176.pack
0 directories, 2 files

(base) lv@lv:files$ cat  .git/packed-refs 
# pack-refs with: peeled fully-peeled sorted 
fa6f3c061b39c1647c4d47dd290ec305bf09e22b refs/heads/master

(base) lv@lv:files$ git verify-pack -v  .git/objects/pack/pack-a7140efc921894dd3a2b218606f05ea70dba7176.idx
fa6f3c061b39c1647c4d47dd290ec305bf09e22b commit 202 142 12
e8a79f611466f84209ff3d47834a2c8403f97596 commit 154 113 154
0c7ca0ce07b920785f123c5cb4bd63e15ec14336 blob   12349857 5794708 267
22c50fbe1b848115f4901c420368dc86f1087b10 blob   576 371 5794975 1 0c7ca0ce07b920785f123c5cb4bd63e15ec14336
3d005bd84deebf46707c799e5e2c1f3f8a9df432 tree   37 47 5795346
8163e2d8ecab117c194bf174525ee7b601ef7283 tree   37 48 5795393
非 delta：5 个对象
链长 = 1: 1 对象
.git/objects/pack/pack-a7140efc921894dd3a2b218606f05ea70dba7176.pack: ok

# unpack

mv .git/objects/pack/pack-a7140efc921894dd3a2b218606f05ea70dba7176.pack .git/
git unpack-objects < .git/pack-a7140efc921894dd3a2b218606f05ea70dba7176.pack
```

#### **git gc**

gabage object : such as gabage blob, generated by `git add` or branch deleted.

use `git gc` pack, but not clean.

use `git fsck` check the unreachable object.

use `git prune` clean the unreachable object.

```shell
(base) lv@lv:files$ git fsck 
检查对象目录中: 100% (256/256), 完成.
检查对象中: 100% (6/6), 完成.
dangling blob 3446ce38d78bd5afd6b4a3f88bba1ba4ab118b40

(base) lv@lv:files$ git prune

(base) lv@lv:files$ git fsck 
检查对象目录中: 100% (256/256), 完成.
检查对象中: 100% (6/6), 完成.
```

[How to remove unused objects from a git repository?](https://stackoverflow.com/questions/3797907/how-to-remove-unused-objects-from-a-git-repository/14729486)

```shell
git -c gc.reflogExpire=0 -c gc.reflogExpireUnreachable=0 \
  -c gc.rerereresolved=0 -c gc.rerereunresolved=0 \
  -c gc.pruneExpire=now gc "$@"
```

#### **fast forward merge**

master: point to `a`.
bugfix: branch form master and point to `b`， fast than master branch.
fast forward merge:  merge master point to `b`.

HEAD: current latest commit.
ORIG_HEAD: commit before mege, can use for rollback, maybe generate file form target mege branch to the working directory.

#### **3 way merge**

```shell
                                master
c0 <- c1 <- c2 <- c4
                        ^
                        |_ c3
                            bugfix
```

will generate a new merge commit.

```shell
                                            master
c0 <- c1 <- c2 <- c4 <- c5
                        ^                   |
                        |-- c3<------|
                            bugfix
```
commit `c5` had two parent node.

Enconter `conflict merge`, we need to handle conflict if git can not handle it.

also create a new merge commit.

#### **git rebase**

```shell
                            master
c0 <- c1 <- c2 <- c4
                        ^
                        |_ c3
                            bugfix

            |
            |

                            master
c0 <- c1 <- c2 <- c4
                                    ^
                                    |_ c3
                                        bugfix

            |
            |

                                            master
c0 <- c1 <- c2 <- c4 - c5
                                            |
                                        bugfix
```

#### **git tag**

.git/refs/tags/xxx

git tag <tag name>: create lightweight tag.
git tag -a <tag name> -m <tag message> create anotated tag.
git tag -a <tag name> <commit SHA1 value> create anotated tag point to a commit.
git tag: list tags.
git tag -d <tag name>: delete tag.


**lightweight tag**

no create tag object.

```shell
(base) lv@lv:files$ git tag v1.0.0

(base) lv@lv:files$ ls .git/refs/tags/v1.0.0 
.git/refs/tags/v1.0.0

(base) lv@lv:files$ cat  .git/refs/tags/v1.0.0 
fa6f3c061b39c1647c4d47dd290ec305bf09e22b

(base) lv@lv:files$ git log
commit fa6f3c061b39c1647c4d47dd290ec305bf09e22b (HEAD -> master, tag: v1.0.0)
Author: lvsoso <wuqize5109@qq.com>
```

**anotated tag**

will create tag object.

```shell
(base) lv@lv:files$ git tag -d v1.0.0
已删除标签 'v1.0.0'（曾为 fa6f3c0）

(base) lv@lv:files$ git tag -a v1.0.0 -m "version v1.0.0"

(base) lv@lv:files$ git tag
v1.0.0

(base) lv@lv:files$ cat .git/refs/tags/v1.0.0 
62b0508dcd61ba4209e5c51726c81efae20157f6

(base) lv@lv:files$ git cat-file -t 62b05
tag

(base) lv@lv:files$ git cat-file -p 62b05
object fa6f3c061b39c1647c4d47dd290ec305bf09e22b
type commit
tag v1.0.0
tagger lvsoso <wuqize5109@qq.com> 1636271684 +0800

version v1.0.0

```

tag object contain timestamp, even have the same tag name, will create different tag object.

#### **remote branch**

remote branch info log in our local file when we `git clone`.
```shell
(base) lv@lv:git-demo$ cat .git/packed-refs 
# pack-refs with: peeled fully-peeled sorted 
a3519dbbfea1b36440d11f11f0add366dea086d9 refs/heads/dev
a3519dbbfea1b36440d11f11f0add366dea086d9 refs/heads/main
a3519dbbfea1b36440d11f11f0add366dea086d9 refs/remotes/origin/main

```

`git fetch` - Download objects and refs from another repository.
`git fetch --prune` - Can clean stale branch.
```shell
(base) lv@lv:git-demo$ git fetch 
remote: Enumerating objects: 4, done.
remote: Counting objects: 100% (4/4), done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
展开对象中: 100% (3/3), 完成.
来自 github.com:lvsoso/git-demo
   a3519db..7e200e3  main       -> origin/main
(base) lv@lv:git-demo$ cat .git/packed-refs 
# pack-refs with: peeled fully-peeled sorted 
a3519dbbfea1b36440d11f11f0add366dea086d9 refs/heads/dev
a3519dbbfea1b36440d11f11f0add366dea086d9 refs/heads/main
a3519dbbfea1b36440d11f11f0add366dea086d9 refs/remotes/origin/main
(base) lv@lv:git-demo$ cat .git/refs/remotes/origin/main 
7e200e31751f5ace8dc8b66fbec20ad5985cc8c8
```

`git remote show origin`
`git remote prune`
```shell
(base) lv@lv:git-demo$ git remote show origin
* 远程 origin
  获取地址：git@github.com:lvsoso/git-demo.git
  推送地址：git@github.com:lvsoso/git-demo.git
  HEAD 分支：main
  远程分支：
    main 已跟踪
  为 'git pull' 配置的本地分支：
    main 与远程 main 合并
  为 'git push' 配置的本地引用：
    main 推送至 main (最新)
```

#### **git pull**

`git pull` - Fetch from and integrate with another repository or a local branch

`git fetch` + `git merge`

```shell
(base) lv@lv:git-demo$ git pull -v
# fetch
remote: Enumerating objects: 5, done.
remote: Counting objects: 100% (5/5), done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 3 (delta 1), reused 0 (delta 0), pack-reused 0
展开对象中: 100% (3/3), 完成.
来自 github.com:lvsoso/git-demo
   68f7144..282c0a5  dev        -> origin/dev
 = [最新]            main       -> origin/main
# merge
更新 68f7144..282c0a5
Fast-forward
 dev.txt | 1 +
 1 file changed, 1 insertion(+)
```

`68f7144` is `ORIG_HEAD` of local dev branch.
`git reset --hard ORIG_HEAD` can rollback.
```shell
(base) lv@lv:git-demo$ cat .git/ORIG_HEAD 
68f714427812bef213e8986886609cacd89d1274
(base) lv@lv:git-demo$ git reset --hard ORIG_HEAD
HEAD 现在位于 68f7144 Update dev.txt
(base) lv@lv:git-demo$ git status
位于分支 dev
您的分支落后 'origin/dev' 共 1 个提交，并且可以快进。
  （使用 "git pull" 来更新您的本地分支）

无文件要提交，干净的工作区
```

```shell
# fast-forward
        A---B---C master on origin
        /
D---E
        ^
        origin/master in your repository


D---E---A---B---C master on origin
                               ^
                origin/master in your repository 
```

```shell
# 3-way merge
        A---B---C master on origin
        /
D---E---F---G master
        ^
        origin/master in your repository

        A---B---C origin/master
        /                 \
D---E---F---G---H master

```



#### **FETCH_HEAD**

create a new branch in remote repository and commit a new file.

fetch it in local.

```shell
(base) lv@lv:git-demo$ git fetch 
remote: Enumerating objects: 4, done.
remote: Counting objects: 100% (4/4), done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
展开对象中: 100% (3/3), 完成.
来自 github.com:lvsoso/git-demo
 * [新分支]          dev        -> origin/dev
(base) lv@lv:git-demo$ tree .git/
.git/
├── branches
├── COMMIT_EDITMSG
├── config
├── description
├── `FETCH_HEAD` <------
├── HEAD
├── hooks
│   ├── applypatch-msg.sample
│   ├── commit-msg.sample
│   ├── fsmonitor-watchman.sample
│   ├── post-update.sample
│   ├── pre-applypatch.sample
│   ├── pre-commit.sample
│   ├── prepare-commit-msg.sample
│   ├── pre-push.sample
│   ├── pre-rebase.sample
│   ├── pre-receive.sample
│   └── update.sample
├── index
├── info
│   ├── exclude
│   └── refs
├── logs
│   ├── HEAD
│   └── refs
│       ├── heads
│       │   ├── dev
│       │   └── main
│       └── remotes
│           └── origin
│               ├── `dev` <------
│               └── main
├── objects
│   ├── 41
│   │   └── `0ca19fdc64ecc1dc128e7fe49227f0e6c3a295` <------
│   ├── 5a
│   │   └── `d28e22767f979da2c198dc6c1003b25964e3da` <------
│   ├── d6
│   │   └── `ab78d363df26d8f95a453b6e7141a4c8f1cf94` <------
│   ├── info
│   │   └── packs
│   └── pack
│       ├── pack-22fd9b810a54502ed9210cb853475cfc8d18c9e1.idx
│       └── pack-22fd9b810a54502ed9210cb853475cfc8d18c9e1.pack
├── ORIG_HEAD
├── packed-refs
└── refs
    ├── heads
    ├── remotes
    │   └── origin
    │       └── `dev` <------
    └── tags

19 directories, 33 files
```

the first line in `FETCH_HEAD` show which local branch we exec `git fetch`.

```shell
(base) lv@lv:git-demo$ cat .git/FETCH_HEAD 
7e200e31751f5ace8dc8b66fbec20ad5985cc8c8		branch 'main' of github.com:lvsoso/git-demo
d6ab78d363df26d8f95a453b6e7141a4c8f1cf94	not-for-merge	branch 'dev' of github.com:lvsoso/git-demo
```
#### **git pull**

```shell
(base) lv@lv:git-demo$ git checkout -b featch-1
切换到一个新分支 'featch-1'

(base) lv@lv:git-demo$ git branch -vv
  dev      282c0a5 [origin/dev] Update dev.txt
* featch-1 282c0a5 Update dev.txt
  main     7e200e3 [origin/main] Create README.md

(base) lv@lv:git-demo$ mkdir feature-1
(base) lv@lv:git-demo$ echo 'feature-1' > feature-1/feature-1.txt

(base) lv@lv:git-demo$ git add feature-1

(base) lv@lv:git-demo$ git commit -m "add feature1"
[featch-1 c2658cb] add feature1
 1 file changed, 1 insertion(+)
 create mode 100644 feature-1/feature-1.txt

```


```shell
# just push our local feature-1 content to origin/feature-1
git push  origin featch-1

# bind local feature1 to origin/feature-1
git push --set-upstream origin featch-1
git push -u origin featch-1
```

```shell
(base) lv@lv:git-demo$ git push --set-upstream origin featch-1
对象计数中: 4, 完成.
Delta compression using up to 6 threads.
压缩对象中: 100% (2/2), 完成.
写入对象中: 100% (4/4), 327 bytes | 327.00 KiB/s, 完成.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local object.
remote: 
remote: Create a pull request for 'featch-1' on GitHub by visiting:
remote:      https://github.com/lvsoso/git-demo/pull/new/featch-1
remote: 
To github.com:lvsoso/git-demo.git
 * [new branch]      featch-1 -> featch-1
分支 'featch-1' 设置为跟踪来自 'origin' 的远程分支 'featch-1'。
```


#### **git hook**

- pre-commit
- prepare-commit-msg
- commit-msg
- post-commit ?
- pre-push
- pre-receive
- update
- post-update
- 

```shell
├── hooks
│   ├── applypatch-msg.sample
│   ├── commit-msg.sample
│   ├── fsmonitor-watchman.sample
│   ├── post-update.sample
│   ├── pre-applypatch.sample
│   ├── pre-commit.sample
│   ├── prepare-commit-msg.sample
│   ├── pre-push.sample
│   ├── pre-rebase.sample
│   ├── pre-receive.sample
│   └── update.sample
...
```

`pre-commit.sample` is a shell template.

**user hook for all branch and remote repo.**

[git-hook-python](https://github.com/lvsoso/git-hook-python)

#### **git submodule**

`git submodule add git@github.com:lvsoso/git-submodule.git` add

- 新文件：   .gitmodules
- 新文件：   git-submodule

`git submodule update --init --recursive` init

`git submodule update --remote` update

```shell

(base) lv@lv:git-demo$ git status
位于分支 featch-1
您的分支与上游分支 'origin/featch-1' 一致。

要提交的变更：
  （使用 "git reset HEAD <文件>..." 以取消暂存）

	新文件：   .gitmodules
	新文件：   git-submodule

(base) lv@lv:git-demo$ cat .gitmodules 
[submodule "git-submodule"]
	path = git-submodule
	url = git@github.com:lvsoso/git-submodule.git

(base) lv@lv:git-demo$ tree  .git/modules/
.git/modules/
└── git-submodule
    ├── branches
    ├── config
    ├── description
    ├── HEAD
    ├── hooks
    │   ├── applypatch-msg.sample
    │   ├── commit-msg.sample
    │   ├── fsmonitor-watchman.sample
    │   ├── post-update.sample
    │   ├── pre-applypatch.sample
    │   ├── pre-commit.sample
    │   ├── prepare-commit-msg.sample
    │   ├── pre-push.sample
    │   ├── pre-rebase.sample
    │   ├── pre-receive.sample
    │   └── update.sample
    ├── index
    ├── info
    │   └── exclude
    ├── logs
    │   ├── HEAD
    │   └── refs
    │       ├── heads
    │       │   └── main
    │       └── remotes
    │           └── origin
    │               └── HEAD
    ├── objects
    │   ├── info
    │   └── pack
    │       ├── pack-5c78dd5d50ee9a4468e58c815de6bbf51d01db6a.idx
    │       └── pack-5c78dd5d50ee9a4468e58c815de6bbf51d01db6a.pack
    ├── packed-refs
    └── refs
        ├── heads
        │   └── main
        ├── remotes
        │   └── origin
        │       └── HEAD
        └── tags

17 directories, 24 files
```

#### **git worktree**

use stash
```shell
git add .
git stash
git checkout master
...
git checkout dev
git stash pop
```

use worktree
```shell
# create worktree for branch save changed
git worktree add ../worktree-master master
```

```shell
git worktree list
git worktree remove worktree-master
```