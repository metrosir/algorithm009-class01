# 1. 分治、回溯的实现和特性

参考：[https://www.jianshu.com/p/3b0656fa21a9](https://www.jianshu.com/p/3b0656fa21a9)<br />
<br />遇到问题找重复性，重复性有**最近重复性**和**最优重复性**，**最优重复性就是动态规划。**最近的重复性根据重复性怎么构造以及怎么分解，就有了分治、回溯…… 本质上就是一种递归（找重复性）<br />

<a name="nLjhW"></a>
# 分治的思想
在递归状态树的时候，把递归的问题化解成各个子问题，分解成问题<br />

<a name="LGdsL"></a>
## 泛型递归模板
```go
func recursion(level, res, param, ...) {
    //递归终结条件
    if level == maxLevel {
        return
    }
    //处理逻辑
    process()
    //递归，下探到下一层
    recursion(level+1, res, param, ...)
    //清理当前层
}
```

<br />

<a name="9XQZ7"></a>
## 分治的思想


<a name="aHZAw"></a>
# 分治代码模板
```go
func recursion(level, res, param, ...) {
    //第1步 递归终结条件
    if level == maxLevel {
    	return
    }
    
    //第2步 准备数据
    data := prepare_data()
    //第3步 解决问题
    sub1 = process(data[1])
    sub2 = process(data[2])
    ...
    
    //第4步 处理并生成最终结果
    result = process_result(sub1,sub2,...)
    
    //第5步 清理当前层
}
```


<a name="FOv1K"></a>
# 回溯的思想
递归的时候去每一层中去尝试，看有没有符合的结果<br />

<a name="tnfcT"></a>
## 回溯的代码模板跟泛型递归的一样



# 1. 递归的实现、特性以及思维要点

<a name="wwrPA"></a>
# 泛型递归代码模板
```go
func recursion(level int, params,params,...) {
    //1.递归终结条件
    if level > MAX_LEVEL {
    	return
    }
    //2.处理当前层的逻辑
    process(level, data...)
    //3.下探到下一层，
    recursion(level + 1, p1, p2)
    //4.最后异步 清理当前层
}
```

<br />思维要点

1. 不要人肉进行递归（最大的的误区）
1. 找到最近最简的方法，将其拆解成可重复的解决的问题（重复子问题）
1. 数学归纳法的思维
