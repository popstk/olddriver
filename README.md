# 前端
* vue
* vue-router
* axios
* webpack

# 后端
* 存储系统 - mongodb
* 爬虫框架 - scrapy
* web服务器 - flask

# 后端项目结构
* backend - restful web服务器
* crawler- scrapy爬虫
* scrapyd - 爬虫调度服务器
* tool - 工具集

# 部署
0. 安装python依赖
> pipenv install
1. 使用docker启动数据存储
> docker-compose up -d
2. 启动爬虫调度服务器和gui
> cd scapyd
>  
> nohup scrapyd & 
>
> nohup scrapydweb &

3. 部署爬虫到服务器
> cd crawler &&  scrapyd-deploy
4. 启动web服务器
> cd backend && nohup python app.py &

# 工具
* mongockient.sh - 存储容器mongodb的客户端
* keeprunning.py - 每天定时自动触发爬虫采集

# 备忘
* mongodb清空爬虫数据
> use spider  
> db.xx.drop  

# 兼容
* python 2.7
