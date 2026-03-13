package calendar

import (
	"github.com/Pycckuu1024/app/app/events"

	"fmt"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(key string, e events.Event) {
	eventsMap[key] = e

}

func ShowEvents() {
	for _, v := range eventsMap {
		fmt.Printf("\n Событие : %s.  Дата : %v.\n", v.Title, v.StartAt)
	}
}
