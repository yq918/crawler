package main

import (
	"net/http"
	"fmt"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io/ioutil"
	"regexp"
)

const URL = "http://www.zhenai.com/zhenghun"

func main() {
	resp, err := http.Get(URL)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("http get code:%d",resp.StatusCode)
		return
	}
	//将获取的编码转为utf-8编码
	en := determineEncoding(resp.Body)
    reader := transform.NewReader(resp.Body,en.NewDecoder())

	bytes, e := ioutil.ReadAll(reader)
	if e != nil{
		panic(e)
	}
	//fmt.Printf("%s",bytes)

	processMatchString(bytes)
}


/**
自动 获取reader 编码
 */
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil{
		panic(e)
	}
	e2, _, _ := charset.DetermineEncoding(bytes, "")
	return e2
}

/**
正则处理HTML，获取链接地址与地区名称
 */
func processMatchString(contents []byte)  {
	 compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/\w+)[^>]*">([^<]+)</a>`)
	 all := compile.FindAllSubmatch(contents, -1)
	 for _,m := range all{
		 fmt.Printf("city:%s,url:%s\n",m[2],m[1])
	 }
	 fmt.Printf("match :%d",len(all))
}
