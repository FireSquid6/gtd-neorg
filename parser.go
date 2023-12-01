package main

import (
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

func parseInboxTask(line string) (GtdTask, error) {
	return GtdTask{
		text:     line,
		tags:     []string{},
		gotoList: Inbox,
		date:     emptyDate(),
	}, nil
}

func tasksAreEqual(task1 GtdTask, task2 GtdTask) bool {
	return reflect.DeepEqual(task1, task2)
}
