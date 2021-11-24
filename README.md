## bing每日一图

### 搭配Plash使用 https://github.com/sindresorhus/Plash

```
go mod tidy
go run main.go
curl http://127.0.0.1:1000/bing
# 重定向至最新每日一图

# or with docker
docker run -itd -p 1000:1000 --name=bing_pic brooke9537/bing_pic
```

### 将 http://127.0.0.1:1000/bing 填至Plash website
