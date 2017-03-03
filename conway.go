// Copyright 2017 Aleksandr Demeshko. All rights reserved.

// conway project conway.go
package conway

type Cell struct {
	x, y int64
}

type Population struct {
	cells     map[Cell]int
	popNumber int64
}

func (c Cell) Neighbors() []Cell {
	return []Cell{
		Cell{c.x - 1, c.y - 1}, Cell{c.x - 1, c.y}, Cell{c.x - 1, c.y + 1},
		Cell{c.x, c.y - 1}, Cell{c.x, c.y + 1},
		Cell{c.x + 1, c.y - 1}, Cell{c.x + 1, c.y}, Cell{c.x + 1, c.y + 1}}
}

func (p *Population) Next() {
	var processed = make(map[Cell]int)
	var np = make(map[Cell]int)

	for cell, num := range p.cells {
		cnt := 0
		// let's process all neighbors of the current cell
		// also count all populated neighbor cells into cnt
		for _, neighbor := range cell.Neighbors() {
			if _, ok := p.cells[neighbor]; ok {
				// this neigbor cell is populated - will not process it here
				cnt++
				continue
			}
			if _, ok := processed[neighbor]; ok {
				// this neigbor cell is already processed
				continue
			}
			processed[neighbor] = 1
			ncnt := 0
			// count all populated neighbors of the neighbor into ncnt
			for _, nn := range neighbor.Neighbors() {
				if _, ok := p.cells[nn]; ok {
					ncnt++
				}
			}
			if ncnt == 3 {
				np[neighbor] = 0
			}
		}
		if cnt == 2 || cnt == 3 {
			np[cell] = num + 1
		}
	}

	p.cells = np
	p.popNumber++
}
