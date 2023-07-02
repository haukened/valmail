package dea

import (
	"bufio"
	"bytes"
	"embed"
	"sync"
)

//go:embed disposable_email_blocklist.conf
var f embed.FS

type DEA interface{}

var list = make(map[string]struct{})
var lock sync.Mutex

func init() {
	lock.Lock()
	data, err := f.ReadFile("disposable_email_blocklist.conf")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		url := scanner.Text()
		list[url] = struct{}{}
	}
	lock.Unlock()
}

func IsDEA(url string) bool {
	_, ok := list[url]
	return ok
}
