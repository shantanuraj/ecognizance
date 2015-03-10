package main

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"strings"

	"github.com/russross/blackfriday"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func ProcessEdition(e *Edition) {
	e.Editorial.Gravatar = GetMD5Hash(e.Editorial.Email)
	e.Featured.Gravatar = GetMD5Hash(e.Featured.Email)

	for i := range e.Articles {
		e.Articles[i].Gravatar = GetMD5Hash(e.Articles[i].Email)
	}
}

func GetHTML(post string) template.HTML {
	return template.HTML(blackfriday.MarkdownBasic([]byte(post)))
}

func GetMD5Hash(text string) string {
	text = strings.TrimSpace(text)
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
