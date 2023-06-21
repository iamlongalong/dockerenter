#!/bin/bash

# 获取所有正在运行的容器名称和ID
containers=$(docker ps --format "{{.Names}} ({{.ID}})")

# 判断是否有正在运行的容器
if [ -z "$containers" ]; then
  echo "没有正在运行的容器"
  exit 1
fi

# 使用fzf工具提供上下选择能力
container=$(echo "$containers" | fzf)

# 获取容器ID
container_id=$(echo "$container" | awk '{print $1}')

# 进入容器
docker exec -it "$container_id" /bin/bash
