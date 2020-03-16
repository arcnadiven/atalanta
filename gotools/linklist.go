package gotools

type ListNode struct {
	Val  int
	Next *ListNode
}

//尾插
func CreateLinkList(nums []int) *ListNode {
	var headPoint, tailPoint *ListNode
	if len(nums) == 0 {
		return headPoint
	}
	for _, v := range nums {
		if headPoint == nil {
			headPoint = &ListNode{
				Val:  v,
				Next: nil,
			}
			tailPoint = headPoint
			continue
		}
		if tailPoint.Next == nil {
			tailPoint.Next = &ListNode{
				Val:  v,
				Next: nil,
			}
		} else {
			for n := tailPoint; n != nil; n = n.Next {
				tailPoint = n
			}
		}
		tailPoint = tailPoint.Next
	}
	return headPoint
}

func (h *ListNode) AddNode(idx, val int) {

}

func (h *ListNode) RangeLinkList() []int {
	list := []int{}
	if h == nil {
		return list
	}
	for node := h; node != nil; node = node.Next {
		list = append(list, node.Val)
	}
	return list
}
