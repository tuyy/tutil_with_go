package main

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
	"time"
)

const (
	DATE_LEN      = 8  // 20190524
	DATETIME_LEN  = 14 // 20190524110000
	TIMESTAMP_LEN = 10 // 1590159600

	DATE_LAYOUT            = "20060102"
	DATETIME_LAYOUT        = "20060102150405"
	DATETIME_RETURN_LAYOUT = "2006/01/02 15:04:05 (MST)"
)

type Result struct {
	Dt string
	Ts int64
}

func getTimeLayoutBySize(baseSize int) string {
	switch baseSize {
	case DATE_LEN:
		return DATE_LAYOUT
	case DATETIME_LEN:
		return DATETIME_LAYOUT
	default:
		panic(fmt.Sprintf("invalid baseSize. size:%d", baseSize))
	}
}

func getDateAndDatetime(base string) time.Time {
	loc, _ := time.LoadLocation("Asia/Seoul")

	t, err := time.ParseInLocation(getTimeLayoutBySize(len(base)), base, loc)
	if err != nil {
		panic(fmt.Sprintf("failed to ParseInLocation(..). err:%s", err))
	}

	return t
}

func process(base string, op int) {
	i, err := strconv.ParseInt(base, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("input is not number. input:%s", base))
	}

	var t time.Time

	switch len(base) {
	case DATE_LEN, DATETIME_LEN:
		t = getDateAndDatetime(base)
	case TIMESTAMP_LEN:
		t = time.Unix(i, 0)
	default:
		panic(fmt.Sprintf("invalid args size. input:%s", base))
	}
	t = t.AddDate(0, 0, op)
	rz := Result{t.Format(DATETIME_RETURN_LAYOUT), t.Unix()}

	tmpl, err := template.New("RZ").Parse("{{.Dt}} // {{.Ts}}\n")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, rz)
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args
	if len(args) == 1 || len(args) > 3 {
		panic("dt {20190524 or 20190524093050 or 1590246000} (optional){+1}")
	}

	if len(args) == 3 {
		op, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			panic(fmt.Sprintf("invalid operator. err:%s", err))
		}

		process(args[1], int(op))
	} else {
		process(args[1], 0)
	}

}
