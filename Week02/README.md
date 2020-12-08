学习笔记

remind
In Go, panic is not exception


《Go进阶训练营第四课问题收集》 https://shimo.im/docs/R6gP8qyvWqJrgRCk/


# 作业
* 题目：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候， 是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
* 回答：应该抛给上层，出错时尽量使用 errors.Wrap 去包裹堆栈信息, 中间层透传就好, 在最上层打印并处理; sql.ErrNoRows 要处理掉, 可以屏蔽掉不同数据库的报错差异性, 方便上层对空记录集的处理.
以下为毛剑老师给的week2作业答案，代码就是每个人有每个人的理解，仅供大家参考～

```go
dao: 

 return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))


biz:

if errors.Is(err, code.NotFound} {

}
```






# recover 野生 goroutine
```go

package main
import(
  "fmt"
  "time"
)

func main(){
  fmt.Println("vim-go")
  Go(func(){
    fmt.Println("hello")
    //panic
    panic("一路向西")
  })
  
  time.Sleep(5*time.Second)

}

func Go(x func()){
  go func(){
    defer func(){
      if err := recover(); err != nil{
        fmt.Println(err)
      }
    }()
    
    //exec x()
    x()
    
  }()
}

```


* Sentinel error: 耦合太大
* Opaque error: 老师认为是最灵活的错误处理策略，要求代码和调用者之间的耦合最少
* Wrap errors：打印堆栈信息
* 最佳实践： pkg/errors
* go1.13: 不够完整
* go2：可以期待下
    
