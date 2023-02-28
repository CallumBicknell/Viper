package models

type dirContents struct {
	Name   string
	Edited string
}

type OpenDir struct {
	Path     string
	Contents []dirContents
}
