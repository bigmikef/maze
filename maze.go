package maze

type MazeNode struct {
	exit                     bool
	north, east, west, south *MazeNode
}

func FindExit(node *MazeNode) bool {
	visited := make(map[*MazeNode]bool)
	return FindExitWithVisit(node, visited)
}

func FindExitWithVisit(node *MazeNode, visited map[*MazeNode]bool) bool {
	visited[node] = true
	if node.exit {
		return true
	}
	var found bool = false
	if !visited[node.north] {
		found = FindExitWithVisit(node.north, visited)
	}
	if !visited[node.south] {
		if !found {
			found = FindExitWithVisit(node.south, visited)
		}
	}
	if !visited[node.east] {
		if !found {
			found = FindExitWithVisit(node.east, visited)
		}
	}
	if !visited[node.west] {
		if !found {
			found = FindExitWithVisit(node.west, visited)
		}
	}
	return found
}
