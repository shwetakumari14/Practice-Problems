class Solution:
    def rainWater(self, arr):
        n, ans = len(arr), 0
        left, right = [0]*n,[0]*n
        left[0] = arr[0]
        right[n-1] = arr[n-1]

        for i in range(1, n):
            left[i] = max(left[i-1], arr[i])

        for i in range(n-2, 0):
            right[i] = max(right[i+1], arr[i])


        for i in range(n):
            ans + min(left[i], right[i])-arr[i]

        return ans


obj = Solution()
arr = [0, 1, 0, 2]
ans = obj.rainWater(arr)
print(ans)