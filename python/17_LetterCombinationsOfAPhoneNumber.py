class Solution(object):
    def letterCombinations(self, digits):
        if len(digits) == 0:
            return []
        alphabets = ['0', '1', 'abc', 'def', 'ghi', 'jkl', 'mno', 'pqrs', 'tuv', 'wxyz']
        res = ['']
        for digit in digits:
            digit = int(digit)
            for i in range(len(res)-1, -1, -1):
                elem = res.pop(i)
                for a in alphabets[digit]:
                    res.append(elem + a)
        return res

#        if len(digits) == 0:
#            return []
#        alphabets_dict = {
#            '1': [],
#            '2': ['a', 'b', 'c'],
#            '3': ['d', 'e', 'f'],
#            '4': ['g', 'h', 'i'],
#            '5': ['j', 'k', 'l'],
#            '6': ['m', 'n', 'o'],
#            '7': ['p', 'q', 'r', 's'],
#            '8': ['t', 'u', 'v'],
#            '9': ['w', 'x', 'y', 'z'],
#            '0': [' ']
#        }
#        res = alphabets_dict[digits[0]].copy()
#        for c in digits[1:]:
#            temp = res.copy()
#            res.clear()
#            for elem in temp:
#                for alphabet in alphabets_dict[c]:
#                    res.append(elem + alphabet)
#        return res


if __name__ == '__main__':
    solution = Solution()
    ans = solution.letterCombinations('43')
    print(ans)