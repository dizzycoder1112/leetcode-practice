Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

 

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
 

Constraints:

3 <= nums.length <= 3000
-105 <= nums[i] <= 105

---

## Real-World Applications (實務應用)

### 1. Financial Trading (金融交易)
**場景：對沖交易組合 (Hedging Portfolio)**

```go
// 找出三個資產的組合，使其總風險暴露為零
assets := []int{-5, 3, 2, -8, 7, -3}  // 每個資產的風險值（正/負）
// 目標：找出三個資產組合，總風險 = 0

// 實際應用：
// [-8, 3, 5] → -8 + 3 + 5 = 0 (平衡的投資組合)
// [-3, -5, 8] → -3 + (-5) + 8 = 0
```

**使用場景：**
- 構建風險中性投資組合
- 平衡多空倉位
- 對沖外匯風險

---

### 2. Chemistry (化學配方)
**場景：化學平衡方程式**

```go
// 找出三種化合物的組合使反應平衡
compounds := []int{-2, 1, 1, -3, 2, 3}  // 化合價/電荷
// 目標：找出三種化合物組合使總電荷為 0

// 實際應用：
// 配製 pH 中性溶液
// 平衡氧化還原反應
```

---

### 3. Team Assignment (團隊分配)
**場景：平衡團隊技能**

```go
// 三個員工的技能互補，總和為理想值
skillLevels := []int{-3, -1, 0, 1, 2, 4}
// 負數：需要支援的領域
// 正數：專長領域
// 目標：組成互補團隊 (總和 = 0)

// 結果：
// [-3, 1, 2] → Junior + Mid + Senior = 平衡團隊
// [-1, -3, 4] → 兩個新手配一個專家
```

**實際應用：**
- 組建項目團隊（前端 + 後端 + DevOps）
- 分配工作負載
- 配對導師與學員

---

### 4. Logistics (物流配送)
**場景：配送路線優化**

```go
// 三個配送點的位置差異總和為 0（回到起點）
locations := []int{-10, -5, 0, 5, 10, 15}  // 相對於倉庫的距離
// 目標：找出三個配送點，總位移為 0（效率最高）

// 結果：
// [-10, 0, 10] → 左10 → 中心 → 右10 → 回倉庫
// [-5, -10, 15] → 往返路線平衡
```

---

### 5. Gaming (遊戲設計)
**場景：角色平衡**

```go
// 組建平衡的三人隊伍
characterStats := []int{-2, -1, 1, 2, 3, -3}
// 負數：防禦型角色
// 正數：攻擊型角色
// 目標：攻防平衡隊伍 (總和 = 0)

// 結果：
// [-3, 1, 2] → 坦克 + DPS + 輔助
// [-2, -1, 3] → 兩個坦克 + 一個高攻擊
```

**實際應用：**
- MOBA 遊戲英雄組合
- RPG 隊伍配置
- 卡牌遊戲套牌平衡

---

### 6. Inventory Management (庫存管理)
**場景：調撥平衡**

```go
// 三個倉庫間的貨物調撥，使總庫存變化為 0
warehouses := []int{-100, -50, 50, 100, 150, -150}
// 負數：需要補貨
// 正數：有多餘庫存
// 目標：三個倉庫間互相調撥，總和為 0

// 結果：
// [-100, -50, 150] → 兩個缺貨倉庫從盈餘倉庫調貨
// [-150, 50, 100] → 大缺貨倉庫從兩個盈餘倉庫調貨
```

---

## Algorithm Pattern Summary (演算法模式總結)

### When to use 3Sum approach (何時使用 3Sum 方法)

✅ **適用場景：**
1. 需要找「三個元素的組合」
2. 有特定的「總和目標」（通常是 0 或固定值）
3. 需要「所有符合條件的組合」
4. 可以接受 O(n²) 時間複雜度

❌ **不適用場景：**
1. 只需要一個答案（用 HashMap 更快）
2. 元素數量超過 10,000（太慢）
3. 需要保持原始順序（排序會破壞順序）

### Core Techniques (核心技巧)

```
1. Sort + Two Pointers (排序 + 雙指針)
   - 時間：O(n²)
   - 空間：O(1)
   - 優點：去重複簡單、程式碼清晰

2. Deduplication (去重複)
   - 跳過重複的第一個元素
   - 跳過重複的左右指針

3. Early Termination (提前終止)
   - 如果當前數字 > 0 且是正數，可以提前結束
   - 優化實際執行時間
```

---

## Related Problems (相關題目)

| Problem | Time Complexity | Key Difference |
|---------|----------------|----------------|
| Two Sum | O(n) | HashMap, 只找一對 |
| **3Sum** | **O(n²)** | **排序 + 雙指針** |
| 3Sum Closest | O(n²) | 找最接近目標的和 |
| 4Sum | O(n³) | 多固定一個數 |
| kSum | O(n^(k-1)) | 遞迴固定 k-2 個數 |

---

