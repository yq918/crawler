package main

import (
	"regexp"
	"fmt"
)

const  text  = `my email is strive965432@gmail.com
email2 is 11111@qq.com
email3 is 89234324@163.com
email4 is 8932483@sina.com
`

func main() {
	compile := regexp.MustCompile(`\w+@\w+\.\w+`)
	s := compile.FindString(text)
	fmt.Printf("%s",s)

	fmt.Println()
	allString := compile.FindAllString(text,-1)
	fmt.Printf("%s---%s",allString,allString[1])




}
