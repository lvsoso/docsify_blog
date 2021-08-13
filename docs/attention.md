## 注意

替换 rm
```shell
rm() {                                                                                                                                            
    #echo $* |sed "s/ /\n/g" |xargs -I {} mv -f {} /home/trash
    tar c $* |(cd /home/trash && tar xf - ); rm -rf $*
}

```