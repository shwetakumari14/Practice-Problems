class Solution:
    def luckyNumbers(self, a):

        factors, count = [0]*(a+1), 0

        for i in range(2, a+1):
            if factors[i] == 0:
                j = 2*i
                while j <= a:
                    factors[j] += 1
                    j += i
        
        for i in range(6, a+1):
            if factors[i] == 2:
                count += 1
        
        return count

obj = Solution()
a = 8
ans = obj.luckyNumbers(a)
print(ans)