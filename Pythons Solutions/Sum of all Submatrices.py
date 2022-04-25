class Solution:
    def sumOfSubMatrices(self, arr):
        n, m, sum = len(arr), len(arr[0]), 0

        for i in range(len(arr)):
            for j in range(len(arr[0])):
                sum += (i+1)*(n-i)*(j+1)*(m-j)*arr[i][j]

        return sum

obj = Solution()
arr = [[1,1],[1,1]]
ans = obj.sumOfSubMatrices(arr)
print(ans)