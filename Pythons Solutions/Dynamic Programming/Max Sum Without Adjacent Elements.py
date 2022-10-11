class Solution:
    def maxSumTopDown(self, arr):
        if len(arr) == 0:
            return arr[0]
        
        if len(arr) == 2:
            return max(arr)
        
        dp = [-1]*len(arr)
        dp[0], dp[1] = arr[0], max(arr[0], arr[1])

        for i in range(2, len(arr)):
            dp[i] = max((arr[i] + dp[i-2]), dp[i-1])

        return dp[len(arr)-1]
    
    def maxSumBottomUp(self, arr):
        if len(arr) == 0:
            return arr[0]
        
        if len(arr) == 2:
            return max(arr)
        
        ones, twos = arr[0], max(arr[0], arr[1])

        for i in range(2, len(arr)):
            temp = twos
            twos = max((arr[i] + ones), twos)
            ones = temp

        return twos

obj = Solution()
arr = [4, 8, 10, 7, 6]
ans = obj.maxSumBottomUp(arr)
print(ans)