package main

func main() {

}

type RecentCounter struct {
	times []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		times: make([]int, 0, 10000),
	}
}

//用来统计一个固定时间段内请求的总数
func (r *RecentCounter) Ping(t int) int {
	r.times = append(r.times, t)

	for len(r.times) > 0 && r.times[0]+3000 < t {
		r.times = r.times[1:]
	}
	return len(r.times)
}
