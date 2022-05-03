class Solution:
    def checkSpecialArray(self, arr):
        count = 0

        for i in range(len(arr)):
            if arr[i]%2 == 1:
                count +=1
        
        if count % 2 == 0:
            return "Yes"
        else:
            return "No"

obj = Solution()
arr = [9, 17]
ans = obj.checkSpecialArray(arr)
print(ans)