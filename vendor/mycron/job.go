package mycron

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "time"
    "os/exec"
    "runtime"
    "os"
    "strings"
    "bytes"
    "mydb"
    "logger"
)

type Job struct {
	Id              int
	Name, Spec, Cmd string
	STime, ETime    string
	Status          int8
	Running         int8
	Modify          int8
	Ip              string
        Singleton       int8
}

type RunRet struct {
    Pid   int
    Out   string
    Err   error
}

var (
    db  mydb.MyDB
    err   error
    log * logger.GwsLogger
)

func init() {
    db, err = mydb.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
                    Mysql_user, Mysql_pwd, Mysql_host, Mysql_prot,Mysql_dbname))
	if err != nil {
		panic(err.Error())
	}
    db.DB.SetMaxOpenConns(30)
    db.DB.SetMaxIdleConns(10)
    //设置服务器可能关闭时间
    db.DB.SetConnMaxLifetime(1*time.Hour)
    db.DB.Ping()
    log = logger.GetDeaultLogger(Log_file)
}

func GetCronList() ( []Job,  error) {
    ut := time.Now().Format("2006-01-02 15:04:05")
    var jobs []Job
    _,err := db.Raw("SELECT * FROM cron where status = 1 and stime < ? and etime > ?", ut, ut).FetchRows(&jobs)
    if err != nil {
        panic(err.Error())
    }
    return jobs,nil
}

func GetModifyList()( []Job, error){
    defer func() {
        if err := recover(); err != nil {
            log.Error(err.(string));
        }
    }()
    ut := time.Now().Format("2006-01-02 15:04:05")
    var jobs []Job
    _,err := db.Raw("SELECT * FROM cron where stime < ? and etime > ? and modify = 1", ut, ut).FetchRows(&jobs)
    if err != nil {
        panic(err.Error())
    }
    return jobs,nil
}
func UpdateModifyList() (int64,error){
    ut :=  time.Now().Format("2006-01-02 15:04:05")
    return db.Raw("update cron set modify = 0 where stime < ? and etime > ? ", ut, ut).Exec()
}

func AtOnce()( []Job, error){
    defer func() {
        if err := recover(); err != nil {
            log.Error(err.(string));
        }
    }()
    var jobs []Job
    _,err := db.Raw("SELECT * FROM cron where once = 1").FetchRows(&jobs)
    if err != nil {
        panic(err.Error())
    }
    return jobs,nil
}

func UpdateAtOnceList() (int64,error){
    return db.Raw("update cron set once = 0").Exec()
}

func (job Job) ChangeRunningStatus(status int) (int64,error) {
    return db.Raw("update cron set running = ? where id = ?", status, job.Id).Exec()
}


func (job Job) JobStep(step int,str string,process_id int) (int64,error) {
    return db.Raw("insert into cron_hist set cid = ?,step = ?,process_id =? ,time = ?,ret=?",
                    job.Id, step,process_id,time.Now().Format("2006-01-02 15:04:05"),str).Insert()
}

func (job Job) Run(){
    job.ChangeRunningStatus(1)
    job.Exec()
    job.ChangeRunningStatus(0)
}

func (job Job) Exec()  {
    var cmd * exec.Cmd
    if runtime.GOOS == "windows"{
        cmd = exec.Command("cmd", "/C", job.Cmd)
    }else {
        shell := os.Getenv("SHELL")
        cmd = exec.Command(shell, "-c", job.Cmd)
    }

    var out bytes.Buffer
    cmd.Stdout = &out
    if err := cmd.Start(); err != nil {
        log.Error("jobid:%d|cmd:%s|errinfo:%s|pid:%d",job.Id,cmd.Path,err.Error(),cmd.Process.Pid)
        job.JobStep(3,err.Error(),cmd.Process.Pid)
        return
    }
    start := "start"
    job.JobStep(0,start,cmd.Process.Pid)
    done := make(chan error)
    go func() {
        done <- cmd.Wait()
    }()
    select {
    case  err :=<-done:
        if err !=nil{
            log.Error("jobid:%d|cmd:%s|errinfo:%s|pid:%d",job.Id,cmd.Path,err.Error(),cmd.Process.Pid)
            job.JobStep(3,err.Error(),cmd.Process.Pid)
            return
        }
    }
    job.JobStep(1,strings.TrimSpace(out.String()),cmd.Process.Pid)
}