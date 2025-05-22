package calendar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Holiday struct {
	Date string `json:"holiday_date"`
	Name string `json:"holiday_name"`
}

func fetchHolidays(month time.Month, year int) ([]Holiday, error) {
	url := fmt.Sprintf("https://api-harilibur.vercel.app/api?month=%d&year=%d", int(month), year)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var holidays []Holiday
	if err := json.NewDecoder(resp.Body).Decode(&holidays); err != nil {
		return nil, err
	}
	return holidays, nil
}

func NewHolidayCalendarWidget(month time.Month, year int) *widgets.QCalendarWidget {
	calendar := widgets.NewQCalendarWidget(nil)
	calendar.SetGridVisible(true)

	format := gui.NewQTextCharFormat()
	format.SetFontWeight(int(gui.QFont__Bold))
	format.SetForeground(gui.NewQBrush3(gui.NewQColor3(255, 0, 0, 255), core.Qt__SolidPattern))

	holidays, err := fetchHolidays(month, year)
	if err == nil {
		for _, h := range holidays {
			date, err := time.Parse("2006-01-02", h.Date)
			if err == nil {
				qdate := core.NewQDate3(date.Year(), int(date.Month()), date.Day())
				calendar.SetDateTextFormat(qdate, format)
			}
		}
	}

	calendar.SetWindowTitle(fmt.Sprintf("Indonesian Holidays - %d/%d", month, year))
	calendar.SetMinimumSize2(400, 300)

	return calendar
}
