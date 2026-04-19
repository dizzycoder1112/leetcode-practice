from collections import Counter
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:

        return self.bucketSort(nums, k)

    def bucketSort(self, nums: List[int], k: int) -> List[int]:
        freqCount = Counter(nums)

        buckets = [[] for _ in range(len(nums)+1)]

        for num, count in freqCount.items():
            buckets[count].append(num)

        result = []

        for nums in buckets[::-1]:
            for num in nums:
                if k>len(result):
                    result.append(num)

        return result