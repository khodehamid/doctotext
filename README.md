# DOC2TEXT Word To Text Convertor . 

#### With this package you can extract __doc__ and __docx__ texts.
---
## Features
---
+ UTF-8 support .
+ Extracting XML file content  .
+ Support custom regex . 
## Installation
---
#### to install this package on your project just run : 
    go get github.com/khodehamid/doc2text
## Quick Start
---
##### the following code show you how to extract doc , docx or xml file content : 

``` go
package main

import (
	"log"
	"github.com/khodehamid/doc2text"
)

func main() {
	// you can send your custom regex pattern to extract from file otherwise send nil .
	content, err := doc2text.GetDocContent("path\to\file", nil)
	if err == nil {
		log.Fatal(err)
	} else {
		log.Println(content)
	}
}

```

## Optional Flags
---
`pattern` in all of this package except `FilterBytesstringWithCustomRegex` are optional and user can send it nil .
## Full Example
--- 
``` go 
package main

import (
	"log"

	"github.com/khodehamid/doc2text"
)

func main() {
// you can send your custom regex pattern to extract from file otherwise send nil .
	docContent, err := doc2text.GetDocContent("path\to\\doc\file.doc", nil)
	if err == nil {
		log.Fatal(err)

	} else {
		log.Println(docContent)
	}
//we want to extract our xml file content with out xml tags. 
	xmlContent, err := doc2text.GetDocContentWithXMLPath("path\to\\xml\file.xml", nil)
	if err == nil {
		log.Fatal(err)

	} else {
		log.Println(xmlContent)
	}
//we want to remove the non-word characters from our string . 
	tmepStringToFilter := "this is temp string with custom characters to filter@!@#!$#$#$%#$%^$ 123213"
	filteredString := doc2text.FilterBytesstringWithCustomRegex("[\\W]", []byte(tmepStringToFilter), " ")
	log.Println(filteredString)
}


```
