
#golang  任务调度系统

##主要特性
- 从mysql读取cron配置,系统坚持job添加和修改后自主更新，开始任务和结束任务完全配置话

- 采用crontab表达式支持到秒级

- 任务运行状态全部透明化

- 支持立即执行调试job

- 支持*nux 和 windows  脚本入口命令分别是 shell -c  和  cmd /c


#使用说明

- 先安装[gocronadmin](https://github.com/widaT/gocronadmin "gocronadmin") 

- linux 下使用 app.sh 脚本   ./app.sh [start|stop]

- crontab 算法原理参照 [crontab表达式和算法](http://blog.csdn.net/u014798316/article/details/46460697 "crontab表达式和算法") 