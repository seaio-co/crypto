package md5

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"os"
	"fmt"
	"io/ioutil"
	"path"
	"io"
)

// Md5Sum 常用的md5摘要算法
func Md5Sum(input string) string {

	h := md5.New()
	h.Write([]byte(input))
	sum := h.Sum(nil)
	sumStr := hex.EncodeToString(sum)
	sumStr = strings.ToLower(sumStr)
	return sumStr
}

var (
	Fa      string
	md5out string
	Type   string
	md5in  string
	help   string
)

func MakeMd5() string {
	fileName := Fa

	var md5file string
	if len(md5out) == 0 {
		md5file = Fa + ".md5"
	} else {
		md5file = md5out
	}
	//fmt.Println("md5file = ", md5file)
	//md5file := "filelist.go.text.md5"
	//bufs := make([]byte, 1024)

	fin, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(fileName, err)
	}
	defer fin.Close()

	fout, foute := os.Create(md5file)
	if foute != nil {
		fmt.Println(md5file, foute)
	}
	defer fout.Close()

	Buf, buferr := ioutil.ReadFile(fileName)
	if buferr != nil {
		fmt.Println(fileName, buferr)
	}

	temp := hex.EncodeToString(byte2string(md5.Sum(Buf))) + " " + "*" + path.Base(fileName)
	_, err = fout.WriteString(temp)
	if err != nil {
		fmt.Println(md5file, err)
	}
	fmt.Printf("md5sum = %s", temp)

	return temp
}

func Verifymd5sum(file string, md5file string) bool {
	if len(md5file) == 0 {
		f, e := os.Stat(file + ".md5")
		if e != nil {
			fmt.Println(file, ".md5", "Doesn't exits")
			return false
		}
		md5file = f.Name()

	}
	md5sumfile, err := ioutil.ReadFile(md5file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(md5sumfile))

	md5sum := string(md5sumfile)
	str := strings.Split(md5sum, " ")
	//fmt.Println(len(str))
	var md5fileName string
	md5string := str[0] //get the md5sum string from file
	//fmt.Println("str[1][:1] = ", str[1][:1])
	if str[1][:1] == string("*") {

		md5fileName = str[1][1:len(str[1])] // get the filename from md5sum.md5 file
	} else {
		md5fileName = str[1]
	}

	//check filename  with md5sum.md5
	if path.Base(md5fileName) == path.Base(file) {
		fmt.Println("FileName has compact")
	} else {
		fmt.Println("Not the same FileName")
	}

	fmt.Println("md5fileName =", md5fileName[:len(md5fileName)])

	h, _ := os.Open(file)
	buf := md5.New()
	io.Copy(buf, h)

	md5fromfile := hex.EncodeToString(buf.Sum(nil))

	if md5fromfile == md5string {
		return true
	}

	//fmt.Printf("Name = %s md5sum = %s\n", md5string, md5fileName)
	return false
}

func byte2string(in [16]byte) []byte {
	return in[:16]
}