package model

import (
	"context"
	"time"
)

type PeriodKey int

const (
	PeriodKey5     PeriodKey = 5
	PeriodKey10    PeriodKey = 10
	PeriodKey30    PeriodKey = 30
	PeriodKey60    PeriodKey = 60
	PeriodKey720   PeriodKey = 720   // 12 hours
	PeriodKey1440  PeriodKey = 1440  // 1 day
	PeriodKey2880  PeriodKey = 2880  // 2 days
	PeriodKey4320  PeriodKey = 4320  // 3 days
	PeriodKey5760  PeriodKey = 10080 // 1 week
	PeriodKey10080 PeriodKey = 20160 // 2 weeks
)

var periodMap = map[PeriodKey]string{
	PeriodKey5:     "5 minutes",
	PeriodKey10:    "10 minutes",
	PeriodKey30:    "30 minutes",
	PeriodKey60:    "1 hour",
	PeriodKey720:   "12 hours",
	PeriodKey1440:  "1 day",
	PeriodKey2880:  "2 days",
	PeriodKey4320:  "3 days",
	PeriodKey5760:  "1 week",
	PeriodKey10080: "2 weeks",
}

func PeriodMap() map[PeriodKey]string {
	return periodMap
}

var periodTimeMap = map[PeriodKey]time.Duration{
	PeriodKey5:     5 * time.Minute,
	PeriodKey10:    10 * time.Minute,
	PeriodKey30:    30 * time.Minute,
	PeriodKey60:    1 * time.Hour,
	PeriodKey720:   12 * time.Hour,
	PeriodKey1440:  24 * time.Hour,
	PeriodKey2880:  48 * time.Hour,
	PeriodKey4320:  72 * time.Hour,
	PeriodKey5760:  7 * 24 * time.Hour,
	PeriodKey10080: 14 * 24 * time.Hour,
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
