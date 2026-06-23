package main

import (
	"strconv"
	"fmt"
	"strings"
)


func ToHex(str string) string{
	tokens := strings.Fields(str)
	for i:=0; i < len(tokens); i++{
		if tokens[i] == "(hex)"{
			if i > 0{
				num, err := strconv.ParseInt(tokens[i-1],16, 64)
				if err == nil{
					tokens[i-1]= strconv.FormatInt(num, 10)
				}
				tokens = append(tokens[:i], tokens[i+1:]...)
				i--
				
			}
		}else if tokens[i] == "bin"{
			if i > 0{
				nums , err := strconv.ParseInt(tokens[i-1],2,64)
				if err == nil{
					tokens[i-1] = strconv.FormatInt(nums,10)
				}
				tokens = append(tokens[:i], tokens[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(tokens, " ")

}

func main(){
	fmt.Println(ToHex("1E (hex) files 111 (bin)"))
}