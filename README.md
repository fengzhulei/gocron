
#golang  任务调度系统

##主要特性
- 从mysql读取cron配置,系统坚持job添加和修改后自主更新，开始任务和结束任务完全配置话

- 采用crontab表达式支持到秒级

- 任务运行状态全部透明化

- 支持立即执行调试job

- 支持一个job同时多个进程跑，和支持各个进程状态监控

- 支持*nux 和 windows  脚本入口命令分别是 shell -c  和  cmd /c

##Tables

###任务配置

    CREATE TABLE `cron` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `uid` int(11) NOT NULL COMMENT '用户id',
      `name` varchar(50) NOT NULL DEFAULT '',
      `time` varchar(50) NOT NULL,
      `cmd` varchar(255) NOT NULL,
      `sTime` int(11) NOT NULL,
      `eTime` int(11) NOT NULL,
      `status` tinyint(1) NOT NULL DEFAULT '0',
      `running` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否在运行中 0不是,1是',
      `modify` tinyint(1) NOT NULL DEFAULT '0',
      `process` tinyint(2) NOT NULL DEFAULT '1' COMMENT '进程数量',
      `ip` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
      `singleton` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否单例执行0非单例，1单例',
      `atonce` tinyint(1) NOT NULL DEFAULT '0' COMMENT '马上执行',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8


###任务执行记录

    CREATE TABLE `cron_hist` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `cid` int(11) NOT NULL,
      `process_id` int(11) NOT NULL DEFAULT '0' COMMENT 'shell进程id',
      `branch` int(11) NOT NULL DEFAULT '0' COMMENT '执行分支',
      `step` tinyint(3) NOT NULL,
      `time` datetime NOT NULL,
      `ret` varchar(255) DEFAULT NULL,
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=60790 DEFAULT CHARSET=utf8;

##示例for windows:

###测试配置
    INSERT INTO `cron` VALUES (1, 'test', '*/1 * * * * ?', 'php e:\\php.php', 1427337701, 1437337701, 1, 0, 0, 1, '', 0);
    INSERT INTO `cron` VALUES (2, 'test2', '*/2 * * * * ?', 'php e:\\test.php', 1427337701, 1447337701, 1, 0, 0, 1, '', 0);
    
###文件 e:\php.php
    <?php
        echo "每秒执行";	 
    ?>

###文件 e:\test.php
    <?php
        echo "每两秒执行一次";
    ?>
        
##示例 for linux ：  
###测试配置
    INSERT INTO `cron` VALUES (1, 'test', '*/1 * * * * ?', '/home/wida/sh.sh', 1427337701, 1437337701, 1, 0, 0, 1, '', 0);
    INSERT INTO `cron` VALUES (2, 'test2', '*/2 * * * * ?', '/home/wida/sh2.sh', 1427337701, 1447337701, 1, 0, 0, 1, '', 0);

###文件 /home/wida/sh.sh

    #!/bin/sh

    php /home/wida/php.php

###文件 /home/wida/sh2.sh

    #!/bin/sh

    php /home/wida/php.php

###文件 /home/wida/php.php

    <?php
        echo "test";
    ?>