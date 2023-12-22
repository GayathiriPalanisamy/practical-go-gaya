package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)
func main(){
	sig, err := sha1sum("http.log.gz")
	if err != nil{
		log.Fatalf("error %s", err)
	}
	fmt.Print(sig)

	sig, err = sha1sum("sha.go")
	if err != nil{
		log.Fatalf("error %s", err)
	}
	fmt.Print("\n", sig)
}

/* 
if filename end swith .gz 
   cat http.log.gz | gunzip |sha1sum
else
   cat http.log.gz | sha1sum
*/
func sha1sum(fileName string)(string, error) {
	file, err := os.Open(fileName)
	if err != nil{
		return "", nil
	}
	defer file.Close() // it happens at function level and deffered are collaed in reverse order LIFO
	
	var r io.Reader = file
	if strings.HasSuffix(fileName, ".gz"){
		file , err := gzip.NewReader(file)
		if err != nil{
			return "", nil
		}
		defer file.Close()
		r = file
	}

	// io.CopyN(os.Stdout, r , 100) //uncompressed

	w := sha1.New()
	if _, err := io.Copy(w,r); err != nil{
		return "", nil
	}
	sig := w.Sum(nil)
	
	return fmt.Sprintf("%x", sig), nil
}