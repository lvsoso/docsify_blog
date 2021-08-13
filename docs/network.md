

### ubuntu 修改 netplan 配置
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

### ubuntu route 

[https://www.cyberciti.biz/faq/linux-route-add/](https://www.cyberciti.biz/faq/linux-route-add/)

```shell
# show list
ip route list
route -n

# add gateway
ip route add 192.168.1.0/24 dev eth0
route add default gw 192.168.1.254 eth0


route add -host 10.0.0.21 dev tun0

# check new route
ping your-router-ip-here
ping your-ISPs-Gateway-ip-here
ping 192.168.1.254
ping www.cyberciti.biz

# persistence static routing
# vi /etc/network/interfaces

up route add -net 192.168.1.0 netmask 255.255.255.0 gw 192.168.1.254
down route del -net 192.168.1.0 netmask 255.255.255.0 gw 192.168.1.254

# vi /etc/rc.local
/sbin/ip route add 192.168.1.0/24 dev eth0
```
