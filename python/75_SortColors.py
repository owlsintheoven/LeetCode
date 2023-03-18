class Solution(object):
    def sortColors(self, nums):
        count = {
            0: 0,
            1: 0,
            2: 0
        }
        for num in nums:
            count[num] += 1
        nums[: count[0]] = [0 for i in range(count[0])]
        nums[count[0]: count[0] + count[1]] = [1 for i in range(count[1])]
        nums[count[0] + count[1]:] = [2 for i in range(count[2])]

        return nums


if __name__ == '__main__':
    solution = Solution()
    nums = [2,0,2,1,1,0]
#    nums = [2,0,1]
    print(nums)
    ans = solution.sortColors(nums)
    print(ans)