select 能在 Channel 上进行非阻塞的收发操作；select 在遇到多个 Channel 同时响应时会随机挑选 case 执行；

当存在可以收发的 Channel 时，直接处理该 Channel 对应的 case；当不存在可以收发的 Channel 是，执行 default 中的语句；

panic 只会触发当前 Goroutine 的延迟函数调用；recover 只有在 defer 函数中调用才会生效；panic 允许在 defer 中嵌套多次调用；