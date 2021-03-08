package main

import (
	"database/sql"
	"github.com/bippio/go-impala"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"fmt"
	"impala/utils"
)

var ini_parser = utils.IniParser{}

var (
	httpHost   string
	ImpalaInfo sync.Map
	err        error
	downDir    string
)

// 加载配置
func init() {
	if err := ini_parser.Load("config.ini"); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	downDir = filepath.ToSlash(rootDir) + "/down"
	fmt.Printf("downDir:%s\r\n", downDir)
	_, err = os.Stat(downDir)
	if err != nil {
		os.MkdirAll(downDir, os.ModePerm)
	}
	fmt.Println("test")
	return

	// 心跳检测
	go checkPing()

	//注册一个处理器函数
	http.HandleFunc("/query", sqlQuery)

	httpHost = ini_parser.GetString("global", "httpHost")
	log.Printf("httpHost:%s\r\n", httpHost)
	err = http.ListenAndServe(httpHost, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

/*心跳检测*/
func checkPing() {
	checkImpalaDb()
	t := time.Tick(60 * time.Minute)
	for now := range t {
		fmt.Printf("%v \n", now)
		checkImpalaDb()
	}
}

/*检测连接，断线重连*/
func checkImpalaDb22() {
	env := []string{"exp", "pro"}
	for _, val := range env {
		if _db, ok := ImpalaInfo.Load(val); ok {
			db, ok := _db.(*sql.DB)
			if ok {
				if db.Ping() != nil {
					db.Close()
					ImpalaInfo.Delete(val)
				}
			}
		}
		// 连接impala
		if _, ok := ImpalaInfo.Load(val); !ok {
			log.Printf(val+"Impala")
			tempConf := strings.Split(ini_parser.GetString("global", val+"Impala"), ":")
			log.Printf(val+"Conf:%s\r\n", tempConf)
			opts := impala.Options{BatchSize: 1024, BufferSize: 4096, Host: tempConf[0], Port: tempConf[1], LogOut: ioutil.Discard}
			ImpalaInfo.Store(val, sql.OpenDB(impala.NewConnector(&opts)))
		}
	}
}

func checkImpalaDb() {
	if _db, ok := ImpalaInfo.Load("exp"); ok {
		db, ok := _db.(*sql.DB)
		if ok {
			if db.Ping() != nil {
				db.Close()
				ImpalaInfo.Delete("exp")
			}
		}
	}

	if _db, ok := ImpalaInfo.Load("pro"); ok {
		db, ok := _db.(*sql.DB)
		if ok {
			if db.Ping() != nil {
				db.Close()
				ImpalaInfo.Delete("pro")
			}
		}
	}

	// 连接impala
	if _, ok := ImpalaInfo.Load("exp"); !ok {
		expConf := strings.Split(ini_parser.GetString("global", "expImpala"), ":")
		log.Printf("expConf:%s\r\n", expConf)
		opts := impala.Options{BatchSize: 1024, BufferSize: 4096, Host: expConf[0], Port: expConf[1], LogOut: ioutil.Discard}
		ImpalaInfo.Store("exp", sql.OpenDB(impala.NewConnector(&opts)))
	}

	if _, ok := ImpalaInfo.Load("pro"); !ok {
		proConf := strings.Split(ini_parser.GetString("global", "proImpala"), ":")
		log.Printf("proConf:%s\r\n", proConf)
		opts := impala.Options{BatchSize: 1024, BufferSize: 4096, Host: proConf[0], Port: proConf[1], LogOut: ioutil.Discard}
		ImpalaInfo.Store("pro", sql.OpenDB(impala.NewConnector(&opts)))
	}
}


func sqlQuery(w http.ResponseWriter, req *http.Request) {
	//resmap := make(map[int][]map[string]interface{})
	//fields := make(map[int][]string)
	//errs := make(map[int]string)

	// 读取报文中所有内容
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		io.WriteString(w, "读取body数据失败！")
		log.Fatal("读取body数据失败: ", err)
	}
	io.WriteString(w, string(bodyBytes))
	fmt.Println(string(bodyBytes))
}
