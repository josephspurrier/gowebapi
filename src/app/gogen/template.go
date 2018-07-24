package gogen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Collection represents a group of templates.
type Collection struct {
	ConfigType       string        `json:"config.type"`
	ConfigCollection []interface{} `json:"config.collection"`
	ItemUpper        string        `json:"itemUpper"`
	ItemLower        string        `json:"itemLower"`
}

// KeyPair represents a key and a value..
type KeyPair struct {
	Key   string
	Value string
}

const (
	// LDelimiter is the left delimiter for templates.
	LDelimiter = "[["
	// RDelimiter is the right delimiter for templates.
	RDelimiter = "]]"
)

func glob(dir string, ext string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// CreateTemplate will generate a template from source code.
func CreateTemplate(args []string, templateName string, projectFolder string, templateFolder string) error {
	// Get a list of all the files in the folder.
	basePath := filepath.Join(projectFolder, args[0]) + "/"
	arr, err := glob(basePath, ".go")
	if err != nil {
		return err
	}

	// Loop through the key pairs.
	fields := make(map[string]string)
	pairs := make([]KeyPair, 0)
	for _, v := range args[2:] {
		pair := strings.Split(v, ":")
		p := KeyPair{}
		p.Key = pair[0]
		p.Value = pair[1]
		pairs = append(pairs, p)
		fields[pair[0]] = fmt.Sprintf("%v.%v%v", LDelimiter, pair[0], RDelimiter)
	}

	// Loop through each file.
	items := make([]interface{}, 0)
	for _, f := range arr {
		filename := strings.TrimPrefix(f, basePath)
		file := strings.TrimSuffix(filename, ".go")
		item := make(map[string]interface{})
		item[templateName+"/"+file+"/default"] = fields
		items = append(items, item)

		genFile := filepath.Join(templateFolder, templateName, file, "default.gen")
		jsonFile := filepath.Join(templateFolder, templateName, file, "default.json")

		// Check if the folder exists.
		dir := filepath.Dir(genFile)
		if !Exists(dir) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
		}
		b, err := ioutil.ReadFile(f)
		if err != nil {
			log.Println("Cannot read file:", filename)
			continue
		}

		s := string(b)

		mSingle := make(map[string]string)
		mSingle["config.type"] = "single"
		outputPath := filepath.Join(args[0], filename)

		for _, kp := range pairs {
			s = strings.Replace(s, kp.Value, fmt.Sprintf("%v.%v%v", LDelimiter, kp.Key, RDelimiter), -1)
			outputPath = strings.Replace(outputPath, kp.Value, fmt.Sprintf("%v.%v%v", LDelimiter, kp.Key, RDelimiter), -1)
			mSingle[kp.Key] = ""
		}

		mSingle["config.output"] = outputPath

		jsonOut, err := json.Marshal(mSingle)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(genFile, []byte(s), 0644)
		if err != nil {
			return err
		}
		fmt.Println("Template generated:", genFile)

		err = ioutil.WriteFile(jsonFile, jsonOut, 0644)
		if err != nil {
			return err
		}
		fmt.Println("Template generated:", jsonFile)
	}

	m := make(map[string]interface{})
	m["config.type"] = "collection"
	m["config.collection"] = items

	out, err := json.Marshal(m)
	if err != nil {
		return err
	}

	collectionFile := filepath.Join(templateFolder, templateName, "default.json")

	// Check if the folder exists.
	dir := filepath.Dir(collectionFile)
	if !Exists(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(collectionFile, out, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Template generated:", collectionFile)

	return nil
}
