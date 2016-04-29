// Def: ii is an inverted-index composed of term - doclist

package main

type docList []int

// Find documents matching query {ink}
func find(ii map[string]docList, term string) []int {
	return ii[term]
}

// Find documents matching query {ink, pink}
func findMultiple(ii map[string]docList, terms []string) []int {
	return nil
}
