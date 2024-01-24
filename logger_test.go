package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
	"testing"
)

func testLogGeneric(t *testing.T, level LogLevel) {
	const (
		structure   string = "struct"
		function           = "function"
		msg                = "my message"
		msgWithVars        = "my message with vars (n=1): %s=%d"
	)
	const id int = 0
	const noId int = -1
	var vars []any = []any{
		"n",
		1,
	}

	var resultAsBytes []byte
	var resultAsString string
	var receiver *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *Logger = NewLogger(level, receiver)
	var levelMap map[LogLevel]string = map[LogLevel]string{
		LogLevelNull:     "NULL",
		LogLevelEmerge:   logLevelEmergePrefix,
		LogLevelAlert:    logLevelAlertPrefix,
		LogLevelCritical: logLevelCriticalPrefix,
		LogLevelError:    logLevelErrorPrefix,
		LogLevelWarning:  logLevelWarningPrefix,
		LogLevelNotice:   logLevelNoticePrefix,
		LogLevelInfo:     logLevelInfoPrefix,
		LogLevelDebug:    logLevelDebugPrefix,
		LogLevelTrace:    logLevelTracePrefix,
	}
	var lf []logFunc = []logFunc{
		nil,
		logger.LogEmerge,
		logger.LogAlert,
		logger.LogCritical,
		logger.LogError,
		logger.LogWarning,
		logger.LogNotice,
		logger.LogInfo,
		logger.LogDebug,
		logger.LogTrace,
	}

	var expectedString []string = []string{
		"",
		fmt.Sprintf(logPattern, logLevelEmergePrefix, structure, function, msg),
		fmt.Sprintf(logPattern, logLevelAlertPrefix, structure, function, msg),
		fmt.Sprintf(logPattern, logLevelCriticalPrefix, structure, function, msg),
		fmt.Sprintf(logPattern, logLevelErrorPrefix, structure, function, msg),
		fmt.Sprintf(logPattern, logLevelWarningPrefix, structure, function, msg),
		fmt.Sprintf(logPattern, logLevelNoticePrefix, structure, function, msg),
		fmt.Sprintf(logPattern, logLevelInfoPrefix, structure, function, msg),
		fmt.Sprintf(logPattern, logLevelDebugPrefix, structure, function, msg),
		fmt.Sprintf(logPattern, logLevelTracePrefix, structure, function, msg),
	}

	var expectedStringWithId []string = []string{
		"",
		fmt.Sprintf(logPatternWithId, logLevelEmergePrefix, structure, function, id, msg),
		fmt.Sprintf(logPatternWithId, logLevelAlertPrefix, structure, function, id, msg),
		fmt.Sprintf(logPatternWithId, logLevelCriticalPrefix, structure, function, id, msg),
		fmt.Sprintf(logPatternWithId, logLevelErrorPrefix, structure, function, id, msg),
		fmt.Sprintf(logPatternWithId, logLevelWarningPrefix, structure, function, id, msg),
		fmt.Sprintf(logPatternWithId, logLevelNoticePrefix, structure, function, id, msg),
		fmt.Sprintf(logPatternWithId, logLevelInfoPrefix, structure, function, id, msg),
		fmt.Sprintf(logPatternWithId, logLevelDebugPrefix, structure, function, id, msg),
		fmt.Sprintf(logPatternWithId, logLevelTracePrefix, structure, function, id, msg),
	}

	var expectedStringWithVars []string = []string{
		"",
		fmt.Sprintf(logPattern, logLevelEmergePrefix, structure, function, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPattern, logLevelAlertPrefix, structure, function, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPattern, logLevelCriticalPrefix, structure, function, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPattern, logLevelErrorPrefix, structure, function, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPattern, logLevelWarningPrefix, structure, function, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPattern, logLevelNoticePrefix, structure, function, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPattern, logLevelInfoPrefix, structure, function, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPattern, logLevelDebugPrefix, structure, function, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPattern, logLevelTracePrefix, structure, function, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
	}

	var expectedStringWithVarsAndId []string = []string{
		"",
		fmt.Sprintf(logPatternWithId, logLevelEmergePrefix, structure, function, id, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPatternWithId, logLevelAlertPrefix, structure, function, id, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPatternWithId, logLevelCriticalPrefix, structure, function, id, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPatternWithId, logLevelErrorPrefix, structure, function, id, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPatternWithId, logLevelWarningPrefix, structure, function, id, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPatternWithId, logLevelNoticePrefix, structure, function, id, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPatternWithId, logLevelInfoPrefix, structure, function, id, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPatternWithId, logLevelDebugPrefix, structure, function, id, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
		fmt.Sprintf(logPatternWithId, logLevelTracePrefix, structure, function, id, fmt.Sprintf(msgWithVars, vars[0], vars[1])),
	}

	for i := LogLevelEmerge; i <= LogLevelTrace; i++ {
		// Test a simple message
		lf[i](structure, function, msg, noId)
		resultAsBytes, _ = io.ReadAll(receiver)
		resultAsString = string(resultAsBytes)
		if i <= level {
			if !strings.HasSuffix(resultAsString, expectedString[i]) {
				t.Errorf("incorrect string returned with the log level %s with a simple message", levelMap[i])
			}
		} else {
			if resultAsString != "" {
				t.Errorf("empty string expected with the log level %s with a simple message", levelMap[i])
			}
		}

		// Test a message with vars
		lf[i](structure, function, msgWithVars, noId, vars...)
		resultAsBytes, _ = io.ReadAll(receiver)
		resultAsString = string(resultAsBytes)
		if i <= level {
			if !strings.HasSuffix(resultAsString, expectedStringWithVars[i]) {
				t.Errorf("incorrect string returned with the log level %s with message with vars", levelMap[i])
			}
		} else {
			if resultAsString != "" {
				t.Errorf("empty string expected with the log level %s with message with vars", levelMap[i])
			}
		}

		// Test a simple message with Id
		lf[i](structure, function, msg, id)
		resultAsBytes, _ = io.ReadAll(receiver)
		resultAsString = string(resultAsBytes)
		if i <= level {
			if !strings.HasSuffix(resultAsString, expectedStringWithId[i]) {
				t.Errorf("incorrect string returned with the log level %s with a message with id", levelMap[i])
			}
		} else {
			if resultAsString != "" {
				t.Errorf("empty string expected with the log level %s with a message with id", levelMap[i])
			}
		}

		// Test a message with vars
		lf[i](structure, function, msgWithVars, id, vars...)
		resultAsBytes, _ = io.ReadAll(receiver)
		resultAsString = string(resultAsBytes)
		if i <= level {
			if !strings.HasSuffix(resultAsString, expectedStringWithVarsAndId[i]) {
				t.Errorf("incorrect string returned with the log level %s with message with vars and id", levelMap[i])
			}
		} else {
			if resultAsString != "" {
				t.Errorf("empty string expected with the log level %s with message with vars and id", levelMap[i])
			}
		}
	}
}

func TestLogNull(t *testing.T) {
	const (
		structure   string = "struct"
		function           = "function"
		msg                = "my message"
		msgWithVars        = "my message with vars (n=1): %s=%d"
	)
	const id int = 0
	const noId int = -1
	var vars []any = []any{
		"n",
		1,
	}

	const level LogLevel = LogLevelNull

	var resultAsBytes []byte
	var resultAsString string
	var receiver *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *Logger = NewLogger(level, receiver)
	var levelMap map[LogLevel]string = map[LogLevel]string{
		LogLevelNull:     "NULL",
		LogLevelEmerge:   logLevelEmergePrefix,
		LogLevelAlert:    logLevelAlertPrefix,
		LogLevelCritical: logLevelCriticalPrefix,
		LogLevelError:    logLevelErrorPrefix,
		LogLevelWarning:  logLevelWarningPrefix,
		LogLevelNotice:   logLevelNoticePrefix,
		LogLevelInfo:     logLevelInfoPrefix,
		LogLevelDebug:    logLevelDebugPrefix,
		LogLevelTrace:    logLevelTracePrefix,
	}
	var lf []logFunc = []logFunc{
		nil,
		logger.LogEmerge,
		logger.LogAlert,
		logger.LogCritical,
		logger.LogError,
		logger.LogWarning,
		logger.LogNotice,
		logger.LogInfo,
		logger.LogDebug,
		logger.LogTrace,
	}

	for i := LogLevelEmerge; i <= LogLevelTrace; i++ {
		// Test a simple message
		lf[i](structure, function, msg, noId)
		resultAsBytes, _ = io.ReadAll(receiver)
		resultAsString = string(resultAsBytes)
		if resultAsString != "" {
			t.Errorf("empty string expected with the log level %s with a simple message", levelMap[i])
		}

		// Test a message with vars
		lf[i](structure, function, msgWithVars, noId, vars...)
		resultAsBytes, _ = io.ReadAll(receiver)
		resultAsString = string(resultAsBytes)
		if resultAsString != "" {
			t.Errorf("empty string expected with the log level %s with message with vars", levelMap[i])
		}

		// Test a simple message with Id
		lf[i](structure, function, msg, id)
		resultAsBytes, _ = io.ReadAll(receiver)
		resultAsString = string(resultAsBytes)
		if resultAsString != "" {
			t.Errorf("empty string expected with the log level %s with a message with id", levelMap[i])
		}

		// Test a message with vars
		lf[i](structure, function, msgWithVars, id, vars...)
		resultAsBytes, _ = io.ReadAll(receiver)
		resultAsString = string(resultAsBytes)
		if resultAsString != "" {
			t.Errorf("empty string expected with the log level %s with message with vars and id", levelMap[i])
		}
	}
}
func TestLogEmerge(t *testing.T) {
	testLogGeneric(t, LogLevelEmerge)
}
func TestLogAlert(t *testing.T) {
	testLogGeneric(t, LogLevelAlert)
}
func TestLogCritical(t *testing.T) {
	testLogGeneric(t, LogLevelCritical)
}
func TestLogError(t *testing.T) {
	testLogGeneric(t, LogLevelError)
}
func TestLogWarning(t *testing.T) {
	testLogGeneric(t, LogLevelWarning)
}
func TestLogNotice(t *testing.T) {
	testLogGeneric(t, LogLevelNotice)
}
func TestLogInfo(t *testing.T) {
	testLogGeneric(t, LogLevelInfo)
}
func TestLogDebug(t *testing.T) {
	testLogGeneric(t, LogLevelDebug)
}
func TestLogTrace(t *testing.T) {
	testLogGeneric(t, LogLevelTrace)
}

func BenchmarkLogSimple(b *testing.B) {
	b.ReportAllocs()
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *Logger = NewLogger(LogLevelNotice, buffer)

	for i := 0; i < b.N; i++ {
		logger.LogNotice("struct", "func", "message", -1)
	}
}
func BenchmarkLogWithId(b *testing.B) {
	b.ReportAllocs()
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *Logger = NewLogger(LogLevelNotice, buffer)

	for i := 0; i < b.N; i++ {
		logger.LogNotice("struct", "func", "message", i)
	}
}
func BenchmarkLogWithVars(b *testing.B) {
	b.ReportAllocs()
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *Logger = NewLogger(LogLevelNotice, buffer)

	for i := 0; i < b.N; i++ {
		logger.LogNotice("struct", "func", "message %s", -1, "test")
	}
}
func BenchmarkLogWithVarsAndId(b *testing.B) {
	b.ReportAllocs()
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *Logger = NewLogger(LogLevelNotice, buffer)

	for i := 0; i < b.N; i++ {
		logger.LogNotice("struct", "func", "message %s", i, "test")
	}
}
func BenchmarkDefaultLogger(b *testing.B) {
	b.ReportAllocs()
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *log.Logger = log.New(buffer, "", log.Ldate|log.Lmicroseconds)

	for i := 0; i < b.N; i++ {
		logger.Printf(logPattern, logLevelCriticalPrefix, "struct", "func", "message")
	}
}

func BenchmarkLogNullSimple(b *testing.B) {
	b.ReportAllocs()
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *Logger = NewLogger(LogLevelNotice, buffer)

	for i := 0; i < b.N; i++ {
		logger.LogDebug("struct", "func", "message", -1)
	}
}
func BenchmarkLogNullWithId(b *testing.B) {
	b.ReportAllocs()
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *Logger = NewLogger(LogLevelNotice, buffer)

	for i := 0; i < b.N; i++ {
		logger.LogDebug("struct", "func", "message", i)
	}
}
func BenchmarkLogNullWithVars(b *testing.B) {
	b.ReportAllocs()
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *Logger = NewLogger(LogLevelNotice, buffer)

	for i := 0; i < b.N; i++ {
		logger.LogDebug("struct", "func", "message %s", -1, "test")
	}
}
func BenchmarkLogNullWithVarsAndId(b *testing.B) {
	b.ReportAllocs()
	var buffer *bytes.Buffer = bytes.NewBuffer([]byte{})
	var logger *Logger = NewLogger(LogLevelNotice, buffer)

	for i := 0; i < b.N; i++ {
		logger.LogDebug("struct", "func", "message %s", i, "test")
	}
}
