package main

type GtdFile struct {
	name  string
	lines []string
}

type GtdInTask struct {
	text    string
	context string
}
