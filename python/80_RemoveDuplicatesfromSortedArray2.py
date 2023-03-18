class Solution(object):
    def removeDuplicates(self, nums):
        k = len(nums)
        prev_num = -10001
        prev_count = 0
        for i in range(len(nums)-1, -1, -1):
            if nums[i] != prev_num:
                prev_num = nums[i]
                prev_count = 1
                continue
            prev_count += 1
            if prev_count > 2:
                nums.pop(i)
                k -= 1
        return k

if __name__ == '__main__':
    solution = Solution()
    nums = [0,0,1,1,1,1,2,3,3]
    k = solution.removeDuplicates(nums)
    print(k)
    print(nums)