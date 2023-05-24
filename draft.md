接口

```bash
dk enter/exec  # 执行 docker exec
dk logs  # 执行 docker logs
dk stop  # 执行 docker stop
dk rm # 执行 docker rm

dk use  # 选择一台远端机器

# 其他命令则完全代理给 docker 命令
```

配置

```yaml
targets:
    - name: longser1
      user: root
      pass: xx
      identity: xx
      proxy:
        - name: xxx

```
配置这部分，可以重写一下 sshw，把相关能力扩展了
