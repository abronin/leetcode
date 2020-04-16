package main

import "fmt"

func main() {
	obj := Constructor()
	obj.CheckIn(1, "A", 1)
	obj.CheckOut(1, "B", 1)
	res := obj.GetAverageTime("A", "B")
	fmt.Println(res)
}

type CI struct {
	Station string
	Time    int
}

type UndergroundSystem struct {
	Travels  map[string]map[string][]int
	Checkins map[int]CI
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		Travels:  map[string]map[string][]int{},
		Checkins: map[int]CI{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.Checkins[id] = CI{Station: stationName, Time: t}
	if _, ok := this.Travels[stationName]; !ok {
		this.Travels[stationName] = map[string][]int{}
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if _, ok := this.Travels[this.Checkins[id].Station][stationName]; !ok {
		this.Travels[this.Checkins[id].Station][stationName] = []int{}
	}
	this.Travels[this.Checkins[id].Station][stationName] = append(this.Travels[this.Checkins[id].Station][stationName], t-this.Checkins[id].Time)
	delete(this.Checkins, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	return calcAverage(this.Travels[startStation][endStation])
}

func calcAverage(times []int) float64 {
	if len(times) == 0 {
		return 0
	}

	sum := 0
	for _, v := range times {
		sum += v
	}
	return float64(float64(sum) / float64(len(times)))
}
