package types

import (
	"time"

	"github.com/golang-module/carbon/v2"
)

type DateTime struct {
	carbon.Carbon
}

func NewDateTime(tt ...string) *DateTime {
	now := carbon.Now()
	if len(tt) != 0 {
		now = carbon.Parse(tt[0])
	}
	return &DateTime{
		Carbon: now,
	}
}

func (d *DateTime) String() string {
	return d.ToIso8601MicroString()
}

func (d *DateTime) MinuteStartTime() time.Time {
	return d.StartOfMinute().StdTime()
}

func (d *DateTime) MinuteEndTime() time.Time {
	return d.EndOfMinute().StdTime()
}

func (d *DateTime) StartOfTenMinutes() time.Time {
	minute := d.Minute() / 10 * 10
	return d.SetMinute(minute).SetSecond(0).SetNanosecond(0).StdTime()
}

func (d *DateTime) EndOfTenMinutes() time.Time {
	minute := (d.Minute()+10)/10*10 - 1
	return d.SetMinute(minute).SetSecond(59).SetNanosecond(999999999).StdTime()
}

func (d *DateTime) HourStartTime() time.Time {
	return d.StartOfHour().StdTime()
}

func (d *DateTime) HourEndTime() time.Time {
	return d.EndOfHour().StdTime()
}

func (d *DateTime) DayStartTime() time.Time {
	return d.StartOfDay().StdTime()
}

func (d *DateTime) DayEndTime() time.Time {
	return d.EndOfDay().StdTime()
}

func (d *DateTime) WeekStartTime() time.Time {
	return d.StartOfWeek().StdTime()
}

func (d *DateTime) WeekEndTime() time.Time {
	return d.EndOfWeek().StdTime()
}

func (d *DateTime) MonthStartTime() time.Time {
	return d.StartOfMonth().StdTime()
}

func (d *DateTime) MonthEndTime() time.Time {
	return d.EndOfMonth().StdTime()
}

func (d *DateTime) QuarterStartTime() time.Time {
	return d.StartOfQuarter().StdTime()
}

func (d *DateTime) QuarterEndTime() time.Time {
	return d.EndOfQuarter().StdTime()
}

func (d *DateTime) SeasonStartTime() time.Time {
	return d.StartOfSeason().StdTime()
}

func (d *DateTime) SeasonEndTime() time.Time {
	return d.EndOfSeason().StdTime()
}

func (d *DateTime) YearStartTime() time.Time {
	return d.StartOfYear().StdTime()
}

func (d *DateTime) YearEndTime() time.Time {
	return d.EndOfYear().StdTime()
}

func (d *DateTime) DecadeStartTime() time.Time {
	return d.StartOfDecade().StdTime()
}

func (d *DateTime) DecadeEndTime() time.Time {
	return d.EndOfDecade().StdTime()
}

func (d *DateTime) CenturyStartTime() time.Time {
	return d.StartOfCentury().StdTime()
}

func (d *DateTime) CenturyEndTime() time.Time {
	return d.EndOfCentury().StdTime()
}

func (d *DateTime) Christmas() time.Time {
	return d.SetDateTimeNano(d.Carbon.Year(), 12, 25, 0, 0, 0, 0).StdTime()
}

func (d *DateTime) FirstWeekdayOfMonth(weekday time.Weekday) time.Time {
	firstDayOfMonth := carbon.CreateFromStdTime(d.MonthStartTime())
	i := 0
	for firstDayOfMonth.AddDays(i).StdTime().Weekday() != weekday {
		i++
	}
	return firstDayOfMonth.AddDays(i).StdTime()
}

func (d *DateTime) LastWeekdayOfMonth(weekday time.Weekday) time.Time {
	lastDayOfMonth := carbon.CreateFromStdTime(d.MonthEndTime())
	i := 0
	for lastDayOfMonth.SubDays(i).StdTime().Weekday() != weekday {
		i++
	}
	return lastDayOfMonth.SubDays(i).StartOfDay().StdTime()
}

func (d *DateTime) FirstWeekdayOfYear(weekday time.Weekday) time.Time {
	firstDayOfYear := carbon.CreateFromStdTime(d.YearStartTime())
	i := 0
	for firstDayOfYear.AddDays(i).StdTime().Weekday() != weekday {
		i++
	}
	return firstDayOfYear.AddDays(i).StdTime()
}

func (d *DateTime) LastWeekdayOfYear(weekday time.Weekday) time.Time {
	lastDayOfYear := carbon.CreateFromStdTime(d.YearEndTime())
	i := 0
	for lastDayOfYear.SubDays(i).StdTime().Weekday() != weekday {
		i++
	}
	return lastDayOfYear.SubDays(i).StartOfDay().StdTime()
}

func (d *DateTime) FirstSundayOfMonth() time.Time {
	return d.FirstWeekdayOfMonth(time.Sunday)
}

func (d *DateTime) LastSundayOfMonth() time.Time {
	return d.LastWeekdayOfMonth(time.Sunday)
}

func (d *DateTime) FirstSundayOfYear() time.Time {
	return d.FirstWeekdayOfYear(time.Sunday)
}

func (d *DateTime) LastSundayOfYear() time.Time {
	return d.LastWeekdayOfYear(time.Sunday)
}

func (d *DateTime) TheDayIsFirstSundayOfMonth() bool {
	return d.FirstSundayOfMonth().Equal(d.DayStartTime())
}
