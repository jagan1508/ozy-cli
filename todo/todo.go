package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	Position int
	Done     bool
}

func SaveItems(filename string, items []Item) error {
	j, err := json.Marshal(items)
	if err != nil {
		return err
	}
	fmt.Println(string(j))
	err = os.WriteFile(filename, j, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(f, &items); err != nil {
		return []Item{}, err
	}
	for i, _ := range items {
		items[i].Position = i + 1
	}
	return items, nil
}

func (i *Item) SetPriority(priority int) {
	switch priority {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	} else if i.Priority == 3 {
		return "(3)"
	} else {
		return ""
	}
}

func (i *Item) Label() string {
	return strconv.Itoa(i.Position) + ". "
}

type ByPri []Item

func (s ByPri) Len() int {
	return len(s)
}

func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	if s[i].Priority != s[j].Priority {
		return s[i].Priority < s[j].Priority
	}
	return s[i].Position < s[j].Position
}
