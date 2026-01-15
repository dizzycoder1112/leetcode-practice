from typing import List


class Solution:
    def separateSquares(self, squares: List[List[int]]) -> float:
        events = []

        for (x, y, l) in squares:
            events.append({"y": y, "delta": 1, "x1": x, "x2": x + l})      # 1 means enter
            events.append({"y": y + l, "delta": -1, "x1": x, "x2": x + l}) # -1 means leave

        # 排序：同 y 時，進入 (delta=1) 排在離開 (delta=-1) 之前
        events.sort(key=lambda e: (e["y"], -e["delta"]))

        # ========== 第一次掃描：計算總面積 ==========
        active = []
        totalArea = 0
        prevY = events[0]["y"]

        i = 0
        while i < len(events):
            currY = events[i]["y"]
            if currY > prevY and len(active) > 0:
                intervalsCopy = [[x1, x2] for [x1, x2] in active]  # 深拷貝
                width = self.unionLength(intervalsCopy)
                totalArea += width * (currY - prevY)

            while i < len(events) and events[i]["y"] == currY:
                e = events[i]
                if e["delta"] == 1:
                    active.append([e["x1"], e["x2"]])
                else:
                    active.remove([e["x1"], e["x2"]])
                i += 1
            prevY = currY

        target = totalArea / 2

        # ========== 第二次掃描：找答案 ==========
        active = []
        area = 0
        prevY = events[0]["y"]

        i = 0
        while i < len(events):
            currY = events[i]["y"]
            if currY > prevY and len(active) > 0:
                intervalsCopy = [[x1, x2] for [x1, x2] in active]  # 深拷貝
                width = self.unionLength(intervalsCopy)
                extraArea = width * (currY - prevY)

                if area + extraArea >= target:
                    return (target - area) / width + prevY

                area += extraArea

            while i < len(events) and events[i]["y"] == currY:
                e = events[i]
                if e["delta"] == 1:
                    active.append([e["x1"], e["x2"]])
                else:
                    active.remove([e["x1"], e["x2"]])
                i += 1
            prevY = currY

        return prevY

    def unionLength(self, intervals: List[List[int]]) -> int:
        """計算區間聯集的總長度（LeetCode 56 的做法）"""
        if not intervals:
            return 0

        # 按起點排序
        intervals.sort(key=lambda x: x[0])
        merged = [intervals[0]]

        for i in range(1, len(intervals)):
            current = intervals[i]
            last = merged[-1]
            if current[0] <= last[1]:  # 重疊
                if current[1] > last[1]:
                    last[1] = current[1]
            else:  # 不重疊
                merged.append(current)

        total = 0
        for (start, end) in merged:
            total += end - start
        return total


# ========== 測試 ==========
if __name__ == "__main__":
    sol = Solution()

    # Example 1: Expected 1.00000
    squares1 = [[0, 0, 1], [2, 2, 1]]
    print(f"Example 1: {sol.separateSquares(squares1):.5f}")

    # Example 2: Expected 1.00000 (3454: 重疊只算一次)
    squares2 = [[0, 0, 2], [1, 1, 1]]
    print(f"Example 2: {sol.separateSquares(squares2):.5f}")

    # Example 3: 之前出錯的 case
    squares3 = [[26, 29, 3], [10, 24, 1]]
    print(f"Example 3: {sol.separateSquares(squares3):.5f}")  # Expected 30.33333

    # Example 4: 之前出錯的 case
    squares4 = [[16, 27, 3], [18, 24, 5]]
    print(f"Example 4: {sol.separateSquares(squares4):.5f}")
