import heapq
from collections import Counter
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:

        return self.heapSort(nums, k)

    def heapSort(self, nums: List[int], k: int) -> List[int]:
        heap = []
        freqCount = Counter(nums)

        for num, count in freqCount.items():
            heapq.heappush(heap, (count, num))
        
        return [num for count, num in heapq.nlargest(k, heap)]