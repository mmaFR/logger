package logger

import (
	"fmt"
	"io"
	"log"
)

const (
	LogLevelNull LogLevel = iota
	LogLevelEmerge
	LogLevelAlert
	LogLevelCritical
	LogLevelError
	LogLevelWarning
	LogLevelNotice
	LogLevelInfo
	LogLevelDebug
	LogLevelTrace
)

const (
	logLevelEmergePrefix   string = "[EMERGE  ]"
	logLevelAlertPrefix    string = "[ALTER   ]"
	logLevelCriticalPrefix string = "[CRITICAL]"
	logLevelErrorPrefix    string = "[ERROR   ]"
	logLevelWarningPrefix  string = "[WARNING ]"
	logLevelNoticePrefix   string = "[NOTICE  ]"
	logLevelInfoPrefix     string = "[INFO    ]"
	logLevelDebugPrefix    string = "[DEBUG   ]"
	logLevelTracePrefix    string = "[TRACE   ]"
)

const (
	logPattern       string = "%s %s -> %s: %s\n"
	logPatternWithId string = "%s %s -> %s-%d: %s\n"
)

type LogLevel uint8
type logFunc func(structure, function, msg string, id int, vars ...any)

type Logger struct {
	logger *log.Logger
	// Log
	LogEmerge   logFunc
	LogAlert    logFunc
	LogCritical logFunc
	LogError    logFunc
	LogWarning  logFunc
	LogNotice   logFunc
	LogInfo     logFunc
	LogDebug    logFunc
	LogTrace    logFunc

	//Fatal
	FatalEmerge   logFunc
	FatalAlert    logFunc
	FatalCritical logFunc
	FatalError    logFunc
	FatalWarning  logFunc
	FatalNotice   logFunc
	FatalInfo     logFunc
	FatalDebug    logFunc
	FatalTrace    logFunc
}

func (l *Logger) SetVerbosity(level LogLevel) {
	for lvl := LogLevelNull; lvl <= LogLevelTrace; lvl++ {
		switch lvl {
		case LogLevelNull:
			l.LogEmerge = func(structure, function, msg string, id int, vars ...any) {}
			l.LogAlert = func(structure, function, msg string, id int, vars ...any) {}
			l.LogCritical = func(structure, function, msg string, id int, vars ...any) {}
			l.LogError = func(structure, function, msg string, id int, vars ...any) {}
			l.LogWarning = func(structure, function, msg string, id int, vars ...any) {}
			l.LogNotice = func(structure, function, msg string, id int, vars ...any) {}
			l.LogInfo = func(structure, function, msg string, id int, vars ...any) {}
			l.LogDebug = func(structure, function, msg string, id int, vars ...any) {}
			l.LogTrace = func(structure, function, msg string, id int, vars ...any) {}

			l.FatalEmerge = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelEmergePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelEmergePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.FatalAlert = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelAlertPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelAlertPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.FatalCritical = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelCriticalPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelCriticalPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.FatalError = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelErrorPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelErrorPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.FatalWarning = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelWarningPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelWarningPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.FatalNotice = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelNoticePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelNoticePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.FatalInfo = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelInfoPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelInfoPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.FatalDebug = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelDebugPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelDebugPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.FatalTrace = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelTracePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelTracePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelEmerge:
			if lvl <= level {
				l.LogEmerge = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelEmergePrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelEmergePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.LogEmerge = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.FatalEmerge = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelEmergePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelEmergePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelAlert:
			if lvl <= level {
				l.LogAlert = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelAlertPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelAlertPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.LogAlert = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.FatalAlert = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelAlertPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelAlertPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelCritical:
			if lvl <= level {
				l.LogCritical = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelCriticalPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelCriticalPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.LogCritical = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.FatalCritical = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelCriticalPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelCriticalPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelError:
			if lvl <= level {
				l.LogError = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelErrorPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelErrorPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.LogError = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.FatalError = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelErrorPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelErrorPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelWarning:
			if lvl <= level {
				l.LogWarning = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelWarningPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelWarningPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.LogWarning = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.FatalWarning = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelWarningPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelWarningPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelNotice:
			if lvl <= level {
				l.LogNotice = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelNoticePrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelNoticePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.LogNotice = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.FatalNotice = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelNoticePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelNoticePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelInfo:
			if lvl <= level {
				l.LogInfo = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelInfoPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelInfoPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.LogInfo = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.FatalInfo = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelInfoPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelInfoPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelDebug:
			if lvl <= level {
				l.LogDebug = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelDebugPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelDebugPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.LogDebug = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.FatalDebug = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelDebugPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelDebugPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelTrace:
			if lvl <= level {
				l.LogTrace = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelTracePrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelTracePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.LogTrace = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.FatalTrace = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelTracePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelTracePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		}
	}
}

func NewLogger(level LogLevel, dst io.Writer) *Logger {
	var l *Logger = new(Logger)
	l.logger = log.New(dst, "", log.Ldate|log.Lmicroseconds)
	l.SetVerbosity(level)
	return l
}
