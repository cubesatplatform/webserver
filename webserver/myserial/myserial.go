package myserial

import (
	"fmt"
)

func Deserialize(_str string) map[string]string {
	Parameters := make(map[string]string)
	if len(_str) == 0 {
		return Parameters
	}
	if _str[len(_str)-1] != '^' { //~
		_str += string('^') //~
	}
	var str1 string

	var components []string

	var c byte
	tok1 := byte('^') //~
	tok2 := byte(';') //:
	stop := 0
	start := 0

	for i := 0; i < len(_str); i++ {
		c = _str[i]
		if c == tok1 {
			stop = i
			components = append(components, _str[start:stop])

			start = i + 1
		}
	}

	for x := 0; x < len(components); x++ {
		str1 = components[x]
		for i := 0; i < len(str1); i++ {
			c = str1[i]

			if c == tok2 {
				stop = i
				Parameters[str1[0:i]] = str1[i+1:]
				break
			}
		}
	}
	//fmt.Println(Parameters)

	for k, v := range Parameters {
		fmt.Printf("%s -> %s\n", k, v)
	}

	return Parameters
}

func Serialize(Parameters map[string]string) string {
	str := ""
	for k, v := range Parameters {
		str += k + ":" + v + "~"
	}
	return str
}
