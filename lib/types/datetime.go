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
	return d.StartOfMinute().ToStdTime()
}

func (d *DateTime) MinuteEndTime() time.Time {
	return d.EndOfMinute().ToStdTime()
}

func (d *DateTime) StartOfTenMinutes() time.Time {
	minute := d.Minute() / 10 * 10
	return d.SetMinute(minute).SetSecond(0).SetNanosecond(0).ToStdTime()
}

func (d *DateTime) EndOfTenMinutes() time.Time {
	minute := (d.Minute()+10)/10*10 - 1
	return d.SetMinute(minute).SetSecond(59).SetNanosecond(999999999).ToStdTime()
}

func (d *DateTime) HourStartTime() time.Time {
	return d.StartOfHour().ToStdTime()
}

func (d *DateTime) HourEndTime() time.Time {
	return d.EndOfHour().ToStdTime()
}

func (d *DateTime) DayStartTime() time.Time {
	return d.StartOfDay().ToStdTime()
}

func (d *DateTime) DayEndTime() time.Time {
	return d.EndOfDay().ToStdTime()
}

func (d *DateTime) WeekStartTime() time.Time {
	return d.StartOfWeek().ToStdTime()
}

func (d *DateTime) WeekEndTime() time.Time {
	return d.EndOfWeek().ToStdTime()
}

func (d *DateTime) MonthStartTime() time.Time {
	return d.StartOfMonth().ToStdTime()
}

func (d *DateTime) MonthEndTime() time.Time {
	return d.EndOfMonth().ToStdTime()
}

func (d *DateTime) QuarterStartTime() time.Time {
	return d.StartOfQuarter().ToStdTime()
}

func (d *DateTime) QuarterEndTime() time.Time {
	return d.EndOfQuarter().ToStdTime()
}

func (d *DateTime) SeasonStartTime() time.Time {
	return d.StartOfSeason().ToStdTime()
}

func (d *DateTime) SeasonEndTime() time.Time {
	return d.EndOfSeason().ToStdTime()
}

func (d *DateTime) YearStartTime() time.Time {
	return d.StartOfYear().ToStdTime()
}

func (d *DateTime) YearEndTime() time.Time {
	return d.EndOfYear().ToStdTime()
}

func (d *DateTime) DecadeStartTime() time.Time {
	return d.StartOfDecade().ToStdTime()
}

func (d *DateTime) DecadeEndTime() time.Time {
	return d.EndOfDecade().ToStdTime()
}

func (d *DateTime) CenturyStartTime() time.Time {
	return d.StartOfCentury().ToStdTime()
}

func (d *DateTime) CenturyEndTime() time.Time {
	return d.EndOfCentury().ToStdTime()
}

func (d *DateTime) Christmas() time.Time {
	return d.SetDateTimeNano(d.Carbon.Year(), 12, 25, 0, 0, 0, 0).ToStdTime()
}

func (d *DateTime) FirstWeekdayOfMonth(weekday time.Weekday) time.Time {
	firstDayOfMonth := carbon.CreateFromStdTime(d.MonthStartTime())
	i := 0
	for firstDayOfMonth.AddDays(i).ToStdTime().Weekday() != weekday {
		i++
	}
	return firstDayOfMonth.AddDays(i).ToStdTime()
}

func (d *DateTime) LastWeekdayOfMonth(weekday time.Weekday) time.Time {
	lastDayOfMonth := carbon.CreateFromStdTime(d.MonthEndTime())
	i := 0
	for lastDayOfMonth.SubDays(i).ToStdTime().Weekday() != weekday {
		i++
	}
	return lastDayOfMonth.SubDays(i).StartOfDay().ToStdTime()
}

func (d *DateTime) FirstWeekdayOfYear(weekday time.Weekday) time.Time {
	firstDayOfYear := carbon.CreateFromStdTime(d.YearStartTime())
	i := 0
	for firstDayOfYear.AddDays(i).ToStdTime().Weekday() != weekday {
		i++
	}
	return firstDayOfYear.AddDays(i).ToStdTime()
}

func (d *DateTime) LastWeekdayOfYear(weekday time.Weekday) time.Time {
	lastDayOfYear := carbon.CreateFromStdTime(d.YearEndTime())
	i := 0
	for lastDayOfYear.SubDays(i).ToStdTime().Weekday() != weekday {
		i++
	}
	return lastDayOfYear.SubDays(i).StartOfDay().ToStdTime()
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
