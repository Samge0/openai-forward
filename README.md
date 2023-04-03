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


### 有疑问请添加微信（备注: openai-forward），不定期通过解答
> 微信号 SamgeApp