package main

import (
	"testing"

	"github.com/jangler/oracles-randomizer/graph"
	"github.com/jangler/oracles-randomizer/logic"
	"github.com/jangler/oracles-randomizer/rom"
)

func TestGraph(t *testing.T) {
	testSeasonsGraph(t)
	// testAgesGraph(t)
}

// check that graph logic is working as expected
func testSeasonsGraph(t *testing.T) {
	rom.Init(rom.GameSeasons)
	r := NewRoute(rom.GameSeasons)
	g := r.Graph

	// test basic start item
	checkReach(t, g, map[string]string{
		"d0 key chest": "feather",
	}, "maku tree", false, false)
	checkReach(t, g, map[string]string{
		"d0 key chest": "sword",
	}, "maku tree", false, true)

	// test hard logic via bombs as weapon
	checkReach(t, g, map[string]string{
		"d0 key chest":           "moosh's flute",
		"d0 rupee chest":         "bombs, 10",
		"horon village SE chest": "gnarled key",
	}, "d1 stalfos drop", false, false)
	checkReach(t, g, map[string]string{
		"d0 key chest":           "moosh's flute",
		"d0 rupee chest":         "bombs, 10",
		"horon village SE chest": "gnarled key",
	}, "d1 stalfos drop", true, true)

	// test key counting
	checkReach(t, g, map[string]string{
		"d0 key chest":     "sword",
		"maku tree":        "gnarled key",
		"d1 stalfos drop":  "d1 small key",
		"d1 stalfos chest": "bombs, 10",
	}, "d1 basement", false, false)
	checkReach(t, g, map[string]string{
		"d0 key chest":     "sword",
		"maku tree":        "gnarled key",
		"d1 stalfos drop":  "d1 small key",
		"d1 stalfos chest": "bombs, 10",
		"d1 railway chest": "d1 small key",
	}, "d1 basement", false, true)
}

// check that graph logic is working as expected
func testAgesGraph(t *testing.T) {
	rom.Init(rom.GameAges)
	r := NewRoute(rom.GameAges)
	g := r.Graph

	// test basic start item
	checkReach(t, g, map[string]string{
		"starting chest": "feather",
	}, "black tower worker", false, false)
	checkReach(t, g, map[string]string{
		"starting chest": "sword",
	}, "black tower worker", false, true)

	// test hard logic via bombs as weapon
	checkReach(t, g, map[string]string{
		"starting chest":     "bombs, 10",
		"nayru's house":      "bracelet",
		"black tower worker": "shovel",
	}, "d2 bombed terrace", false, false)
	checkReach(t, g, map[string]string{
		"starting chest":     "bombs, 10",
		"nayru's house":      "bracelet",
		"black tower worker": "shovel",
	}, "d2 bombed terrace", true, true)

	// test key counting
	checkReach(t, g, map[string]string{
		"starting chest":      "sword 1",
		"nayru's house":       "bombs, 10",
		"black tower worker":  "dimitri's flute",
		"d3 pols voice chest": "d3 small key",
	}, "d3 bush beetle room", false, false)
	checkReach(t, g, map[string]string{
		"starting chest":      "sword 1",
		"nayru's house":       "bombs, 10",
		"black tower worker":  "dimitri's flute",
		"d3 pols voice chest": "d3 small key",
		"d3 statue drop":      "d3 small key",
		"d3 armos drop":       "d3 small key",
	}, "d3 bush beetle room", false, true)

	// test bombs from head thwomp in hard logic
	headThwompBombMap := map[string]string{
		"starting chest":        "bracelet",
		"nayru's house":         "harp 1",
		"black tower worker":    "harp 2",
		"lynna city chest":      "switch hook 1",
		"fairies' woods chest":  "iron shield",
		"symmetry city brother": "sword 1",
		"d2 moblin drop":        "feather",
		"d2 basement drop":      "d2 small key",
		"d2 thwomp tunnel":      "d2 small key",
		"d2 thwomp shelf":       "d2 small key",
		"d2 moblin platform":    "d2 small key",
		"d2 rope room":          "d2 small key",
		"d2 statue puzzle":      "d2 boss key",
	}
	checkReach(t, g, headThwompBombMap, "d2 bombed terrace", false, false)
	checkReach(t, g, headThwompBombMap, "d2 bombed terrace", true, true)
}

func BenchmarkGraphExplore(b *testing.B) {
	// init graph
	r := NewRoute(rom.GameSeasons)
	b.ResetTimer()

	// explore all items from the d0 sword chest
	for name := range logic.SeasonsExtraItems() {
		r.Graph.Explore(make(map[*graph.Node]bool), false, r.Graph[name])
	}
}

// helper function for testing whether a node is reachable given a certain
// slotting
func checkReach(t *testing.T, g graph.Graph, links map[string]string,
	target string, hard, expect bool) {
	t.Helper()

	// add parents at the start of the function, and remove them at the end. if
	// a parent is blank, remove it at the start and add it at the end (only
	// useful for default seasons).
	for parent, child := range links {
		if parent == "" {
			g[child].ClearParents()
		} else {
			g[child].AddParents(g[parent])
		}
	}
	defer func() {
		for parent, child := range links {
			if parent == "" {
				g[child].AddParents(g["start"])
			} else {
				g[child].ClearParents()
			}
		}
	}()
	g.ClearMarks()

	if (g[target].GetMark(g[target], hard) == graph.MarkTrue) != expect {
		if expect {
			t.Errorf("expected to reach %s, but could not", target)
		} else {
			t.Errorf("expected not to reach %s, but could", target)
		}
	}
}
