class Solution:
    def openDoors(self, a):
        if a == 0:
            return 0

        left, right = 0, a

        while left < (right-1):
            mid = left + (right-left)//2
            if mid <= a // mid:
                left = mid
            else:
                right = mid - 1

        if right <= a // right:
            return right

        return left

obj = Solution()
a = 6
ans = obj.openDoors(a)
print(ans)