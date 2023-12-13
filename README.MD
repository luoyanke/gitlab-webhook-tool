
# 访问路径
- http://ip:6710/web-hook


# 编译
制品包在 ./dist/linux/

window下交叉编译
```bash
$ENV:CGO_ENABLED=0
$ENV:GOOS="linux"
$ENV:GOARCH="amd64"
go build -o ./dist/linux/gitlab-webhook-tool ./cmd/main.go
```


# 启动参数
```bash
    ./gitlab-webhook-tool -port 6711 -feishuWebhook https://open.feishu.cn/open-apis/bot/v2/hook/11111111111111111
```
- port  监听的端口
-feishuWebhook 飞书机器人
