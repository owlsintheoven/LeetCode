class Solution(object):
    def romanToInt(self, s):
        roman_dict = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000
        }
        count = 0
        prev = -1
        op = 1
        for c in s[::-1]:
            curr = roman_dict[c]
            if curr > prev:
                count += curr
                op = 1
            elif curr == prev:
                count += curr * op
            else:
                op = -1
                count += curr * op
            prev = curr
        return count


if __name__ == '__main__':
    solution = Solution()
    roman_str = 'MCMXCIV'
    roman_str = 'III'
    ans = solution.romanToInt(roman_str)
    print(ans)