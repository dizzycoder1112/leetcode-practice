from collections import deque
class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        queue = deque()
        fresh = 0

        for y in range(rows):
            for x in range(cols):
                if grid[y][x] == 2:
                    queue.append((x, y))
                elif grid[y][x] == 1:
                    fresh += 1
        
        minutes = 0
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        while queue and fresh > 0:
            levelSize = len(queue)

            while levelSize > 0:
                current = queue.popleft()
                for dx, dy in directions:
                    x0, y0 = current
                    x1 = x0+dx
                    y1 = y0+dy
                    if 0<=x1<cols and 0<=y1<rows and grid[y1][x1] == 1:
                        grid[y1][x1] = 2
                        fresh -= 1
                        queue.append((x1, y1))
                levelSize -= 1



            minutes += 1
        
        return minutes if fresh == 0 else -1

        
        


        


