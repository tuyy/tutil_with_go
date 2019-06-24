package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	DATE_LEN      = 8  // ex) 20190524
	DATETIME_LEN  = 14 // ex) 20190524110000
	TIMESTAMP_LEN = 10 // ex) 1590159600

	DATE_LAYOUT            = "20060102"
	DATETIME_LAYOUT        = "20060102150405"
	DATETIME_RETURN_LAYOUT = "2006/01/02 15:04:05 (MST)"
)

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

func Process(base string, op int) string {
	i, err := strconv.ParseInt(base, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("input is not number. input:%s err:%s", base, err))
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

	return fmt.Sprintf("%s // %d", t.Format(DATETIME_RETURN_LAYOUT), t.Unix())
}

func main() {
	args := os.Args
	if len(args) == 1 || len(args) > 3 {
		panic("dt {20190524 or 20190524093050 or 1590246000} (optional){+1}")
	}

	var op int64
	var err error

	if len(args) == 3 {
		op, err = strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			panic(fmt.Sprintf("invalid operator. err:%s", err))
		}
	}

	fmt.Println(Process(args[1], int(op)))
}
