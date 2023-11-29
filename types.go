package main

// we're using our own Date implementation since Go's time packace includes random stuff about specific times we don't care about
type Date struct {
	year  int
	month int
	day   int
}

type Event struct {
	event string
	date  Date
}

type Weekday int

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
	Next
	Done
	Future
	Trash
	Projects
	Waiting
)

func emptyDate() Date {
	return Date{-1, -1, -1}
}
func emptyEvent() Event {
	return Event{"", emptyDate()}
}

type GtdFile struct {
	name  string
	lines []string
}

type GtdTask struct {
	text      string
	contexts  []string
	project   string
	gotoList  GtdListName
	waitingOn Event
}
