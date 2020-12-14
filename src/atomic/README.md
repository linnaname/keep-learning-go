和Java的atomic半斤八两都是CAS原理实现，主要的几个方法

Load：该类方法主要负责从相应的内存地址中获取对应的值

Store：该类主要负责将对应的值保存在相应的内存地址中

Add：该类方法可以理解是Load和Store的结合，也就是先Load然后Add

Swap：该类方法可以理解为先Load，在Store新值，然后返回旧值

CompareAndSwap：该类方法可以这样理解：先比较旧数据和地址中保存数据的值，如果相同的话，执行Swap，把新的数值保存在地址中，返回true，如果不同，那么直接返回false


对于bean类型的操作可以使用atomic.Value，其实就是把bean或者struct当作interface来操作，对interface多做了解就明白大概原理了，可以看看[这篇文章](https://studygolang.com/articles/23242?fr=sidebar)
