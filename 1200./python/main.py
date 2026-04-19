from typing import List
class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        arr.sort()
        minDiff = float('inf')
        result = []

        for i in range(len(arr)-1):
            diff = arr[i+1] - arr[i]
            if diff < minDiff:
                minDiff = diff
                result = [[arr[i], arr[i+1]]]
            elif diff==minDiff:
                result.append([arr[i], arr[i+1]])
            
        return result 
