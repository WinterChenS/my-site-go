package initialize

import (
	"fmt"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"winterchen.com/my-site-go/src/global"
)

const (
	out_path = "/foot.log"
	err_path = "/foot.err"
)

func InitLogger() {
	// init zap config
	_, err := os.Stat(global.Configs.LogsPath)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(global.Configs.LogsPath, os.ModePerm)
			if err != nil {
				panic(fmt.Sprintf("create log path error: %v", err))
			}
		}
	}
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		CallerKey:     "caller",
		StacktraceKey: "trace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2022-09-05 13:04:05"))
		},
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	infoHook_1 := os.Stdout
	infoHook_2 := getWriter(out_path)
	errorHook := getWriter(err_path)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoHook_1), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoHook_2), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorHook), warnLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	global.Log = logger
	defer logger.Sync()

}

func getWriter(filename string) io.Writer {

	hook, err := rotatelogs.New(
		global.Configs.LogsPath+filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
