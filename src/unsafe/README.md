1.指针不能进行数学运算，不能进行不同类型转换、赋值、相等不等的的比较，golang的指针更像是为了一种获得指针传递外的妥协，golang中并不存在严格的指针，算是个还不错的的设计，但也让人有点迷惑
2.uintptr本质是一个整数类型，并没有指针的语义，意思就是uintptr 所指向的对象会被 gc 无情地回收。而 unsafe.Pointer 有指针语义，可以保护它所指向的对象在“有用”的时候不会被垃圾回收，这点需要注意

unsafe 包提供了 2 点重要的能力：
任何类型的指针和 unsafe.Pointer 可以相互转换。
uintptr 类型和 unsafe.Pointer 可以相互转换。

主要是Sizeof（结构体大小）、Offsetof偏移量、Alignof对齐信息

功能相对于Java的unsafe包少很多，Golang的源码中很多地方都用到