## dockerenter 

有时候使用 docker 时，发现要执行 exec 很不方便，得先 docker ps，拿到 containerID 之后得 copy paste 等等，于是做了这个工具。

这是一个简单的 docker 工具，用来替代  `docker exec -it xxxx bash`  这个命令。

🎉🎉🎉 有了这个工具，节省下来的时间又可以摸 5 秒钟 🐟 啦啦啦~  🎉🎉🎉

灵感来源于 devspace enter 命令，之前写的 sshw scp 扩展也是这个处理方法，还是挺方便的。

![demo](demo.gif)


## 如何安装

1. 到 [release](https://github.com/iamlongalong/dockerenter/releases) 页面，下载符合自己版本的文件

2. 将二进制文件移动到 bin 下
```bash
sudo mv denter_mac_vx.x /usr/local/bin/
```

3. 输入 `denter` (当然了，你得打开 docker 先~)


## todo

- [ ] 扩展为 docker 子命令的方式

## 其他

有任何建议，欢迎提哈 ~ 😁
