def gcd(a, b):
    
    if a < b:
        a, b = b, a 

    if b == 0:
        return a
    
    return gcd(b, a%b)


class Solution:
    def deleteEle(self, arr):
        num = arr[0]

        for i in range(1, len(arr)):
            num = gcd(num, arr[i])
        
        if num == 1:
            return 0
        
        return -1

obj = Solution()
arr = [7, 2, 5]
ans = obj.deleteEle(arr)
print(ans)