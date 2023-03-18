# c0b0b0d
# 1010101
# 1012101
# max_index = 3
# iterate through dp, dp_index -> s_index -> left == right? -> dp[i] += 2
class Solution(object):
    def longestPalindrome(self, s):
        if s == s[::-1]:
            return s
        dp = [1 if i % 2 == 0 else 0 for i in range(2 * len(s) - 1)]
        max_dp_index = 0
        for i in range(len(dp)):
            s_index = float(i / 2)
            # the length of possible palindrome is an odd number
            if dp[i] == 1:
                dst = 1
            # the length of possible palindrome is an even number
            else:
                dst = 0.5
            while int(s_index - dst) >= 0 and int(s_index + dst) < len(s):
                if s[int(s_index - dst)] == s[int(s_index + dst)]:
                    dst += 1
                    dp[i] += 2
                else:
                    break
            if dp[i] > dp[max_dp_index]:
                max_dp_index = i
        return s[int(max_dp_index / 2 - (dp[max_dp_index] - 1) / 2): int(max_dp_index / 2 + (dp[max_dp_index] - 1) / 2) + 1]


if __name__ == '__main__':
    solution = Solution()
#    s = 'babad'
    s = 'cbbd'
#    s = 'aaa'
    print(s)
    ans = solution.longestPalindrome(s)
    print(ans)

