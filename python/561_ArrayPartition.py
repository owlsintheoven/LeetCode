class Solution(object):
    def arrayPairSum(self, nums):
        all = []
        self.helper(nums, [], all)
        print('len', len(all))
        for a in all:
            print(a)
        max_n = -1
        for a in all:
            total = sum([min(c) for c in a])
            max_n = total if total > max_n else max_n
        return max_n

    def helper(self, nums, current, all):
        if len(nums) == 0:
            all.append(current.copy())
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                b = nums.pop(j)
                a = nums.pop(i)
                current.append([a, b])
                self.helper(nums, current, all)
                nums.insert(i, a)
                nums.insert(j, b)
                current.pop()

#            current.append(nums.pop(i))
#            self.helper(nums, current, all)
#            nums.insert(i, current.pop())


if __name__ == '__main__':
    solution = Solution()
#    nums = [1, 4, 3, 2]
    nums = [6, 2, 6, 5, 1, 2]
    ans = solution.arrayPairSum(nums)
    print(ans)