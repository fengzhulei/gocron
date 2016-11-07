package mydb

import (
    "testing"
    "fmt"
)
type Job struct {
    Id              int
    Name, Time, Cmd string
    STime, ETime    int
    Status          uint8
    Running         uint8
    Modify          uint8
    Process         uint8
    Ip              string
    Singleton       uint8
}

func (job Job) String() {
    return fmt.Sprintf("id:%d name:%s time:%s cmd:%s stime:%d etime:%d status:%d running:%d modify:%d process:%d ip:%s singleton:%d",
    job.Id,job.Name,job.Time,job.Cmd,job.STime,job.ETime,job.Status,job.Running,job.Modify,job.Process,job.Ip,job.Singleton)
}

func TestFetchRow(t *testing.T) {
    db, _ := Open("mysql", "wida:wida@tcp(127.0.0.1:3306)/mycron?charset=utf8")
    db.DB.SetMaxOpenConns(30)
    db.DB.SetMaxIdleConns(10)
    db.DB.Ping()

    defer  db.Close()
    d := Item{}
    err := db.Raw("SELECT * FROM cron where id=?",2).FetchRow(&d)
    if (err != nil ){
        t.Error(err.Error())
    }
    if  d["name"] != "test2"{
        t.Error("data no right")
    }
    var data Job
    err = db.Raw("SELECT * FROM cron where id=?",2).FetchRow(&data)
    if (err != nil ){
        t.Error(err.Error())
    }
    if  data.Name != "test2"{
        t.Error("data no right")
    }
}

func TestFetchRows(t * testing.T){
    db, _ := Open("mysql", "wida:wida@tcp(127.0.0.1:3306)/mycron?charset=utf8")
    db.DB.SetMaxOpenConns(30)
    db.DB.SetMaxIdleConns(10)
    db.DB.Ping()

    defer  db.Close()
    s := []Item{}
    i,err:= db.Raw("SELECT * FROM cron").FetchRows(&s)
    if err != nil{
        t.Error(err.Error())
    }
    if i != int64(len(s)){
        t.Error("data no right")
    }
    if s[0]["name"] != "test"{
        t.Error("data no right")
    }
    var sdata []Job
    _,err = db.Raw("SELECT * FROM cron").FetchRows(&sdata)
    if (err != nil ){
        t.Error(err.Error())
    }
    if i != int64(len(sdata)){
        t.Error("data no right")
    }
    if sdata[1].Name != "test2"{
        t.Error("data no right")
    }
}