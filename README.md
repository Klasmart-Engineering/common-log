# Common Log

本Logger是对zap logger的封装，并且保留了替换成其他logger的能力

使用方法
```
import (
    "gitlab.badanamu.com.cn/calmisland/common-log/log"
)

log.Info(ctx, "just a test message", log.Time("now", time.Now()), log.String("abc", "123"))
```

如果需要重定向logger的输出到某个io.Writer，可以这样写
```
import (
    "gitlab.badanamu.com.cn/calmisland/common-log/log"
)

logger := log.New(WithWriter(YourWriter))
log.ReplaceGlobals(logger)

log.Debug(ctx,  "writer to io.Writer")
```
