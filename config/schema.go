package config

var Conf schema

type schema struct {
	Document Document `yaml:"document"`
	Dir      Dir      `yaml:"dir"`
	Pattern  Pattern  `yaml:"pattern"`
}

type Dir struct {
	Html   string `yaml:"html"`
	Target string `yaml:"target"`
}

type Pattern struct {
	Title   string `yaml:"title"`
	Date    string `yaml:"date"`
	Content string `yaml:"content"`
	Tags    string `yaml:"tags"`
}

type Document struct {
	Draft string `yaml:draft`
}
