//by days solution

//.go

func ByDaysPeriod(period DaysPeriod) func(Record) bool {
	return func(rec Record) bool {
		return period.From <= rec.Day && rec.Day <= period.To
	}
}

func ByDaysPeriod(pp DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= pp.From && r.Day <= pp.To
	}
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		return p.From <= record.Day && record.Day <= p.To
	}
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return p.From <= r.Day && r.Day <= p.To
	}
	// panic("Please implement the ByDaysPeriod function")
}