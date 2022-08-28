package markdown

import (
	"fmt"
	"strings"
)

const (
	listPrefix              = "- "
	numListPrefix           = "1. "
	taskListPrefix          = "- [ ]"
	completedTaskListPrefix = "- [x]"
)

type listItem any

func RootList(items ...listItem) string {
	return strings.Join(List(items...), "\n")
}

func RootNumList(items ...listItem) string {
	return strings.Join(NumList(items...), "\n")
}

type listItems []string

func List(items ...listItem) listItems {
	return list(listPrefix, items...)
}

func NumList(items ...listItem) listItems {
	return list(numListPrefix, items...)
}

func list(prefix string, items ...listItem) listItems {
	list := make(listItems, 0, len(items))
	for _, item := range items {
		switch item := item.(type) {
		case listItems:
			for _, v := range item {
				list = append(list, "     "+v)
			}
		case string:
			list = append(list, prefix+item)
		default:
			list = append(list, prefix+fmt.Sprint(item))
		}
	}

	return list
}

type taskItem struct {
	completed bool
	item      listItem
}

func Task(completed bool, item string) taskItem {
	return taskItem{
		completed: completed,
		item:      item,
	}
}

func RootTaskList(tasks ...taskItem) string {
	v, _ := TaskList(tasks...).item.(listItems)
	return strings.Join(v, "\n")
}

func TaskList(tasks ...taskItem) taskItem {
	items := make([]listItem, 0, len(tasks))
	for _, task := range tasks {
		items = append(items, task.item)
	}

	list := make(listItems, 0, len(items))
	for i, item := range items {
		prefix := taskListPrefix
		if tasks[i].completed {
			prefix = completedTaskListPrefix
		}

		switch item := item.(type) {
		case listItems:
			for _, v := range item {
				list = append(list, "     "+v)
			}
		case string:
			list = append(list, prefix+item)
		default:
			list = append(list, prefix+fmt.Sprint(item))
		}
	}

	return taskItem{item: list}
}
