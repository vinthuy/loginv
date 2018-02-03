package main

type rbnode struct {
	Key    interface{}
	Val    interface{}
	color  string //黑还是红
	parent *rbnode
	left   *rbnode
	right  *rbnode
}

const (
	Red   = "red"
	Black = "black"
)

type rbtree struct {
	root *rbnode
}

type rbtreeI interface {
	insert(key interface{})
	remove(key interface{})
	foreach()
	compare(k1 interface{}, k2 interface{}) int
}

func (this *rbtree) compare(k1 int, k2 int) int {
	if k1 > k2 {
		return 1
	} else if k1 < k2 {
		return -1
	}
	return 0
}

func (this *rbtree) insert(key int) {
	var current *rbnode
	rbnode := &rbnode{
		color: Red,
		Key:   key,
	}
	rootnode := this.root

	current = rootnode
	if current == nil {
		this.root = rbnode
		rbnode.color = Black
		return
	}

	for {

		leftNode := current.left
		rightNode := current.right

		if this.compare(rbnode.Key.(int), current.Key.(int)) > 0 { //find insert posion for right
			if rightNode != nil {
				current = rightNode
			} else {
				current.right = rbnode
				rbnode.parent = current
				break
			}
		} else {
			if leftNode != nil {
				current = leftNode
			} else {
				current.left = rbnode
				rbnode.parent = current
				break
			}
		}
	}
	this.fixRbtree(rbnode)
}

func (this *rbtree) fixRbtree(cur *rbnode) {
	var optRbNode = cur
	var ancleNode *rbnode
	for {
		//case1
		// 插入节点的父节点和其叔叔节点（祖父节点的另一个子节点）均为红色的
		// 1.如果插入的是左节点
		parentNode := optRbNode.parent
		grandPaNode := parentNode.parent
		if grandPaNode == nil {
			//说明这个只有根节点，也就是说height = 2
			return
		}

		if parentNode == grandPaNode.left {
			ancleNode = grandPaNode.right
		} else {
			ancleNode = grandPaNode.left
		}
		//1-1换色
		if optRbNode.color == Red && parentNode.color == Red && ancleNode.color == Red {
			grandPaNode.color = Black
			ancleNode.color = Black
			grandPaNode.color = Red
			optRbNode = grandPaNode
		} else if optRbNode.color == Red && parentNode.color == Red && optRbNode == parentNode.right && ancleNode.color == Black && ancleNode == grandPaNode.right {
			//2.插入节点的父节点是红色，叔叔节点是黑色，且插入节点是其父节点的右子节点
			// 需要左旋
			this.rotateLeft(optRbNode)
			optRbNode = optRbNode.left
		} else if optRbNode.color == Red && parentNode.color == Red && optRbNode == parentNode.left && ancleNode.color == Black && ancleNode == grandPaNode.right {
			//3.插入节点的父节点是红色，叔叔节点是黑色，且插入节点是其父节点的左子节点。
			// parent==red grandPa == right
			parentNode.color = Black
			grandPaNode.color = Red
			this.rotateRight(optRbNode)
		}
		break
	}
}

func (this *rbtree) rotateRight(optRbNode *rbnode) {
	parentNode := optRbNode.parent
	grandPaNode := parentNode.parent

	slingNode := parentNode.right

	grandPaNode.parent = parentNode
	parentNode.right = grandPaNode

	grandPaNode.left = slingNode
	slingNode.parent = grandPaNode.left
}

func (this *rbtree) rotateLeft(optRbNode *rbnode) {
	parentNode := optRbNode.parent
	grandPaNode := parentNode.parent

	grandPaNode.left = optRbNode
	optRbNode.parent = grandPaNode

	parentNode.right = optRbNode.left
	optRbNode.left.parent = parentNode

	parentNode.parent = optRbNode
	optRbNode.left = parentNode

}
