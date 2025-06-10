package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	for current != nil && pos > 0 {
		current = current.Next
		pos--
	}
	return current
}
