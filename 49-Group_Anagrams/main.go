package main

import "sort"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	groupAnagrams(strs)

}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	groups := make(map[string][]string)
	for _, str := range strs {
		key := sortString(str)
		groups[key] = append(groups[key], str)
	}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	// sort.Slice(result, func(i, j int) bool {
	//     return len(result[i]) < len(result[j])
	// })
	return result
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] > runes[j]
	})
	return string(runes)
}
