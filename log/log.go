package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	logger      = logrus.New()
	initialized = false
	logFile     *io.Writer
)

func InitLogger(level logrus.Level, logFilePath string) {
	if initialized {
		Errorln("logger already initialized")
		return
	}
	logger.SetLevel(level)
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		ErrorPrintln("failed to open log file %s: %s", logFile, err)
	}
	logger.SetOutput(logFile)
	initialized = true
}

func Debugln(args ...any) {
	logger.Debugln(args)
}

func DebugPrintln(args ...any) {
	if logger.Level < logrus.DebugLevel {
		return
	}
	fmt.Println(args...)
	Debugln(args)
}

func Debugf(format string, args ...any) {
	logger.Debugf(format, args...)
}

func DebugPrintf(format string, args ...any) {
	if logger.Level < logrus.DebugLevel {
		return
	}
	fmt.Printf(format, args...)
	Debugf(format, args)
}

func Infoln(args ...any) {
	logger.Infoln(args)
}

func InfoPrintln(args ...any) {
	if logger.Level < logrus.InfoLevel {
		return
	}
	fmt.Println(args...)
	Infoln(args)
}

func Infof(format string, args ...any) {
	logger.Infof(format, args...)
}

func InfoPrintf(format string, args ...any) {
	if logger.Level < logrus.InfoLevel {
		return
	}
	fmt.Printf(format, args...)
	Infof(format, args)
}

func Errorln(args ...any) {
	logger.Errorln(args)
}

func ErrorPrintln(args ...any) {
	if logger.Level < logrus.ErrorLevel {
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, args...)
	Errorln(args)
}

func Errorf(format string, args ...any) {
	logger.Errorf(format, args...)
}

func ErrorPrintf(format string, args ...any) {
	if logger.Level < logrus.ErrorLevel {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(format, args...))
	Errorf(fmt.Sprintf(format, args...))
}

func Fatalln(args ...any) {
	logger.Fatalln(args)
}

func FatalPrintln(args ...any) {
	if logger.Level < logrus.FatalLevel {
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, args...)
	Fatalln(args)
}

func Fatalf(format string, args ...any) {
	logger.Fatalf(format, args...)
}

func FatalPrintf(format string, args ...any) {
	if logger.Level < logrus.FatalLevel {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(format, args...))
	Fatalf(format, args)
}

func Panicln(args ...any) {
	logger.Panicln(args)
}

func PanicPrintln(args ...any) {
	if logger.Level < logrus.PanicLevel {
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, args...)
	Panicln(args)
}

func Panicf(format string, args ...any) {
	logger.Panicf(format, args...)
}

func PanicPrintf(format string, args ...any) {
	if logger.Level < logrus.PanicLevel {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(format, args...))
	Panicf(format, args)
}

func DryErrorln(args ...any) {
	if logger.Level < logrus.ErrorLevel {
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, args...)
}

func DryErrorf(format string, args ...any) {
	if logger.Level < logrus.ErrorLevel {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(format, args...))
}
