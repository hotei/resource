// resource.go (c) 2015 David Rook - all rights reserved
//
// tool to convert files to go code - see also go generate for other ways
//
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"strings"
	//
	//"github.com/hotei/mdr"
)

var (
	flagSourceFile   string
	flagResourceFile string
	flagPackageName  string
	flagVarName      string
	flagLineLength   int

	testData *string // saves handle to the base64 string for comparison later
)

func init() {
	flag.StringVar(&flagSourceFile, "source", "", "Source file (eg pic.jpg) ")
	flag.StringVar(&flagResourceFile, "rc", "", "Resource (output.go) file")
	flag.StringVar(&flagPackageName, "package", "main", "Package name")
	flag.StringVar(&flagVarName, "var", "", "Variable being created")
	flag.IntVar(&flagLineLength, "line", 80, "Max line length for created strings")
}

func flagSetup() {
	if len(flagSourceFile) < 4 {
		fmt.Printf("Source flag missing\n")
		flag.Usage()
		os.Exit(0)
	}
	if len(flagResourceFile) < 1 {
		fmt.Printf("Resource flag missing\n")
		flag.Usage()
		os.Exit(0)
	}
	if len(flagVarName) < 1 {
		fmt.Printf("Variable flag missing\n")
		flag.Usage()
		os.Exit(0)
	}
}

// Encode64File returns error if :
//		cant read input file
//		output file cant be opened
func Encode64File(fname string, rcName string) error {
	bites, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Printf("While reading %s got %v\n", fname, err)
		return err
	}
	of, err := os.Create(rcName)
	if err != nil {
		log.Printf("While opening %s got %v\n", rcName, err)
		return err
	}
	defer of.Close()

	stuff := base64.StdEncoding.EncodeToString(bites)
	testData = &stuff
	fmt.Printf("Stuff length = %d\n", len(stuff))
	fmt.Fprintf(of, "package main\n")
	fmt.Fprintf(of, "import (\n")
	fmt.Fprintf(of, "%q\n", "encoding/base64")
	fmt.Fprintf(of, "%q\n", "log")
	fmt.Fprintf(of, "%q\n", "strings")
	fmt.Fprintf(of, ")\n")
	fmt.Fprintf(of, "var (\n")
	fmt.Fprintf(of, "\t%s []byte\n", flagVarName)
	fmt.Fprintf(of, ")\n\n")
	fmt.Fprintf(of, "func init() {\n")
	fmt.Fprintf(of, "var tmp64 []string = []string {\n")
	lineLen := flagLineLength
	start := 0
	done := false
	for {
		end := start + lineLen
		if end >= len(stuff) {
			end = len(stuff)
			done = true
		}
		//fmt.Printf("stuff[%4d %4d]= %q ,\n", start, end-1, stuff[start:end])
		fmt.Fprintf(of, "%q ,\n", stuff[start:end])
		if done {
			break
		}
		start += lineLen
	}
	fmt.Fprintf(of, "}\n")
	fmt.Fprintf(of, "var err error\n")
	fmt.Fprintf(of, "// decode it here...\n")
	fmt.Fprintf(of, `tmpOne := strings.Join(tmp64,"")`+"\n")
	fmt.Fprintf(of, "%s,err = base64.StdEncoding.DecodeString(tmpOne)\n",
		flagVarName)
	fmt.Fprintf(of, "if err != nil {\n")
	//fmt.Fprintf(of, `fmt.Printf(of,"error:%%v", err)`+"\n")
	fmt.Fprintf(of, `log.Panicf("error: %%v",err)`+"\n")
	fmt.Fprintf(of, "}\n")
	fmt.Fprintf(of, "}\n")
	return err
}

/*
func decodeTest(fname string) {
	sha, _ := mdr.FileSHA256(fname)
	bites, err := base64.StdEncoding.DecodeString(*testData)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	sha2 := mdr.BufSHA256(bites)
	if sha != sha2 {
		fmt.Printf("Stuff sha256 file:%s\n", sha)
		fmt.Printf("Stuff sha256Bufr :%s\n", sha2)
		log.Panicf("decoded version and original don't match")
	}
}
*/

func main() {
	paranoid := false
	flag.Parse()
	flagSetup()
	err := Encode64File(flagSourceFile, flagResourceFile)
	if err != nil {
		log.Panicf("%v", err)
	}
	if paranoid {
		// decodeTest(flagSourceFile)
	}
}
