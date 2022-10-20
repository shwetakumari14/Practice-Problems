class Solution:
    def longestSubsequence(self, nums):
        dp = [1]*len(nums)

        for i in range(1, len(nums)):
            for j in range(i):
                if nums[i] > nums[j]:
                    dp[i] = max(dp[i], dp[j]+1)
        
        return max(dp)

obj = Solution()
nums = [10,9,2,5,3,7,101,18]
ans = obj.longestSubsequence(nums)
print(ans)