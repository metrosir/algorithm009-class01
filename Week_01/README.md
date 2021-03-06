# 刷题大法--五毒神掌
### 刷题第一遍
1、五分钟：读题+思考（如果5分钟想不出来最多15分钟，如果还没想出来直接跳到第2步）

2、直接看解法：注意！多解法，比较解法优劣

3、背诵、默写好的解法

### 刷题第二遍
1、马上自己写 -》LeetCode 提交

2、多种解法比较、体会-》优化！

### 刷题第三遍
1、过了一天后，再重复做题

2、不同的解法的熟悉程度->专项训练

### 刷题第四遍
1、过了一周：反复回来联系相同的题目

### 刷题第五遍
1、面试前一周恢复训练



# 程序的时间、空间复杂度

### 时间复杂度
#### O(1)：常数复杂度
#### O(log(n))：对数复杂度
#### O(n)：线性时间复杂度
#### O(n^2)：平方
#### O(n^3)：立方
#### O(2^n)：指数
#### O(n!)：阶乘

### 空间复杂度
#### 数组的长度 就是空间复杂度
    一维数组的空间复杂度是 0(n)
    二维数组的空间复制度是 0(n^2)
#### 递归的深度
    递归最深的深度就是空间复杂度的最大值

### 时间、空间复杂度个人理解
    时间复杂度就是程序在运行过程中，操作当前代码的次数，如果操作了常数次时间复杂度就是0(1)；如果操作了可变参数次就是O(n)；如果操作了操作了可变参数一半或其他倍数的次数则是O(log(n))；如果在程序中操作了两个嵌套，每次嵌套循环可变参n次 则时间复杂度是O(n^2)，如果嵌套了三层则是O(n^3)；如果用到的是递归函数则时间复杂度是2^n。空间复杂度和时间复杂度有点类似；程序中有一个一位数组，数组的长度是n，则这个数组的空间复杂度是O(n)，如果是一个二维数组空间复杂度则是(n^2)；同时也了解了fibonacci数列的时间、空间复杂度的计算，在实现fibonacci数列的时候，在循环到相同节点进行结算的时候取上一个已经运行过的节点的结果，不需要再继续循环，这样既减少了时间复杂度也减少了空间复杂度。


# 数组
    数组里面的原始的类型没有严格的要求，比较多样化（泛型），任何一个单元的类型都可以放进去。数组底层的实现有一个内存管理器的东西，每当申请一个数组计算机实际上是在内存中开辟了一段连续的地址，每一个地址可以通过内存管理器进行访问。在访问数组中的任意元素的时候时间复杂度都是一样的O(1)，随机的访问，访问的时间特别快。数组这种结构的问题关键在于 在新增和删除操作数组的时候会进行一些相应的操作。在新增、删除数组元素的时候都是O(n)的时间复杂度。

### 增加元素 O(n)
### 删除元素 O(n)
### 查询元素 O(1)

# 链表
    1、在添加、删除比较频繁的情况下，数组在这个时候并不好用
    2、linked list的原始在定义好之后有一个value 和 一个next，next指向下一个元素，串在一起就变成了类似于数组的结构，它的每一个元素必须用class来定义，这个class一般叫node，一个成员变量是value（这个成员变量也可以是一个类 ，比较丰富一点，它可以有很多的值），一个成员变量是next （指针，指向下一个元素），串在一起就是一个链表。
    3、如果只有一个next指针叫做单链表
    4、next指针同时往前和后面指这个链表就叫做双向链表，前面个next叫做它的先前指针 （prev或者previous）
    5、头指针一般用head来表示，尾指针用tail来表示
    6、最后一个元素它的next指针指向空，因为没有next指针了
    7、如果tail指针的next指向了head这个叫循环链表

### 增加元素 O(1)
### 删除元素 O(1)
### 查询元素 O(n)


# Skip List（跳表）
    1、跳表是基于链表的，跳表的使用只能用于链表中的元素有序的情况下（跳表中的元素始终是有序的，不然就不叫跳表），所以跳表对标的是平衡树（也就是二叉搜索树中的平衡树和二分查找），跳表指的是一种插入、删除、搜索都是O(log(n))的数据结构。
    2、跳表最大的优势是原理简单、容易实现、方便扩展、效率高效。Redis、LevelDB使用的是跳表（替代了平衡树）。
    3、跳表是用来替代 平衡树、二分查找的；平衡树和二分查找也是在元素有序的情况下怎么有效的查找需要的元素。
    4、在跳表中增加、删除元素的时候，需要调整索引，这种情况下时间复杂度是 O(log(n))。

### 跳表查询的时间复杂度 O(log(n))


# 栈和队列
### 栈 stack
    先入后出
    添加删除皆为 O(1)
    查询为O(n)，因为无序

### 队列
    先入先出
    添加删除皆为 O(1)
    查询为O(n)，因为无序

### 双端队列 deque （double-end queue）
    栈和队列的结合体
    可以往前面push进去，也可以从前面pop出来
    可以往后面push进去，也可以从后面pop出来
    插入和删除都是O(1)
    查询为O(n)，因为无序

### 优先队列（priority queue）
    插入操作：O(1)
    取出操作：O(log n) 按照元素的优先级取出（如VIP先行）
    底层具体实现的数据结构较为多样和复杂：Heap（堆）、bst、treap
