学习笔记

# 作业
👉Week03 作业题目：
1.基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
<br/>*以上作业，要求提交到Github上面，Week03作业提交地址：
https://github.com/Go-000/Go-000/issues/69*

**👉请务必按照示例格式进行提交，不要复制其他同学的格式，以免格式错误无法抓取作业。**
<br/><br/><br/>

# 毛剑面试必考并发，cow、atomic.value
* redis面试：cow必考，bgsave
* 技术视野，大小公司的区别
<br/>
<br/>
<br/>
<br/>

# 高频使用锁的情况，
* fast path是spinning
* slow path是barging

<br/>
<br/>
<br/>
<br/>

# 空转
* intel pause
* nginx对空转也有优化

<br/>
<br/>
<br/>
<br/>

# 可能的策略
* 一个报错，全部取消
* 一个报错，降级处理
* ……
* *g.Wait()，能拿到第一个报错的error*

<br/>
<br/>
<br/>
<br/>

# interface的坑
* 搜索“go interface nil”，与interface的实现相关（type+value）
*https://www.cnblogs.com/chris-cp/p/7596089.html*

<br/>
<br/>
<br/>
<br/>

# channel最大的坑
* 一定是交给发送者，close
* 实践中，可能无法找到管理chan的地方，比如tcp请求会开很多chan
* 实践中，老师会发送nil给channel，接收者知道要安全退出
* 但是发送nil的时候，可能还有人在往chan中发数据
