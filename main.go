package doctotext

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

const DEFAULT_FILE_REGEX = "<[\\?a-zA-Z0-9:\"\\/= .-]*>"

// This function get path of file and tries to extract it's content .
// ATTENTION: this funciton just extract every plain text from doc files ,
func GetDocContent(filePath string, pattern any) (string, error) {
	if pattern == nil {
		pattern = DEFAULT_FILE_REGEX
	}

	reader, err := zip.OpenReader(filePath)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer reader.Close()
	for _, f := range reader.File {
		fileName := f.Name
		if fileName == "word/document.xml" {
			fc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer fc.Close()
			bytes, _ := ioutil.ReadAll(fc)
			return FilterBytesstringWithCustomRegex(pattern.(string), bytes, nil), err
		}
	}
	return "", err
}

// this function provides content of xml file .
// ATTENTION: this funciton just extract every plain text from XML files ,
func GetDocContentWithXMLPath(xmlPath string, pattern any) (string, error) {
	if pattern == nil {
		pattern = DEFAULT_FILE_REGEX
	}
	file, err := os.Open(xmlPath)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer file.Close()
	fileBytes := make([]byte, 1024)
	_, err = file.Read(fileBytes)
	return FilterBytesstringWithCustomRegex(pattern.(string), fileBytes, nil), err
}

// This function takes regex pattern and removing other charactors that doesn't match to that pattern.
// if the third parameter was nil then it replace the character to empty string ("").
func FilterBytesstringWithCustomRegex(regex string, fileBytes []byte, replacement any) string {
	if replacement == nil {
		replacement = ""
	}
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(string(fileBytes), replacement.(string))
}
