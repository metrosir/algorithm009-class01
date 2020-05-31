# 总结

<a name="597aM"></a>
# 1.树、二叉树、二叉搜索树
<a name="LSE9Q"></a>
## 树
<a name="Uyod3"></a>
### 定义


1. **树**是元素的集合，树有**多个节点**用来存储元素。某些节点直接存在关系用连线表示，连线称之为**边**，边的上端节点称为**父节点**，下端称为**子节点**
1. 每个节点可以有多个子节点，而该节点是相应节点的父节点，树节点中有一个没有父节点的节点我们成为**根节点，**没有子节点的节点成为**叶节点**。
1. 树中节点的最大层次我们称之为**深度**，
1. 树是元素的集合，该集合可以为空，这时候树中没有元素，我们称之为**空树**。
1. 如果该集合不为空，那么该集合有一个根节点，以及0个或者多个子树



<a name="homkq"></a>
### 树的实现


1. linux的文件管理系统
1. git
1. ……



<a name="yzHc2"></a>
## 二叉树

<br />

<a name="KHSQH"></a>
### 定义


1. 二叉树是一种特殊的树
1. **二叉树的每个节点最多只能有2个子节点**
1. 二叉树的每个节点有一个左子节点，和右子节点（左子节点是左子树的根节点，右子节点是右子树的根节点）


<br />

<a name="iBro2"></a>
### 代码定义


```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func New() *TreeNode  {
    return &TreeNode{
        Val: 1,
        Left: &TreeNode{},
        Right: &TreeNode{},
    }
}
```

<br />

<a name="wFdq0"></a>
### 二叉树的遍历模板

<br />

<a name="UhMx6"></a>
#### 前序遍历


```go
type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}
func qx(root *TreeNode, *res) {
    if (root == nil) {
    	return
    }
    res := append(res, root.Val)
    qx(root.Left,res)
    qx(root.Right,res)
}
```

<br />

<a name="uj9lR"></a>
#### 中序遍历


```go
type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}
func qx(root *TreeNode, *res) {
    if (root == nil) {
    	return
    }
	qx(root.Left,res)
    res := append(res, root.Val)
    qx(root.Right,res)
}
```

<br />

<a name="564On"></a>
#### 后序遍历


```go
type TreeNode struct {
	Val int
    Left *TreeNode
    Right *TreeNode
}
func qx(root *TreeNode, *res) {
    if (root == nil) {
    	return
    }
	qx(root.Left,res)
    qx(root.Right,res)
    res := append(res, root.Val)
}
```

<br />

<a name="mapTK"></a>
## 二叉搜索树


<a name="imSM0"></a>
### 定义


1. 二叉搜索树是特殊的二叉树
1. **二叉搜索树的每个节点都比左子树的任意节点大，比右子树的任意节点小**
1. ……
<a name="n4gYU"></a>
# 2.堆「Heap」
<a name="JhHtB"></a>
## 堆「Heap」的定义

1. 可以迅速找到一个堆数中的最大值或者最小值的数据结构（两者一般来说只能取其一）
1. 将根结点最大的堆叫做**大顶堆**或**大根堆**，根节点最小的堆叫**小顶堆**或**小根堆**；
1. 我们可以理解Heap（堆）只是定义了一个接口，这个接口定义了查找最大值、删除最大值、新增加值这些操作，每个操作都是高效的。
1. 堆的实现一共有9种实现方式：查看百科定义->[入口](https://en.wikipedia.org/wiki/Heap_(data_structure)#Comparison_of_theoretic_bounds_for_variants)



<a name="JqBtO"></a>
## 常见的堆

1. 二叉堆
1. 斐波拉切堆
1. ……



<a name="tTfAv"></a>
## 大顶堆常见操作

1. find-max（找最大值）：O（1）
1. delete-max（删除最大值）：O(log n)
1. insert(create)（增加）：O(log n) or O (1)



<a name="iVOwi"></a>
## 堆的实现
<a name="Zjqrh"></a>
### 二叉堆

1. 二叉堆实现相对容易，但是时间复杂度跟其他堆相比较差
1. 通过完全二叉树进行实现（注意：不是二叉搜索树）
<a name="N5QYf"></a>
#### 二叉堆（大顶堆）满足下列特性

1. 是一棵完全二叉树
1. 树中任意节点的值总是>=其子节点的值
1. 根据二叉堆的特细，我们可以把二叉堆直接放到数组中进行运算
1. 根节点是：a[0]
1. 索引为i的左孩子索引是：**2*i+1**
1. 索引为i的右孩子索引是：**2*i+2**
1. 索引为i的父亲节点的索引是：**floor((i-1)/2)**
<a name="v5qe7"></a>
#### insert 插入操作

1. 新元素一律先插入到堆的尾部
1. 依次调整整个堆的结构（一直到根即可）HeapifyUp
1. 时间复杂度，数的深度就是最大的时间复杂度，数是N个且是二叉树结构，为：**0(log2n)**


<br />**具体操作步骤**<br />例：我们要在原始的大根堆（[90,80,70,60,40,30,20,10,50]）中添加一个元素85

1. 将85添加到末尾
1. 85大于父节点40，将它和父节点交换
1. 85大于父节点（80）将它和父节点交换

见下图：<br />
<br />![image.png](https://cdn.nlark.com/yuque/0/2020/png/963603/1590900658367-57eb785e-097e-4d8c-bde6-cb8c2e6c0694.png#align=left&display=inline&height=211&margin=%5Bobject%20Object%5D&name=image.png&originHeight=422&originWidth=946&size=87627&status=done&style=shadow&width=473)<br />![image.png](https://cdn.nlark.com/yuque/0/2020/png/963603/1590900707783-910c29de-9dee-49c0-927c-6cf67684db49.png#align=left&display=inline&height=285&margin=%5Bobject%20Object%5D&name=image.png&originHeight=570&originWidth=662&size=119068&status=done&style=shadow&width=331)<br />
<br />![image.png](https://cdn.nlark.com/yuque/0/2020/png/963603/1590900916467-c7af0587-1f8e-4cef-8b80-77b471c17a66.png#align=left&display=inline&height=241&margin=%5Bobject%20Object%5D&name=image.png&originHeight=482&originWidth=838&size=104490&status=done&style=shadow&width=419)<br />

<a name="LkI3I"></a>
#### Delete max删除堆顶操作

1. 将堆尾的元素替换到顶部
1. 依次从根部向下调整整个堆的结构，一直到堆尾即可；HeapifyDown
1. 时间复杂度0(log n)
1. <br />

**具体操作步骤**<br />例：我们要在原始的大根堆（[90,80,70,60,40,30,20,10,50]）中将对顶90删除

1. 将90删除
1. 用户末尾元素替换被删除的数据，同时数组长度减一
1. 40小于子节点，选择最大的子节点交换位置（一直重复这个过程，直到进行到最后）


<br />![image.png](https://cdn.nlark.com/yuque/0/2020/png/963603/1590901360463-62775e8d-87a5-45e7-9e4e-b5264eee7b1e.png#align=left&display=inline&height=355&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1136&originWidth=2384&size=627422&status=done&style=shadow&width=746)<br />

<a name="wBE98"></a>
#### 注意
二叉堆是堆（**日常工程项目中使用优先队列** **priority_queue **就可以）的一种常见且简单的实现，但是并不是最优的实现
<a name="qjQtA"></a>
####
