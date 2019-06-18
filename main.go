package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var s []byte
	fmt.Scan(&s)

	start := time.Now()

	l, i1, i2, err := compareLetters(s)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(l, i1, i2)

	end := time.Now()
	fmt.Printf("total processing time: %f s", end.Sub(start).Seconds())

	// if err := makeRandomLetters(); err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }
}

func compareLetters(s []byte) (length, i1, i2 int, err error) {
	length = 0

	// 部分列なのでi=0はダメ
	for i := 1; i < len(s); i++ {
		l, j, err := compareLettersSub(s, i)
		if err != nil {
			return 0, 0, 0, err
		}
		if l > length {
			length = l
			i1 = j
			i2 = i
		}
	}

	return length, i1, i2, nil
}

func compareLettersSub(s []byte, offset int) (length, idx int, err error) {
	// 流石に1文字は一致するでしょう
	length = 1
	idx = 0

	flag := false
	start := 0

	l := len(s) - offset
	if l < 0 {
		return 0, 0, errors.New("offset is larger than len(s)")
	}

	for i := 0; i < l; i++ {
		if s[i] == s[i+offset] {
			if !flag {
				start = i
				flag = true
			}
		} else {
			if flag {
				if _l := i - start + 1; _l > length {
					length = _l
					idx = start
				}
				flag = false
			}
		}

		if i == l-1 && flag {
			length = i - start + 1
			idx = start
		}
	}

	return length, idx, nil
}

func makeRandomLetters() error {
	var length int
	fmt.Scan(&length)

	s := make([]byte, length)
	letters := map[int]byte{
		0: 'A',
		1: 'T',
		2: 'C',
		3: 'G',
	}

	for i := 0; i < length; i++ {
		s[i] = letters[rand.Intn(4)]
	}

	if _, err := os.Stat("./in"); os.IsNotExist(err) {
		os.Mkdir("./in", 0777)
	}
	file, err := os.OpenFile(fmt.Sprintf("./in/%d.in", length), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.New(fmt.Sprintf("cannnot open ./in/%d.in:", length) + err.Error())
	}

	file.Write(s)

	return nil
}
