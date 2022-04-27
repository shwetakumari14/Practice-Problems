class Solution:
    def primeSum(self, a):
        prime = [True]*a
        prime[0], prime[1] = False, False

        for i in range(2,a+1):
            if not prime[i]:
                continue
            if i > a/i:
                break

            j = i*i
            while j <a:
                prime[j] = False
                j += i
        
        ans = [0]*2
        
        for i in range(2, a+1):
            if prime[i] and prime[a-i]:
                ans[0] = i
                ans[1] = a-i
                break


        return ans

obj = Solution()
a = 6
ans = obj.primeSum(a)
print(ans)