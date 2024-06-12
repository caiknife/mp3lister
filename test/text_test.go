package test

import (
	"testing"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func TestCurrency(t *testing.T) {
	t1799, err := time.Parse(time.DateOnly, "1799-01-01")
	if err != nil {
		t.Error(err)
		return
	}

	for it := currency.Query(currency.Date(t1799)); it.Next(); {
		from := ""
		if t, ok := it.From(); ok {
			from = t.Format(time.DateOnly)
		}
		t.Logf("%v is used in %v since: %v\n", it.Unit(), it.Region(), from)
	}
}

func TestNumber(t *testing.T) {
	p := message.NewPrinter(language.English)
	_, _ = p.Printf("%v bottles of beer on the wall.\n", number.Decimal(1234))

	_, _ = p.Printf("%v of gophers lose too much fur.\n", number.Percent(0.12))

	p = message.NewPrinter(language.Dutch)
	_, _ = p.Printf("There are %v bikes per household.\n", number.Decimal(1.2))
}
