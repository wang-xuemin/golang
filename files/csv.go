package files

import (
	"encoding/csv"
	"log"
	"os"
)

func CreateCsv(filename string, data [][]string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err.Error())
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF")
	v := csv.NewWriter(f)
	err = v.WriteAll(data)
	if err != nil {
		log.Println(err.Error())
	}
}

func ReadCsv(filename string) (r [][]string, err error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Println(err.Error())
	}
	v := csv.NewReader(f)
	r, err = v.ReadAll()
	if err != nil {
		log.Println(err.Error())
	}
	return
}
