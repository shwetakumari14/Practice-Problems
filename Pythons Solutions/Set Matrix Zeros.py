class Solution:
    def setToZero(self, arr):
        row, col = [False]*len(arr), [False]*len(arr[0])

        for i in range(len(arr)):
            for j in range(len(arr[0])):
                if arr[i][j] == 0:
                    row[i] = True
                    col[j] = True

        for i in range(len(arr)):
            for j in range(len(arr[0])):
                if row[i] == True or col[j] == True:
                    arr[i][j] == 0

        return arr

obj = Solution()
arr = [[1, 0, 0],[1, 1, 1],[1, 1, 1]]
ans = obj.setToZero(arr)
print(ans)