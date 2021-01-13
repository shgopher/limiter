package limiter
// 限流器
type Limiter struct {
	ch chan struct{}
}

// 限流器的初始阶段是可以承受 number个数据的流量的。
func NewLimiter(number int) *Limiter {
	l := &Limiter{make(chan struct{}, number)}
	return l

}

// 限流器工作，取出令牌，每次取出一个。chan就是每次加入一个。
func (l *Limiter) Limit() {
	l.ch <- struct{}{}
}

// 增加令牌，chan就是每次取出一个
func (l *Limiter) Add() {
	for i := 0; i < len(l.ch); i++ {
		<-l.ch
	}
}
