# 二分法查找

> 二分查找也称折半查找（Binary Search），它是一种效率较高的查找方法。
但是，折半查找要求线性表必须采用顺序存储结构，而且表中元素按关键字有序排列

## 核心代码

`mid = (start + end) >> 1`

```
mid = (start + end) >> 1 //求中间下位值, golang整形除取最小值,相当于math.Floor
if val > lst[mid] {      //如果目标值大于中间值,则将中间mid + 1赋值给start
    start = mid + 1
} else if val < lst[mid] { //如果目标值小于中间值,则将中间mid-1赋值给end
    end = mid - 1
} else { //表示值相等,找到目标下标
    break
}
```


