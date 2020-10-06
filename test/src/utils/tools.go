package utils

import (
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

// MustGet will return the env or panic if it is not present
func MustGet(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panic().Msgf("ENV missing, key: %s", k)
	}
	return v
}

// MustGetBool will return the env as boolean or panic if it is not present
func MustGetBool(k string) bool {
	v := os.Getenv(k)
	if v == "" {
		log.Panic().Msgf("ENV missing, key: %s", k)
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panic().Msgf("ENV key: %s - error: %q", k, err)
	}
	return b
}

// MustGetInt32 will return the env as int32 or panic if it is not present
func MustGetInt32(k string) int {
	v := os.Getenv(k)
	if v == "" {
		log.Panic().Msgf("ENV missing, key: %s", k)
	}
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		log.Panic().Msgf("ENV key: %s - error: %q", k, err)
	}
	return int(i)
}

// MustGetInt64 will return the env as int64 or panic if it is not present
func MustGetInt64(k string) int64 {
	v := os.Getenv(k)
	if v == "" {
		log.Panic().Msgf("ENV missing, key: %s", k)
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Panic().Msgf("ENV key: %s - error: %q", k, err)
	}
	return i
}
