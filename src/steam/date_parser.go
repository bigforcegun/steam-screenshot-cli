package steam

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	timeFormat = "2 Jan 2006 9:57am"
	longForm   = "Jan 2, 2006 at 3:04pm (MST)"
)

// ParseSteamDate
//
//	i'm sorry - i do not know how to do it another way
func ParseSteamDate(dateString string) (time.Time, error) {

	var yearPart, monthPart, dayPart, timePart string

	dateParts := strings.Split(dateString, " ")
	if len(dateParts) < 4 {
		// FIXME - ну тут точно хуита а не проверка и ретурн кривой
		return time.Time{}, errors.New("bad string 🤷‍♂️")
	}
	// with yearPart == 8 Aug, 2023 @ 9:57pm
	if strings.Contains(dateString, ",") {
		yearPart = dateParts[2]
		monthPart = strings.Replace(dateParts[1], ",", "", -1)
		dayPart = dateParts[0]
		timePart = dateParts[4]
	} else { // without yearPart == 6 Jan @ 4:07am
		yearPart = strconv.Itoa(time.Now().Year())
		monthPart = dateParts[1]
		dayPart = dateParts[0]
		timePart = dateParts[3]
	}

	// parsableDateString := fmt.Sprintf("%v %v, %v at %v", monthPart,dayPart , yearPart, timePart)
	parsableDateString := fmt.Sprintf("%v %v, %v at %v (UTC)", monthPart, dayPart, yearPart, timePart)

	resultTime, err := time.Parse(longForm, parsableDateString)
	return resultTime, err
}
