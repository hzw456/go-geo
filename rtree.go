package geo

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type Rtree struct {
	MinChildren int
	MaxChildren int
	root        *node
	size        int
	height      int
}

func NewRandomIDTree(objs ...Geometry) *Rtree {
	minChildren, maxChildren := 0, 0
	if len(objs) < 10 {
		minChildren = 2
		maxChildren = 10
	} else {
		minChildren = len(objs) / 6
		maxChildren = len(objs) / 2
	}
	rid := time.Now().UnixNano()
	var sps []Spatial
	for i, v := range objs {
		sps = append(sps, Spatial{fmt.Sprintf("%d_0_%d", rid, i), nil, v})
	}
	return NewMinMaxTree(minChildren, maxChildren, sps...)
}

func NewTree(objs ...Spatial) *Rtree {
	minChildren, maxChildren := 0, 0
	if len(objs) < 10 {
		minChildren = 2
		maxChildren = 10
	} else {
		minChildren = len(objs) / 6
		maxChildren = len(objs) / 2
	}
	return NewMinMaxTree(minChildren, maxChildren, objs...)
}

func NewTreeWith(objs ...Spatial) *Rtree {
	minChildren, maxChildren := 0, 0
	if len(objs) < 10 {
		minChildren = 2
		maxChildren = 10
	} else {
		minChildren = len(objs) / 6
		maxChildren = len(objs) / 2
	}
	return NewMinMaxTree(minChildren, maxChildren, objs...)
}

// 需指定 min max孩子数的tree
func NewMinMaxTree(min, max int, objs ...Spatial) *Rtree {
	rt := &Rtree{
		MinChildren: min,
		MaxChildren: max,
		height:      1,
		root: &node{
			entries: []entry{},
			leaf:    true,
			level:   1,
		},
	}

	if len(objs) <= rt.MaxChildren {
		for _, obj := range objs {
			rt.Insert(obj)
		}
	} else {
		rt.bulkLoad(objs)
	}

	return rt
}

// Size returns the number of objects currently stored in tree.
func (tree *Rtree) Size() int {
	return tree.size
}

// Depth returns the maximum depth of tree.
func (tree *Rtree) Depth() int {
	return tree.height
}

type dimSorter struct {
	dim  int
	objs []entry
}

func (s *dimSorter) Len() int {
	return len(s.objs)
}

func (s *dimSorter) Swap(i, j int) {
	s.objs[i], s.objs[j] = s.objs[j], s.objs[i]
}

func (s *dimSorter) Less(i, j int) bool {
	if s.dim == 0 {
		return s.objs[i].bb.MinX < s.objs[j].bb.MinX
	} else {
		return s.objs[i].bb.MaxX < s.objs[j].bb.MaxX
	}
}

// walkPartitions splits objs into slices of maximum k elements and
// iterates over these partitions.
func walkPartitions(k int, objs []entry, iter func(parts []entry)) {
	n := (len(objs) + k - 1) / k // ceil(len(objs) / k)

	for i := 1; i < n; i++ {
		iter(objs[(i-1)*k : i*k])
	}
	iter(objs[(n-1)*k:])
}

func sortByDim(dim int, objs []entry) {
	sort.Sort(&dimSorter{dim, objs})
}

// bulkLoad bulk loads the Rtree using OMT algorithm. bulkLoad contains special
// handling for the root node.
func (tree *Rtree) bulkLoad(objs []Spatial) {
	n := len(objs)

	// create entries for all the objects
	entries := make([]entry, n)
	for i := range objs {
		entries[i] = entry{
			bb:  BoundingBox(objs[i].Geom),
			obj: objs[i],
		}
	}

	// following equations are defined in the paper describing OMT
	var (
		N = float64(n)
		M = float64(tree.MaxChildren)
	)
	// Eq1: height of the tree
	// use log2 instead of log due to rounding errors with log,
	// eg, math.Log(9) / math.Log(3) > 2
	h := math.Ceil(math.Log2(N) / math.Log2(M))

	// Eq2: size of subtrees at the root
	nsub := math.Pow(M, h-1)

	// Inner Eq3: number of subtrees at the root
	s := math.Ceil(N / nsub)

	// Eq3: number of slices
	S := math.Floor(math.Sqrt(s))

	// sort all entries by first dimension
	sortByDim(0, entries)

	tree.height = int(h)
	tree.size = n
	tree.root = tree.omt(int(h), int(S), entries, int(s))
}

// omt is the recursive part of the Overlap Minimizing Top-loading bulk-
// load approach. Returns the root node of a subtree.
func (tree *Rtree) omt(level, nSlices int, objs []entry, m int) *node {
	// if number of objects is less than or equal than max children per leaf,
	// we need to create a leaf node
	if len(objs) <= m {
		// as long as the recursion is not at the leaf, call it again
		if level > 1 {
			child := tree.omt(level-1, nSlices, objs, m)
			n := &node{
				level: level,
				entries: []entry{{
					bb:    child.computeBoundingBox(),
					child: child,
				}},
			}
			child.parent = n
			return n
		}
		return &node{
			leaf:    true,
			entries: objs,
			level:   level,
		}
	}

	n := &node{
		level:   level,
		entries: make([]entry, 0, m),
	}

	// maximum node size given at most M nodes at this level
	k := (len(objs) + m - 1) / m // = ceil(N / M)

	// In the root level, split objs in nSlices. In all other levels,
	// we use a single slice.
	vertSize := len(objs)
	if nSlices > 1 {
		vertSize = nSlices * k
	}

	// create sub trees
	walkPartitions(vertSize, objs, func(vert []entry) {
		// sort vertical slice by a different dimension on every level
		sortByDim((tree.height-level+1)%2, vert)

		// split slice into groups of size k
		walkPartitions(k, vert, func(part []entry) {
			child := tree.omt(level-1, 1, part, tree.MaxChildren)
			child.parent = n

			n.entries = append(n.entries, entry{
				bb:    child.computeBoundingBox(),
				child: child,
			})
		})
	})
	return n
}

// node represents a tree node of an Rtree.
type node struct {
	parent  *node
	leaf    bool
	entries []entry
	level   int // node depth in the Rtree
}

func (n *node) String() string {
	return fmt.Sprintf("node{leaf: %v, entries: %v}", n.leaf, n.entries)
}

// entry represents a spatial index record stored in a tree node.
type entry struct {
	bb    *Box // bounding-box of all children of this entry
	child *node
	obj   Spatial
}

func (e entry) String() string {
	if e.child != nil {
		return fmt.Sprintf("entry{bb: %v, child: %v}", e.bb, e.child)
	}
	return fmt.Sprintf("entry{bb: %v, obj: %v}", e.bb, e.obj)
}

// Spatial is an interface for objects that can be stored in an Rtree and queried.
type Spatial struct {
	ID         string
	Properties interface{}
	Geom       Geometry
}

// 插入rtree
func (tree *Rtree) Insert(obj Spatial) {
	e := entry{BoundingBox(obj.Geom), nil, obj}
	tree.insert(e, 1)
	tree.size++
}

// 插入rtree
func (tree *Rtree) InsertWithRandomId(geom Geometry) {
	e := entry{BoundingBox(geom), nil, Spatial{fmt.Sprintf("%d_1_%d", time.Now().UnixNano(), tree.Size), nil, geom}}
	tree.insert(e, 1)
	tree.size++
}

// insert adds the specified entry to the tree at the specified level.
func (tree *Rtree) insert(e entry, level int) {
	leaf := tree.chooseNode(tree.root, e, level)
	leaf.entries = append(leaf.entries, e)

	// update parent pointer if necessary
	if e.child != nil {
		e.child.parent = leaf
	}

	// split leaf if overflows
	var split *node
	if len(leaf.entries) > tree.MaxChildren {
		leaf, split = leaf.split(tree.MinChildren)
	}
	root, splitRoot := tree.adjustTree(leaf, split)
	if splitRoot != nil {
		oldRoot := root
		tree.height++
		tree.root = &node{
			parent: nil,
			level:  tree.height,
			entries: []entry{
				{bb: oldRoot.computeBoundingBox(), child: oldRoot},
				{bb: splitRoot.computeBoundingBox(), child: splitRoot},
			},
		}
		oldRoot.parent = tree.root
		splitRoot.parent = tree.root
	}
}

// chooseNode finds the node at the specified level to which e should be added.
func (tree *Rtree) chooseNode(n *node, e entry, level int) *node {
	if n.leaf || n.level == level {
		return n
	}

	// find the entry whose bb needs least enlargement to include obj
	diff := math.MaxFloat64
	var chosen entry
	for _, en := range n.entries {
		bb := en.bb.Union(e.bb)
		d := bb.Size() - en.bb.Size()
		if d < diff || (d == diff && en.bb.Size() < chosen.bb.Size()) {
			diff = d
			chosen = en
		}
	}

	return tree.chooseNode(chosen.child, e, level)
}

// adjustTree splits overflowing nodes and propagates the changes upwards.
func (tree *Rtree) adjustTree(n, nn *node) (*node, *node) {
	// Let the caller handle root adjustments.
	if n == tree.root {
		return n, nn
	}

	// Re-size the bounding box of n to account for lower-level changes.
	en := n.getEntry()
	en.bb = n.computeBoundingBox()

	// If nn is nil, then we're just propagating changes upwards.
	if nn == nil {
		return tree.adjustTree(n.parent, nil)
	}

	// Otherwise, these are two nodes resulting from a split.
	// n was reused as the "left" node, but we need to add nn to n.parent.
	enn := entry{nn.computeBoundingBox(), nn, Spatial{}}
	n.parent.entries = append(n.parent.entries, enn)

	// If the new entry overflows the parent, split the parent and propagate.
	if len(n.parent.entries) > tree.MaxChildren {
		return tree.adjustTree(n.parent.split(tree.MinChildren))
	}

	// Otherwise keep propagating changes upwards.
	return tree.adjustTree(n.parent, nil)
}

// getEntry returns a pointer to the entry for the node n from n's parent.
func (n *node) getEntry() *entry {
	var e *entry
	for i := range n.parent.entries {
		if n.parent.entries[i].child == n {
			e = &n.parent.entries[i]
			break
		}
	}
	return e
}

// computeBoundingBox finds the MBR of the children of n.
func (n *node) computeBoundingBox() (bb *Box) {
	childBoxes := make([]*Box, len(n.entries))
	for i, e := range n.entries {
		childBoxes[i] = e.bb
	}
	bb = BoxUnion(childBoxes...)
	return
}

// split splits a node into two groups while attempting to minimize the
// bounding-box area of the resulting groups.
func (n *node) split(minGroupSize int) (left, right *node) {
	// find the initial split
	l, r := n.pickSeeds()
	leftSeed, rightSeed := n.entries[l], n.entries[r]

	// get the entries to be divided between left and right
	remaining := append(n.entries[:l], n.entries[l+1:r]...)
	remaining = append(remaining, n.entries[r+1:]...)

	// setup the new split nodes, but re-use n as the left node
	left = n
	left.entries = []entry{leftSeed}
	right = &node{
		parent:  n.parent,
		leaf:    n.leaf,
		level:   n.level,
		entries: []entry{rightSeed},
	}

	// TODO
	if rightSeed.child != nil {
		rightSeed.child.parent = right
	}
	if leftSeed.child != nil {
		leftSeed.child.parent = left
	}

	// distribute all of n's old entries into left and right.
	for len(remaining) > 0 {
		next := pickNext(left, right, remaining)
		e := remaining[next]

		if len(remaining)+len(left.entries) <= minGroupSize {
			assign(e, left)
		} else if len(remaining)+len(right.entries) <= minGroupSize {
			assign(e, right)
		} else {
			assignGroup(e, left, right)
		}

		remaining = append(remaining[:next], remaining[next+1:]...)
	}

	return
}

// getAllBoundingBoxes traverses tree populating slice of bounding boxes of non-leaf nodes.
func (n *node) getAllBoundingBoxes() []*Box {
	var Boxs []*Box
	if n.leaf {
		return Boxs
	}
	for _, e := range n.entries {
		if e.child == nil {
			return Boxs
		}
		BoxsInter := append(e.child.getAllBoundingBoxes(), e.bb)
		Boxs = append(Boxs, BoxsInter...)
	}
	return Boxs
}

func assign(e entry, group *node) {
	if e.child != nil {
		e.child.parent = group
	}
	group.entries = append(group.entries, e)
}

// assignGroup chooses one of two groups to which a node should be added.
func assignGroup(e entry, left, right *node) {
	leftBB := left.computeBoundingBox()
	rightBB := right.computeBoundingBox()
	leftEnlarged := leftBB.Union(e.bb)
	rightEnlarged := rightBB.Union(e.bb)

	// first, choose the group that needs the least enlargement
	leftDiff := leftEnlarged.Size() - leftBB.Size()
	rightDiff := rightEnlarged.Size() - rightBB.Size()
	if diff := leftDiff - rightDiff; diff < 0 {
		assign(e, left)
		return
	} else if diff > 0 {
		assign(e, right)
		return
	}

	// next, choose the group that has smaller area
	if diff := leftBB.Size() - rightBB.Size(); diff < 0 {
		assign(e, left)
		return
	} else if diff > 0 {
		assign(e, right)
		return
	}

	// next, choose the group with fewer entries
	if diff := len(left.entries) - len(right.entries); diff <= 0 {
		assign(e, left)
		return
	}
	assign(e, right)
}

// pickSeeds chooses two child entries of n to start a split.
func (n *node) pickSeeds() (int, int) {
	left, right := 0, 1
	maxWastedSpace := -1.0
	for i, e1 := range n.entries {
		for j, e2 := range n.entries[i+1:] {
			d := e1.bb.Union(e2.bb).Size() - e1.bb.Size() - e2.bb.Size()
			if d > maxWastedSpace {
				maxWastedSpace = d
				left, right = i, j+i+1
			}
		}
	}
	return left, right
}

// pickNext chooses an entry to be added to an entry group.
func pickNext(left, right *node, entries []entry) (next int) {
	maxDiff := -1.0
	leftBB := left.computeBoundingBox()
	rightBB := right.computeBoundingBox()
	for i, e := range entries {
		d1 := leftBB.Union(e.bb).Size() - leftBB.Size()
		d2 := rightBB.Union(e.bb).Size() - rightBB.Size()
		d := math.Abs(d1 - d2)
		if d > maxDiff {
			maxDiff = d
			next = i
		}
	}
	return
}

// Deletion
type Condition func(object ...Spatial) bool

// Delete removes an object from the tree.  If the object is not found, returns
// false, otherwise returns true. Uses the default comparator when checking
// equality.
//
func (tree *Rtree) Delete(obj Spatial) bool {
	condition := func(object ...Spatial) bool {
		if len(object) != 2 {
			return false
		}
		return object[0] == object[1]
	}
	return tree.DeleteWithComparator(obj, condition)
}

func (tree *Rtree) DeleteBySpatialID(id string) bool {
	condition := func(object ...Spatial) bool {
		if len(object) != 2 {
			return false
		}
		return object[0].ID == object[1].ID
	}
	sps := tree.SearchBySpatialID(id)
	if len(sps) == 0 {
		return false
	}
	for _, sp := range sps {
		tree.DeleteWithComparator(sp, condition)
	}
	return true
}

// todo 删除box内的geom
// func (tree *Rtree) DeleteWhereInBox(bb *Box, condition Condition) bool {
// 	for _, e := range tree.root.entries {
// 		if !e.bb.Intersect(bb) {
// 			continue
// 		}

// 		if !n.leaf {
// 			results = tree.searchIntersectWithCond(results, e.child, bb, cond)
// 			continue
// 		}
// 		if cond != nil {
// 			if cond(e.obj) {
// 				results = append(results, e.obj)
// 			}
// 		} else {
// 			results = append(results, e.obj)
// 		}
// 	}

// }

// DeleteWithComparator removes an object from the tree using a custom
// comparator for evaluating equalness. This is useful when you want to remove
// an object from a tree but don't have a pointer to the original object
// anymore.
func (tree *Rtree) DeleteWithComparator(obj Spatial, cond Condition) bool {
	n := tree.findLeaf(tree.root, obj, cond)
	if n == nil {
		return false
	}

	ind := -1
	for i, e := range n.entries {
		if cond(e.obj, obj) {
			ind = i
		}
	}
	if ind < 0 {
		return false
	}

	n.entries = append(n.entries[:ind], n.entries[ind+1:]...)

	tree.condenseTree(n)
	tree.size--

	if !tree.root.leaf && len(tree.root.entries) == 1 {
		tree.root = tree.root.entries[0].child
	}

	tree.height = tree.root.level

	return true
}

// findLeaf finds the leaf node containing obj.
func (tree *Rtree) findLeaf(n *node, obj Spatial, cmp Condition) *node {
	if n.leaf {
		return n
	}
	// if not leaf, search all candidate subtrees
	for _, e := range n.entries {
		if e.bb.Contain(BoundingBox(obj.Geom)) {
			leaf := tree.findLeaf(e.child, obj, cmp)
			if leaf == nil {
				continue
			}
			// check if the leaf actually contains the object
			for _, leafEntry := range leaf.entries {
				if cmp(leafEntry.obj, obj) {
					return leaf
				}
			}
		}
	}
	return nil
}

// condenseTree deletes underflowing nodes and propagates the changes upwards.
func (tree *Rtree) condenseTree(n *node) {
	deleted := []*node{}

	for n != tree.root {
		if len(n.entries) < tree.MinChildren {
			// remove n from parent entries
			entries := []entry{}
			for _, e := range n.parent.entries {
				if e.child != n {
					entries = append(entries, e)
				}
			}
			if len(n.parent.entries) == len(entries) {
				panic(fmt.Errorf("Failed to remove entry from parent"))
			}
			n.parent.entries = entries

			// only add n to deleted if it still has children
			if len(n.entries) > 0 {
				deleted = append(deleted, n)
			}
		} else {
			// just a child entry deletion, no underflow
			n.getEntry().bb = n.computeBoundingBox()
		}
		n = n.parent
	}

	for _, n := range deleted {
		// reinsert entry so that it will remain at the same level as before
		e := entry{n.computeBoundingBox(), n, Spatial{}}
		tree.insert(e, n.level+1)
	}
}

// Searching

func (tree *Rtree) SearchIntersect(bb *Box) []Spatial {
	return tree.searchIntersectWithCond([]Spatial{}, tree.root, bb, nil)
}

func (tree *Rtree) SearchIntersectWithCond(bb *Box, cond Condition) []Spatial {
	return tree.searchIntersectWithCond([]Spatial{}, tree.root, bb, cond)
}

func (tree *Rtree) searchIntersectWithCond(results []Spatial, n *node, bb *Box, cond Condition) []Spatial {
	for _, e := range n.entries {
		if !e.bb.Intersect(bb) {
			continue
		}

		if !n.leaf {
			results = tree.searchIntersectWithCond(results, e.child, bb, cond)
			continue
		}
		if cond != nil {
			if cond(e.obj) {
				results = append(results, e.obj)
			}
		} else {
			results = append(results, e.obj)
		}
	}
	return results
}

func (tree *Rtree) SearchBySpatialID(id string) []Spatial {
	condition := func(object ...Spatial) bool {
		return id == object[0].ID
	}
	res := tree.SearchIntersectWithCond(tree.root.computeBoundingBox(), condition)
	return res
}

// GetAllBoundingBoxes returning slice of bounding boxes by traversing tree. Slice
// includes bounding boxes from all non-leaf nodes.
func (tree *Rtree) GetAllBoundingBoxes() []*Box {
	var Boxs []*Box
	if tree.root != nil {
		Boxs = tree.root.getAllBoundingBoxes()
	}
	return Boxs
}

// utilities for sorting slices of entries

type entrySlice struct {
	entries []entry
	dists   []float64
}

func (s entrySlice) Len() int { return len(s.entries) }

func (s entrySlice) Swap(i, j int) {
	s.entries[i], s.entries[j] = s.entries[j], s.entries[i]
	s.dists[i], s.dists[j] = s.dists[j], s.dists[i]
}

func (s entrySlice) Less(i, j int) bool {
	return s.dists[i] < s.dists[j]
}

func sortEntries(p Point, entries []entry) ([]entry, []float64) {
	sorted := make([]entry, len(entries))
	dists := make([]float64, len(entries))
	return sortPreallocEntries(p, entries, sorted, dists)
}

func sortPreallocEntries(p Point, entries, sorted []entry, dists []float64) ([]entry, []float64) {
	// use preallocated slices
	sorted = sorted[:len(entries)]
	dists = dists[:len(entries)]

	for i := 0; i < len(entries); i++ {
		sorted[i] = entries[i]
		dists[i] = p.minDist(entries[i].bb)
	}
	sort.Sort(entrySlice{sorted, dists})
	return sorted, dists
}

func pruneEntriesMinDist(d float64, entries []entry, minDists []float64) []entry {
	var i int
	for ; i < len(entries); i++ {
		if minDists[i] > d {
			break
		}
	}
	return entries[:i]
}

func (p *Point) minDist(r *Box) float64 {

	sum := 0.0
	if p.X < r.MinX {
		d := p.X - r.MinX
		sum += d * d
	} else if p.X > r.MaxX {
		d := p.X - r.MaxX
		sum += d * d
	} else {
		sum += 0
	}
	return sum
}

// NearestNeighbors gets the closest Spatials to the Point.
func (tree *Rtree) NearestNeighbors(k int, p Point) []Spatial {
	maxBufSize := tree.MaxChildren * tree.Depth()
	branches := make([]entry, maxBufSize)
	branchDists := make([]float64, maxBufSize)

	// allocate the buffers for the results
	dists := make([]float64, 0, k)
	objs := make([]Spatial, 0, k)

	objs, _, _ = tree.nearestNeighbors(k, p, tree.root, dists, objs, branches, branchDists)
	return objs
}

// insert obj into nearest and return the first k elements in increasing order.
func insertNearest(k int, dists []float64, nearest []Spatial, dist float64, obj Spatial) ([]float64, []Spatial, bool) {
	i := sort.SearchFloat64s(dists, dist)
	for i < len(nearest) && dist >= dists[i] {
		i++
	}
	if i >= k {
		return dists, nearest, false
	}

	// no resize since cap = k
	if len(nearest) < k {
		dists = append(dists, 0)
		nearest = append(nearest, Spatial{})
	}

	left, right := dists[:i], dists[i:len(dists)-1]
	copy(dists, left)
	copy(dists[i+1:], right)
	dists[i] = dist

	leftObjs, rightObjs := nearest[:i], nearest[i:len(nearest)-1]
	copy(nearest, leftObjs)
	copy(nearest[i+1:], rightObjs)
	nearest[i] = obj

	return dists, nearest, false
}

func (tree *Rtree) nearestNeighbors(k int, p Point, n *node, dists []float64, nearest []Spatial, b []entry, bd []float64) ([]Spatial, []float64, bool) {
	var abort bool
	if n.leaf {
		for _, e := range n.entries {
			dist := p.minDist(e.bb)
			dists, nearest, abort = insertNearest(k, dists, nearest, dist, e.obj)
			if abort {
				break
			}
		}
	} else {
		branches, branchDists := sortPreallocEntries(p, n.entries, b, bd)
		// only prune if buffer has k elements
		if l := len(dists); l >= k {
			branches = pruneEntriesMinDist(dists[l-1], branches, branchDists)
		}
		for _, e := range branches {
			nearest, dists, abort = tree.nearestNeighbors(k, p, e.child, dists, nearest, b[len(n.entries):], bd[len(n.entries):])
			if abort {
				break
			}
		}
	}
	return nearest, dists, abort
}

// 非原子性操作
func (tree *Rtree) UpsertBySpatialID(objects ...Spatial) {
	for _, object := range objects {
		tree.DeleteBySpatialID(object.ID)
		tree.Insert(object)
	}
}
