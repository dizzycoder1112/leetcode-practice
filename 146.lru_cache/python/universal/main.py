from typing import Dict

class Node:
    def __init__(self, key: int=0, value: int=0):
        self.key: int = key
        self.value: int = value
        self.prev: Node = None
        self.next: Node = None

class LRUCache:
    def __init__(self, capacity: int):
        self.capacity:int = capacity
        self.cache: Dict[int, Node] = {}

        self.head = Node()
        self.tail = Node()

        self.head.next = self.tail
        self.tail.prev = self.head
    
    def _remove(self, node: Node) -> None:
        prev_node = node.prev
        next_node = node.next
        prev_node.next = next_node
        next_node.prev = prev_node

    def _add_to_end(self, node: Node) -> None:
        current_latest_node = self.tail.prev
        current_latest_node.next = node
        node.next = self.tail
        node.prev = current_latest_node
        self.tail.prev = node

        

    def get(self, key: int) -> int:
        if key not in self.cache.keys():
            return -1
        
        node = self.cache[key]
        self._remove(node)
        self._add_to_end(node)

        return node.value
        

    def put(self, key: int, value: int) -> None:
        if key in self.cache.keys():
            node = self.cache.pop(key)
            self._remove(node)
        elif len(self.cache.keys()) >= self.capacity:
            oldest_node = self.head.next
            del self.cache[oldest_node.key]
            self._remove(oldest_node)

        new_node = Node(key, value)
        self._add_to_end(new_node)
        self.cache[key] = new_node
        

