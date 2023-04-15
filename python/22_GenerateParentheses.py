class Solution(object):
    def generateParenthesis(self, n):
        dp = [[] for _ in range(n + 1)]
        dp[1] = ['()']
        for i in range(2, n + 1):
            prev = dp[i - 1]
            all = []
            for s in prev:
                lefts = []
                for idx in range(len(s)):
                    c = s[idx]
                    if c == '(':
                        lefts.append(idx)
                    else:
                        for left_idx in lefts:
                            all.append(s[: left_idx + 1] + '(' + s[left_idx + 1: idx + 1] + ')' + s[idx + 1:])
                all.append('()' + s)
                all.append(s + '()')
                all.append('(' + s + ')')
            dp[i] = list(set(all))
        return dp[-1]


if __name__ == '__main__':
    solution = Solution()
    ans = solution.generateParenthesis(6)
    print(ans)

