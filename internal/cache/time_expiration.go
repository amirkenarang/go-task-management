package cache

import (
	"time"
)

type Expiration time.Duration

// Duration constants
const (
	ONE_MINUTE    = 1 * time.Minute
	TEN_MINUTE    = 10 * time.Minute
	THIRTY_MINUTE = 30 * time.Minute
	ONE_HOUR      = 1 * time.Hour
	TWO_HOUR      = 2 * time.Hour
	THREE_HOUR    = 3 * time.Hour
	FOUR_HOUR     = 4 * time.Hour
	FIVE_HOUR     = 5 * time.Hour
	SIX_HOUR      = 6 * time.Hour
	SEVEN_HOUR    = 7 * time.Hour
	EIGHTH_HOUR   = 8 * time.Hour
	NINE_HOUR     = 9 * time.Hour
	TEN_HOUR      = 10 * time.Hour
	ELEVEN_HOUR   = 11 * time.Hour
	TWELVE_HOUR   = 12 * time.Hour
	ONE_DAY       = 24 * time.Hour
	THREE_DAY     = 3 * ONE_DAY
	ONE_WEEK      = 7 * ONE_DAY
	ONE_MOUNTH    = 30 * ONE_DAY
)
