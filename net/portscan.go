//GO语言  实现端口扫描
//缺陷
//port  无法设置成全局变量不知道怎么设置的
//var l = list.New()   这个是数组操作并不是消息队列    跟消息队列功能类似
//实现功能
//实现生成 IP段
//实现端口扫描
//实现参数传入
//写入文件到本地
//main.go 58.215.20.30 58.215.201.30 80
//文件名 开始IP 结束IP 扫描端口
//QQ29295842   希望认识更多的朋友技术交流
//QQ群367196336   go   golang WEB安全开发
package main

import (
	"container/list"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func ip2num(ip string) int {
	canSplit := func(c rune) bool { return c == '.' }
	lisit := strings.FieldsFunc(ip, canSplit) //[58 215 20 30]
	//fmt.Println(lisit)
	ip1StrInt, _ := strconv.Atoi(lisit[0])
	ip2StrInt, _ := strconv.Atoi(lisit[1])
	ip3StrInt, _ := strconv.Atoi(lisit[2])
	ip4StrInt, _ := strconv.Atoi(lisit[3])
	return ip1StrInt<<24 | ip2StrInt<<16 | ip3StrInt<<8 | ip4StrInt
}
func num2ip(num int) string {
	ip1Int := (num & 0xff000000) >> 24
	ip2Int := (num & 0x00ff0000) >> 16
	ip3Int := (num & 0x0000ff00) >> 8
	ip4Int := num & 0x000000ff
	//fmt.Println(ip1Int)
	data := fmt.Sprintf("%d.%d.%d.%d", ip1Int, ip2Int, ip3Int, ip4Int)
	return data
}
func gen_ip(Aip1 int, Aip2 int) {
	index := Aip1
	for index < Aip2 {
		//fmt.Println(num2ip(index))
		// 入队, 压栈
		ip_data := num2ip(index)
		//fmt.Println(ip_data)
		l.PushBack(ip_data)
		index++
	}
}
func text_add(name string, data string) { //向文件中写入数据   text_add("file2.txt", "qqqqqqqqqqqqqqqqqqqqqqq")
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error open file")
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(data)
	_, err = f.WriteString("\r\n")
	if err != nil {
		panic(err)
	}
}

//text_add("file2.txt", "qqqqqqqqqqqqqqqqqqqqqqq")
var l = list.New()

func socketIp(host string, port string) bool {
	var (
		remote = host + ":" + port
	)
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", remote) //转换IP格式
	//fmt.Printf("%s", tcpAddr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr) //查看是否连接成功
	if err != nil {
		fmt.Printf("no==%s:%s\r\n", host, port)
		return false
	}
	defer conn.Close()
	fmt.Printf("ok==%s:%s\r\n", host, port)
	return true
}
func forIp(port string, ch chan int) {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	filePath := "temp/"
	file_name := fmt.Sprintf("%d-%d-%d_%s", year, mon, day, port)
	for { //死循环
		if l.Len() <= 0 {
			fmt.Println("jump")
			break //#跳出
		}
		// 出队  从前读取
		i1 := l.Front()
		l.Remove(i1)
		IP, _ := i1.Value.(string)
		if socketIp(IP, port) {
			//OK
			//获取当前  日期作为文件名  在把IP写入进去
			text_add(filePath+file_name+"_ok.txt", IP)
		} //else {
		//  text_add(file_name+"_no.txt", IP)
		// }
		time.Sleep(time.Millisecond * 500) //纳秒为单位
	}
	ch <- 0
}
func main() {
	argsLen := len(os.Args)
	//fmt.Println(argsLen)
	if argsLen != 4 {
		fmt.Println("portscan 58.215.20.30 58.215.201.30 80")
	} else {
		gen_ip(ip2num(os.Args[1]), ip2num(os.Args[2]))
		ch := make(chan int)
		go forIp(os.Args[3], ch)
		<-ch
	}
}
