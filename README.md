# Goland 工具轮子
>`轮子制造者` 😊

## 结构目录
```json
├─examples
└─src
    └─gcoroutine
```

## 编码规范
* effectively go
* 变量
* struct
* 方法
* 注释
* 包名 -- 全小写
* slice
* error


### 并发包(gcoroutine)--线程安全
* slice
  * 读写锁实现
* queue
  * 读写锁实现
  * 采用双向链表实现
    * 本来想用map+slice结构实现
    * 考虑到map遍历比较繁琐,所以用循环链表实现
    * 有其它实现方案的可以说下哈😊
* map
  * 读写锁实现

### 时间包(gtime)
* 时间戳转换
* 时间操作
* 时区转换



### 测试例子
