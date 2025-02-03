package logger

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func Info(msg string) {
	log.Info().Msg(msg)
}

func Error(err error) {
	log.Error().Stack().Err(err).Msg("")
}

func Debug(msg any) {
	message := fmt.Sprint(msg)
	log.Debug().Msg(message)
}
