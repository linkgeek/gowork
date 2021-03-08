package main

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"impala2/utils"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
	"github.com/bippio/go-impala"
	"github.com/tidwall/gjson"
)

// 加载配置
var conf = utils.IniParser{}


// 数组栈，后进先出
/*
type ImpalaInfo struct {
	impalaDb map[string]*sql.DB   // 底层切片
	lock  sync.Mutex // 为了并发安全使用的锁

}*/

var ImpalaInfo sync.Map

//var impalaInfo  *ImpalaInfo;

var httpHost string

var downDir string;

//返回结构体
type QInfo struct {
	Data  []map[string]interface{}
	Cols  []string
	Err error
	Index int
}

// 初始化配置相关
func init() {
	conf.Load("run.ini")
	//impalaInfo=&ImpalaInfo{impalaDb:make(map[string]*sql.DB)};
}
func main() {
	rootDir,err:= filepath.Abs(filepath.Dir(os.Args[0]))
	downDir=filepath.ToSlash(rootDir)+"/down"
	fmt.Printf("downDir:%s\r\n",downDir)
	_, err = os.Stat(downDir)
	if err != nil {
		os.MkdirAll(downDir,os.ModePerm)
	}
	//启动心跳检测
	go checkPing()
	//初始化http
	httpHost = conf.GetString("global", "httpHost")
	log.Printf("httpHost:%s\r\n", httpHost)
	http.HandleFunc("/query", Find)
	http.HandleFunc("/", api)
	http.Handle("/down/", http.StripPrefix("/down/", http.FileServer(http.Dir(downDir))))
	err = http.ListenAndServe(httpHost, nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

/*
检查连接
如果断线重连
*/
func checkImpalaDb(){

	if _db, ok := ImpalaInfo.Load("exp"); ok {
		db, ok := _db.(*sql.DB)
		if ok {
			//有错误直接置为nil
			if(db.Ping()!=nil){
				db.Close();
				ImpalaInfo.Delete("exp")
			}
		}

	}
	if _db, ok := ImpalaInfo.Load("pro"); ok {
		db, ok := _db.(*sql.DB)
		if ok {
			//有错误直接置为nil
			if (db.Ping() != nil) {
				db.Close();
				ImpalaInfo.Delete("pro")
			}
		}
	}

	if _, ok := ImpalaInfo.Load("exp"); !ok {
		//连接impala
		expConf := strings.Split(conf.GetString("global", "expImpala"), ":")
		log.Printf("expConf:%s\r\n", expConf)
		opts :=impala.Options{BatchSize:1024,BufferSize:4096,Host:expConf[0],Port:expConf[1],LogOut: ioutil.Discard}
		ImpalaInfo.Store("exp",sql.OpenDB(impala.NewConnector(&opts)))
	}

	if _, ok := ImpalaInfo.Load("pro"); !ok {
		//连接impala
		proConf := strings.Split(conf.GetString("global", "proImpala"), ":")
		log.Printf("proConf:%s\r\n", proConf)
		opts :=impala.Options{BatchSize:1024,BufferSize:4096,Host:proConf[0],Port:proConf[1],LogOut: ioutil.Discard}
		ImpalaInfo.Store("pro",sql.OpenDB(impala.NewConnector(&opts)))
	}
}




/*检测心跳*/
func checkPing(){
	//首先执行一次避免。。
	checkImpalaDb();
	c := time.Tick(60 * time.Minute)
	for now := range c {
		fmt.Printf("%v \n", now)
		checkImpalaDb();
	}
}


/*查询*/
func queryData(_sql string,env string)([]map[string]interface{}, [] string,error){
	_impalaDb, ok := ImpalaInfo.Load(env);
	if !ok {
		log.Println("no found env:"+env)
		return nil,nil,errors.New("no found env:"+env)
	}
	impalaDb, ok := _impalaDb.(*sql.DB)
	if !ok {
		log.Println("db type err:"+env)
		return nil,nil,errors.New("db type err:"+env)
	}


	//判断数据格式
	if strings.HasPrefix(strings.ToLower(strings.TrimSpace(_sql)), "drop") || strings.HasPrefix(strings.ToLower(strings.TrimSpace(_sql)), "delete") {
		return nil,nil,errors.New("sql delete drop")
	}


	rows, err := impalaDb.QueryContext(context.Background(), _sql)
	if err != nil {
		log.Println(err)
		return nil,nil,err
	}
	cols, _ := rows.Columns()
	values := make([]interface{}, len(cols))
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}
	//最后得到的map
	var results  []map[string]interface{}
	i := 0
	for rows.Next() { //循环，让游标往下推
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			return nil,nil,err
		}
		row := make(map[string]interface{}) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = v
		}
		results= append(results,row) //装入结果集中
		i++
	}
	//mjson, _ := json.Marshal(results)
	return results,cols,nil
}
/*处理http*/
func Find(w http.ResponseWriter, req *http.Request) {
	resmap := make(map[int][]map[string]interface{})
	fields := make(map[int][]string)
	errs := make(map[int]string)
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		io.WriteString(w, "读取body数据失败！")
		log.Fatal("read body error :", err)
	}
	sqlJson := gjson.ParseBytes(bodyBytes).Get("sql")
	v1 := gjson.ParseBytes(bodyBytes).Get("v1")
	env := gjson.ParseBytes(bodyBytes).Get("env")
	if !sqlJson.Exists() {
		io.WriteString(w, "param err!\r\n")
		return
	}
	var queryNum = 0       //计数器
	ch := make(chan QInfo) //通道获取数据
	result := gjson.Parse(sqlJson.String())
	for k, value := range result.Array() {
		queryNum++
		//log.Println("value.String():" + value.String())
		go impalaExec(value.String(),env.String(), k, ch)
	}

	// 遍历接收通道数据
	for {
		data, ok := <-ch
		if ok {
			resmap[data.Index] = data.Data
			fields[data.Index]=data.Cols;
			if(data.Err!=nil){
				errs[data.Index]=data.Err.Error();
			}else{
				errs[data.Index]="";
			}
			queryNum--
		}
		if queryNum <= 0 {
			break
		}
	}

	defer close(ch)
	var mjson []byte;
	//新版返回更多内容
	if(v1.Num==1) {
		back := make(map[string]interface{})
		back["data"]=resmap
		back["field"]=fields;
		back["errs"]=errs;
		mjson, _ = json.Marshal(back)
	}else{
		mjson, _ = json.Marshal(resmap)
	}
	//返回结果
	io.WriteString(w, string(mjson))
	return
}
/*并发执行*/
func impalaExec(sql string,env string, index int, ch chan QInfo) bool {
	data,cols,err:=queryData(sql,env);
	if(err!=nil){
		return queryExit(index,nil,nil, err,ch)
	}
	return queryExit(index,data,cols,nil, ch)
}

func queryExit(index int, data []map[string]interface{}, cols  []string,err error,ch chan QInfo) bool {
	ch <- QInfo{Index: index, Data: data,Cols:cols,Err:err}
	return true
}


/*处理http*/
func api(w http.ResponseWriter, req *http.Request) {
	//跨域处理
	cors(w);
	var err error;
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		io.WriteString(w, "读取body数据失败！")
		log.Fatal("read body error :", err)
	}
	var env string = "exp";
	_env := gjson.ParseBytes(bodyBytes).Get("env")
	if (_env.String() != "") {
		env = _env.String();
	}

	var data []map[string]interface{};
	var csvUrl string;
	var field []string;
	fmt.Printf("path:%s\r\n", req.URL.Path)

	switch req.URL.Path {
	case "/databases":
		databases, _, err := queryData("show DATABASES", env);
		if (err == nil) {
			for _, item := range databases {
				//乱七八糟的数据库过滤
				if(len(item["name"].(string))>64){
					continue;
				}
				data1, _, _ := queryData("show tables IN "+item["name"].(string), env);
				item["tables"] = data1;
				//没有表的数据库过滤
				if(len(data1)>0) {
					data = append(data, item)
				}
			}
		}
	case "/tableField":
		//data,_,err=queryData("describe formatted ads.a_coin_consume_stat","exp");
		dbName := gjson.ParseBytes(bodyBytes).Get("dbName")
		table := gjson.ParseBytes(bodyBytes).Get("table")
		data, _, err = queryData("DESCRIBE "+dbName.String()+"."+table.String(), env);
		break;


	case "/exData":
		//data,_,err=queryData("describe formatted ads.a_coin_consume_stat","exp");
		dbName := gjson.ParseBytes(bodyBytes).Get("dbName")
		table := gjson.ParseBytes(bodyBytes).Get("table")
		_type := gjson.ParseBytes(bodyBytes).Get("type").Int()
		feleds:=gjson.ParseBytes(bodyBytes).Get("feleds").Array()
		wheres:=gjson.ParseBytes(bodyBytes).Get("wheres").Array();

		//处理字段
		var feled string;
		if(len(feleds)>0) {
			for  _,value := range feleds {
				if (len(feled)>0) {
					feled = feled +","+ value.String();
				}else {
					feled = value.String();
				}
			}
		}else{
			feled="*"
		}

		//处理查询条件
		var where string;
		if(len(wheres)>0) {
			for  _,_item := range wheres {
				item:=_item.Map()
				_where:=item["left"].String()+item["opera"].String()+"'"+item["right"].String()+"'";
				if (len(where)>0) {
					where = where +" AND "+_where
				}else {
					where ="WHERE "+ _where
				}
			}
		}else{
			where=""
		}
		var limit string=" LIMIT 50";
		//导出不需要limit
		if(_type==1){
			limit=""
		}
		var sql string;
		sql="SELECT "+feled+" FROM "+dbName.String()+"."+table.String()+" "+where+limit;
		fmt.Printf("sql:%s\r\n",sql)
		//导出
		if(_type==1){
			scheme := "http://"
			if req.TLS != nil {
				scheme = "https://"
			}
			fileName:=dbName.String()+"."+table.String()+"_"+time.Now().Format("2006-01-02-15-04-05")+fmt.Sprintf("%x", md5.Sum([]byte(RandString(20))))+".csv"
			csvUrl=scheme+req.Host+"/down/"+fileName;
			err = queryWriteCsv(sql,downDir+"/"+fileName, env);
		}else {
			data, _, err = queryData(sql, env);
		}
		fmt.Printf("error:%v\r\n",err)
		break;
	}

	var mjson []byte;
	back := make(map[string]interface{})
	back["data"]=data
	back["csvUrl"]=csvUrl
	back["field"]=field;
	back["errs"]=err;
	mjson, _ = json.Marshal(back)
	//返回结果
	io.WriteString(w, string(mjson))
	return
}

/*查询写入csv*/
func queryWriteCsv(_sql string,fileName string,env string)(error){
	_impalaDb, ok := ImpalaInfo.Load(env);
	if !ok {
		log.Println("no found env:"+env)
		return errors.New("no found env:"+env)
	}
	impalaDb, ok := _impalaDb.(*sql.DB)
	if !ok {
		log.Println("db type err:"+env)
		return errors.New("db type err:"+env)
	}


	//判断数据格式
	if strings.HasPrefix(strings.ToLower(strings.TrimSpace(_sql)), "drop") || strings.HasPrefix(strings.ToLower(strings.TrimSpace(_sql)), "delete") {
		return errors.New("sql delete drop")
	}

	rows, err := impalaDb.QueryContext(context.Background(), _sql)
	if err != nil {
		log.Println(err)
		return err
	}
	if err != nil {
		log.Println(err)
		return err
	}
	cols, _ := rows.Columns()
	values := make([]interface{}, len(cols))
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}
	//生成csv文件
	fout,err := os.Create(fileName)
	defer fout.Close()
	if err != nil {
		fmt.Println(fileName,err)
		return err
	}
	w := csv.NewWriter(fout)
	w.Write(cols)
	w.Flush()
	for rows.Next() { //循环，让游标往下推
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			return err
		}
		row := make(map[string]interface{}) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = v
		}
		var csvLine []string;
		for _, v := range cols {     //每行数据是放在values里面，现在把它挪到row里
			csvLine=append(csvLine,Strval(row[v]))
		}
		w.Write(csvLine)
		w.Flush()
	}
	return nil;
}


/*跨域处理*/
func cors(w http.ResponseWriter){
	w.Header ().Set ("Access-Control-Allow-Origin", "*")       // 允许访问源
	w.Header ().Set ("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    // 允许 post 访问
	w.Header ().Set ("Access-Control-Allow-Headers", "Content-Type,Authorization") //header 的类型
	w.Header ().Set ("Access-Control-Max-Age", "1728000")
	w.Header ().Set ("Access-Control-Allow-Credentials", "true")
	w.Header ().Set ("content-type", "application/json") // 返回数据格式是 json
}


func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}


func RandString(len int) string {
	r:=rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}