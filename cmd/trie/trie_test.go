package main

import (
	"strings"
	"testing"
)

func TestGetIndex(t *testing.T) {
	// `a` == 97
	cases := []struct {
		Name         string
		Str          string
		ExpectIndex0 int
	}{
		{
			Name:         "1",
			Str:          "a",
			ExpectIndex0: 0,
		},
		{
			Name:         "2",
			Str:          "A",
			ExpectIndex0: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			str := Lower(c.Str)
			index := GetIndex(rune(str[0]))
			if index != c.ExpectIndex0 {
				t.Errorf("Expect %d while get %d", c.ExpectIndex0, index)
			}
		})
	}
}

func BenchmarkExpendTrieTree(b *testing.B) {
	tree := NewNode(true)
	for i := 0; i < b.N; i++ {
		words := "BenchmarkExpendTrieTree"
		ExpendTrieTree(tree, words)
	}
}

func BenchmarkCreateTree(b *testing.B) {

	str := `
he Old Man and the Sea recounts an epic battle between an old, experienced fisherman and a giant marlin said to be the largest catch of his life. It opens by explaining that the fisherman, who is named Santiago, has gone 84 days without catching any fish at all (although a comment made at some point in the book reveals that he had previously gone 87 days without catching one). He is apparently so unlucky that his young apprentice, Manolin, has been forbidden by his parents to sail with the old man and been ordered to fish with more successful fishermen. Still dedicated to the old man, however, the boy visits Santiago's shack each night, hauling back his fishing gear, feeding him and discussing American baseball — most notably Santiago's idol, Joe DiMaggio. Santiago tells Manolin that on the next day, he will venture far out into the Gulf to fish, confident that his unlucky streak is near its end.

Thus on the eighty-fifth day, Santiago sets out alone, taking his skiff far into the Gulf. He sets his lines and, by noon of the first day, a big fish that he is sure is a marlin takes his bait. Unable to pull in the great marlin, Santiago instead finds the fish pulling his skiff. Two days and two nights pass in this manner, during which the old man bears the tension of the line with his body. Though he is wounded by the struggle and in pain, Santiago expresses a compassionate appreciation for his adversary, often referring to him as a brother. He also determines that because of the fish's great dignity, no one will be worthy of eating the marlin.

On the third day of the ordeal, the fish begins to circle the skiff, indicating his tiredness to the old man. Santiago, now completely worn out and almost in delirium, uses all the strength he has left in him to pull the fish onto its side and stab the marlin with a harpoon, thereby ending the long battle between the old man and the tenacious fish.

Santiago straps the marlin to his skiff and heads home, thinking about the high price the fish will bring him at the market and how many people he will feed.

While Santiago continues his journey back to the shore, sharks are attracted to the trail of blood left by the marlin in the water. The first, a great mako shark, Santiago kills with his harpoon, losing that weapon in the process. He makes a new harpoon by strapping his knife to the end of an oar to help ward off the next line of sharks; in total, five sharks are slain and many others are driven away. But by night, the sharks have almost devoured the marlin's entire carcass, leaving a skeleton consisting mostly of its backbone, its tail and its head, the latter still bearing the giant spear. The old man castigates himself for sacrificing the marlin. Finally reaching the shore before dawn on the next day, he struggles on the way to his shack, carrying the heavy mast on his shoulder. Once home, he slumps onto his bed and enters a very deep sleep.

A group of fishermen gather the next day around the boat where the fish's skeleton is still attached. One of the fishermen measures it to be eighteen feet from nose to tail. Tourists at the nearby café mistakenly take it for a shark. Manolin, worried during the old man's endeavor, cries upon finding him safe asleep. The boy brings him newspapers and coffee. When the old man wakes, they promise to fish together once again. Upon his return to sleep, Santiago dreams of lions on the African beach.
`
	var w strings.Builder

	for i := 0; i < b.N; i++ {
		w.WriteString(str)
		r := strings.NewReader(w.String())
		CreateTree(r)
	}
}
