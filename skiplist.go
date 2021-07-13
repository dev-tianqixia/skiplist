package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// reference:
// https://www.ics.uci.edu/~pattis/ICS-23/lectures/notes/Skip%20Lists.pdf
// https://azrael.digipen.edu/~mmead/www/Courses/CS280/SkipLists.html
// https://opendatastructures.org/newhtml/ods/latex/skiplists.html

const maxLevel = 16

type node struct {
	// temporarily use string as the key type
	key string
	val interface{}
	// forward[i] points to node's successor in list Li
	// the max number of lists is 16 for now, TODO: make maxLevel configurable
	forward [maxLevel]*node
}

type SkipList struct {
	// level <= maxLevel
	level    int
	size     int
	sentinel *node
}

func NewSkipList() *SkipList {
	return &SkipList{
		level:    0,
		sentinel: &node{},
	}
}

// TODO
func NewSkipListFrom() *SkipList {
	return nil
}

// implements search
func (l *SkipList) Get(key string) (interface{}, bool) {
	n := l.sentinel

	for i := l.level; i >= 0; i-- {
		// if the key of next element in current level(n.forward[i]) is less than the given key, iterate horizontally
		for n.forward[i] != nil && n.forward[i].key < key {
			n = n.forward[i]
		}
	}

	if n = n.forward[0]; n != nil && n.key == key {
		return n.val, true
	}
	return nil, false
}

func (l *SkipList) Insert(key string, val interface{}) {
	update := make([]*node, maxLevel)
	for i := range update {
		update[i] = l.sentinel
	}

	n := l.sentinel
	for i := l.level; i >= 0; i-- {
		for n.forward[i] != nil && n.forward[i].key < key {
			n = n.forward[i]
		}
		update[i] = n
	}

	if n = n.forward[0]; n != nil && n.key == key {
		// update
		n.val = val
	} else {
		node := &node{
			key: key,
			val: val,
		}
		update[0].forward[0], node.forward[0] = node, update[0].forward[0]
		l.size++
		// try to increase the number of lists randomly, once a time
		// TODO: cache log values
		limit := int(math.Log2(float64(l.size)))
		if limit > maxLevel {
			limit = maxLevel
		}
		for i := 1; i < limit; i++ {
			if rand.Float64() < 0.5 {
				update[i].forward[i], node.forward[i] = node, update[i].forward[i]
				if i > l.level {
					l.level = i
				}
			} else {
				break
			}
		}
	}
}

func (l *SkipList) Delete(key string) {
	update := make([]*node, maxLevel)

	n := l.sentinel
	for i := l.level; i >= 0; i-- {
		for n.forward[i] != nil && n.forward[i].key < key {
			n = n.forward[i]
		}
		update[i] = n
	}

	if n = n.forward[0]; n != nil && n.key == key {
		// n is the target to be deleted
		for i := 0; i <= l.level; i++ {
			if update[i].forward[i] != n {
				break
			}
			// assert: update[i].forward[i] == n
			update[i].forward[i] = n.forward[i]
		}
		if l.sentinel.forward[l.level] == nil && l.level > 0 {
			l.level--
		}
	}
}

func (l *SkipList) String() string {
	var builder strings.Builder

	for i := l.level; i >= 0; i-- {
		builder.WriteString(fmt.Sprintf("level %d: ", i))
		n := l.sentinel.forward[i]
		for n != nil {
			builder.WriteString(fmt.Sprintf("%s -> ", n.key))
			n = n.forward[i]
		}
		builder.WriteString("nil\n")
	}

	return builder.String()
}

// Size returns the number of elements in the list
func (l *SkipList) Size() int {
	return l.size
}

// Level returns the number of lists
func (l *SkipList) Level() int {
	return l.level + 1
}
