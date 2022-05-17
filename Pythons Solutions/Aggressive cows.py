class Solution:
    def aggressiveCows(self, arr, b):
        arr.sort()
        left, right, ans = 1, arr[len(arr)-1], 0

        while left <= right:
            mid = left + (right+left)//2
            if self.isPossible(mid, arr, b):
                ans = mid
                left = mid + 1
            else:
                right = mid - 1

        
        return ans
    
    def isPossible(self, x, arr, c):
        j, count = 0, 1

        for i in range(len(arr)):
            if arr[i] - arr[j] >= x:
                j = i
                count += 1
        
        return count >= c

obj = Solution()
arr, b = [1, 2, 3, 4, 5], 3
ans = obj.aggressiveCows(arr, b)
print(ans)