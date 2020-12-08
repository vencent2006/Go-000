学习笔记


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
