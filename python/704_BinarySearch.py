class Solution(object):
    def search(self, nums, target):
        l, r = 0, len(nums) - 1
        m = int((r + l) / 2)
        while l <= m and l <= r:
            curr = nums[m]
            if target == curr:
                return m
            elif target < curr:
                r = m - 1
            else:
                l = m + 1
            m = int((r + l) / 2)
        return -1


if __name__ == '__main__':
    solution = Solution()
#    nums = [-1, 0, 3, 5, 9, 12]
#    target = 9
#    nums = [-1, 0, 3, 5, 9, 12]
#    target = 2
    nums = [5]
    target = -5
    ans = solution.search(nums, target)
    print(ans)

