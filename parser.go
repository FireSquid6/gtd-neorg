package main

import (
	//"fmt"
	"reflect"
	"regexp"
	"strings"
)

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

func parseInboxTask(line string, currentDate Date) (GtdTask, error) {
	task := GtdTask{
		text:     "",
		tags:     []string{},
		date:     emptyDate(),
		gotoList: Inbox,
	}

	split := splitInboxLine(line)
	task.text = split.text

	switch split.predata {
	case "":
		task.gotoList = Inbox
	case "_":
		task.gotoList = Trash
	case "?":
		task.gotoList = Backlog
	default:
		date, err := parseRelativeDate(split.predata, currentDate)
		if err != nil {
			return GtdTask{}, err
		}
		task.gotoList = Agenda
		task.date = date
	}

	tagSplit := strings.Split(split.postdata, ",")
	for _, tag := range tagSplit {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			task.tags = append(task.tags, tag)
		}
	}

	return task, nil
}

func tasksAreEqual(task1 GtdTask, task2 GtdTask) bool {
	return reflect.DeepEqual(task1, task2)
}
