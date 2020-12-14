
分成两大类,一：ip、udp、unix（DGRAM）无链接的协议，二：tcp、unix（STREAM），需要建立链接

每一种通信方式都使用 XXConn 结构体来表示，诸如IPConn、TCPConn等，这些结构体都实现了Conn接口，Conn接口实现了基本的读、写、关闭、获取远程和本地地址、设置timeout等功能
基本都有DialXX和ListenXX方法，当然每个协议还是有些不同的方法，比如tcp有accept过程

网络编程范式都差不多，无法是接口不同而已，但还接触过具体的网络开源库，需要做些了解