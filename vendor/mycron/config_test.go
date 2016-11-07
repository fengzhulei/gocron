package mycron

import (
    "testing"
)

func TestConfig(t *testing.T)   {
    if Mysql_host == ""{
        t.Error("config read Mysql_host err")
    }
    if Mysql_prot == 0{
        t.Error("config read Mysql_prot err")
    }
    if Mysql_user == ""{
        t.Error("config read Mysql_user err")
    }
    if Mysql_pwd == ""{
        t.Error("config read Mysql_pwd err")
    }
    if Mysql_dbname == "" {
        t.Error("config read mysql_dbname err")
    }
    if Log_filename == ""{
        t.Error("config read Log_filename err")
    }
    if Log_Path == ""{
        t.Error("config read Log_Path err")
    }
}