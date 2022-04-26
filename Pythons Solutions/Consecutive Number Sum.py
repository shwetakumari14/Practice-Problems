class Solution:
    def numberSum(self, n):
        val, ans, i = 2*n, 0, 1

        while i*i <= n:
            if val % i == 0:
                temp = val/i
                if (temp-i-1)%2 == 0 and (temp-i-1)//2 >=0:
                    ans += 1
            i += 1
        return ans

obj = Solution()
ans = obj.numberSum(9)
print(ans)