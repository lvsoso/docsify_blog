
**inspect**
```shell
docker inspect <containerid>
```

**namespace**

```shell

lsns -t <type>

ls -la /proc/<pid>/ns/

```

**nsenter**

```shell


nsenster -t <pid> -n ip addr

(base) lv@lv:docsify_blog$ nsenter --help

用法：
 nsenter [选项] [<程序> [<参数>...]]

以其他程序的名字空间运行某个程序。

选项：
 -a, --all              enter all namespaces
 -t, --target <pid>     要获取名字空间的目标进程
 -m, --mount[=<文件>]   进入 mount 名字空间
 -u, --uts[=<文件>]     进入 UTS 名字空间(主机名等)
 -i, --ipc[=<文件>]     进入 System V IPC 名字空间
 -n, --net[=<文件>]     进入网络名字空间
 -p, --pid[=<文件>]     进入 pid 名字空间
 -C, --cgroup[=<文件>]  进入 cgroup 名字空间
 -U, --user[=<文件>]    进入用户名字空间
 -S, --setuid <uid>     设置进入空间中的 uid
 -G, --setgid <gid>     设置进入名字空间中的 gid
     --preserve-credentials 不干涉 uid 或 gid
 -r, --root[=<目录>]     设置根目录
 -w, --wd[=<dir>]       设置工作目录
 -F, --no-fork          执行 <程序> 前不 fork
 -Z, --follow-context  根据 --target PID 设置 SELinux 环境

 -h, --help             display this help
 -V, --version          display version


PID=$(docker inspect --format "{{ .State.Pid}}" <container>) 

nsenter --target $PID --mount --uts --ips --net --pid
```

**unshare**

```shell
(base) lv@lv:docsify_blog$ unshare --help

用法：
 unshare [选项] [<程序> [<参数>...]]

以某些未与父(进程)共享的名字空间运行某个程序。

选项：
 -m, --mount[=<文件>]      取消共享 mounts 名字空间
 -u, --uts[=<文件>]        取消共享 UTS 名字空间(主机名等)
 -i, --ipc[=<文件>]        取消共享 System V IPC 名字空间
 -n, --net[=<file>]        取消共享网络名字空间
 -p, --pid[=<文件>]        取消共享 pid 名字空间
 -U, --user[=<文件>]       取消共享用户名字空间
 -C, --cgroup[=<文件>]     取消共享 cgroup 名字空间
 -f, --fork                在启动<程序>前 fork
     --mount-proc[=<目录>] 先挂载 proc 文件系统(连带打开 --mount)
 -r, --map-root-user       将当前用户映射为 root (连带打开 --user)
     --propagation slave|shared|private|unchanged
                           修改 mount 名字空间中的 mount 传播
 -s, --setgroups allow|deny  控制用户名字空间中的 setgroups 系统调用

 -h, --help                display this help
 -V, --version             display version

```

**cgroups**

control groups

```shell
# cpu

/sys/fs/cgroup/cpu/mydir
/sys/fs/cgroup/cpu/mydir/cgroup.procs
/sys/fs/cgroup/cpu/mydir/cpu.cfs_quota_us

# cpuacct
/sys/fs/cgroup/cpu/mydir/cpuacct.stat
/sys/fs/cgroup/cpu/mydir/cpuacct.usage

# memory
/sys/fs/cgroup/memory/mydir/
/sys/fs/cgroup/memory/mydir/memory.usage_in_bytes
/sys/fs/cgroup/memory/mydir/memory.max_usage_in_bytes
/sys/fs/cgroup/memory/mydir/memory.limit_in_bytes
/sys/fs/cgroup/memory/mydir/memory.soft_limit_in_bytes
/sys/fs/cgroup/memory/mydir/memory.oom_control
```

**cgroups driver**

systemd driver
cgroup driver


