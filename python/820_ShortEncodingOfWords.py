class Solution(object):
    def minimumLengthEncoding(self, words):
        min_len = 0
        words.sort(key=len, reverse=True)
        for i in range(len(words)):
            word = words[i]
            is_ending_substring = False
            for prev_word in words[: i]:
                if prev_word.endswith(word):
                    is_ending_substring = True
                    break
            if not is_ending_substring:
                min_len += len(word) + 1
        return min_len


if __name__ == '__main__':
    solution = Solution()
#    words = ["time", "me", "bell"]
#    words = ["t"]
    words = ["feipyxx","e"]
    ans = solution.minimumLengthEncoding(words)
    print(ans)