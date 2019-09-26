package main

type Conn string
type Result string

func (c Conn) DoQuery(q string) Result {
	return ""
}

func main() {

}

// 假设程序从多个复制的数据库同时读取。只需要接收首先到达的一个答案，
//Query 函数获取数据库的连接切片并请求。并行请求每一个数据库并返回收到的第一个响应
func Query(conns []Conn, query string) Result {
	ch := make(chan Result, 1)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query):
			}
		}(conn)
	}
	//查询到了第一条后，立即退出
	return <-ch
}
