class Solution:
    def longestSubstring(self, a, b):
        m, n = len(a), len(b)
        L = [[0 for i in range(n+1)] for j in range(n+1)]

        for i in range(0, m+1):
            for j in range(0, n+1):
                if i == 0 or j == 0:
                    L[i][j] = 0
                elif a[i-1] == b[j-1]:
                    L[i][j] = L[i-1][j-1] + 1
                else:
                    L[i][j] = max(L[i-1][j], L[i][j-1])
        
        return L[m][n]

obj = Solution()
a, b = "abbcdgf", "bbadcgf"
ans = obj.longestSubstring(a, b)
print(ans)