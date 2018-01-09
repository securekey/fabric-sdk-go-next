/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package modulledlogger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"sync"

	"io"

	"sync/atomic"

	"github.com/hyperledger/fabric-sdk-go/api/apilogging"
	"github.com/hyperledger/fabric-sdk-go/pkg/logging/decorator"
	"github.com/hyperledger/fabric-sdk-go/pkg/logging/utils"
)

var rwmutex = &sync.RWMutex{}
var moduleLevels = &decorator.ModuleLeveled{}
var callerInfos = &decorator.CallerInfo{}
var useCustomLogger int32

// default logger factory singleton
var loggerProviderInstance apilogging.LoggerProvider
var loggerProviderOnce sync.Once

type loggingProvider struct {
}

//GetLogger returns SDK logger implementation
func (p *loggingProvider) GetLogger(module string) apilogging.Logger {
	newDefLogger := log.New(os.Stdout, fmt.Sprintf(logPrefixFormatter, module), log.Ldate|log.Ltime|log.LUTC)
	var newCustomLogger apilogging.Logger
	if atomic.LoadInt32(&useCustomLogger) > 0 {
		newCustomLogger = loggerProviderInstance.GetLogger(module)
	}
	return &DefLogger{deflogger: newDefLogger, customLogger: newCustomLogger, module: module}
}

//LoggerProvider returns logging provider for SDK logger
func LoggerProvider() apilogging.LoggerProvider {
	return &loggingProvider{}
}

//InitModulledLogger sets custom logger which will be used over deflogger.
//It is required to call this function before making any loggings.
func InitModulledLogger(l apilogging.LoggerProvider) {
	loggerProviderOnce.Do(func() {
		loggerProviderInstance = l
		atomic.StoreInt32(&useCustomLogger, 1)
	})
}

//DefLogger standard SDK logger
type DefLogger struct {
	deflogger    *log.Logger
	customLogger apilogging.Logger
	module       string
	once         sync.Once
}

//LoggerOpts  for all logger customization options
type loggerOpts struct {
	levelEnabled        bool
	callerInfoEnabled   bool
	customLoggerEnabled bool
}

const (
	logLevelFormatter   = "UTC %s-> %4.4s "
	logPrefixFormatter  = " [%s] "
	callerInfoFormatter = "- %s "
)

//SetLevel - setting log level for given module
func SetLevel(module string, level apilogging.Level) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	moduleLevels.SetLevel(module, level)
}

//GetLevel - getting log level for given module
func GetLevel(module string) apilogging.Level {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return moduleLevels.GetLevel(module)
}

//IsEnabledFor - Check if given log level is enabled for given module
func IsEnabledFor(module string, level apilogging.Level) bool {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return moduleLevels.IsEnabledFor(module, level)
}

//ShowCallerInfo - Show caller info in log lines for given log level
func ShowCallerInfo(module string, level apilogging.Level) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	callerInfos.ShowCallerInfo(module, level)
}

//HideCallerInfo - Do not show caller info in log lines for given log level
func HideCallerInfo(module string, level apilogging.Level) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	callerInfos.HideCallerInfo(module, level)
}

//getLoggerOpts - returns LoggerOpts which can be used for customization
func getLoggerOpts(module string, level apilogging.Level) *loggerOpts {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return &loggerOpts{
		levelEnabled:        moduleLevels.IsEnabledFor(module, level),
		callerInfoEnabled:   callerInfos.IsCallerInfoEnabled(module, level),
		customLoggerEnabled: atomic.LoadInt32(&useCustomLogger) > 0,
	}
}

// Fatal is CRITICAL log followed by a call to os.Exit(1).
func (l *DefLogger) Fatal(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.CRITICAL)
	if opts.customLoggerEnabled {
		l.customLogger.Fatal(args...)
		return
	}
	l.log(opts, apilogging.CRITICAL, args...)
	l.deflogger.Fatal(args...)
}

// Fatalf is CRITICAL log formatted followed by a call to os.Exit(1).
func (l *DefLogger) Fatalf(format string, args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.CRITICAL)
	if opts.customLoggerEnabled {
		l.customLogger.Fatalf(format, args...)
		return
	}
	l.logf(opts, apilogging.CRITICAL, format, args...)
	l.deflogger.Fatalf(format, args...)
}

// Fatalln is CRITICAL log ln followed by a call to os.Exit(1).
func (l *DefLogger) Fatalln(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.CRITICAL)
	if opts.customLoggerEnabled {
		l.customLogger.Fatalln(args...)
		return
	}
	l.logln(opts, apilogging.CRITICAL, args...)
	l.deflogger.Fatalln(args...)
}

// Panic is CRITICAL log followed by a call to panic()
func (l *DefLogger) Panic(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.CRITICAL)
	if opts.customLoggerEnabled {
		l.customLogger.Panic(args...)
		return
	}
	l.log(opts, apilogging.CRITICAL, args...)
	l.deflogger.Panic(args...)
}

// Panicf is CRITICAL log formatted followed by a call to panic()
func (l *DefLogger) Panicf(format string, args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.CRITICAL)
	if opts.customLoggerEnabled {
		l.customLogger.Panicf(format, args...)
		return
	}
	l.logf(opts, apilogging.CRITICAL, format, args...)
	l.deflogger.Panicf(format, args...)
}

// Panicln is CRITICAL log ln followed by a call to panic()
func (l *DefLogger) Panicln(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.CRITICAL)
	if opts.customLoggerEnabled {
		l.customLogger.Panicln(args...)
		return
	}
	l.logln(opts, apilogging.CRITICAL, args...)
	l.deflogger.Panicln(args...)
}

// Print calls go log.Output.
// Arguments are handled in the manner of fmt.Print.
func (l *DefLogger) Print(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.CRITICAL)
	if opts.customLoggerEnabled {
		l.customLogger.Print(args...)
		return
	}
	l.deflogger.Print(args...)
}

// Printf calls go log.Output.
// Arguments are handled in the manner of fmt.Printf.
func (l *DefLogger) Printf(format string, args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.CRITICAL)
	if opts.customLoggerEnabled {
		l.customLogger.Printf(format, args...)
		return
	}
	l.deflogger.Printf(format, args...)
}

// Println calls go log.Output.
// Arguments are handled in the manner of fmt.Println.
func (l *DefLogger) Println(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.CRITICAL)
	if opts.customLoggerEnabled {
		l.customLogger.Println(args...)
		return
	}
	l.deflogger.Println(args...)
}

// Debug calls go log.Output.
// Arguments are handled in the manner of fmt.Print.
func (l *DefLogger) Debug(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.DEBUG)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Debug(args...)
		return
	}
	l.log(opts, apilogging.DEBUG, args...)
}

// Debugf calls go log.Output.
// Arguments are handled in the manner of fmt.Printf.
func (l *DefLogger) Debugf(format string, args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.DEBUG)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Debugf(format, args...)
		return
	}
	l.logf(opts, apilogging.DEBUG, format, args...)
}

// Debugln calls go log.Output.
// Arguments are handled in the manner of fmt.Println.
func (l *DefLogger) Debugln(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.DEBUG)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Debugln(args...)
		return
	}
	l.logln(opts, apilogging.DEBUG, args...)
}

// Info calls go log.Output.
// Arguments are handled in the manner of fmt.Print.
func (l *DefLogger) Info(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.INFO)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Info(args...)
		return
	}
	l.log(opts, apilogging.INFO, args...)
}

// Infof calls go log.Output.
// Arguments are handled in the manner of fmt.Printf.
func (l *DefLogger) Infof(format string, args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.INFO)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Infof(format, args...)
		return
	}
	l.logf(opts, apilogging.INFO, format, args...)
}

// Infoln calls go log.Output.
// Arguments are handled in the manner of fmt.Println.
func (l *DefLogger) Infoln(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.INFO)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Infoln(args...)
		return
	}
	l.logln(opts, apilogging.INFO, args...)
}

// Warn calls go log.Output.
// Arguments are handled in the manner of fmt.Print.
func (l *DefLogger) Warn(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.WARNING)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Warn(args...)
		return
	}
	l.log(opts, apilogging.WARNING, args...)
}

// Warnf calls go log.Output.
// Arguments are handled in the manner of fmt.Printf.
func (l *DefLogger) Warnf(format string, args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.WARNING)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Warnf(format, args...)
		return
	}
	l.logf(opts, apilogging.WARNING, format, args...)
}

// Warnln calls go log.Output.
// Arguments are handled in the manner of fmt.Println.
func (l *DefLogger) Warnln(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.WARNING)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Warnln(args...)
		return
	}
	l.logln(opts, apilogging.WARNING, args...)
}

// Error calls go log.Output.
// Arguments are handled in the manner of fmt.Print.
func (l *DefLogger) Error(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.ERROR)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Error(args...)
		return
	}
	l.log(opts, apilogging.ERROR, args...)
}

// Errorf calls go log.Output.
// Arguments are handled in the manner of fmt.Printf.
func (l *DefLogger) Errorf(format string, args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.ERROR)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Errorf(format, args...)
		return
	}
	l.logf(opts, apilogging.ERROR, format, args...)
}

// Errorln calls go log.Output.
// Arguments are handled in the manner of fmt.Println.
func (l *DefLogger) Errorln(args ...interface{}) {
	opts := getLoggerOpts(l.module, apilogging.ERROR)
	if !opts.levelEnabled {
		return
	}
	if opts.customLoggerEnabled {
		l.customLogger.Errorln(args...)
		return
	}
	l.logln(opts, apilogging.ERROR, args...)
}

//ChangeOutput for changing output destination for the logger.
func (l *DefLogger) ChangeOutput(output io.Writer) {
	l.deflogger.SetOutput(output)
}

func (l *DefLogger) logf(opts *loggerOpts, level apilogging.Level, format string, args ...interface{}) {
	//Format prefix to show function name and log level and to indicate that timezone used is UTC
	customPrefix := fmt.Sprintf(logLevelFormatter, l.getCallerInfo(opts), utils.LogLevelString(level))
	l.deflogger.Output(2, customPrefix+fmt.Sprintf(format, args...))
}

func (l *DefLogger) log(opts *loggerOpts, level apilogging.Level, args ...interface{}) {
	//Format prefix to show function name and log level and to indicate that timezone used is UTC
	customPrefix := fmt.Sprintf(logLevelFormatter, l.getCallerInfo(opts), utils.LogLevelString(level))
	l.deflogger.Output(2, customPrefix+fmt.Sprint(args...))
}

func (l *DefLogger) logln(opts *loggerOpts, level apilogging.Level, args ...interface{}) {
	//Format prefix to show function name and log level and to indicate that timezone used is UTC
	customPrefix := fmt.Sprintf(logLevelFormatter, l.getCallerInfo(opts), utils.LogLevelString(level))
	l.deflogger.Output(2, customPrefix+fmt.Sprintln(args...))
}

func (l *DefLogger) getCallerInfo(opts *loggerOpts) string {

	if !opts.callerInfoEnabled {
		return ""
	}

	const MAXCALLERS = 6                           // search MAXCALLERS frames for the real caller
	const SKIPCALLERS = 4                          // skip SKIPCALLERS frames when determining the real caller
	const DEFAULTLOGPREFIX = "apilogging.(Logger)" // LOGPREFIX indicates the upcoming frame contains the real caller and skip the frame
	const LOGPREFIX = "logging.(*Logger)"          // LOGPREFIX indicates the upcoming frame contains the real caller and skip the frame
	const LOGBRIDGEPREFIX = "logbridge."           // LOGBRIDGEPREFIX indicates to skip the frame due to being a logbridge
	const NOTFOUND = "n/a"

	fpcs := make([]uintptr, MAXCALLERS)

	n := runtime.Callers(SKIPCALLERS, fpcs)
	if n == 0 {
		return fmt.Sprintf(callerInfoFormatter, NOTFOUND)
	}

	frames := runtime.CallersFrames(fpcs[:n])
	funcIsNext := false
	for f, more := frames.Next(); more; f, more = frames.Next() {
		_, funName := filepath.Split(f.Function)
		if f.Func == nil || f.Function == "" {
			funName = NOTFOUND // not a function or unknown
		}

		if strings.HasPrefix(funName, LOGPREFIX) || strings.HasPrefix(funName, LOGBRIDGEPREFIX) || strings.HasPrefix(funName, DEFAULTLOGPREFIX) {
			funcIsNext = true
		} else if funcIsNext {
			return fmt.Sprintf(callerInfoFormatter, funName)
		}
	}

	return fmt.Sprintf(callerInfoFormatter, NOTFOUND)
}
