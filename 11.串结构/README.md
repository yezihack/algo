## 串的数据结构
### 存储结构
1. 顺序存储方式
1. 链式存储方式,需要考虑存储密度 = 串值所占的存储/实际分配的存储

### 如何提高链式存储密度
1. 使用块链结构
```c
typedef struct Chunk{
    char ch[CHUNKSIZE];//使用了块链结构
    struct Chunk *next;
}Chunk
```

## BF算法
1. 字符串比较,使用两个指针逐一比较,不匹配则回溯
1. 时间复杂度: O(m*n

## KMP算法
> KMP由三位科学家提出 .D.E.Knuth, J.H.Morris AND V.R.Pratt