class Solution:
    def maxHeighrOfStairCase(self, A):
        low, high, ans = 1, 1000000000, 0

        while low <= high:
            mid = low + (high-low)//2
            val = mid * (mid+1)//2
            if val == A:
                return mid
            if val < A:
                ans = max(ans, mid)
                low = mid + 1
            else:
                high = mid - 1
        
        return ans

obj = Solution()
ans = obj.maxHeighrOfStairCase(10)
print(ans)