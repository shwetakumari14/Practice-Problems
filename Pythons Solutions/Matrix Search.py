class Solution:
    def findElementInSortedMatrix(self, arr, b):
        n, m = len(arr)-1, len(arr[0])-1
        i, j = 0, m

        while i <=n and j >=0:
            if arr[i][j] == b:
                return 1
            if arr[i][j] > b:
                j -= 1
            else:
                i += 1
        
        return 0

obj = Solution()
arr, b = [[1,   3,  5,  7], [10, 11, 16, 20], [23, 30, 34, 50]], 3
ans = obj.findElementInSortedMatrix(arr, b)
print(ans)