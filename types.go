package main

// we're using our own Date implementation since Go's time packace includes random stuff about specific times we don't care about
type Date struct {
	year  int
	month int
	day   int
}

type GtdListName int

const (
	Inbox GtdListName = iota
	Next
	Done
	Future
	Trash
	Waiting
)

type GtdFile struct {
	name  string
	lines []string
}

type GtdTask struct {
	text      string
	contexts  []string
	project   string
	gotoList  GtdListName
	date      Date
	waitingOn string
}
