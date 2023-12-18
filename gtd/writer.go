package gtd

import "github.com/firesquid6/negtd/date"

func WriteAgendaFile(tasks *[]GtdTask) []string {
	return []string{}
}

func mapTasks(tasks *[]GtdTask) map[date.Date][]GtdTask {
	return map[date.Date][]GtdTask{}
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
