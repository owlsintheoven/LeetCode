class Solution(object):
    def shortestToChar(self, s, c):
        pos = []
        for i in range(len(s)):
            if s[i] == c:
                pos.append(i)
        dst = []
        pos_i = 0
        for i in range(len(s)):
            if pos_i != 0:
                dst.append(min(abs(i - pos[pos_i]), abs(i - pos[pos_i - 1])))
            else:
                dst.append(abs(i - pos[pos_i]))
            if i == pos[pos_i] and pos_i < len(pos) - 1:
                pos_i += 1
        return dst


if __name__ == '__main__':
    solution = Solution()
    s = 'aaba'
    c = 'b'
    ans = solution.shortestToChar(s, c)
    print(ans)
    # [3,2,1,0,1,0,0,1,2,2,1,0]

