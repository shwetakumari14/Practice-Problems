class Solution:
    def check(self, x, arr, b):
        n, j, cnt = len(arr), 0, 1

        for i in range(1, n):
            if arr[i] - arr[j] >= x:
                j = i
                cnt += 1
        return cnt >= b


    def aggressiveCows(self, arr, b):
        low, high, ans = 0, 1e9, 1

        while low <= high:
            mid = int(low + (high - low)//2)
            if self.check(mid, arr, b):
                ans = mid
                low = mid + 1
            else:
                high = mid - 1
        
        return ans

obj = Solution()
arr, b = [1, 2, 3, 4, 5], 3
ans = obj.aggressiveCows(arr, b)
print(ans)