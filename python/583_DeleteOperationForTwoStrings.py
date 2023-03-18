class Solution(object):
    def minDistance(self, word1, word2):
        # cannot change order
        shorter = ''
        longer = ''
        if len(word1) > len(word2):
            longer = word1
            shorter = word2
        else:
            longer = word2
            shorter = word1
        if len(shorter) == 0:
            return len(longer)
        first_appearance = [[] for i in range(len(shorter))]
        char_count = {}
        for i in range(len(shorter)):
            if shorter[i] not in char_count:
                char_count[shorter[i]] = 1
            else:
                char_count[shorter[i]] += 1
            encountered = 0
            for j in range(len(longer)):
                if shorter[i] == longer[j]:
                    first_appearance[i].append(j)
                    encountered += 1
                    if encountered == char_count[shorter[i]]:
                        break
        for i in range(len(first_appearance)-1, -1, -1):
            if len(first_appearance[i]) == 0:
                first_appearance.pop(i)
        if len(first_appearance) == 0:
            return len(longer) + len(shorter)
        longest = 0
        dp = [0 for i in range(len(first_appearance))]
        if len(first_appearance[0]) != 0:
            dp[0] = 1
            longest = 1
        for i in range(1, len(first_appearance)):
            if first_appearance[i][-1] > first_appearance[i-1][0]:
                dp[i] = dp[i-1] + 1
            else:
                dp[i] = 1
                if i > 1:
                    reverse_i = i - 2
                    while reverse_i >= 0:
                        if first_appearance[i][-1] > first_appearance[reverse_i][0]:
                            dp[i] = dp[reverse_i] + 1
                            break
                        reverse_i -= 1
            longest = max(longest, dp[i])
        return len(word1) + len(word2) - 2*longest


if __name__ == '__main__':
    solution = Solution()
    word1 = 'abc' #10
    word2 = 'e' #5
    ans = solution.minDistance(word1, word2)
    print(ans)