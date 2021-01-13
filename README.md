# limiter
使用channel制作的限流算法，代码量极少。基本的原理就是把channel的那个队列当成令牌筒，channel天生的多线程安全。

`_test.go`中的test是一个例子：2秒钟，限流10次数，然后往复运动。
