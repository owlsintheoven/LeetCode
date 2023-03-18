class Solution(object):
    def countSubstrings(self, s):
        count = 0
        pre_part = []
        for i in range(2 * len(s) - 1):
            center = i / 2
            if i % 2 == 0:
                # center is a character
                count += 1
                dst = 1
                pre_part.append(s[int(i / 2)])
            else:
                # center is an empty space
                center = i / 2
                dst = 0.5
            while center - dst >= 0 and center + dst <= len(s) - 1:
                if s[int(center - dst)] == s[int(center + dst)]:
                    count += 1
                    dst += 1
                else:
                    break
        return count


if __name__ == '__main__':
    solution = Solution()
    s = 'aaa'
    ans = solution.countSubstrings(s)
    print(ans)