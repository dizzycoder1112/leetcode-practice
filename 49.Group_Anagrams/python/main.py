class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = {}
        for string in strs:
            sorted_s = ''.join(sorted(string))
            print(sorted_s)
        
        return []