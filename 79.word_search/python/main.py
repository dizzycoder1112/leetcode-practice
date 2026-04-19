class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
                # Optimization 1: Character count check
        board_count = Counter(char for row in board for char in row) 
        word_count = Counter(word)  

        for char, count in word_count.items(): 
            if count > board_count.get(char, 0):  
                return False

        # Optimization 2: Start from less common character
        if board_count[word[0]] > board_count[word[-1]]:  
            word = word[::-1]
            
        rows, cols = len(board), len(board[0])
        for y in range(rows):
            for x in range(cols):
                if self.dfs(board, word, 0, x, y):
                    return True

        return False
        


    def dfs(self, board: List[List[str]], word: str, i: int, x:int, y:int) -> bool:
        rows, cols = len(board), len(board[0])
        if i == len(word):
            return True
        
        if x<0 or x>=cols or y<0 or y>=rows or board[y][x] != word[i]:
            return False
        
        temp = board[y][x]
        board[y][x] = '0'

        found = (
            self.dfs(board, word, i+1, x+1, y) or
            self.dfs(board, word, i+1, x-1, y) or
            self.dfs(board, word, i+1, x, y+1) or
            self.dfs(board, word, i+1, x, y-1)
            )
        
        board[y][x] = temp

        return found
        

        