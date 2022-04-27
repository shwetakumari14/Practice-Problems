class Solution:
    def oldFibonacci(self, a, b):
        numbers = b - a + 1
        mulOf3 = (b//3) - ((a-1)//3)
        return numbers - mulOf3

obj = Solution()
a, b = 6, 20
ans = obj.oldFibonacci(a, b)
print(ans)