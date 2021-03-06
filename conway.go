// Copyright 2017 Aleksandr Demeshko. All rights reserved.

// conway project conway.go
package conway

import (
	"fmt"
	"io"
	"os"
)

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

func (c Cell) ToString() string {
	return fmt.Sprintf("( %v %v )\n", c.x, c.y)
}

func ReadCell(r io.Reader) (*Cell, error) {
	var x, y int64
	n, err := fmt.Fscanf(r, "( %d %d )\n", &x, &y)
	if err != nil {
		return nil, err
	}
	if n != 2 {
		return nil, fmt.Errorf("Wrong string format")
	}
	c := Cell{x, y}
	return &c, nil
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

func (p *Population) SaveToFile(fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	for cell := range p.cells {
		s := cell.ToString()
		f.WriteString(s)
	}
	return nil
}

func NewPopulation(fname string) (*Population, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p := Population{}
	for {
		c, err := ReadCell(f)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		p.cells[*c] = 0
	}
	return &p, nil
}
