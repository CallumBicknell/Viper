package main

import "github.com/sirupsen/logrus"

func main() {
	url := "google.com"
	log := logrus.New()
	log.Infof("Found opendir at `%s`\n", url)
}
