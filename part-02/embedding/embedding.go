//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

type Dashboard struct {
	bandwidth BandwidthUsage
	temp      CpuTemp
	memory    MemoryUsage
}

func (b BandwidthUsage) Average() float32 {
	var sum float32
	for _, v := range b.amount {
		sum += float32(v)
	}
	return sum / float32(len(b.amount))
}

func (c CpuTemp) Average() float32 {
	var sum float32
	for _, v := range c.temp {
		sum += float32(v)
	}
	return sum / float32(len(c.temp))
}

func (m MemoryUsage) Average() float32 {
	var sum float32
	for _, v := range m.amount {
		sum += float32(v)
	}
	return sum / float32(len(m.amount))
}

func (d Dashboard) AverageBandwidth() float32 {
	return d.bandwidth.Average()
}

func (d Dashboard) AverageCpuTemp() float32 {
	return d.temp.Average()
}

func (d Dashboard) AverageMemoryUsage() float32 {
	return d.memory.Average()
}

func (d Dashboard) String() string {
	return fmt.Sprintf("Bandwidth: %f, CpuTemp: %f, MemoryUsage: %f", d.AverageBandwidth(), d.AverageCpuTemp(), d.AverageMemoryUsage())
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{bandwidth, temp, memory}

	fmt.Println(dash)

}
