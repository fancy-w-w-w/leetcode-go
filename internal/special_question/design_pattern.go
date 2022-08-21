package specialquestion

// 1
// // 选项设计模式
// // 问题：有一个结构体，定义一个函数，给结构体初始化

// // 结构体
// type Options struct {
//     str1 string
//     str2 string
//     int1 int
//     int2 int
// }

// // 声明一个函数类型的变量，用于传参
// type Option func(opts *Options)

// func InitOptions(opts ...Option) {
//     options := &Options{}
//     for _, opt := range opts {
//         opt(options)
//     }
//     fmt.Printf("options:%#v\n", options)
// }

// func WithStringOption1(str string) Option {
//     return func(opts *Options) {
//         opts.str1 = str
//     }
// }

// func WithStringOption2(str string) Option {
//     return func(opts *Options) {
//         opts.str2 = str
//     }
// }
// func WithStringOption3(int1 int) Option {
//     return func(opts *Options) {
//         opts.int1 = int1
//     }
// }
// func WithStringOption4(int1 int) Option {
//     return func(opts *Options) {
//         opts.int2 = int1
//     }
// }
// func main() {
//     InitOptions(WithStringOption1("5lmh.com"), WithStringOption2("topgoer.com"), WithStringOption3(5), WithStringOption4(6))
// }

// 2 修饰器模式
// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
//     "strings"
// )

// func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         log.Println("--->WithServerHeader()")
//         w.Header().Set("Server", "HelloServer v0.0.1")
//         h(w, r)
//     }
// }

// func hello(w http.ResponseWriter, r *http.Request) {
//     log.Printf("Recieved Request %s from %s\n", r.URL.Path, r.RemoteAddr)
//     fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
// }

// func main() {
//     http.HandleFunc("/v1/hello", WithServerHeader(hello))
//     err := http.ListenAndServe(":8080", nil)
//     if err != nil {
//         log.Fatal("ListenAndServe: ", err)
//     }
// }

// 观察者模式
// 定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。

// package observer

// import "fmt"

// // 报社 —— 客户
// type Customer interface {
//     update()
// }

// type CustomerA struct {
// }

// func (*CustomerA) update() {
//     fmt.Println("我是客户A, 我收到报纸了")
// }

// type CustomerB struct {
// }

// func (*CustomerB) update() {
//     fmt.Println("我是客户B, 我收到报纸了")
// }

// // 报社 （被观察者)
// type NewsOffice struct {
//     customers []Customer
// }

// func (n *NewsOffice) addCustomer(customer Customer) {
//     n.customers = append(n.customers, customer)
// }

// func (n *NewsOffice) newspaperCome() {
//     // 通知所有客户
//     n.notifyAllCustomer()
// }

// func (n *NewsOffice) notifyAllCustomer() {
//     for _, customer := range n.customers {
//         customer.update()
//     }
// }
// package observer

// import "testing"

// func TestObserver(t *testing.T) {

//     customerA := &CustomerA{}
//     customerB := &CustomerB{}

//     office := &NewsOffice{}
//     // 模拟客户订阅
//     office.addCustomer(customerA)
//     office.addCustomer(customerB)
//     // 新的报纸
//     office.newspaperCome()

// }

// 3
// 单例模式
// package main

// import (
//     "fmt"
//     "sync"
// )

// var lock = &sync.Mutex{}

// type single struct {
// }

// var singleInstance *single

// func getInstance() *single {
//     if singleInstance == nil {
//         lock.Lock()
//         defer lock.Unlock()
//         if singleInstance == nil {
//             fmt.Println("Creating single instance now.")
//             singleInstance = &single{}
//         } else {
//             fmt.Println("Single instance already created.")
//         }
//     } else {
//         fmt.Println("Single instance already created.")
//     }

//     return singleInstance
// }

// 4策略模式

// // 实现此接口，则为一个策略
// type IStrategy interface {
//     do(int, int) int
// }

// // 加
// type add struct{}

// func (*add) do(a, b int) int {
//     return a + b
// }

// // 减
// type reduce struct{}

// func (*reduce) do(a, b int) int {
//     return a - b
// }

// // 具体策略的执行者
// type Operator struct {
//     strategy IStrategy
// }

// // 设置策略
// func (operator *Operator) setStrategy(strategy IStrategy) {
//     operator.strategy = strategy
// }

// // 调用策略中的方法
// func (operator *Operator) calculate(a, b int) int {
//     return operator.strategy.do(a, b)
// }
