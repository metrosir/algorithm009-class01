package Week_02_test

//【题意】
//给定一个链表，判断链表中是否有环。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
// https://leetcode-cn.com/problems/linked-list-cycle/

//输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。

//【解题方式】
//1.双指针，快慢指针，慢指针一次前进一步，快指针一次前进两步，如果快指针和慢指针相遇说明链表带环
//2.哈希表

type ListNode struct {
	Val  int
	Next *ListNode
}

//快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

//哈希表
func hasCycleHashMap(head *ListNode) bool {

	//创建一个哈希结构，key存指针，val存循环到的值
	hash_ := map[*ListNode]int{}
	for head != nil {
		if _, ok := hash_[head]; ok {
			return true
		}
		hash_[head] = head.Val
		head = head.Next
	}
	return false
}
