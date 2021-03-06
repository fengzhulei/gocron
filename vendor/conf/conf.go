package conf

import (
	"bufio"
	"io"
	"os"
	"strings"
	"strconv"
)

const middle = "========="

type config struct {
	mymap  map[string]string
	strcet string
}

var confMap map[string]*config

func init() {
	confMap = make(map[string]*config)
}

func NewConfig(path string) config {
	if _, ok := confMap[path]; !ok {
		confMap[path] = new(config)
		(*confMap[path]).initConfig(path)
	}
	return *confMap[path]
}

func (c *config) initConfig(path string) {
	c.mymap = make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}
		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}
		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}
		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}
		if len(second) == 0 {
			continue
		}
		key := c.strcet + middle + frist
		c.mymap[key] = strings.TrimSpace(second)
	}
}

func (c config) read(section, key string) string {
	key = section + middle + key
	v, found := c.mymap[key]
	if !found {
		return ""
	}
	return v
}

func (c config) GetString(section, key string) string {
	return c.read(section,key)
}

func (c config) GetMap(section string) map[string]string{
	tmap := make(map[string]string , 10)
	for k,v:= range c.mymap {
		if strings.HasPrefix(k,section+middle){
			tmap[strings.TrimPrefix(k,section+middle)] = v
		}
	}
	return tmap
}

func (c config)GetInt(section, key string) int {
	 value,_ := strconv.Atoi(c.read(section, key ))
	 return value
}

func (c config)GetInt64(section,key string) int64 {
    value,_ := strconv.ParseInt(c.read(section, key),10,64)
    return value
}

func (c config)GetBool(section ,key string) bool {
    value,_ := strconv.ParseBool(c.read(section, key ))
    return value
}


func (c config)GetFloat64(section ,key string ) float64 {
    value,_ := strconv.ParseFloat(c.read(section, key ),64)
    return value
}

func (c config) GetArray(section, key, delim string) []string {
    value := c.read(section, key)
    if value == "" {
        return []string{}
    }
    values := strings.Split(value, delim)
    for i := range values {
        values[i] = strings.TrimSpace(values[i])
    }
    return values
}