class Solution:
    def maxProductSubArray(self, arr):
        maxProd, minProd, ans = arr[0], arr[0], arr[0]

        for i in range(1, len(arr)):
            x = max(arr[i], maxProd*arr[i], minProd*arr[i])
            y = min(arr[i], maxProd*arr[i], minProd*arr[i])
            maxProd, minProd = x, y
            ans = max(ans, maxProd)
        
        return ans

obj = Solution()
arr = [4, 2, -5, 1]
ans = obj.maxProductSubArray(arr)
print(ans)