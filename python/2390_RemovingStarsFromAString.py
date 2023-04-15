class Solution:
    def removeStars(self, s: str) -> str:
        res = ''
        count = 0
        for i in range(len(s) - 1, -1, -1):
            c = s[i]
            if c == '*':
                count += 1
            else:
                if count == 0:
                    res += c
                else:
                    count -= 1
        return res[::-1]


if __name__ == '__main__':
    solution = Solution()
    s = 'leet**cod*e'
    s = 'erase*****'
    ans = solution.removeStars(s)
    print(ans)