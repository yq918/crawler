package parser

import (
	"engine"
	"regexp"
	"strconv"
	"model"
)

var (
	ageRe  =  regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
    marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
    heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+CM)</td>`)
	educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	addrRe   = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
	hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
    carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
    houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
)


func ParseProfile(contents []byte,name string) engine.ParserResult  {
     profile := model.Profile{}
     profile.Name = name
	 profile.Age =  extractInt(contents,ageRe)
	 profile.Marriage =  extractString(contents,marriageRe)
	 profile.Height = extractInt(contents,heightRe)
	 profile.Education = extractString(contents,educationRe)
	 profile.Income = extractString(contents,incomeRe)
	 profile.Address = extractString(contents,addrRe)
	 profile.Hokou = extractString(contents,hokouRe)
	 profile.Gender = extractString(contents,genderRe)
	 profile.XinZuo = extractString(contents,xinzuoRe)
	 profile.Car = extractString(contents,carRe)
	 profile.House = extractString(contents,houseRe)

	 result := engine.ParserResult{
	 	Items:[]interface{} {profile},
	 }
     return result
}



func extractString( contents []byte, re *regexp.Regexp)  string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2{
         return  string(match[1])
	}
	return ""
}

func extractInt(content []byte,re *regexp.Regexp) int  {
	submatch := re.FindSubmatch(content)
	if len(submatch) >=2 {
		i, e := strconv.Atoi( string(submatch[1]) )
		if e != nil {
			return i
		}
	}
	return 0
}


