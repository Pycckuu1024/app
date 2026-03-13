package main

import (
	"fmt"

	"github.com/Pycckuu1024/app/app/calendar"
	"github.com/Pycckuu1024/app/app/events"
)

func main() {
	e, err := events.NewEvent("Встреча", "2025/06/12 16:33")
	b, err := events.NewEvent("Новый Год", "2025/12/31 00:00")
	c, err := events.NewEvent("1 января", "2026/01/01 00:01")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return

	}
	calendar.AddEvent("event1", e)
	calendar.AddEvent("event2", b)
	calendar.AddEvent("event3", c)

	calendar.ShowEvents()
}
