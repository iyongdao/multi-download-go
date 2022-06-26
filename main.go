package main

import (
	"log"
	"time"


	dl "github.com/iyongdao/multi-download-go/download"
)

func main()  {
	url := "要下载文件的地址"
	//https://issuepcdn.baidupcs.com/issue/netdisk/MACguanjia/BaiduNetdisk_mac_4.10.1.dmg
	fileName := "文件名"
	start := time.Now()
	d := dl.NewDownloader(url, fileName, "",3)
	if err := d.Download(); err != nil {
		log.Fatal(err)
	}
	log.Printf("下载%v 耗时：%v s", fileName, time.Since(start).Seconds())
}