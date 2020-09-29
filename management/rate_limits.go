package management

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type RateLimitError struct {
	error
	RateLimit
}

type RateLimit struct {
	Limit     uint64
	Remaining uint64
	Reset     time.Time
}

func newRateLimit(h http.Header) (*RateLimit, error) {
	msg := "error parsing %s header value: %s"

	reset, err := strconv.Atoi(h.Get(RateLimitReset))
	if err != nil {
		return nil, fmt.Errorf(msg, RateLimitReset, err)
	}

	limit, err := strconv.Atoi(h.Get(RateLimitLimit))
	if err != nil {
		return nil, fmt.Errorf(msg, RateLimitLimit, err)
	}

	remaining, err := strconv.Atoi(h.Get(RateLimitRemaining))
	if err != nil {
		return nil, fmt.Errorf(msg, RateLimitRemaining, err)
	}

	return &RateLimit{
		Limit:     uint64(limit),
		Remaining: uint64(remaining),
		Reset:     time.Unix(int64(reset), 0),
	}, nil
}
