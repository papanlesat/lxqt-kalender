package main

import (
	"os"
	"time"

	"github.com/papanlesat/lxqt-kalender/calendar"
	"github.com/papanlesat/lxqt-kalender/lxqtsettings"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Apply LXQt settings globally
	settings := lxqtsettings.NewApplier()
	settings.Apply()

	// Optional: start a watcher to reapply settings on change
	go settings.Watch()

	// Show holiday calendar
	calendarWidget := calendar.NewHolidayCalendarWidget(time.Now().Month(), time.Now().Year())
	calendarWidget.Show()

	widgets.QApplication_Exec()
	app.Exec()

}
