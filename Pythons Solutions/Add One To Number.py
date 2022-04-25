class Solution:
    def addOneToNum(self, arr):
        temp, carry, ans = [], 1, []

        for i in range(len(arr)-1, -1, -1):
            sum = arr[i] + carry
            carry = sum // 10
            temp.append(sum%10)
            sum = 0
        
        if carry > 0:
            temp.append(carry)

        idx = len(temp)-1
        for i in range(len(arr)-1, -1, -1):
            if temp[i] == 0:
                idx -=1
        
        for i in range(idx, -1, -1):
            ans.append(temp[i])

        return ans

obj = Solution()
arr = [0, 1, 2, 3]
ans = obj.addOneToNum(arr)
print(ans)