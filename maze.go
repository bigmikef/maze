package maze

// Room a maze node or room with 4 doors and a possible exit.
type Room struct {
	Exit  bool
	North *Room
	East  *Room
	South *Room
	West  *Room
}

// FindExit will traverse MazeNode looking for an exit.
func FindExit(node *Room) bool {
	visited := make(map[*Room]bool)
	return findExitWithVisit(node, visited)
}

func findExitWithVisit(node *Room, visited map[*Room]bool) bool {
	visited[node] = true
	if node.Exit {
		return true
	}
	var found = false
	if node.North != nil && !visited[node.North] {
		found = findExitWithVisit(node.North, visited)
	}
	if node.South != nil && !visited[node.South] {
		if !found {
			found = findExitWithVisit(node.South, visited)
		}
	}
	if node.East != nil && !visited[node.East] {
		if !found {
			found = findExitWithVisit(node.East, visited)
		}
	}
	if node.West != nil && !visited[node.West] {
		if !found {
			found = findExitWithVisit(node.West, visited)
		}
	}
	return found
}
