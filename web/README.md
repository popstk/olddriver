## web
## 构建速度优化
[使用DllPlugin优化第三方库](https://github.com/lin-xin/blog/issues/10)
每次加入第三方库应该执行npm run build:dll 重新生成static/js/vendor.dll.js

## 部署构建
``` bash
cnpm install             # 安装依赖
cnpm run-script build    # 生成文件
```

## 开发
``` bash
cnpm run dev
```
