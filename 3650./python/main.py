class Solution:
    def minCost(self, n: int, edges: List[List[int]]) -> int:
        graph = [[] for _ in range(n)]
        for u, v, w in edges:
            graph[u].append((v, w))
            graph[v].append((u, 2*w))
        
        dist = [inf] * n
        dist[0] = 0
        heap = [(0,0)]

        while heap:
            cost, node = heapq.heappop(heap)

            if cost > dist[node]:
                continue
            
            if node == n-1:
                return cost
            
            for neighbor, weight in graph[node]:
                newCost = cost+weight
                if newCost < dist[neighbor]:
                    dist[neighbor] = newCost
                    heapq.heappush(heap, (newCost, neighbor))
        return -1