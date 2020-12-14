
深入点的分析：https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/

官方的建议

1.Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named ctx.
2.Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use.
3.Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
4.The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.


1.这种显式在函数第一个参数传参的行为我觉得挺傻的，如果有很多函数要做这个操作，那这种传参不是满天飞了，简直是灾难

2.即使是要做tracing或者令牌传递我个人觉得也不要用context.Value来做，只用来做goroutine的超时取消控制就可以

3.总得来说我觉得这是个糟糕的设计，golang对context的核心期待应该是控制goroutine的超时、取消等等，却把参数的传递等等糅杂了进来，目前觉得并不是很好的设计

