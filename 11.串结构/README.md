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