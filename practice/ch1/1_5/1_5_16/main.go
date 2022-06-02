package main

import (
	"algo/cases/ch1/1_5/dynamic_connectivity/unionfind"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type unionData struct {
	p, q int
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <readfile>", os.Args[0])
	}
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("open file: %v", err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	unionDatas := make([]unionData, 0)
	for s.Scan() {
		t := s.Text()
		pq := strings.Split(t, " ")
		if len(pq) != 2 {
			log.Fatalf("read incorrect line: %s", t)
		}
		p, err := strconv.Atoi(pq[0])
		if err != nil {
			log.Fatalf("invalid number: %s", pq[0])
		}
		q, err := strconv.Atoi(pq[1])
		if err != nil {
			log.Fatalf("invalid number: %s", pq[1])
		}
		unionDatas = append(unionDatas, unionData{p, q})
	}

	qf := unionfind.QuickFindUFWithVisit(len(unionDatas))
	ufDraw(qf, unionDatas, "quick_find")

	qu := unionfind.QuickUnionUFWithVisit(len(unionDatas))
	ufDraw(qu, unionDatas, "quick_union")

	wqu := unionfind.WeightQuickUnionUFWithVisit(len(unionDatas))
	ufDraw(wqu, unionDatas, "weight_quick_union")
}

func ufDraw(uf unionfind.UnionFindWithVisit, unionDatas []unionData, name string) {
	total := 0
	costs := make(plotter.XYs, 0)
	averages := make(plotter.XYs, 0)
	for i, v := range unionDatas {
		cost := uf.UnionWithVisit(v.p, v.q)
		total += cost
		averages = append(averages, plotter.XY{
			X: float64(i),
			Y: float64(total) / float64(i+1),
		})
		costs = append(costs, plotter.XY{
			X: float64(i),
			Y: float64(cost),
		})
	}
	draw := plot.New()
	draw.Title.Text = name
	draw.X.Label.Text = "union times"
	draw.Y.Label.Text = "inner data visit times"
	err := plotutil.AddScatters(draw,
		"costs", costs,
		"averages", averages,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = draw.Save(10*vg.Inch, 10*vg.Inch,
		"./practice/ch1/1_5/1_5_16/"+name+".png")
	if err != nil {
		log.Fatal(err)
	}
}
