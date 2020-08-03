package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type IPAddr struct {
	start string
	end   string
	addr  string
}

var (
	ipMap   = make(map[int]IPAddr)
	ipIndex = make([]int, 0)
)

//查找指定ip段与地址映射文件中的ip地址
func main() {
	ip := "172.168.10.210"
	addr := query(ip)
	fmt.Println(addr)

}
func query(ip string) string {
	ipInt := ipToInt(ip)
	fmt.Println(ipInt)
	for i := len(ipIndex) - 1; i >= 0; i-- {
		if ipInt >= ipIndex[i] {
			addr, _ := ipMap[ipIndex[i]]
			return addr.addr
		}
	}
	return ""
}

func initIPAddress() error {
	f, err := os.Open("./testdata/ip.data")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		item := strings.Split(line, " ")

		start := item[0]
		ipint := ipToInt(start)
		ipMap[ipint] = IPAddr{
			start: start,
			end:   item[1],
			addr:  item[2],
		}
		ipIndex = append(ipIndex, ipint)
	}
	return nil
}

func ipToInt(ip string) int {
	ips := strings.Split(ip, ".")
	ipInt := 0
	var pos uint = 24
	for _, ipItem := range ips {
		ipint, _ := strconv.Atoi(ipItem)
		ipint = ipint << pos
		ipInt = ipInt | ipint
		pos -= 8
	}
	return ipInt
}
func init() {
	var err error
	err = initIPAddress()
	if err != nil {
		log.Fatalln(err)
	}
	sort.Ints(ipIndex)
	fmt.Println(ipIndex)
}
