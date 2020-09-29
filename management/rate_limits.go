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
	reset, err := strconv.Atoi(h.Get(RateLimitReset))
	if err != nil {
		return nil, fmt.Errorf("error parsing x-ratelimit-reset header value: %s", err)
	}

	limit, err := strconv.Atoi(h.Get(RateLimitLimit))
	if err != nil {
		return nil, fmt.Errorf("error parsing x-ratelimit-limit header value: %s", err)
	}

	remaining, err := strconv.Atoi(h.Get(RateLimitRemaining))
	if err != nil {
		return nil, fmt.Errorf("error parsing x-ratelimit-remaining header value: %s", err)
	}

	return &RateLimit{
		Limit:     uint64(limit),
		Remaining: uint64(remaining),
		Reset:     time.Unix(int64(reset), 0),
	}, nil
}
