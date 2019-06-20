package lib

import (
	"bufio"
	"github.com/emirpasic/gods/lists/arraylist"
	"iads/lib/base"
	"log"
	"os"
	"regexp"
)

type ErrMsgStruct struct {
	Id       int
	ErrMsg   string
	ErrSolve string
}

var (
	fp     *os.File
	err    error
	logger *log.Logger

	logFiles arraylist.List //日志文件路径数组
	errMsgs  arraylist.List //错误信息数组
)

func InitEnv() {
	//添加日志文件路径
	logFiles.Clear()
	errMsgs.Clear()
	logFiles.Add("/var/log/messages", "/var/log/mcelog", "/var/log/kerl", "/var/log/syslog")
	errMsgs.Add("error", "failed")
}

func StartTest() {
	logger.Println("Start testing......")
	fp, err = os.OpenFile("testlog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("create logfile failed.")
	}
	logger = log.New(fp, "", log.LstdFlags|log.Llongfile)
}

func StopTest() {
	logger.Println("test stopped.")
	_ = fp.Sync()
	fp.Close()
}

/********************* logfile **********************/
type RowLog struct {
	Index int64
	Data  string
}

func NewRowLog() *RowLog {
	return &RowLog{
		Index: -1,
		Data:  "",
	}
}

func AnalysisLogFile(filename string) (arraylist.List, error) {
	it := errMsgs.Iterator()
	regStr := "(?i)"
	for it.Next() {
		if regStr == "(?i)" {
			regStr += it.Value().(string)
		} else {
			regStr = regStr + "|" + it.Value().(string)
		}
	}
	arr := arraylist.List{}
	file, err := os.Open(filename)
	if err != nil {
		return arr, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var index int64 = 1
	for scanner.Scan() {
		nowRowStr := scanner.Text()
		tmp, err := regexp.Match(regStr, []byte(nowRowStr))
		if err != nil {
			logger.Println("regex log file error.")
		}
		if true == tmp {
			tmpRow := NewRowLog()
			tmpRow.Index = index
			tmpRow.Data = nowRowStr
			arr.Add(tmpRow)
		}
		index++
	}
	return arr, err
}

/*****************************************************************/

func DownloadFromServer(remotePath string, localPath string) error {
	ssh := base.NewSsh("192.168.1.111", "root", "000000")
	err := ssh.SftpConnect()
	f, _ := os.Stat(remotePath)
	if f.IsDir() {
		err = ssh.DownloadDir(remotePath, localPath)
	} else {
		err = ssh.DownloadFile(remotePath, localPath)
	}
	return err
}
