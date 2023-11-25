package main

import (
	"errors"
	"reflect"
)

func parseInboxTask(line string) (GtdTask, error) {
	// TODO
	task := GtdTask{}
	if line == "" {
		return task, errors.New("Empty line")
	}

	return task, nil
}

func tasksAreEqual(task1 GtdTask, task2 *GtdTask) bool {
	return reflect.DeepEqual(task1, task2)
}
