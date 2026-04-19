class UnionFind:
    """並查集資料結構"""
    def __init__(self, n: int):
        self.parent = list(range(n))  # 每個節點的老大（初始是自己）
        self.size = [1] * n           # 每個集合的大小

    def find(self, x: int) -> int:
        """找到 x 的根節點（老大）"""
        if self.parent[x] != x:
            # 路徑壓縮：讓路徑上所有節點直接指向根
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x: int, y: int) -> bool:
        """合併 x 和 y 所屬的集合"""
        root_x = self.find(x)
        root_y = self.find(y)

        if root_x == root_y:
            return False  # 已經在同一集合

        # 按秩合併：小樹接到大樹下面
        if self.size[root_x] < self.size[root_y]:
            self.parent[root_x] = root_y
            self.size[root_y] += self.size[root_x]
        else:
            self.parent[root_y] = root_x
            self.size[root_x] += self.size[root_y]

        return True

    def get_size(self, x: int) -> int:
        """取得 x 所屬集合的大小"""
        return self.size[self.find(x)]


class Board:
    def __init__(self):
        self.rows = 9
        self.cols = 9
        self.grid = [['+'] * self.cols for _ in range(self.rows)]
        self.currentPlayer = 'R'

        # 每個顏色各自一個 UnionFind
        n = self.rows * self.cols
        self.uf = {
            'R': UnionFind(n),
            'G': UnionFind(n)
        }
        self.territory_count = {'R': 0, 'G': 0}  # 領地數量

    def _to_id(self, x: int, y: int) -> int:
        """2D 座標轉 1D 編號"""
        return y * self.cols + x

    def print_board(self):
        """印出棋盤"""
        print("  " + " ".join(str(i) for i in range(self.cols)))
        for y, row in enumerate(self.grid):
            print(f"{y} " + " ".join(row))
        print(f"當前玩家: {self.currentPlayer}")
        print(f"R 領地數: {self.territory_count['R']}, G 領地數: {self.territory_count['G']}")

    def place_stone(self, x: int, y: int) -> bool:
        """落子"""
        if not (0 <= x < self.cols and 0 <= y < self.rows):
            return False
        if self.grid[y][x] != '+':
            return False

        color = self.currentPlayer
        self.grid[y][x] = color

        # === Union Find 核心邏輯 ===
        current_id = self._to_id(x, y)
        uf = self.uf[color]

        # 1. 新增一塊領地
        self.territory_count[color] += 1

        # 2. 檢查四個方向，合併同色棋子
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]
        for dx, dy in directions:
            nx, ny = x + dx, y + dy
            if 0 <= nx < self.cols and 0 <= ny < self.rows:
                if self.grid[ny][nx] == color:
                    neighbor_id = self._to_id(nx, ny)
                    # Union 成功表示合併了兩塊不同的領地
                    if uf.union(current_id, neighbor_id):
                        self.territory_count[color] -= 1

        # 換人
        self.currentPlayer = 'G' if color == 'R' else 'R'
        return True

    def pass_turn(self):
        """跳過"""
        self.currentPlayer = 'G' if self.currentPlayer == 'R' else 'R'

    def area_size(self, x: int, y: int) -> int:
        """查詢 (x, y) 所屬領地的大小 - O(1)!"""
        if not (0 <= x < self.cols and 0 <= y < self.rows):
            return 0

        color = self.grid[y][x]
        if color == '+':
            return 0

        current_id = self._to_id(x, y)
        return self.uf[color].get_size(current_id)

    def debug_uf(self, color: str):
        """除錯：顯示 UnionFind 狀態"""
        print(f"\n=== {color} 的 UnionFind 狀態 ===")
        uf = self.uf[color]

        # 找出所有該顏色的棋子
        groups = {}
        for y in range(self.rows):
            for x in range(self.cols):
                if self.grid[y][x] == color:
                    id = self._to_id(x, y)
                    root = uf.find(id)
                    if root not in groups:
                        groups[root] = []
                    groups[root].append((x, y))

        for root, members in groups.items():
            print(f"領地 (根={root}): {members}, 大小={len(members)}")


if __name__ == "__main__":
    board = Board()

    print("=== 展示 Union Find 合併過程 ===\n")

    # R 連續下三子形成一條線，觀察合併過程
    moves = [
        (0, 0),  # R - 第一子
        (5, 5),  # G - 隨便下
        (1, 0),  # R - 與左邊合併！
        (5, 6),  # G
        (2, 0),  # R - 再與左邊合併！
        (5, 7),  # G
        (1, 1),  # R - 與上方合併！形成 L 型
    ]

    for i, (x, y) in enumerate(moves):
        color = board.currentPlayer
        print(f"步驟 {i+1}: {color} 落子 ({x}, {y})")
        board.place_stone(x, y)

        # 只顯示左上角 3x3
        print("  0 1 2")
        for row in range(3):
            print(f"{row} " + " ".join(board.grid[row][:3]))

        print(f"R 領地數: {board.territory_count['R']}")

        if color == 'R':
            size = board.area_size(x, y)
            print(f"→ ({x},{y}) 所屬領地大小: {size}")

        print()

    # 最終 UnionFind 狀態
    board.debug_uf('R')

    print("\n=== 重點：DFS vs Union Find ===")
    print("DFS: 每次查詢領地大小需要 O(領地大小) 遍歷")
    print("Union Find: 查詢只需要 O(1)！")
    print(f"\n現在查詢 (0,0) 的領地大小: {board.area_size(0, 0)}")
    print(f"現在查詢 (2,0) 的領地大小: {board.area_size(2, 0)}")
    print("↑ 兩者屬於同一領地，所以大小相同")
