package xds

import "fmt"

type STreeNode struct {
	Val   string
	Left  *STreeNode
	Right *STreeNode
}

//LeetCode用，levelOrder序列化，节点缺失时使用null代替
func SMarshalBinaryTree(src []interface{}) *STreeNode {
	if len(src) == 0 {
		return nil
	}
	if src[0] == nil {
		return nil
	}
	root := &STreeNode{}
	if val, ok := src[0].(string); ok {
		root.Val = val
	} else {
		panic("root node value is not a string")
	}
	currentLevel := []*STreeNode{root} //level Order需要维护一个临时队列
	srcIdx := 1
	for srcIdx < len(src) {
		maxIdx := 0
		for curIdx, node := range currentLevel {
			//尝试分配左子节点
			if srcIdx == len(src) { //判断是否越界
				break
			}
			{
				if val, ok := src[srcIdx].(string); ok { //判断是否为int,有int则补一位
					lChild := &STreeNode{Val: val} //那我怎么关联上源地址？？？
					currentLevel = append(currentLevel, lChild)
					node.Left = lChild
				} else {
					if src[srcIdx] == nil {
					} else {
						panic(fmt.Sprintf("node: %s left child value is not number", node.Val))
					}
				}
			}
			srcIdx++

			//尝试分配右子节点
			if srcIdx == len(src) { //判断是否越界
				break
			}
			{
				if val, ok := src[srcIdx].(string); ok { //判断是否为int,有int则补一位
					rChild := &STreeNode{Val: val} //那我怎么关联上源地址？？？
					currentLevel = append(currentLevel, rChild)
					node.Right = rChild
				} else {
					if src[srcIdx] == nil {
					} else {
						panic(fmt.Sprintf("node: %s right child value is not number", node.Val))
					}
				}
			}
			srcIdx++

			//更新currentLevel临时队列的下标
			maxIdx = curIdx
		}
		currentLevel = currentLevel[maxIdx+1:]
	}
	return root
}

//自动填充null
func SLevelOrder(root *STreeNode) []interface{} {
	result := []interface{}{}
	if root == nil {
		return nil
	}
	lastLevel := []*STreeNode{root}
	for len(lastLevel) != 0 {
		lastIndex := 0
		for idx, node := range lastLevel {
			if node == nil {
				result = append(result, NULL)
				continue
			}
			result = append(result, node.Val)
			lastLevel = append(lastLevel, node.Left)
			lastLevel = append(lastLevel, node.Right)
			lastIndex = idx
		}
		lastLevel = lastLevel[lastIndex+1:]
	}
	idx := 0
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != NULL {
			idx = i
			break
		}
	}
	return result[:idx+1]
}

func SPreOrder(root *STreeNode) []string {
	if root == nil {
		return nil
	}
	result := []string{}
	result = append(result, root.Val)
	result = append(result, SPreOrder(root.Left)...)
	result = append(result, SPreOrder(root.Right)...)
	return result
}

func SInOrder(root *STreeNode) []string {
	if root == nil {
		return nil
	}
	result := []string{}
	result = append(result, SInOrder(root.Left)...)
	result = append(result, root.Val)
	result = append(result, SInOrder(root.Right)...)
	return result
}

func SPostOrder(root *STreeNode) []string {
	if root == nil {
		return nil
	}
	result := []string{}
	result = append(result, SPostOrder(root.Left)...)
	result = append(result, SPostOrder(root.Right)...)
	result = append(result, root.Val)
	return result
}
