class Solution(object):
    def validateStackSequences(self, pushed, popped):
        pushed_idx, popped_idx = 0, 0
        for elem in pushed:
            pushed[pushed_idx] = elem
            pushed_idx += 1
            while pushed_idx > 0 and pushed[pushed_idx - 1] == popped[popped_idx]:
                pushed_idx -= 1
                popped_idx += 1
            pushed_idx += 1
        return pushed_idx == 0
#
#        stack = []
#        popped_idx = 0
#        for elem in pushed:
#            stack.append(elem)
#            while len(stack) != 0:
#                if stack[-1] == popped[popped_idx]:
#                    stack.pop()
#                    popped_idx += 1
#                else:
#                    break
#        return len(stack) == 0
#        while pushed_idx != length or popped_idx != length:
#            if len(stack) == 0:
#                if pushed_idx == length:
#                    return False
#                stack.append(pushed[pushed_idx])
#                pushed_idx += 1
#            if popped_idx == length:
#                return False
#            if stack[-1] != popped[popped_idx]:
#                if pushed_idx < length:
#                    stack.append(pushed[pushed_idx])
#                    pushed_idx += 1
#                else:
#                    return False
#            else:
#                stack.pop()
#                popped_idx += 1
#        return True


if __name__ == '__main__':
    solution = Solution()
#    pushed = [1, 2, 3, 4, 5]
#    popped = [4, 5, 3, 2, 1]
#    pushed = [1, 2, 3, 4, 5]
#    popped = [4, 3, 5, 1, 2]
    pushed = [4, 0, 1, 2, 3]
    popped = [4, 2, 3, 0, 1]
    ans = solution.validateStackSequences(pushed, popped)
    print(ans)

