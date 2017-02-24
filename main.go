package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	banner = `

  /$$$$$$   /$$$$$$  /$$$$$$$   /$$$$$$  /$$  /$$  /$$
 /$$__  $$ /$$__  $$| $$__  $$ /$$__  $$| $$ | $$ | $$
| $$  \ $$| $$  \ $$| $$  \ $$| $$  \ $$| $$ | $$ | $$
| $$  | $$| $$  | $$| $$  | $$| $$  | $$| $$ | $$ | $$
|  $$$$$$$|  $$$$$$/| $$  | $$|  $$$$$$/|  $$$$$/$$$$/
 \____  $$ \______/ |__/  |__/ \______/  \_____/\___/
 /$$  \ $$
|  $$$$$$/
 \______/

Print current time use below

$ gonow

options

`
	formatHelp = `time format string (ex:yyyy-MM-dd'T'HH:mm:ss.SSS)
	or golang time format https://golang.org/pkg/time/#pkg-constants
	default option print
	time.Now()
	time.Now().Format("2006-01-02T15:04:05.000")
	time.Now().Format("20060102150405.000")
`
	yyyy = "2006"
	yy   = "06"
	MM   = "01"
	dd   = "02"
	HH   = "15"
	mm   = "04"
	ss   = "05"
	S    = "0"
)

func main() {
	format := flag.String("f", "", formatHelp)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, banner)
		flag.PrintDefaults()
	}
	flag.Parse()
	if *format == "" {
		fmt.Println(time.Now())
		fmt.Println(time.Now().Format("2006-01-02T15:04:05.000"))
		fmt.Println(time.Now().Format("20060102150405.000"))
	} else {
		var layout = strings.Replace(*format, "yyyy", yyyy, -1)
		layout = strings.Replace(layout, "yy", yy, -1)
		layout = strings.Replace(layout, "MM", MM, -1)
		layout = strings.Replace(layout, "dd", dd, -1)
		layout = strings.Replace(layout, "HH", HH, -1)
		layout = strings.Replace(layout, "mm", mm, -1)
		layout = strings.Replace(layout, "ss", ss, -1)
		layout = strings.Replace(layout, "S", S, -1)
		fmt.Printf("formatting %s >> %s\n", *format, layout)
		fmt.Println(time.Now().Format(layout))
	}
}
