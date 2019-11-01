# 跳表
> Redis里的有序集合就是使用一个字典和一个跳跃表
```
typedef struct zset {
    zskiplist *zsl;
    dict *dict;
} zset;
```