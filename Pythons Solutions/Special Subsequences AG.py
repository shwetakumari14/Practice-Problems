class Solution:
    def findAG(self, str):
        aCount, agCount = 0, 0

        for char in str:
            if char == "A":
                aCount +=1
            elif char == "G":
                agCount += aCount

        return agCount

obj = Solution()
str = "GAB"
ans = obj.findAG(str)
print(ans)