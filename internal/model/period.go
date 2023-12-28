package model

import (
	"context"
	"time"
)

type PeriodKey int

var PeriodMap = map[PeriodKey]string{
	5:  "5 minutes",
	10: "10 minutes",
	30: "30 minutes",
	60: "1 hour",
}

var PeriodTimeMap = map[PeriodKey]time.Duration{
	5:  5 * time.Minute,
	10: 10 * time.Minute,
	30: 30 * time.Minute,
	60: 1 * time.Hour,
}

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (r *RealClock) Now() time.Time {
	return time.Now()
}

type clockCtxKey struct{}

func SetClock(ctx context.Context, c Clock) context.Context {
	return context.WithValue(ctx, clockCtxKey{}, c)
}

func GetClock(ctx context.Context) Clock {
	clock, ok := ctx.Value(clockCtxKey{}).(Clock)
	if ok {
		return clock
	}

	return &RealClock{}
}

type FrozenClock struct {
	Time time.Time
}

func (f *FrozenClock) Now() time.Time {
	return f.Time
}
