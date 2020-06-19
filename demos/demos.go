package demos

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"math/big"
	manda "math/rand"
	"os"
	"reflect"
	"time"
)

func DemoChan() {
	c := make(chan string, 1024)
	c <- "www"
	c <- "eee"
	c <- "rrr"
	c <- "ttt"
	s := <- c
	defer close(c)
	log.Println(s)
}

func DemoSlice() {
	var slice []string = []string{"slice-0", "slice-1", "slice-2", "slice-3", "slice-4", "slice-5", "slice-6", "slice-7"}
	log.Println(slice[:2])
	log.Println(slice[:3])
	log.Println(slice[:4])
	log.Println(slice[1:])
	log.Println(slice[2:])
	log.Println(slice[5:])
	log.Println(slice[:])
}

func DemoSelect() {
	var a, b int = 2, 5
	var c = make(chan int, 20)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	select {
	case a = <- c:
		log.Println(a, "----")
	case b = <- c:
		log.Println(b, "``````")
	default:
		log.Println(0)
	}
}

func DemoRand() {
	n1, _ := rand.Int(rand.Reader, big.NewInt(100))
	n2, _ := rand.Int(rand.Reader, big.NewInt(100))
	n3, _ := rand.Int(rand.Reader, big.NewInt(100))
	n4, _ := rand.Int(rand.Reader, big.NewInt(100))
	n5, _ := rand.Int(rand.Reader, big.NewInt(100))
	log.Println(n1)
	log.Println(n2)
	log.Println(n3)
	log.Println(n4)
	log.Println(n5)
	log.Println()
	manda.Seed(time.Now().UnixNano())
	log.Println(manda.Intn(100))
	log.Println(manda.Intn(100))
	log.Println(manda.Intn(100))
	log.Println(manda.Intn(100))
	log.Println(manda.Intn(100))

	r := []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	n := make([]rune, 4)
	for i := range n {
		n[i] = r[manda.Intn(len(r))]
	}
	log.Println(string(n))

}

func DemoMd5() {
	data := []byte("golang")
	m := md5.New()
	m.Write(data)
	cipherStr := m.Sum(nil)
	fmt.Printf("%x\n", cipherStr)
	fmt.Printf("%x\n", md5.Sum(data))
	fmt.Println(hex.EncodeToString(cipherStr))
}

func DemoXml() {
	xmlDoc := `<?xml version="1.0" encoding="UTF-8"?>
				<note>
				  <to>Tove</to>
				  <from>Jani</from>
				  <heading>Reminder</heading>
				  <body>Don't forget me this weekend!</body>
				</note>`
	type xmlStruct struct {
		XMLName  xml.Name `xml:"note"`
		To string `xml:"to"`
		From string `xml:"from"`
		Heading string `xml:"heading"`
		Body string `xml:"body"`
	}
	x := xmlStruct{}
	err := xml.Unmarshal([]byte(xmlDoc), &x)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(x)
	}

	type Address struct {
		City, State string
	}

	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42, Height: 3.2}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "	", "	")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(string(output))

}

func DemoJson() {
	var err error
	// json to struct
	type List struct {
		Id int  `json:"id"`
		Name string `json:"name"`
		Mobile string   `json:"mobile"`
		IsDelete bool   `json:"is_delete"`
	}
	type Data struct {
		Code int    `json:"code"`
		List []List `json:"list"`
	}
	jsonStr := `{"code": 1, "list": [
                {"id": 1, "name": "刘备", "mobile": "13399999999", "is_delete": true},
                {"id": 2, "name": "张飞", "mobile": "13388888888", "is_delete": false},
                {"id": 3, "name": "关羽", "mobile": "13366666666", "is_delete": true}
                ]}`
	fmt.Println("--------------------------------------- json to struct")
	var data Data
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err == nil {
		fmt.Println(data)
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println("--------------------------------------- json to map")
	// json to map
	mapData := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonStr), &mapData)
	if err == nil {
		fmt.Println(mapData)
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println("--------------------------------------- struct to json")
	// struct to json
	structJson, err := json.Marshal(data)
	if err == nil {
		fmt.Println(string(structJson))
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println("--------------------------------------- map to json")
	// map to json
	mapJson, err := json.Marshal(mapData)
	if err == nil {
		fmt.Println(string(mapJson))
	} else {
		fmt.Println(err.Error())
	}
}

func DemoPath() {
	dir, _ := os.Getwd()
	log.Println(dir)
}

func DemoTypeOf(p interface{}) {
	switch p.(type) {
	case string:
		fmt.Println(p, reflect.TypeOf(p), "string")
		break
	case int:
		fmt.Println(p, reflect.TypeOf(p), "int")
		break
	case float64:
		fmt.Println(p, reflect.TypeOf(p), "float64")
		break
	case bool:
		fmt.Println(p, reflect.TypeOf(p), "bool")
		break
	default:
		fmt.Println(p, reflect.TypeOf(p), "other")
		break
	}
}

func DemoTimes() {
	// 东八区
	cstZone := time.FixedZone("CST", 8*3600)
	n := time.Now().In(cstZone)
	// 获取时间
	t := n.Format("2006-01-02 15:04:05")
	// 年
	year := n.Year()
	// 月
	month := n.Month()
	// 日
	day := n.Day()
	// 时
	hour := n.Hour()
	// 分
	minute := n.Minute()
	// 秒
	second := n.Second()
	// 时间戳
	u := n.Unix()
	log.Printf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second)
	log.Println(t)
	log.Println(u)
}

func DemoTimeFormat() {
	// 时间戳格式化
	t := time.Unix(1592544444, 0)
	ft := t.Format("2006-01-02 15:04:05")
	log.Println(ft)
	// 日期转时间戳
	t1, _ := time.Parse("2006-01-02 15:04:05", "2020-06-19 13:27:22")
	t2, _ := time.Parse("2006-01-02 15:04:05", "2020-06-19 13:27:23")
	t3, _ := time.Parse("2006-01-02 15:04:05", "2020-06-19 13:27:24")
	log.Println(t1.Unix() - 8 * 3600)
	log.Println(t2.Unix() - 8 * 3600)
	log.Println(t3.Unix() - 8 * 3600)
}
























