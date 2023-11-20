package sqldumptocsv

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

func strip(str string) string {
	return strings.Trim(str, "\r\n\t ,)()")
}

func parseInsertValue(str string) []string {
	result := []string{}
	v := strip(str)
	e := strings.Split(v, ",")
	for i := 0; i < len(e); i++ {
		result = append(result, strings.Trim(e[i], " '"))
	}
	// fmt.Printf(v)
	// fmt.Printf("%s", e)
	return result
}

func ParseInsert(str string, topK int) Result {
	result := Result{}
	v := strings.Split(str, "VALUES\r\n\t")
	id := strings.Split(v[0], "`")
	values := strings.Join(v[1:], "VALUES\r\n\t")

	result.TableName = strip(id[1])
	for i := 3; i < len(id); i += 2 {
		result.Columns = append(result.Columns, strip(id[i]))
	}
	spt := split([]byte(values), '\t')

	// fmt.Printf("%s", id)
	// fmt.Printf("%s", values)
	// fmt.Printf("%s", spt)
	if topK != -1 {
		if topK > len(spt) {
			topK = len(spt)
		}
	} else {
		topK = len(spt)
	}

	for i := 0; i < topK; i += 1 {
		result.Values = append(result.Values, parseInsertValue(spt[i]))
	}

	return result
}

// parses a SQL dump file and converts it to CSV
func Parse(sqldump []byte, topK int) []Result {
	// TODO

	// split by ;
	texts := split(sqldump, ';')
	for i := 0; i < len(texts); i++ {
		texts[i] = strings.Trim(texts[i], "\n\r\t ")
	}
	texts = filter(texts)

	result := []Result{}

	for _, v := range texts {
		result = append(result, ParseInsert(v, topK))
	}

	return result
}
