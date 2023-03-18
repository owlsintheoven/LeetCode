class Solution(object):
    #def longestValidParentheses(self, s):
    #    stack = []
    #    i = 0
    #    valid_parentheses = []
    #    for c in s:
    #        if c == '(':
    #            stack.append(i)
    #        else:
    #            if len(stack) != 0:
    #                left_bracket = stack.pop()
    #                #continuous
    #                #outer
    #                if len(valid_parentheses) != 0:
    #                    prev_parentheses = valid_parentheses[-1]
    #                    if prev_parentheses[1]+1 == left_bracket:
    #                        valid_parentheses[-1] = (prev_parentheses[0], i)
    #                    elif left_bracket < prev_parentheses[0]:
    #                        if len(valid_parentheses) > 1:
    #                            prev_prev_parentheses = valid_parentheses[-2]
    #                            if prev_prev_parentheses[1]+1 == left_bracket:
    #                                valid_parentheses.pop()
    #                                valid_parentheses.pop()
    #                                valid_parentheses.append((prev_prev_parentheses[0], i))
    #                            else:
    #                                valid_parentheses[-1] = (left_bracket, i)
    #                        else:
    #                            valid_parentheses[-1] = (left_bracket, i)
    #                    else:
    #                        valid_parentheses.append((left_bracket, i))
    #                else:
    #                    valid_parentheses.append((left_bracket, i))
    #        i += 1
    #    longest = 0
    #    for parentheses in valid_parentheses:
    #        length = parentheses[1] - parentheses[0] + 1
    #        if length > longest:
    #            longest = length
    #    return longest
    def longestValidParentheses(self, s):
        longest = 0
        dp = [0 for i in range(len(s))]
        for i in range(len(s)):
            if s[i] == '(':
                continue
            if i != 0:
                if s[i-1] == '(':
                    dp[i] = 2 if i < 2 else 2 + dp[i-2]
                elif i-dp[i-1] >= 1 and s[i-dp[i-1]-1] == '(':
                    print(i, '(')
                    dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2 if i-dp[i-1] >= 2 else dp[i-1] + 2
            longest = max(longest, dp[i])
        return longest




if __name__ == '__main__':
#    test_str = '(()()()((()'
    test_str = '(()()()((()'
#    test_str = '()((()))'
#    test_str = ')()())'
#    test_str = '(()())'
    solution = Solution()
    ans = solution.longestValidParentheses(test_str)
    print(ans)

'''
public class Solution {
    public int longestValidParentheses(String s) {
        int maxans = 0;
        int dp[] = new int[s.length()];
        for (int i = 1; i < s.length(); i++) {
            if (s.charAt(i) == ')') {
                if (s.charAt(i - 1) == '(') {
                    dp[i] = (i >= 2 ? dp[i - 2] : 0) + 2;
                } else if (i - dp[i - 1] > 0 && s.charAt(i - dp[i - 1] - 1) == '(') {
                    dp[i] = dp[i - 1] + ((i - dp[i - 1]) >= 2 ? dp[i - dp[i - 1] - 2] : 0) + 2;
                }
                maxans = Math.max(maxans, dp[i]);
            }
        }
        return maxans;
    }
}'''