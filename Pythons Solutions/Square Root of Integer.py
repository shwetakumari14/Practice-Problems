class Solution:
    def sqrtOfNum(self, n):
        left, right = 1, n

        while left < right - 1:
            mid = left + (right - left)//2
            if mid <= n//mid:
                left = mid
            else:
                right = mid - 1
        
        if right <= n//right:
            return right
        
        return left

obj = Solution()
ans = obj.sqrtOfNum(256)
print(ans)