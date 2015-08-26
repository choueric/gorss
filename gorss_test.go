package gorss_test

import (
	"fmt"
	"github.com/choueric/gorss"
	//"io/ioutil"
	"os"
	"testing"
)

func Test_GenerateFeed(t *testing.T) {
	fmt.Printf("test GenerateFeed\n")
	gorss.GenerateFeed(nil, os.Stdout, 0)
}

/*
func Test_ParsePage(t *testing.T) {
	fmt.Printf("test ParsePage\n")
	f, err := os.OpenFile("demo.html", os.O_RDWR, os.ModePerm)
	if err != nil {
		t.Errorf("open failed")
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Errorf("read failed")
	}
	gorss.ParsePage(data)
}
*/
