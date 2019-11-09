package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/kr/pretty"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir/metadata"
)

var (
	outputFormat string
)

func init() {
	flag.StringVar(&outputFormat, "format", "csv", "output file format (csv,tsv,json,yaml)")
}

type FuncInfo struct {
	Id          int    `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	LinkageName string `json:"linkage_name" yaml:"linkage_name"`
	Directory   string `json:"directory" yaml:"directory"`
	Filename    string `json:"filename" yaml:"filename"`
	Line        int    `json:"line" yaml:"line"`
}

type FuncInfoWriter interface {
	Write(w io.Writer, datas []*FuncInfo) (err error)
}

type FuncInfoCSVWriter struct {
}

func (f *FuncInfoCSVWriter) Write(w io.Writer, datas []*FuncInfo) (err error) {
	for _, data := range datas {
		_, err = fmt.Fprintf(w, "%d,%s,%s,%s,%s,%d\n", data.Id, data.Name, data.LinkageName, data.Directory, data.Filename, data.Line)
		if err != nil {
			return
		}
	}
	return
}

type FuncInfoTSVWriter struct {
}

func (f *FuncInfoTSVWriter) Write(w io.Writer, datas []*FuncInfo) (err error) {
	for _, data := range datas {
		_, err = fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%d\n", data.Id, data.Name, data.LinkageName, data.Directory, data.Filename, data.Line)
		if err != nil {
			return
		}
	}
	return
}

type FuncInfoJSONWriter struct {
}

func (f *FuncInfoJSONWriter) Write(w io.Writer, datas []*FuncInfo) (err error) {
	out, err := json.Marshal(datas)
	if err != nil {
		return
	}
	_, err = w.Write(out)
	if err != nil {
		return
	}
	return
}

type FuncInfoYAMLWriter struct {
}

func (f *FuncInfoYAMLWriter) Write(w io.Writer, datas []*FuncInfo) (err error) {
	out, err := yaml.Marshal(datas)
	if err != nil {
		return
	}
	_, err = w.Write(out)
	if err != nil {
		return
	}
	return
}

func main() {
	flag.Parse()

	var inputFiles []string
	if flag.NArg() == 0 {
		inputFiles = append(inputFiles, os.Stdin.Name())
	} else {
		inputFiles = append(inputFiles, flag.Args()...)
	}
	var funcInfoWriter FuncInfoWriter
	switch outputFormat {
	case "csv":
		funcInfoWriter = &FuncInfoCSVWriter{}
	case "tsv":
		funcInfoWriter = &FuncInfoTSVWriter{}
	case "json":
		funcInfoWriter = &FuncInfoJSONWriter{}
	case "yaml":
		funcInfoWriter = &FuncInfoYAMLWriter{}
	default:
		fmt.Fprintf(os.Stderr, "Unknown output format '%s'\n", outputFormat)
		flag.Usage()
		os.Exit(1)
	}

	funcInfos := make([]*FuncInfo, 0)
	for _, inputFile := range inputFiles {
		m, err := asm.ParseFile(inputFile)
		if err != nil {
			log.Fatal("file read error", err)
		}

		for _, f := range m.Funcs {
			for _, attachment := range f.Metadata.MDAttachments() {
				node := attachment.Node
				switch node := node.(type) {
				case *metadata.DISubprogram:
					funcInfo := &FuncInfo{
						Id:          int(node.ID()),
						Name:        node.Name,
						LinkageName: node.LinkageName,
						Directory:   node.File.Directory,
						Filename:    node.File.Filename,
						Line:        int(node.Line),
					}
					funcInfos = append(funcInfos, funcInfo)
				default:
					pretty.Logln(node)
					os.Exit(1)
				}
			}
		}
	}
	err := funcInfoWriter.Write(os.Stdout, funcInfos)
	if err != nil {
		log.Fatal(err)
	}
}
