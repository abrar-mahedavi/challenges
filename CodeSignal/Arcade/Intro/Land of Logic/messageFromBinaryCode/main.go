package main

import "fmt"
import "strconv"

func messageFromBinaryCode(code string) string {
    s := ""
    for i:=0; i < len(code); i +=8 {
        c,_ := strconv.ParseUint(code[i:i+8],2, 8)
        s = fmt.Sprintf("%s%c",s,c)
    }
    return s
}


func main() {
	fmt.Printf("%v", messageFromBinaryCode("010010000110010101101100011011000110111100100001")
}
