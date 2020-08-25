package basicinfo
//Todo
//2.ip
//3.socks proxy
import (
	"bufio"
	"fmt"
	"os"
)
type respinfo struct {// 信息结构体
	Statuscode int // 状态码
	Server string // Server
	Title string // Ttile
	//Proto string // 协议
	Location string //跳转地址
	url string //url
	ip []string  //ip
	port int //端口
}
type saveres struct {
	infos []respinfo //每个url的信息结构体
	len int //访问的url数量
}
type domaininfo struct {
	url string
	port []int
	ip []string
	ipflag int
 }
func Main(argv map[string]string)  {
	if _,ok := argv["filepath"];ok {
		filepath := argv["filepath"]

		domain,flag := getinfo(filepath)
		if flag==0 {
			parseresult := domainHandle(domain)
		}
		else {
			parseresult := ipHandle(domain)
		}
		result := sendrequest(parseresult)
		handleresult(result)
	}
	if _,ok := argv["iprange"];ok{
		iprange := argv["iprange"]
		println(iprange)
	}
}


//从文件中读取域名或者ip
//ip 1
//domain 0
func getinfo (filedir string) ([]string, int){//
	var urls  []string//创建str类型的切片
	file, _ := os.Open(filedir)
	if file == nil{
		fmt.Println("Open domain file fail")
		os.Exit(0)
	}
	defer file.Close()//在函数返回时执行
	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanWords)
	for filescanner.Scan(){
		urls = append(urls, filescanner.Text())
	}
	return urls,0//返回url切片
}

//func process (now int,total int){
//	pgb := pgbar.New("开始扫描")
//	b := pgb.NewBar("1st", 20000)
//}//进度条