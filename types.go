package main

// we're using our own Date implementation since Go's time packace includes random stuff about specific times we don't care about
type Date struct {
	year  int
	month int
	day   int
}

type Weekday int

type Tag struct {
	name    string
	tagType TagType
}

type TagType int

const (
	Project = iota
	General
	Who
	Where
)

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type GtdListName int

const (
	Inbox GtdListName = iota
	Agenda
	Backlog
	Trash
	Waiting
)

func emptyDate() Date {
	return Date{-1, -1, -1}
}

type GtdFile struct {
	name  string
	lines []string
}

type GtdTask struct {
	text     string
	contexts []string
	project  string
	gotoList GtdListName
	date     Date
}

type GtdEvent struct{}
