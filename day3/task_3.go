package main

import (
	"bufio"
	"fmt"
	"os"
)

type MapCell int

const (
	Tree MapCell = iota
	Open
)

var cellTranslate = map[rune]MapCell{
	'.': Open,
	'#': Tree,
}

type Map struct {
	Content            [][]MapCell
	CurrentX, CurrentY int
	MeetedTrees        int
}

func (m *Map) Step(right int, down int) {
	m.CurrentY += down
	m.CurrentX += right
	m.CurrentX %= len(m.Content[0])
	if m.Content[m.CurrentY][m.CurrentX] == Tree {
		m.MeetedTrees++
	}
}

func (m Map) IsOver() bool {
	return m.CurrentY == len(m.Content)-1
}

func (m Map) tryPath(right int, down int) int {
	for !m.IsOver() {
		m.Step(right, down)
	}
	return m.MeetedTrees
}

func task1(forestMap Map) {
	fmt.Printf("Task 1 result: %s\n", forestMap.tryPath(3, 1))
}

func task2(m Map) {
	result := m.tryPath(1, 1) * m.tryPath(3, 1) * m.tryPath(5, 1) * m.tryPath(7, 1) * m.tryPath(1, 2)
	fmt.Printf("Task 1 result: %s\n", result)
}

func main() {
	var forestMap Map
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var mapLine []MapCell
		for _, c := range scanner.Text() {
			mapLine = append(mapLine, cellTranslate[c])
		}
		forestMap.Content = append(forestMap.Content, mapLine)
	}
	task1(forestMap)
	task2(forestMap)
}
