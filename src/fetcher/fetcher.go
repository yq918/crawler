package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
)

func Fetch(URL string) ([]byte,error) {
	resp, err := http.Get(URL)
	if err != nil{
		 return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("http get code:%d",resp.StatusCode)
	}
	//将获取的编码转为utf-8编码
	en := determineEncoding(resp.Body)
	reader := transform.NewReader(resp.Body,en.NewDecoder())

    return ioutil.ReadAll(reader)
}


/**
自动 获取reader 编码
 */
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil{
	    return unicode.UTF8
	}
	e2, _, _ := charset.DetermineEncoding(bytes, "")
	return e2
}
