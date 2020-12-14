协程安全，然而受到GC的影响无法保证put之后一定能get到，不适合于做连接池而连接池需要自己管理对象的生命周期。

底层使用切片加链表来实现双端队列，并将缓存的对象存储在切片中。在加入 victim 机制前，sync.Pool 里对象的最⼤缓存时间是一个 GC 周期，当 GC 开始时，没有被引⽤的对象都会被清理掉；加入 victim 机制后，最大缓存时间为两个 GC 周期。

fmt包的print使用了pool

对开源的connection pool和object pool还需要做些了解 