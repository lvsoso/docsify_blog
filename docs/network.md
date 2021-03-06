

ubuntu 修改 netplan 配置
```shell
# This file is generated from information provided by
# the datasource.  Changes to it will not persist across an instance.
# To disable cloud-init's network configuration capabilities, write a file
# /etc/cloud/cloud.cfg.d/99-disable-network-config.cfg with the following:
# network: {config: disabled}
network:
    ethernets:
        ens6:
                addresses: [192.168.1.53/24]
                gateway4: 192.168.1.104
                nameservers:
                         addresses: [8.8.8.8, 114.114.114.114]
                dhcp4: no
                dhcp6: no
                #dhcp4: yes
    version: 2

```