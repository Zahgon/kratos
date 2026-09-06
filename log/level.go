package log

import (
	"log/slog"
)

type Level = slog.Level

type Leveler = slog.Leveler

type LevelVar = slog.LevelVar

const LevelKey = slog.LevelKey

const (
	LevelDebug Level = slog.LevelDebug

	LevelInfo Level = slog.LevelInfo

	LevelWarn Level = slog.LevelWarn

	LevelError Level = slog.LevelError

	LevelFatal Level = slog.LevelError + 4
)

func ParseLevel(s string) Level { _ = "STUB: not implemented"; return *new(Level) }
