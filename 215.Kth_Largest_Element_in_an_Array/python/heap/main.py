from typing import List, Optional
import heapq

class MaxHeap:
    def __init__(self, nums: Optional[List[int]] = None):
      if nums:
        self._heap = [-num for num in nums]
        heapq.heapify(self._heap)
      else:
        self._heap = []
      
    def push(self, value: int):
      heapq.heappush(self._heap, -value)
    
    def pop(self) -> int:
      value = heapq.heappop(self._heap)
      return -value
    
    def peek(self) -> int:
      return -self._heap[0]
       
    
    def __len__(self) -> int:
      return len(self._heap)
    


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
      maxHeap = MaxHeap(nums)
      result: Optional[int] = None
      for _ in range(k):
        result = maxHeap.pop()
      
      return result