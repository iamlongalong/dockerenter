#!/bin/bash

dlogs() {
    CONTAINER=$(docker ps --format '{{.Names}}' | fzf)  # 使用 fzf 选择容器
    if [ -n "$CONTAINER" ]; then
        docker logs -f "$CONTAINER"  # 如果选择了容器，查看日志
    else
        echo "没有选择容器"
    fi
}

dlogs