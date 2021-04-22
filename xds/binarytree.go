package xds

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) AddLeftNode(val int) {
	if root.Left == nil {
		root.Left = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	} else {
		root.Left.Val = val
	}
}

func (root *TreeNode) AddRightNode(val int) {
	if root.Right == nil {
		root.Right = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	} else {
		root.Right.Val = val
	}
}

//允许输入数字与nil

// func coverCreateBTtoList(str string) []int {

// }

// func CreateBinaryTree(list []int) (root *TreeNode, err error) {
// 	if list == nil {
// 		return nil, nil
// 	}
// 	for _, v := range list {

// 	}
// 	return
// }

// func (tn *TreeNode) add(val int) (isSuccess bool) {
// 	if tn == nil {
// 		if tn.Left == nil {

// 		} else if tn.Right == nil {

// 		}
// 	}
// }

//中序

//先序

//后序

/*
层级遍历(队列实现方案)：
1.准备一个队列，元素为节点地址，初始放入根节点
2.每次遍历会遍历整个队列，每次遍历一层，从队列开头拿出节点，取其非空左右孩子推入队列末尾，由于golang slice的range副本机制，并不会发生队列死循环
3.遍历结束后整理队列，移除上层的所有节点
4.当queue为空时，结束遍历
*/

func LevelOrderByQueue(node *TreeNode) []int {
	list := []int{}
	if node == nil {
		return list
	}
	queue := []*TreeNode{node}
	for len(queue) != 0 {
		max := 0
		for idx, node := range queue {
			max = idx
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[max+1:]
	}
	return list
}
