
def gcd(a, b):
    
    if a < b:
        a, b = b, a 

    if b == 0:
        return a
    
    return gcd(b, a%b)

class Solution:
    def specialNum(self, a,b,c):
        lcm = (b*c)/gcd(b,c)
        count, i = 0, lcm

        while i<= a:
            if i%lcm == 0:
                count+=1
            i+=lcm
        
        return count

obj = Solution()
ans = obj.specialNum(12,3,2)
print(ans)