package logger

//默认的文件log
func GetDeaultLogger(logFile string)  * GwsLogger{
	log := NewLogger(1000)
	log.SetLogger("file",`{"filename":"`+logFile+`"}`)
	return log
}