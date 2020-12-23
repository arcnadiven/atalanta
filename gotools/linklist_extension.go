package gotools

//必须是两个有序链表

func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l1node := l1
	l2node := l2
	for l2node != nil {
		if l2node.Val < l1node.Val { //该节点比l1node值小
			l2 = l2.Next
			l2node.Next = l1
			l1 = l2node
			l1node = l1
		} else {                //该节点比l1node值大
			for l1node != nil { //遍历l1	每次遍历结束之时只会变动一个节点
				if l1node.Next == nil {
					l1node.Next = l2
					return l1
				}
				if l2node.Val >= l1node.Val && l2node.Val < l1node.Next.Val {
					l2 = l2.Next
					l2node.Next = l1node.Next
					l1node.Next = l2node
					l1node = l1node.Next
					break
				}
				l1node = l1node.Next
			}
		}
		l2node = l2
	}
	return l1
}
