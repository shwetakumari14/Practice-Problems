class Solution:
    def columnNumber(self, str):
        ans = 0

        for char in str:
            ans = ans*26 + ord(char) - 64

        return ans

obj = Solution()
str = "ABCD"
ans = obj.columnNumber(str)
print(ans)