#!/bin/bash

# 切换到 steam 用户并执行命令 steamcmd +login anonymous +app_update 2394010 validate +quit

su - steam << EOF

# 清空日志
echo "" > ./command.log

# 执行 steamcmd 命令并将输出重定向到文件
/usr/games/steamcmd +login anonymous +quit | tee -a ./command.log
EOF