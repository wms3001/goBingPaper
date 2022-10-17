# goBingPaper
## 简介
获取bing最新更新的壁纸
## 使用
```
go get github.com/wms3001/goBingPaper
```
## 实例
1. 测试token
```go
bingPaper := BingPaper{}
res := bingPaper.GetPaper()
log.Println(res)
```
