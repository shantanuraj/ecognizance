package main

import "html/template"

type Article struct {
	Title  string
	Author string
	Email  string
	Post   string

	//Derived fields
	Gravatar string        `schema:"-"`
	PostHTML template.HTML `schema:"-"`
}

type Ecohustle struct {
	Image   string
	Caption string
}

type Edition struct {
	Title string
	Cover string

	Editorial Article
	Featured  Article
	Articles  []Article
	Ecohustle []Ecohustle
}
