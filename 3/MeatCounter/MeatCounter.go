package MeatCounter

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func OnPage(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func MeatCount() string {
	link := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	baconipsum := OnPage(link)
	return MeatFormat(baconipsum)
}

func MeatFormat(baconipsum string) string {
	if baconipsum == "" {
		return ""
	}
	arrSplit := strings.FieldsFunc(strings.ToLower(baconipsum), mySplit)
	var meatBag = make(map[string]int)
	for _, s := range arrSplit {
		meatBag[s] += 1
	}
	// for meat, value := range meatBag {
	// 	fmt.Println(meat, " ", value)
	// }
	json_marshalled, _ := json.Marshal(meatBag)

	return string(json_marshalled)
}

func mySplit(r rune) bool {
	return r == '.' || r == ',' || r == ' ' || r == '\n'
}
