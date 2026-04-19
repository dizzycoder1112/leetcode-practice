class Solution:
    def minimumCost(self, source: str, target: str, original: List[str], changed: List[str], cost: List[int]) -> int:
        # 建立 26x26 距离矩阵
        INF = float('inf')
        dist = [[INF] * 26 for _ in range(26)]

        # 对角线为 0
        for i in range(26):
            dist[i][i] = 0

        # 填入边（取最小值，因为可能有重复边）
        for o, c, w in zip(original, changed, cost):
            i, j = ord(o) - ord('a'), ord(c) - ord('a')
            dist[i][j] = min(dist[i][j], w)

        # Floyd-Warshall
        for k in range(26):
            for i in range(26):
                for j in range(26):
                    dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])

        # 计算总成本
        total = 0
        for s, t in zip(source, target):
            if s == t:
                continue
            i, j = ord(s) - ord('a'), ord(t) - ord('a')
            if dist[i][j] == INF:
                return -1
            total += dist[i][j]

        return total