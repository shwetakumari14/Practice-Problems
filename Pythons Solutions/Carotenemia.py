class Solution:
    def carotenemia(self, arr, b):
        sum = 0
        for i in range(len(arr)):
            sum += arr[i]
            if sum >= b:
                return i

obj = Solution()
arr, b = [1, 3, 2, 4], 5
ans = obj.carotenemia(arr, b)
print(ans)