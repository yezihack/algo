## 树与图的查找

### 平衡二叉树AVL (Balanced Binary Tree)
1. 可以是空树
1. 必须符合二叉树的性质.
1. 左右子树每一个结点高度之差的绝对值不大于1,否则称为失衡.需要进行调整.
#### 平衡因子
1. 平衡因子=结点左子树的高度-结点右子树的高度
1. 每一个结点必须是(-1, 0, 1)
![](https://mubu.com/document_image/25e2a2c1-1404-4a5d-b143-9b457763f682-2746950.jpg)

#### 失衡的四种状态
1. LL型
1. LR型
1. RR型
1. RL型
![](https://mubu.com/document_image/51e756e4-1325-4a15-b62b-3c5feb92ce92-2746950.jpg)
![](https://mubu.com/document_image/19bb645f-f239-43ef-a45b-0a3c5bf7ddf9-2746950.jpg)

#### 如何调整失衡的二叉树
> 两个原则
1. 降低高度
1. 保持二叉排序树的性质.