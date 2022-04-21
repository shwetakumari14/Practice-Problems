class Solution:
    def digitOne(self, num):
        ans  = 0
        i = 1
        while i <= num:
            div = i * 10
            ans += (num//div)*i + min(max(num%div - i + 1, 0), i)
            i = i * 10

        return int(ans)

obj = Solution()
print(obj.digitOne(10))