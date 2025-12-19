# 424. Longest Repeating Character Replacement

## Problem

You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

**Example 1:**
```
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
```

**Example 2:**
```
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
```

**Constraints:**
- 1 <= s.length <= 10^5
- s consists of only uppercase English letters.
- 0 <= k <= s.length

---

## Approach: Sliding Window

### Key Insight

For a window to be valid (can be converted to all same characters with at most k replacements):

```
Characters to replace = Window Length - Max Frequency Character Count

windowLen - maxFreq <= k  →  Valid Window
```

### Algorithm

1. Use two pointers (left, right) to maintain a sliding window
2. Track frequency of each character in the window
3. Track the maximum frequency of any single character
4. Expand right pointer, shrink left when window becomes invalid

### Visual Walkthrough

```
Input: s = "AABABBA", k = 1

Valid window condition: windowLen - maxFreq <= k

═══════════════════════════════════════════════════════════

Step 1: right=0
┌───┬───┬───┬───┬───┬───┬───┐
│ A │ A │ B │ A │ B │ B │ A │
└───┴───┴───┴───┴───┴───┴───┘
  L
  R

freq: {A:1}  maxFreq=1
windowLen=1, replace: 1-1=0 <= 1 ✓
maxLen = 1

═══════════════════════════════════════════════════════════

Step 2: right=1
┌───┬───┬───┬───┬───┬───┬───┐
│ A │ A │ B │ A │ B │ B │ A │
└───┴───┴───┴───┴───┴───┴───┘
  L   R
  └───┘

freq: {A:2}  maxFreq=2
windowLen=2, replace: 2-2=0 <= 1 ✓
maxLen = 2

═══════════════════════════════════════════════════════════

Step 3: right=2
┌───┬───┬───┬───┬───┬───┬───┐
│ A │ A │ B │ A │ B │ B │ A │
└───┴───┴───┴───┴───┴───┴───┘
  L       R
  └───────┘

freq: {A:2, B:1}  maxFreq=2
windowLen=3, replace: 3-2=1 <= 1 ✓
maxLen = 3

═══════════════════════════════════════════════════════════

Step 4: right=3
┌───┬───┬───┬───┬───┬───┬───┐
│ A │ A │ B │ A │ B │ B │ A │
└───┴───┴───┴───┴───┴───┴───┘
  L           R
  └───────────┘

freq: {A:3, B:1}  maxFreq=3
windowLen=4, replace: 4-3=1 <= 1 ✓
maxLen = 4  ← Can replace B with A to get "AAAA"

═══════════════════════════════════════════════════════════

Step 5: right=4
┌───┬───┬───┬───┬───┬───┬───┐
│ A │ A │ B │ A │ B │ B │ A │
└───┴───┴───┴───┴───┴───┴───┘
  L               R
  └───────────────┘

freq: {A:3, B:2}  maxFreq=3
windowLen=5, replace: 5-3=2 > 1 ✗ Invalid!

Shrink left:
┌───┬───┬───┬───┬───┬───┬───┐
│ A │ A │ B │ A │ B │ B │ A │
└───┴───┴───┴───┴───┴───┴───┘
      L           R
      └───────────┘

freq: {A:2, B:2}  maxFreq=3 (not updated, see optimization below)
windowLen=4, replace: 4-3=1 <= 1 ✓
maxLen = 4

═══════════════════════════════════════════════════════════

Step 6: right=5
┌───┬───┬───┬───┬───┬───┬───┐
│ A │ A │ B │ A │ B │ B │ A │
└───┴───┴───┴───┴───┴───┴───┘
      L               R
      └───────────────┘

freq: {A:2, B:3}  maxFreq=3
windowLen=5, replace: 5-3=2 > 1 ✗ Invalid!

Shrink left:
┌───┬───┬───┬───┬───┬───┬───┐
│ A │ A │ B │ A │ B │ B │ A │
└───┴───┴───┴───┴───┴───┴───┘
          L           R
          └───────────┘

freq: {A:1, B:3}  maxFreq=3
windowLen=4, replace: 4-3=1 <= 1 ✓
maxLen = 4

═══════════════════════════════════════════════════════════

Step 7: right=6
┌───┬───┬───┬───┬───┬───┬───┐
│ A │ A │ B │ A │ B │ B │ A │
└───┴───┴───┴───┴───┴───┴───┘
          L               R
          └───────────────┘

freq: {A:2, B:3}  maxFreq=3
windowLen=5, replace: 5-3=2 > 1 ✗ Invalid!

Shrink left:
┌───┬───┬───┬───┬───┬───┬───┐
│ A │ A │ B │ A │ B │ B │ A │
└───┴───┴───┴───┴───┴───┴───┘
              L           R
              └───────────┘

freq: {A:2, B:2}  maxFreq=3
windowLen=4, replace: 4-3=1 <= 1 ✓
maxLen = 4

═══════════════════════════════════════════════════════════

Result: maxLen = 4
Best window: [A,A,B,A] → Replace B → "AAAA"
```

---

## Key Optimization: maxFreq Never Decreases

### Why don't we update maxFreq when shrinking?

After shrinking, the actual max frequency in window might be smaller than `maxFreq`. But we intentionally **don't decrease** it.

### Reason

```
Goal: Find the LONGEST valid substring

- If we previously achieved maxFreq=3, we found a valid window of length 4
- To find a LONGER window, maxFreq must be >= 3
- Keeping stale maxFreq=3 won't miss any longer answers
- It only means the current window might be "virtually valid"
  but maxLen already recorded the best result
```

### Benefit

- Avoids recalculating actual max frequency in window (would be O(26) each time)
- Time complexity stays O(n) instead of O(26n)

---

## Complexity

- **Time:** O(n) - each character visited at most twice
- **Space:** O(1) - frequency map has at most 26 entries

---

## Comparison with Similar Problems

| Problem | Goal | Window Valid Condition | Track |
|---------|------|------------------------|-------|
| #3 Longest Substring Without Repeating | Longest unique chars | No duplicates | Char positions |
| #424 This Problem | Longest same char (with k replacements) | windowLen - maxFreq <= k | Char frequencies |

Both use sliding window, but #3 tracks positions to detect duplicates, while #424 tracks frequencies to calculate replacements needed.
