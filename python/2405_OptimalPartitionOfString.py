class Solution(object):
    def partitionString(self, s):
        if len(s) == 0:
            return 0
        prev = ''
        count = 1
        for c in s:
            if c not in prev:
                prev += c
            else:
                prev = c
                count += 1
        return count


if __name__ == '__main__':
    solution = Solution()
    s = 'abacaba'
    s = ''
    ans = solution.partitionString(s)
    print(ans)

