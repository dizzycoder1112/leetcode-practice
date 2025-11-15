/*
** NOTE: The key point of implementation **

1. In a proper LRU Cache, we want O(1) operations.
get -> O(1)
put	-> O(1)

2. We can use a Map to store the key-value pairs.
3. We can use a doubly linked list to store the keys in order of their usage.
4. get(key): if key exists, return the value and move the key to the end of the list. Otherwise, return -1.
5. put(key, value): if key exists, update the value and move the key to the end of the list. Otherwise, add the key-value pair to the end of the list. If the list is full, remove the first key from the list.

 */

class LRUCache<K, V> {
  private cache: Map<K, V>;
  private capacity: number;

  constructor(capacity: number = 3) {
    this.cache = new Map();
    this.capacity = capacity;
  }

  get(key: K): V | undefined {
    if (!this.cache.has(key)) return undefined;

    const value = this.cache.get(key)!;
    // Move to end to show it's recently used
    this.cache.delete(key);
    this.cache.set(key, value);

    console.log(`read key ${String(key)}:`, value);
    this.readState();
    return value;
  }

  put(key: K, value: V): void {
    if (this.cache.has(key)) {
      this.cache.delete(key); // Refresh position
    }

    this.cache.set(key, value);

    if (this.cache.size > this.capacity) {
      // Remove least recently used (first inserted)
      const oldestKey = this.cache.keys().next().value;
      this.cache.delete(oldestKey);
    }

    console.log('insert or update item:', key, value);
    this.readState();
  }

  private readState(): void {
    console.log("cache state:", Array.from(this.cache.entries()));
  }

  checkCapacity(): number {
    console.log("max capacity:", this.capacity);
    return this.capacity;
  }
}


async function main() {
  const cache = new LRUCache<string, number>(3);

  // Put some values in the cache
  cache.put("a", 1);
  cache.put("b", 2);
  cache.put("c", 3);
  
  // Access a value to mark it as recently used
  cache.get("a"); // moves "a" to the most recently used position

  console.log(cache)
  
  // Add a new value, triggering eviction of the least recently used ("b")
  cache.put("d", 4);
  
  // Try to get an evicted value



}

main();