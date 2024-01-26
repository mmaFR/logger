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
	logLevelEmergeName   string = "EMERGE"
	logLevelAlertName    string = "ALERT"
	logLevelCriticalName string = "CRITICAL"
	logLevelErrorName    string = "ERROR"
	logLevelWarningName  string = "WARNING"
	logLevelNoticeName   string = "NOTICE"
	logLevelInfoName     string = "INFO"
	logLevelDebugName    string = "DEBUG"
	logLevelTraceName    string = "TRACE"
)

const (
	logLevelEmergePrefix   string = "[" + logLevelEmergeName + "  ]"
	logLevelAlertPrefix    string = "[" + logLevelAlertName + "   ]"
	logLevelCriticalPrefix string = "[" + logLevelCriticalName + "]"
	logLevelErrorPrefix    string = "[" + logLevelErrorName + "   ]"
	logLevelWarningPrefix  string = "[" + logLevelWarningName + " ]"
	logLevelNoticePrefix   string = "[" + logLevelNoticeName + "  ]"
	logLevelInfoPrefix     string = "[" + logLevelInfoName + "    ]"
	logLevelDebugPrefix    string = "[" + logLevelDebugName + "   ]"
	logLevelTracePrefix    string = "[" + logLevelTraceName + "   ]"
)

const (
	logPattern       string = "%s %s -> %s: %s\n"
	logPatternWithId string = "%s %s -> %s-%d: %s\n"
)

var levelMap map[LogLevel]string = map[LogLevel]string{
	LogLevelNull:     "NONE",
	LogLevelEmerge:   logLevelEmergeName,
	LogLevelAlert:    logLevelAlertName,
	LogLevelCritical: logLevelCriticalName,
	LogLevelError:    logLevelErrorName,
	LogLevelWarning:  logLevelWarningName,
	LogLevelNotice:   logLevelNoticeName,
	LogLevelInfo:     logLevelInfoName,
	LogLevelDebug:    logLevelDebugName,
	LogLevelTrace:    logLevelTraceName,
}

func GetLevelName(level LogLevel) string {
	var name string
	var exists bool
	name, exists = levelMap[level]
	if !exists {
		return "invalid log level"
	} else {
		return name
	}
}

type LogLevel uint8

func NewLogLevel(l uint8) LogLevel {
	if LogLevel(l) > LogLevelTrace {
		return LogLevelTrace
	} else {
		return LogLevel(l)
	}
}

type logFunc func(structure, function, msg string, id int, vars ...any)

type Logger struct {
	logger *log.Logger
	// Log
	logEmerge   logFunc
	logAlert    logFunc
	logCritical logFunc
	logError    logFunc
	logWarning  logFunc
	logNotice   logFunc
	logInfo     logFunc
	logDebug    logFunc
	logTrace    logFunc

	//Fatal
	fatalEmerge   logFunc
	fatalAlert    logFunc
	fatalCritical logFunc
	fatalError    logFunc
	fatalWarning  logFunc
	fatalNotice   logFunc
	fatalInfo     logFunc
	fatalDebug    logFunc
	fatalTrace    logFunc
}

func (l *Logger) LogEmerge(structure, function, msg string, id int, vars ...any) {
	l.logEmerge(structure, function, msg, id, vars...)
}
func (l *Logger) LogAlert(structure, function, msg string, id int, vars ...any) {
	l.logAlert(structure, function, msg, id, vars...)
}
func (l *Logger) LogCritical(structure, function, msg string, id int, vars ...any) {
	l.logCritical(structure, function, msg, id, vars...)
}
func (l *Logger) LogError(structure, function, msg string, id int, vars ...any) {
	l.logError(structure, function, msg, id, vars...)
}
func (l *Logger) LogWarning(structure, function, msg string, id int, vars ...any) {
	l.logWarning(structure, function, msg, id, vars...)
}
func (l *Logger) LogNotice(structure, function, msg string, id int, vars ...any) {
	l.logNotice(structure, function, msg, id, vars...)
}
func (l *Logger) LogInfo(structure, function, msg string, id int, vars ...any) {
	l.logInfo(structure, function, msg, id, vars...)
}
func (l *Logger) LogDebug(structure, function, msg string, id int, vars ...any) {
	l.logDebug(structure, function, msg, id, vars...)
}
func (l *Logger) LogTrace(structure, function, msg string, id int, vars ...any) {
	l.logTrace(structure, function, msg, id, vars...)
}
func (l *Logger) FatalEmerge(structure, function, msg string, id int, vars ...any) {
	l.fatalEmerge(structure, function, msg, id, vars...)
}
func (l *Logger) FatalCritical(structure, function, msg string, id int, vars ...any) {
	l.fatalCritical(structure, function, msg, id, vars...)
}
func (l *Logger) FatalAlert(structure, function, msg string, id int, vars ...any) {
	l.fatalAlert(structure, function, msg, id, vars...)
}
func (l *Logger) FatalError(structure, function, msg string, id int, vars ...any) {
	l.fatalError(structure, function, msg, id, vars...)
}
func (l *Logger) FatalWarning(structure, function, msg string, id int, vars ...any) {
	l.fatalWarning(structure, function, msg, id, vars...)
}
func (l *Logger) FatalNotice(structure, function, msg string, id int, vars ...any) {
	l.fatalNotice(structure, function, msg, id, vars...)
}
func (l *Logger) FatalInfo(structure, function, msg string, id int, vars ...any) {
	l.fatalInfo(structure, function, msg, id, vars...)
}
func (l *Logger) FatalDebug(structure, function, msg string, id int, vars ...any) {
	l.fatalDebug(structure, function, msg, id, vars...)
}
func (l *Logger) FatalTrace(structure, function, msg string, id int, vars ...any) {
	l.fatalTrace(structure, function, msg, id, vars...)
}

func (l *Logger) SetVerbosity(level LogLevel) {
	for lvl := LogLevelNull; lvl <= LogLevelTrace; lvl++ {
		switch lvl {
		case LogLevelNull:
			l.logEmerge = func(structure, function, msg string, id int, vars ...any) {}
			l.logAlert = func(structure, function, msg string, id int, vars ...any) {}
			l.logCritical = func(structure, function, msg string, id int, vars ...any) {}
			l.logError = func(structure, function, msg string, id int, vars ...any) {}
			l.logWarning = func(structure, function, msg string, id int, vars ...any) {}
			l.logNotice = func(structure, function, msg string, id int, vars ...any) {}
			l.logInfo = func(structure, function, msg string, id int, vars ...any) {}
			l.logDebug = func(structure, function, msg string, id int, vars ...any) {}
			l.logTrace = func(structure, function, msg string, id int, vars ...any) {}

			l.fatalEmerge = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelEmergePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelEmergePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.fatalAlert = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelAlertPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelAlertPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.fatalCritical = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelCriticalPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelCriticalPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.fatalError = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelErrorPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelErrorPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.fatalWarning = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelWarningPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelWarningPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.fatalNotice = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelNoticePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelNoticePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.fatalInfo = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelInfoPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelInfoPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.fatalDebug = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelDebugPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelDebugPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
			l.fatalTrace = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelTracePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelTracePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelEmerge:
			if lvl <= level {
				l.logEmerge = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelEmergePrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelEmergePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.logEmerge = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.fatalEmerge = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelEmergePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelEmergePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelAlert:
			if lvl <= level {
				l.logAlert = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelAlertPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelAlertPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.logAlert = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.fatalAlert = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelAlertPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelAlertPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelCritical:
			if lvl <= level {
				l.logCritical = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelCriticalPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelCriticalPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.logCritical = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.fatalCritical = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelCriticalPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelCriticalPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelError:
			if lvl <= level {
				l.logError = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelErrorPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelErrorPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.logError = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.fatalError = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelErrorPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelErrorPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelWarning:
			if lvl <= level {
				l.logWarning = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelWarningPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelWarningPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.logWarning = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.fatalWarning = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelWarningPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelWarningPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelNotice:
			if lvl <= level {
				l.logNotice = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelNoticePrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelNoticePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.logNotice = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.fatalNotice = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelNoticePrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelNoticePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelInfo:
			if lvl <= level {
				l.logInfo = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelInfoPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelInfoPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.logInfo = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.fatalInfo = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelInfoPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelInfoPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelDebug:
			if lvl <= level {
				l.logDebug = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelDebugPrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelDebugPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.logDebug = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.fatalDebug = func(structure, function, msg string, id int, vars ...any) {
				if id < 0 {
					l.logger.Fatalf(logPattern, logLevelDebugPrefix, structure, function, fmt.Sprintf(msg, vars...))
				} else {
					l.logger.Fatalf(logPatternWithId, logLevelDebugPrefix, structure, function, id, fmt.Sprintf(msg, vars...))
				}
			}
		case LogLevelTrace:
			if lvl <= level {
				l.logTrace = func(structure, function, msg string, id int, vars ...any) {
					if id < 0 {
						l.logger.Printf(logPattern, logLevelTracePrefix, structure, function, fmt.Sprintf(msg, vars...))
					} else {
						l.logger.Printf(logPatternWithId, logLevelTracePrefix, structure, function, id, fmt.Sprintf(msg, vars...))
					}
				}
			} else {
				l.logTrace = func(structure, function, msg string, id int, vars ...any) {}
			}
			l.fatalTrace = func(structure, function, msg string, id int, vars ...any) {
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
