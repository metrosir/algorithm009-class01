package Week_02_test

import (
	"fmt"
	"testing"
)

//表示二叉堆
var d int = 2

//存放堆的数组
var heap []int

//堆的长度
var heapSize int

func binaryHeap(capacity int) {
	heapSize = 0
	//heap = make([]int, capacity)
	heap = []int{}
}

//判断堆是否为空
func isEmpty() bool {
	return false
	return heapSize == 0
}

//判断堆是否满了
func isFull() bool {
	return false
	return heapSize == len(heap)
}

//获取父亲节点
func parent(i int) int {
	return (i - 1) / d
}

//获取儿子节点
func children(i int, k int) int {
	return d*i + k
}

//执行插入操作
//步骤：
//1.把新值存放在原堆的尾部
//2.堆数组的长度加1
//3.依次调整堆的结构，知道根即可
func insert(x int) {
	if isFull() {
		panic("heap is full")
	}
	heap = append(heap, x)
	heapSize++
	heapifyUp(heapSize - 1)
}

//插入元素时重新维护堆的顺序
func heapifyUp(i int) {
	insertval := heap[i]
	for insertval > heap[parent(i)] {
		//子节点比父节点大，则替换父子节点的值
		heap[i] = heap[parent(i)]
		i = parent(i)
	}
	heap[i] = insertval
}

//执行删除操作
//步骤
//1.将堆顶的元素删除
//2.将堆尾的数据填充到堆顶，堆数组长度减一
//3.依次调整堆的结构，将子节点较大的值替换到将要调整的位置
func delete(x int) int {
	if isEmpty() {
		panic("heap is empty")
	}
	maxVal := heap[x]
	//fmt.Println(heap[x])
	heap[x] = heap[heapSize-1]
	//heap = heap[:heapSize - 1]
	heapSize--
	heapifyDown(x)
	heap = heap[:heapSize]
	return maxVal
}

//删除元素时重新维护堆的顺序
func heapifyDown(i int) {
	var childr int
	//需要处理的节点值
	tmpVal := heap[i]
	for children(i, 1) < heapSize {
		//获取儿子节点中较大的节点
		childr = maxChildren(i)
		if tmpVal >= heap[childr] {
			break
		}
		heap[i] = heap[childr]
		i = childr
	}
	heap[i] = tmpVal
}

//获取子节点中较大的值的下标
func maxChildren(i int) int {
	leftC := children(i, 1)
	rightC := children(i, 2)
	if heap[leftC] > heap[rightC] {
		return leftC
	} else {
		return rightC
	}
}

func printHeap() {
	fmt.Println(heap)
}

func findMax() int {
	if isEmpty() {
		panic("head is empty")
	}
	return heap[0]
}

func TestHeap(t *testing.T) {
	binaryHeap(10)
	insert(10)
	insert(4)
	insert(9)
	insert(1)
	insert(7)
	insert(5)
	insert(3)

	printHeap()
	delete(0)
	printHeap()
	//delete(2)
	//printHeap()
}
