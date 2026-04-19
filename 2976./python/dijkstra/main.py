from heapq import heappush, heappop

class Solution:
    def minimumCost(self, source: str, target: str, original: List[str], changed: List[str], cost: List[int]) -> int:
        # 建立邻接表
        graph = [[] for _ in range(26)]
        for o, c, w in zip(original, changed, cost):
            i, j = ord(o) - ord('a'), ord(c) - ord('a')
            graph[i].append((j, w))

        # Dijkstra: 从起点 src 到所有点的最短距离
        def dijkstra(src: int) -> list[int]:
            dist = [float('inf')] * 26
            dist[src] = 0
            pq = [(0, src)]  # (距离, 节点)

            while pq:
                d, u = heappop(pq)
                if d > dist[u]:
                    continue
                for v, w in graph[u]:
                    if dist[u] + w < dist[v]:
                        dist[v] = dist[u] + w
                        heappush(pq, (dist[v], v))

            return dist

        # 缓存：避免重复计算同一个起点
        cache = {}

        # 计算总成本
        total = 0
        for s, t in zip(source, target):
            if s == t:
                continue
            i, j = ord(s) - ord('a'), ord(t) - ord('a')

            if i not in cache:
                cache[i] = dijkstra(i)

            if cache[i][j] == float('inf'):
                return -1
            total += cache[i][j]

        return total