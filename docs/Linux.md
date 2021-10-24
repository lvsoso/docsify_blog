### libusb
```shell
sudo apt install libusb-dev libudev-dev
```

```shell
watch -n 1 "cat /proc/vmstat | grep dirty"

perf record -a -g -p <pid>

perf report

cd /sys/kernel/debug/tracing
echo vfs_write >> set_ftrace_filter
echo xfs_file_write_iter >> set_ftrace_filter
echo xfs_file_buffered_aio_write >> set_ftrace_filter
echo iomap_file_buffered_write
echo iomap_file_buffered_write >> set_ftrace_filter
echo pagecache_get_page >> set_ftrace_filter
echo try_to_free_mem_cgroup_pages >> set_ftrace_filter
echo try_charge >> set_ftrace_filter
echo mem_cgroup_try_charge >> set_ftrace_filter
echo function_graph > current_tracer
echo 1 > tracing_on
```

### sourcegraph
```shell
docker run -d --publish 7080:7080 --publish 127.0.0.1:3370:3370 --name sg --restart=always --volume ~/.sourcegraph/config:/etc/sourcegraph --volume ~/.sourcegraph/data:/var/opt/sourcegraph sourcegraph/server:3.23.0

curl -L https://sourcegraph.com/.api/src-cli/src_linux_amd64 -o /usr/local/bin/src

chmod +x /usr/local/bin/src

nohup src serve-git >> /dev/null  $2>1 &
```
### Ovirt

agent,https://computingforgeeks.com/install-ovirt-guest-agent-linux/

```shell
# ubuntu 18.04
sudo apt -y update
sudo apt-get install -y ovirt-guest-agent

# ubuntu 16.04
sudo tee /etc/apt/sources.list.d/ovirt-guest-agent.list<<EOF
deb http://download.opensuse.org/repositories/home:/evilissimo:/ubuntu:/16.04/xUbuntu_16.04/ /
EOF

wget http://download.opensuse.org/repositories/home:/evilissimo:/ubuntu:/16.04/xUbuntu_16.04//Release.key
sudo apt-key add - < Release.key 
sudo apt-get update

sudo apt-get install ovirt-guest-agent
```



```shell
sudo yum -y install epel-release
sudo yum -y install qemu-guest-agent
sudo systemctl start qemu-guest-agent
sudo systemctl enable --now qemu-guest-agent
```



spice virt-manager, https://virt-manager.org/download/

```shell

sudo apt-get install virt-manager 

```



### QEMU

```shell
1765  cd code/
 1766  mkdir riscv64-linux
 1767  cd riscv64-linux/
 1768  git clone https://github.com/qemu/qemu
 1769  git clone https://github.com/torvalds/linux
 1770  ls
 1771  git clone https://git.busybox.net/busybox
 1772  proxychains4 git clone https://git.busybox.net/busybox
 1773  cd qemu/
 1774  git checkout v5.0.0
 1775  ./configure  --target-list=riscv64-softmmu
 1776  sudo apt-get install glib
 1777  sudo apt-get install libglib2.0-dev
 1778  ./configure  --target-list=riscv64-softmmu
 1779  sudo apt-get install pixman-dev
 1780  sudo apt-get install libpixman-dev
 1781  sudo apt-get install libpixman-1.0-dev
 1782  sudo apt-get install libpixman
 1783  sudo apt-get install libpixman-1-dev
 1784  ./configure  --target-list=riscv64-softmmu
 1785  make -j $(nproc)
 1786  ./configure  --target-list=riscv64-softmmu
 1787  make -j $(nproc)
 1788  vim  docs/qemu-option-trace.rst.inc
 1789  make -j $(nproc)
 1790  cd ../linux/
 1791  make ARCH=riscv CROSS_COMPILE=riscv64-unknown-linux-gnu- defconfig
 1792  cd -
 1793  sudo make install
 1794  cd -
 1795  make ARCH=riscv CROSS_COMPILE=riscv64-unknown-linux-gnu- defconfig
 1796  history

```

### OSI

- 物理层：数据位（Bit）
- 数据链路层：数据帧（Frame）
- 网络层：数据包（Packet）
- 传输层：数据段（Segment）
- 五层以上：数据（Data）


```shell
# show gateway
ip route show


# show MTU
netstat -i

```

### IP
IP 命令主要用来显示或操纵 Linux 主机的路由、网络设备、策略路由和隧道。

> ip [option][object][command][arguments]

object

- link	网络设备
- address	设备的协议地址
- route	路由表条目
- neighbour	ARP/NDISC 缓冲区条目


command

- 对象的增加（add）、删除（delete）和展示（show 或者 list）。


```shell
# 显示网络设备运行状态
ip link list

# 输出详细网络信息
ip -s link list

# 显示邻居表（ARP 表）
ip neigh list
```

ip 分类
- A类	1.0.0.0 ~ 127.255.255.255
- B类	128.0.0.0 ~ 191.255.255.255
- C类	192.0.0.0 ~ 223.255.255.255
- D类	224.0.0.0 ~ 239.255.255.255
- E类	240.0.0.0 ~ 255.255.255.254

特殊ip

- 0.0.0.0/8	本网络（仅作为源地址时合法）	RFC 5735
- 10.0.0.0/8	专用网络	RFC 1918
- 127.0.0.0/8	环回	RFC 5735
- 172.16.0.0/12	专用网络	RFC 1918
- 192.168.0.0/16	专业网络	RFC 1918
- 255.255.255.255	广播	RFC 919


### socket


domain：协议域也称为协议族，决定了 socket 的地址类型，在通信中必须采用对应的地址。常用的协议族包括：

- AF_INET IPv4 协议
- AF_INET6 IPv6 协议

type：指定 socket 类型。常见的类型包括：

- SOCK_STREAM 字节流套接字
- SOCK_DGRAM 数据报套接字
- SOCK_RAW 原始套接字

protocol：指定协议。常用的协议包括：

- IPPROTO_TCP TCP 传输协议
- IPPTOTO_UDP UDP 传输协议

### telnet

```shell
telnet 192.168.12.2 22

telnet 192.168.12.2 3389
```

### ping

```shell
ping IP地址
```

- -n	输出数值
- -c	指定发送一定数目的包
- -i	设定间隔几秒送一个网络封包给一台机器，预设值是一秒送一次
- -t	设置存活数值 TTL 的大小
- -v	详细显示执行过程
- -R	记录路由过程

### traceroute

Traceroute 是用来侦测主机到目的主机之间经过的路由情况。

1)发送一份TTL 为 1 的 IP 数据报给目的主机，经过第一个路由器时，TTL 值被减为 0，则第一个路由器丢弃该数据报，并返回一份超时 ICMP 报文，于是就得到了路径中第一个路由器的地址；

2)再发送一份 TTL 值为 2的数据报，就可以得到第二个路由器的地址；

3)以此类推，一直到到达目的主机为止，这样便记录下了路径上所有的路由 IP。

> traceroute [参数][主机]

- -f	设置第一个检测数据包的存活数值TTL的大小
- -g	设置来源网关，最多可设置8个
- -i	使用指定的网络界面送出数据包
- -m	设置检测数据包的最大存活数值TTL的大小
- -p	设置 UDP 传输协议的通信端口
- -v	详细显示指令的执行过程
- -w	设置等待远端主机回报的时间


### ARP

通过目标设备的 IP地址（32 位），查询目标设备的 MAC 地址（48 位），以保证通信的顺利进行。

### RARP

RARP：逆地址解析协议，是将局域网中某个主机的物理地址转换为 IP 地址，和 ARP 刚好相反。

### netstat

netstat 命令用于显示与 IP、TCP、UDP 和 ICMP 协议相关的统计数据，一般用于检验本机各端口的网络连接情况。

> netstat [参数][IP]

- -a	显示所有选项
- -A	列出网络类型连线中的相关地址
- -n	直接使用 IP 地址，而不通过域名服务器
- -t	显示 TCP 传输协议的连线状况
- -u	显示 UDP 传输协议的连线状况
- -l	显示监控中的服务器的 Socket
- -c	持续列出网络状态
- -i	显示网络界面信息表单


```shell
# 选出 tcp 和 udp 的
netstat -lunat

# 显示所有端口
netstat -a
# 显示所有设备
netstat -i
# 显示路由表信息
netstat -r
```

### route
用于显示和操作 IP 路由表。

> route [参数]

- -c	显示更多信息
- -n	不解析名字
- -f	请求相互所有网关入口的路由表
- -p	和 add 命令一起使用路由具有永久性

### mtr
判断网络连通性的工具。

- -s	指定 ping 数据包的大小
- -a	用来设置发送数据包的 IP 地址
- -r	报告模式显示
- -v	显示 mtr 的版本信息
- -h	帮助命令
- -n	不对 IP 地址做域名解析
- -4	IPv4
- -6	IPv6


```shell
(base) lv@lv-tp:docsify_blog$ mtr -r xxx.xxx.xxx.xxx
Start: 2021-08-18T06:30:58+0800
HOST: lv-tp                       Loss%   Snt   Last   Avg  Best  Wrst StDev
  1.|-- 192.168.2.1                0.0%    10    3.4   3.6   3.2   5.9   0.8
  2.|-- 192.168.1.1                0.0%    10    3.3   2.8   1.4   4.1   0.9
  3.|-- 10.70.0.1                  0.0%    10    3.3   6.1   2.9  13.8   3.9
  4.|-- 125.33.184.237             0.0%    10    3.8   4.4   3.0   5.6   0.7
  5.|-- 123.126.0.38               0.0%    10    5.5   5.4   3.9   5.8   0.6
  6.|-- 61.49.142.158              0.0%    10    6.6   7.3   4.9  11.2   1.9
  7.|-- ???                       100.0    10    0.0   0.0   0.0   0.0   0.0
  8.|-- xxx.xxx.xxx.xxx            0.0%    10    7.6   7.2   6.0   7.7   0.6
```
- 第一列：显示 IP 地址和本机域名。
- 第二列：显示对应 IP 的丢包率。
- 第三列：snt:10 设置每秒发送数据包的数量，默认值是 10 。
- 第四列：显示最近一次的返回时延。
- 第五列：平均值，是发送 ping 包的平均时延。
- 第六列：最短时延。
- 第七列：最长时延。
- 第八列：标准偏差。

### tcpdump

根据使用者的定义对网络上的数据包进行截获的包分析的工具。 tcpdump　可以将网络中传送的数据报的“头”完全截获下来提供分析。支持针对网络层、协议、主机、网络或端口的过滤，并提供　and、or、not　等逻辑语句来帮助你去掉无用的信息。

> tcpdump [参数]

- -nn	直接以 IP 和端口号显示
- -i	监听网络界面
- -c	监听的封包数
- -w	将监听的封包数保存下来

```shell
sudo tcpdump -i eth0 -nn
```

### ifconfig

获取网络接口配置信息并对此进行修改。

> ifconfig [网络设备][参数]

- -a	显示全部接口信息
- -s	显示摘要信息
- up	启动指定网络设备/网卡
- down	关闭指定网络设备/网卡
- add	向指定网卡配置 IPv6 地址
- del	删除指定网卡的 IPv6 地址

```shell
(base) lv@lv-tp:docsify_blog$ ifconfig
docker0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255
        inet6 fe80::42:10ff:fec5:a8dd  prefixlen 64  scopeid 0x20<link>
        ether 02:42:10:c5:a8:dd  txqueuelen 0  (以太网)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 314  bytes 59091 (59.0 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

enp0s31f6: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        ether 54:e1:ad:e5:36:f7  txqueuelen 1000  (以太网)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
        device interrupt 16  memory 0xf2200000-f2220000  

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (本地环回)
        RX packets 48629  bytes 117368946 (117.3 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 48629  bytes 117368946 (117.3 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

```
一般情况下：

- eth0 表示第一块网卡， （硬件地址）HWaddr 表示当前网卡的物理地址。

- inet addr 表示网卡的 IPv4 地址，此网卡的 IP 地址是 192.168.42.4，广播地址： 0.0.0.0，掩码地址 Mask:255.255.255.0

- lo 表示主机的回环地址，一般用来测试一个网络程序，但又不想让局域网或外网的用户能够查看，只能在此台主机上运行和查看所用的网络接口。

- 第一行：连接类型：Ethernet（以太网）HWaddr（硬件 MAC 地址）

- 第二行：网卡的 IP 地址、子网、掩码

- 第三行：UP（代表网卡开启状态）、RUNNING（代表网卡的网线被接上）、MULTICAST、MTU:1500（最大传输单元）：1500字节 跃点数（Metric）：1

- 第四、五行：接收、发送数据包情况统计

- 第七行：接收、发送数据字节数统计信息