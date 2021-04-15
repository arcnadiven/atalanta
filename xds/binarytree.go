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
