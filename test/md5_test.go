package test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"testing"
)

func TestMD5Method(t *testing.T) {

	//str := "晓智科技"
	//md := md5.New()
	//io.WriteString(md, str)
	//strings := hex.EncodeToString(md.Sum(nil))
	//fmt.Println(strings)

	str := "晓智科技"
	md := md5.New()
	io.WriteString(md, str)
	strings := hex.EncodeToString(md.Sum(nil))
	fmt.Println(strings)
}

func TestMd5Salt(t *testing.T) {
	salt, encodedPwd := password.Encode("generic password", nil)
	check := password.Verify("generic password", salt, encodedPwd, nil)
	fmt.Println(check) // true
	fmt.Println(encodedPwd)

	// Using custom options
	options := &password.Options{10, 10000, 50, md5.New}
	salt, encodedPwd = password.Encode("generic password", options)
	check = password.Verify("generic password", salt, encodedPwd, options)
	fmt.Println(check) // true
	fmt.Println(encodedPwd)
}
