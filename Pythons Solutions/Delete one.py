def gcd(a, b):
    
    if a < b:
        a, b = b, a 

    if b == 0:
        return a
    
    return gcd(b, a%b)

class Solution:
    def deleteOne(self, arr):
        left, right, n, ans = [0]*len(arr),[0]*len(arr), len(arr), 0
        left[0] = arr[0]
        right[n-1] = arr[n-1]

        for i in range(1, len(arr)):
            left[i] = gcd(left[i-1], arr[i])

        for i in range(len(arr)-2, 0):
            right[i] = gcd(right[i+1], arr[i])

        for i in range(len(arr)):
            if i == 0:
                if right[i+1] > ans:
                    ans = right[i+1]
            else:
                if i == n-1:
                    if left[i-1] > ans:
                        ans = left[i-1]
                elif gcd(left[i-1], right[i+1]) > ans:
                    ans = gcd(left[i-1], right[i+1])
                    

        return ans

obj = Solution()
arr = [5, 15, 30]
ans = obj.deleteOne(arr)
print(ans)