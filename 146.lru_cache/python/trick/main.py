from typing import Dict

class LRUCache:
    def __init__(self, capacity: int):
        self.capacity:int = capacity
        self.cache: Dict[int, int] = {}
        

    def get(self, key: int) -> int:
        if key not in self.cache.keys():
            return -1
        result = self.cache.pop(key)  # Remove and get value in one step                             
        self.cache[key] = result
        return result
        

    def put(self, key: int, value: int) -> None:
        if key in self.cache.keys():
            self.cache.pop(key)
        elif len(self.cache.keys()) >= self.capacity:
            oldest_key = next(iter(self.cache))
            del self.cache[oldest_key]

        self.cache[key] = value
        
