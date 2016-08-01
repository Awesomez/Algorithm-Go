package tree

/*
根节点的左子树上所有的节点的值都小于根节点的值
根节点的右子树上所有的节点的值都大于根节点的值
每一棵子树本身也是一个二叉查找树，也就是说，左子树上的节点的值都小于该子树根节点的值，右子树上的节点的值都大于该子树根节点的值。

还有一个特性就是”中序遍历“可以让结点有序
 */

type TreeNode struct {
	data int
	left *TreeNode
	right *TreeNode
	parent *TreeNode
}

type BSTree struct {
	root *TreeNode
}

func (tree *BSTree) Add(data int){
	if tree.root==nil {
		tree.root=&TreeNode{data:data}
	}else{
		current:=tree.root
		for current!=nil{
			if current.data>data {
				current=current.left
			}else if current.data<data{
				current=current.right
			}else{
				return
			}
		}
		if current.data>data{
			current.left=&TreeNode{data:data,parent:current}
		}else{
			current.right=&TreeNode{data:data,parent:current}
		}
	}
}

/**
前驱
我们可以对一个二叉搜索树直接进行中序遍历，立马可以得到节点的前驱和后继节点，
但是这样的方法时间复杂度为O(N),显然不是最好的方法。
 */
func (tree BSTree) GetPredecessor(data int) *TreeNode{
	cur:=tree.Search(data)
	if cur==nil{
		return
	}
	//1. 如果x存在左孩子，则"x的前驱结点"为 "以其左孩子为根的子树的最大结点"。
	if cur.left != nil{
		return tree.GetMax(cur.left)
	}
	//2.1 若该节点是其父节点的右边孩子，那么该节点的前驱结点即为其父节点。
	//2.2 若该节点是其父节点的左边孩子，那么需要沿着其父亲节点一直向树的顶端寻找，
	// 直到找到一个节点P，P节点也是其父节点Q的右孩子,那么Q就是前驱节点
	for cur!=nil {
		if cur==cur.parent.right{
			return cur.parent
		}
		cur=cur.parent
	}
	return nil
}

func (tree *BSTree) Delete(data int) {
	node :=tree.Search(data)

	if node==nil{
		return
	}

	//1.如果要删除的节点没有孩子，直接删掉它就可以
	if node.left==nil && node.right==nil {
		if(node==tree.root){
			tree.root=nil
		}else{
			if node.parent.right==node {
				node.parent.right=nil
			}else{
				node.parent.left=nil
			}
		}
	//2.如果要删除的节点既有左孩子又有右孩子,就把这个节点的直接后继或直接前继的值赋给这个节点，然后删除直接后继节点即可
	}else if node.left!=nil && node.right!=nil {
		successor:=tree.GetSuccessor(node.data)
		node.data=successor.data
		tree.Delete(successor)
	//3.单子树,有左孩子那就把左孩子顶上去，有右孩子就把右孩子顶上去
	}else{
		needLinkedNode := node.right
		if node.left!=nil {
			needLinkedNode=node.left
		}

		//若该节点为根
		if node.root==node {
			node.root=needLinkedNode
			return true
		}

		//该节点为父节点右孩子
		if node==node.parent.right {
			node.parent.right=needLinkedNode
		}else{
			node.parent.left=needLinkedNode
		}
		return true
	}

}

//后继
func (tree BSTree) GetSuccessor(data int) *TreeNode{
	cur:=tree.Search(data)
	if cur==nil{
		return
	}
	if cur.right!=nil{
		return tree.GetMin(cur.right)
	}

	for cur!=nil{
		if cur==cur.parent.left{
			return cur.parent
		}
		cur=cur.parent
	}
}

func (tree BSTree) GetMax(node *TreeNode) *TreeNode{
	if node == nil {
		return
	}
	current := node
	for {
		if current.right == nil {
			return current
		}else{
			current=current.right
		}
	}
}

func (tree BSTree) GetMin(node *TreeNode) int{
	if node == nil {
		return
	}
	current :=node
	for {
		if current.left == nil{
			return current.data
		}else{
			current=current.left
		}
	}
}

func (tree BSTree) Search(data int) *TreeNode{
	if tree.root == nil {
		return
	}
	cur:=tree.root
	for {
		if cur==nil{
			return
		}
		if data<cur.data {
			cur=cur.left
		} else if data>cur.data{
			cur=cur.right
		}else {
			return cur
		}
	}
}

func (tree BSTree) IsEmpty() bool{
	return tree.root==nil
}

func (tree *BSTree) Clear(){
	tree.root=nil
}

/**
前序遍历：根节点->左子树->右子树
中序遍历：左子树->根节点->右子树
后序遍历：左子树->右子树->根节点
 */
/**
中序遍历
 */
func (tree BSTree) InOrder(node *TreeNode){
    if node!=nil {
        tree.InOrder(node.left)
        println(node.data)
        tree.InOrder(node.right)
    }
}

func (tree BSTree) PreOrder(node *TreeNode){
    if node!=nil {
        println(node.data)
        tree.PreOrder(node.left)
        tree.PreOrder(node.right)
    }
}

func (tree BSTree) PostOrder(node *TreeNode){
    if node!=nil {
        tree.PostOrder(node.left)
        tree.PostOrder(node.right)
        println(node.data)
    }
}

