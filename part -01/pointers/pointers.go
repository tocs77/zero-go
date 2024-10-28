package main

import (
	"fmt"
	"slices"
)

//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

type Item struct {
	name string
	tag  bool
}

func changeTag(items []*Item, name string, tag bool) {
	index := slices.IndexFunc(items, func(item *Item) bool {
		return item.name == name
	})
	if index == -1 {
		panic("item not found")
	}
	items[index].tag = tag
}

func checkout(items []*Item) {
	for _, item := range items {
		item.tag = false
	}
}

func print(items []*Item) {
	for _, item := range items {
		fmt.Printf("%v: %v\n", item.name, item.tag)
	}

	fmt.Println("---")
}

func main() {
	items := []*Item{
		{"Shirt", true},
		{"Pants", true},
		{"Purse", true},
		{"Watch", true},
	}

	print(items)

	changeTag(items, "Shirt", false)
	print(items)

	checkout(items)
	print(items)

}
