from typing import List


class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        if not matrix or not matrix[0]:
            return 0

        rows = len(matrix)
        cols = len(matrix[0])
        heights = [0] * cols
        max_area = 0

        for row in range(rows):
            # Step 1: Update heights for current row
            for col in range(cols):
                if matrix[row][col] == "1":
                    heights[col] += 1
                else:
                    heights[col] = 0

            # Step 2: Find largest rectangle in this histogram
            max_area = max(max_area, self.largestRectangleInHistogram(heights))

        return max_area

    def largestRectangleInHistogram(self, heights: List[int]) -> int:
        stack: List[tuple[int, int]] = []  # (start_index, height)
        max_area = 0

        for i, h in enumerate(heights):
            start = i
            # Pop taller bars and calculate their area
            while stack and stack[-1][1] > h:
                idx, height = stack.pop()
                width = i - idx
                max_area = max(max_area, height * width)
                start = idx
            stack.append((start, h))

        # Handle remaining bars in stack
        for idx, height in stack:
            width = len(heights) - idx
            max_area = max(max_area, height * width)

        return max_area


if __name__ == "__main__":
    solution = Solution()

    # Example 1
    matrix1 = [
        ["1", "0", "1", "0", "0"],
        ["1", "0", "1", "1", "1"],
        ["1", "1", "1", "1", "1"],
        ["1", "0", "0", "1", "0"]
    ]
    print(f"Example 1: {solution.maximalRectangle(matrix1)}")  # Expected: 6

    # Example 2
    matrix2 = [["0"]]
    print(f"Example 2: {solution.maximalRectangle(matrix2)}")  # Expected: 0

    # Example 3
    matrix3 = [["1"]]
    print(f"Example 3: {solution.maximalRectangle(matrix3)}")  # Expected: 1
