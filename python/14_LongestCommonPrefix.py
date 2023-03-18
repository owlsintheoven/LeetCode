class Solution(object):
    def longestCommonPrefix(self, strs):
        prefix = ''
        for i in range(len(strs[0])):
            c = strs[0][i]
            is_diff = False
            for str in strs[1:]:
                if i == len(str) or c != str[i]:
                    is_diff = True
                    break
            if not is_diff:
                prefix += c
            else:
                break
        return prefix


if __name__ == '__main__':
    solution = Solution()
#    strs = ["flower", "flow", "flight"]
    strs = ["dog", "racecar", "car"]
    print(strs)
    ans = solution.longestCommonPrefix(strs)
    print(ans)