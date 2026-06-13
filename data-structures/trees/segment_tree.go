package trees

type TreeOperation func(first int, second int) int

type SegenmentTree struct {
	slice []int
	op    TreeOperation
}

func NewSegmentTree(nodes []int, op TreeOperation) *SegenmentTree {
	size := padToPowerOfTwo(len(nodes))
	tree := &SegenmentTree{slice: make([]int, 2*size), op: op}

	copy(tree.slice[size:], nodes)
	tree.populate()

	return tree
}

func (tree *SegenmentTree) Update(position int, value int) {
	mid := len(tree.slice) / 2
	child := mid + position

	tree.slice[child] = value

	for parent := child / 2; parent > 0; parent >>= 1 {
		tree.slice[parent] = tree.slice[parent<<1] + tree.slice[(parent<<1)|1]
	}
}

func (tree *SegenmentTree) Query(left, right int) int {
	result := 0
	mid := len(tree.slice) / 2

	for left, right = left+mid, right+mid; left <= right; left, right = left>>1, right>>1 {
		if left&1 == 1 {
			result += tree.slice[left]
			left++
		}

		if right&1 == 0 {
			result += tree.slice[right]
			right--
		}
	}

	return result
}

func (tree *SegenmentTree) populate() {
	for parent := len(tree.slice)/2 - 1; parent > 0; parent-- {
		tree.slice[parent] = tree.op(tree.slice[parent<<1], tree.slice[parent<<1|1])
	}
}

func padToPowerOfTwo(num int) int {
	normalized := 1

	for normalized < num {
		normalized <<= 1
	}

	return normalized
}
