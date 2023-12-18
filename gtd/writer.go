package gtd

import (
	"sort"

	"github.com/firesquid6/negtd/date"
)

func WriteAgendaFile(tasks *[]GtdTask) []string {
	dateMap := mapTasks(tasks)
	file := []string{}
	keys := make([]date.Date, 0, len(dateMap))
	for d := range dateMap {
		keys = append(keys, d)
	}

	keys = sortDates(keys)

	for _, key := range keys {
		sectionOutput := outputAgendaSection(key, dateMap[key])
		file = append(file, sectionOutput...)
		file = append(file, "")
	}

	return file
}

func sortDates(dates []date.Date) []date.Date {
	sort.Slice(dates, func(i, j int) bool {
		if dates[i].Year != dates[j].Year {
			return dates[i].Year < dates[j].Year
		} else if dates[i].Month != dates[j].Month {
			return dates[i].Month < dates[j].Month
		} else {
			return dates[i].Day < dates[j].Day
		}
	})

	return dates
}

func outputAgendaSection(sectionDate date.Date, tasks []GtdTask) []string {
	section := []string{}
	section = append(section, "* "+date.DateToString(sectionDate))
	for _, task := range tasks {
		section = append(section, " - ( ) "+task.Text)
	}

	return section
}

func mapTasks(tasks *[]GtdTask) map[date.Date][]GtdTask {
	dateMap := map[date.Date][]GtdTask{}
	for _, task := range *tasks {
		if task.Date != date.EmptyDate() {
			if _, ok := dateMap[task.Date]; !ok {
				dateMap[task.Date] = []GtdTask{}
			}
			dateMap[task.Date] = append(dateMap[task.Date], task)
		}
	}

	return dateMap
}

func WriteInboxFile(tasks *[]GtdTask) []string {
	return writeInboxTypeFile(tasks, Inbox)
}

func WriteBacklogFile(tasks *[]GtdTask) []string {
	return writeInboxTypeFile(tasks, Backlog)
}

func writeInboxTypeFile(tasks *[]GtdTask, list GtdListName) []string {
	file := []string{}
	for _, task := range *tasks {
		if task.GotoList == list {
			file = append(file, task.Text)
		}
	}

	return file
}
