from typing import List


class UnionFind:
    def __init__(self, rows: int, cols: int):
        self.rows = rows
        self.cols = cols
        n = rows * cols
        self.parent = list(range(n))  # 每個節點的老大
        self.size = [1] * n           # 每個集合的大小

    def to_id(self, x: int, y: int) -> int:
        """二維座標轉一維編號"""
        return y * self.cols + x

    def find(self, x: int) -> int:
        """找老大（帶路徑壓縮）"""
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x: int, y: int):
        """合併兩個集合"""
        root_x = self.find(x)
        root_y = self.find(y)

        if root_x == root_y:
            return  # 已經同一團

        # 小的併入大的
        if self.size[root_x] < self.size[root_y]:
            self.parent[root_x] = root_y
            self.size[root_y] += self.size[root_x]
        else:
            self.parent[root_y] = root_x
            self.size[root_x] += self.size[root_y]

    def get_size(self, x: int) -> int:
        """取得 x 所屬集合的大小"""
        return self.size[self.find(x)]


class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        uf = UnionFind(rows, cols)

        # 遍歷每個格子
        for y in range(rows):
            for x in range(cols):
                if grid[y][x] == 0:
                    continue

                current_id = uf.to_id(x, y)

                #動態落子需要檢查四邊[(1,0), (-1,0), (0,1), (0,-1)]
                #但這個case可以優化成只檢查右邊和下邊，避免不必要的遍歷
                for dx, dy in [(1, 0), (0, 1)]:
                    nx, ny = x + dx, y + dy

                    # 檢查邊界 & 是否是陸地
                    if 0 <= nx < cols and 0 <= ny < rows and grid[ny][nx] == 1:
                        uf.union(current_id, uf.to_id(nx, ny))

        # 找最大的島嶼
        max_area = 0
        for y in range(rows):
            for x in range(cols):
                if grid[y][x] == 1:
                    max_area = max(max_area, uf.get_size(uf.to_id(x, y)))

        return max_area


if __name__ == "__main__":
    sol = Solution()

    # Test 1
    grid1 = [
        [0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
        [0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0],
        [0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
        [0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0],
    ]
    print(f"Test 1: {sol.maxAreaOfIsland(grid1)}")  # Expected: 6

    # Test 2
    grid2 = [[0, 0, 0, 0, 0, 0, 0, 0]]
    print(f"Test 2: {sol.maxAreaOfIsland(grid2)}")  # Expected: 0

    # Test 3: 簡單例子
    grid3 = [
        [1, 1, 0],
        [1, 1, 0],
        [0, 0, 1],
    ]
    print(f"Test 3: {sol.maxAreaOfIsland(grid3)}")  # Expected: 4
