package sqldumptoobject

import (
	"strings"
)

func split(str []byte, sep byte) []string {
	// TODO
	is_quate := false
	result := []string{}

	tmp := []byte{}
	for i := 0; i < len(str); i++ {
		tmp = append(tmp, str[i])

		if i != 0 {
			if str[i-1] == '\\' && str[i] == '\'' {

			} else if str[i] == '\'' {
				is_quate = !is_quate
			}
		}

		if !is_quate && str[i] == sep {
			result = append(result, string(tmp))
			tmp = []byte{}
		}
	}

	return result
}

func filter(str []string) []string {
	result := []string{}
	for i := 0; i < len(str); i++ {
		if strings.HasPrefix(str[i], "//") || strings.HasPrefix(str[i], "--") || strings.HasPrefix(str[i], "/*") || strings.HasPrefix(str[i], "#") || !strings.HasPrefix(str[i], "INSERT") {
			continue
		}
		result = append(result, str[i])
	}
	return result
}

func ParseInsert(sql string) string {
	return ""
}

func ParseCreate(sql string) string {
	return ""
}

// parses a SQL dump file and converts it to CSV
func Parse(sqldump []byte, topK int) []string {
	// TODO

	// split by ;
	texts := split(sqldump, ';')
	for i := 0; i < len(texts); i++ {
		texts[i] = strings.Trim(texts[i], "\n\r\t ")
	}
	texts = filter(texts)

	return texts
}
