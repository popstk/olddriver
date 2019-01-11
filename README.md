## 前端
- vue
- vue-router
- axios
- webpack

## 后端
- 存储系统 - mongodb
- 爬虫框架 - scrapy
- web服务器 - flask

## 后端项目结构
- backend - restful web服务器
- crawler- scrapy爬虫
- scrapyd - 爬虫调度服务器
- tool - 工具集

## 部署
#### 依赖环境
``` bash
docker-compose up -d    # 数据存储
pipenv install   # 安装依赖
pipenv shell    # 进入环境
```

#### 启动查询服务
```
nohup backend/python app.py &
```

#### 启动爬虫服务
```bash
cd scapyd
nohup scrapyd &   # 启动爬虫调度服务器 
nohup scrapydweb & # 启动爬虫调度gui
```

#### 部署爬虫
```bash
cd crawler
scrapyd-deploy
```

## 工具
* mongockient.sh - 存储容器mongodb的客户端
* keeprunning.py - 每天定时自动触发爬虫采集

##备忘
* mongodb清空爬虫数据
> use spider  
> db.xx.drop  

## 兼容
* python 2.7
