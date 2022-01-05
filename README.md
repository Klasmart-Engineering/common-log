# Common Log

本Logger是对zap logger的封装，并且保留了替换成其他logger的能力

### 安装

```
go get gitlab.badanamu.com.cn/calmisland/common-log/log
```

### 使用方法

```
import (
    "gitlab.badanamu.com.cn/calmisland/common-log/log"
)

log.Info(ctx, "just a test message", log.Time("now", time.Now()), log.String("abc", "123"))
```

### 功能

#### 输出重定向

如果需要重定向logger的输出到某个io.Writer，可以这样写
```
logger := log.New(log.WithWriter(YourWriter))
log.ReplaceGlobals(logger)

log.Debug(ctx,  "write log to io.Writer")
```

#### 按级别过滤日志

```
logger := log.New(WithLogLevel(log.LevelInfo))
log.ReplaceGlobals(logger)

log.Debug(ctx,  "hide in the log")
log.Info(ctx, "show in the log")
```

#### 静态字段

```
logger := log.New(Log.WithStaticFields([]log.Field{
	log.String("service", "test"),
	log.String("hello", "world"),
}))
log.ReplaceGlobals(logger)

log.Info(ctx,  "message with static fields")
```

#### 动态字段

```
logger := log.New(log.WithDynamicFields(func(ctx context.Context) []log.Field {
	value, ok := ctx.Value("aabbccddKK").(string)

	return []log.Field{
		log.String("aabbccddKK", value),
		log.Bool("ok", ok),
	}
}))
log.ReplaceGlobals(logger)

ctx := context.Background()
log.Info(ctx,  "message without dynamic fields")

ctx := context.WithValue(ctx, "aabbccddKK", "665544332211")
log.Info(ctx, "message with dynamic fields".)
```
