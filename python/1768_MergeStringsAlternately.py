class Solution(object):
    def mergeAlternately(self, word1, word2):
        l1, l2 = len(word1), len(word2)
        limit = min(l1, l2)
        res = ''
        for i in range(limit):
            res += (word1[i] + word2[i])
        res += (word1[limit:] + word2[limit:])
        return res


if __name__ == '__main__':
    solution = Solution()
    word1 = 'abc'
    word2 = 'pqr'
    ans = solution.mergeAlternately(word1, word2)
    print(ans)

