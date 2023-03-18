class Solution(object):
    def isValid(self, s):
        bracket_dict = {}
        bracket_dict['('] = ')'
        bracket_dict['{'] = '}'
        bracket_dict['['] = ']'
        stack = []
        for c in s:
            if len(stack) == 0:
                if c in bracket_dict:
                    stack.append(c)
                else:
                    return False
            elif c in bracket_dict:
                stack.append(c)
            elif c == bracket_dict[stack[-1]]:
                stack.pop(-1)
            else:
                return False
        return len(stack) == 0


if __name__ == '__main__':
    solution = Solution()
    answer = solution.isValid('{[]}')
    print(answer)
