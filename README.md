## 转发动态原理是爬取某个用户的页面，获取数据跟着转发，前提条件是这个人只转发抽奖，并且重复度不高
## 活动抽奖原理是抓取up主转发抽奖娘更新的最新抽奖列表

#使用方法
## 更改config/bilibili.go下的配置参数：SESSDATA，BUVID3，CSRF，BILIBILI_UID，ANIO，MY，更改config/db.go下的数据库


# 命令
### 抽奖只能一天获取一次，并且根据获取到的抽奖次数进行抽奖，一天运行一次即可
`go run main.go -get  获取活动抽奖页面数据` 
`go run main.go -doupdate  更新抽奖机会` 
`go run main.go -dorun  开始抽奖` 

### 转发
`go run main.go -zhuanfa` 开始获取数据，转发，评论，关注，点赞
`go run main.go -del` 删除动态，该功能默认删除第二页数据第一条，原因是觉得有些抽奖工具是根据动态数量判断的，所以定期删掉一些动态 
