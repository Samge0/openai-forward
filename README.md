# openai-forward
openai接口转发


### Docker运行
```docker
docker run -itd \
--name openai-forward \
-p 8080:8080 \
--restart=always \
--pull=always \
samge/openai-forward:latest
```

### 其他说明
上面命令在`国外服务器`（之所以用国外机子是因为openai对ip限制较多，你懂的）运行后，访问`ip:8080`端口看到相应json提示则说明转发成功，接下来说说`域名配置`及使用`caddy2`反代+自动https证书；
- 域名配置
  - 在[NameSilo](https://www.namesilo.com/)花1美元左右注册一个便宜的域名。目前`top域名`低至$1.88首年，折扣码填：`ONESAVING`立减1美元（折扣码随时时效，如果时效了可尝试自己google搜一下有没有最新折扣码），如果已有域名直接跳过；
  - 创建`ai`二级域名指向你服务器ip（这里二级域名自己根据喜好配置）；
- caddy2配置
  - 创建目录并填写配置
```shell
sudo mkdir -p ~/caddy2/config \
&& sudo mkdir -p ~/caddy2/data \
&& sudo vim ~/caddy2/Caddyfile
```
  - Caddyfile的配置如下：
```caddyfile
(ai) {
    reverse_proxy 你服务器ip:8080 {
      flush_interval -1
	  header_up -Origin
	  header_up -X-Real-IP
    }
}
ai.你的域名 {
    import ai
}
```
  - 运行Caddy2，-p端口可不指向，根据实际需求配置
```shell
docker run -d \
-v ~/caddy2/Caddyfile:/etc/caddy/Caddyfile \
-v ~/caddy2/config:/config \
-v ~/caddy2/data:/data \
-p 80:80 \
-p 443:443 \
--name caddy2 \
--pull=always \
--restart always \
-e LANG=C.UTF-8 \
samge/caddy2:v2
```
  - 使用`ai.你的域名`替换`api.openai.com`进行请求测试；


### 有疑问请添加微信（备注: openai-forward），不定期通过解答
> 微信号 SamgeApp