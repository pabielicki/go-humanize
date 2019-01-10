package humanize

import (
	// "fmt"
	"time"
	"testing"
)

// func EsxampleTurkish() {
// 	SetLanguage(Turkish)
// 	fmt.Println(LocalTime(time.Now()))
// 	SetLanguage(English)
// 	// Output: şimdi
// }


func TestLocal(t *testing.T) {
	LoadLanguages("pl_PL", "locals/en_US.json")
	now := time.Now()
	l := "pl_PL"
	testList{
		{"now", LocalTime(now, l), "przed chwilą"},
		{"now", LocalTime(now, "en_US"), "now"},
		{"1 second ago", LocalTime(now.Add(-1 * time.Second), l), "1 sekunda temu"},
		{"12 seconds ago", LocalTime(now.Add(-12 * time.Second), l), "12 sekund temu"},
		{"30 seconds ago", LocalTime(now.Add(-30 * time.Second), l), "30 sekund temu"},
		{"45 seconds ago", LocalTime(now.Add(-45 * time.Second), l), "45 sekund temu"},
		{"1 minute ago", LocalTime(now.Add(-63 * time.Second), l), "1 minuta temu"},
		{"15 minutes ago", LocalTime(now.Add(-15 * time.Minute), l), "15 minut temu"},
		{"1 hour ago", LocalTime(now.Add(-63 * time.Minute), l), "1 godzina temu"},
		{"2 hours ago", LocalTime(now.Add(-2 * time.Hour), l), "2 godzin temu"},
		{"21 hours ago", LocalTime(now.Add(-21 * time.Hour), l), "21 godzin temu"},
		{"1 day ago", LocalTime(now.Add(-26 * time.Hour), l), "1 dzień temu"},
		{"2 days ago", LocalTime(now.Add(-49 * time.Hour), l), "2 dni temu"},
		{"3 days ago", LocalTime(now.Add(-3 * Day), l), "3 dni temu"},
		{"1 week ago ", LocalTime(now.Add(-7 * Day), l), "1 tydzień temu"},
		{"2 weeks ago", LocalTime(now.Add(-15 * Day), l), "2 tygodnie temu"},
		{"1 month ago", LocalTime(now.Add(-39 * Day), l), "1 miesiąc temu"},
		{"3 months ago", LocalTime(now.Add(-99 * Day), l), "3 miesięcy temu"},
		{"1 year ago ", LocalTime(now.Add(-365 * Day), l), "1 rok temu"},
		{"2 years ago ", LocalTime(now.Add(-548 * Day), l), "2 lat temu"},
		{"long ago", LocalTime(now.Add(-LongTime), l), "dawno temu"},
	}.validate(t)
}
