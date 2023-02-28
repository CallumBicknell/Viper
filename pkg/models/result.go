package models

type Result struct {
	Links      []string
	Subdomains []string
	TotalLinks int
	OpenPorts  []Port
	Os         string
	OpenDirs   []OpenDir
}
