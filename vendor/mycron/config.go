package mycron

import (
    "conf"
)

const (
    ConfigPath = "gocron.conf"
)

var(
    Mysql_host,Mysql_user,Mysql_pwd,Mysql_dbname,Log_file  string
    Mysql_prot int
)

func init() {
    c := conf.NewConfig(ConfigPath)
    //mysql
    Mysql_host =c.GetString("mysql", "host")
    Mysql_prot = c.GetInt("mysql", "port")
    Mysql_user = c.GetString("mysql", "user")
    Mysql_pwd = c.GetString("mysql", "pwd")
    Mysql_dbname =c.GetString("mysql", "dbname")

    //logger
    Log_file = c.GetString("log","LogPath")
}

