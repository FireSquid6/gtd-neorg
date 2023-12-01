package main

import (
	//"fmt"
	"reflect"
	"regexp"
	"strings"
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
	date     Date
}

func parseInboxPredata(predata string, currentDate Date) (ParsedPredata, error) {
	parsed := ParsedPredata{
		gotoList: Inbox,
		date:     emptyDate(),
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
		date, err := parseRelativeDate(predata, currentDate)
		if err != nil {
			return ParsedPredata{}, err
		}
		parsed.gotoList = Agenda
		parsed.date = date
	}

	return parsed, nil
}

// INBOX SPECIFIC PARSING
func parseInboxTask(line string, currentDate Date) (GtdTask, error) {
	task := GtdTask{
		text:     "",
		tags:     []string{},
		date:     emptyDate(),
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
