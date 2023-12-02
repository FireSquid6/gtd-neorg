package parser

import (
	//"fmt"
	"github.com/firesquid6/negtd/date"
	"reflect"
	"regexp"
	"strings"
)

type GtdTask struct {
	text     string
	tags     []string
	gotoList GtdListName
	date     date.Date
}

type GtdListName int

const (
	Inbox GtdListName = iota
	Agenda
	Backlog
	Trash
	Waiting
)

func parseTags(postdata string) []string {
	tags := []string{}
	tagSplit := strings.Split(postdata, ",")

	for _, tag := range tagSplit {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			tags = append(tags, tag)
		}
	}

	return tags
}

// SPLITTING LINES
type splitLine struct {
	predata  string
	text     string
	postdata string
}

func splitInboxLine(line string) (splits splitLine) {
	re := regexp.MustCompile(`\[(.*?)\]`)
	matches := re.FindAllStringSubmatch(line, -1)

	predata := ""
	postdata := ""
	text := line

	if len(matches) > 0 {
		if strings.HasPrefix(line, matches[0][0]) {
			predata = matches[0][1]
			text = strings.TrimPrefix(line, matches[0][0])
		} else {
			postdata = matches[0][1]
			text = strings.TrimSuffix(line, matches[0][0])
		}
		text = strings.TrimSpace(text)
	}

	return splitLine{
		predata:  predata,
		text:     text,
		postdata: postdata,
	}
}

// PARSING PREDATA
type ParsedPredata struct {
	gotoList GtdListName
	date     date.Date
}

func parseInboxPredata(predata string, currentDate date.Date) (ParsedPredata, error) {
	parsed := ParsedPredata{
		gotoList: Inbox,
		date:     date.EmptyDate(),
	}

	if predata == "" {
		return parsed, nil
	}

	switch predata {
	case "_":
		parsed.gotoList = Trash
	case "?":
		parsed.gotoList = Backlog
	default:
		date, err := date.ParseRelativeDate(predata, currentDate)
		if err != nil {
			return ParsedPredata{}, err
		}
		parsed.gotoList = Agenda
		parsed.date = date
	}

	return parsed, nil
}

// INBOX SPECIFIC PARSING
func parseInboxTask(line string, currentDate date.Date) (GtdTask, error) {
	task := GtdTask{
		text:     "",
		tags:     []string{},
		date:     date.EmptyDate(),
		gotoList: Inbox,
	}

	split := splitInboxLine(line)
	task.text = split.text

	parsedPredata, err := parseInboxPredata(split.predata, currentDate)
	if err != nil {
		return GtdTask{}, err
	}
	task.gotoList = parsedPredata.gotoList
	task.date = parsedPredata.date

	task.tags = parseTags(split.postdata)

	return task, nil
}

func tasksAreEqual(task1 GtdTask, task2 GtdTask) bool {
	return reflect.DeepEqual(task1, task2)
}
