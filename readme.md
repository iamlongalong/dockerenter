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

其实不仅 docker exec 需要知道容器的名字，还有一些其他场景同样需要 容器名 或者 镜像名 或者 网络名 等等，按照这个工具所提供价值的基本原因 (把 填空题 变为 选择题)，很多场景都可以使用这种方式进行优化。

docker 中的例子: `docker cp` `docker export` `docker inspect` `docker kill` `docker logs` `docker rm` `docker rmi` `docker save` `docker export` `docker tag` `docker attach` `docker network` `docker run` `docker logs`
helm 中的例子: `fetch` `install`
kubectl 中的例子: `kubectl cp` `kubectl exec` `kubectl logs` `kubectl describe` ……

从这种交互方式更近一步，就是类似于 k9s、glances、top 这种直接以 terminal 为 UI 界面的终端交互方式。

之后可以考虑下更近一步

## 其他

有任何建议，欢迎提哈 ~ 😁
