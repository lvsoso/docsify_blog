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

