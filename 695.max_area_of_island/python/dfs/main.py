class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        maxArea = 0
        self.directions = [(1,0), (-1,0), (0,1), (0, -1)]

        for y in range(rows):
            for x in range(cols):
                maxArea = max(maxArea, self.dfs(grid, x, y))
        
        return maxArea



    def dfs(self, grid: List[List[int]], x: int, y: int) -> int:
        rows, cols = len(grid), len(grid[0])
        count = 0
        if not 0<=x<cols or not 0<=y<rows or grid[y][x] == 0:
            return count

        grid[y][x] = 0
        count += 1
        for dx, dy in self.directions:
            count += self.dfs(grid, x+dx, y+dy)
        
        return count
        
