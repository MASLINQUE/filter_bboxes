package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var widht_thresh = flag.Float64("widht_thresh", 0, "Widht filter 0-1")
var height_thresh = flag.Float64("height_thresh", 0, "Height filter 0-1")
var diag_thresh = flag.Float64("diag_thresh", 0, "Diag filter 0-1.5")

func main() {

	// open output file
	flag.Parse()
	// make a write buffer
	logrus.SetLevel(logrus.DebugLevel)
	s := bufio.NewScanner(os.Stdin)
	bufsize := 10 << 20
	buf := make([]byte, bufsize)
	s.Buffer(buf, bufsize)
	for {
		if s.Scan() {
			reqdata := s.Bytes()
			expand(reqdata)

		}
	}

}

func expand(reqdata []byte) {

	threshes := map[string]float64{"w": *widht_thresh, "h": *height_thresh, "d": *diag_thresh}

	items := gjson.ParseBytes(reqdata).Get("items")
	if len(items.Array()) == 0 {
		fmt.Println(string(reqdata))
		return
	}
	itemsArray := []gjson.Result{}
	itemsArray = append(itemsArray, items.Array()...)
	responseStringArray := []string{}

	for _, item := range itemsArray {
		item_str := item.String()
		bbox := []float64{}
		item.Get("bbox").ForEach(func(key, value gjson.Result) bool {
			bbox = append(bbox, value.Num)
			return true
		})

		if Filter(bbox, threshes) {
			responseStringArray = append(responseStringArray, item_str)
		}
	}

	responseString := fmt.Sprintf("[%s]", strings.Join(responseStringArray, ", "))

	reqdata_st, _ := sjson.SetRaw(string(reqdata), "items", responseString)
	fmt.Println(string(reqdata_st))

}
