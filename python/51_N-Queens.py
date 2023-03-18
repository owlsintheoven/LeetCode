class Solution(object):
    # queens cannot attack each other if they are not in the same row / column / diagonal
    # should be recursive
    answer = []
    size = 0
    def solveNQueens(self, n):
        self.answer = []
        self.size = n
        m = [value for value in range(n)]
        self.helper(m, 0, [])
        #for solution in self.answer:
        #    solution.sort()
        #initial_place = ['.' for i in range(n-1)]
        #initial_place.insert(0, 'Q')
        #place_dict = {}
        #for i in range(n):
        #    place = ''
        #    for left in range(i):
        #        place += '.'
        #    place += 'Q'
        #    for right in range(i+1, n):
        #        place += '.'
        #    place_dict[i] = place
        #answer_transformed = []
        #for solution in self.answer:
        #    transformed = []
        #    for num in solution:
        #        transformed.append(place_dict[num%n])
        #    answer_transformed.append(transformed)

        #return answer_transformed
        return len(self.answer)

    def helper(self, m, start_c, placed):
        if len(placed) != 0:
            if start_c == self.size:
                self.answer.append(list(placed))
        for j in range(len(m)):
            newly_placed = start_c + m[j] * self.size
            valid = True
            for prev_place in placed:
                if (prev_place > newly_placed and (newly_placed - prev_place) % (self.size - 1) == 0) \
                    or (prev_place < newly_placed and (newly_placed - prev_place) % (self.size + 1) == 0):
                    valid = False
                    break
            if valid:
                new_m = m[:j]
                new_m.extend(m[j + 1:])
                placed.append(newly_placed)
                self.helper(new_m, start_c + 1, placed)
                placed.pop(-1)


if __name__ == '__main__':
    solution = Solution()
    n = 4
    ans = solution.solveNQueens(n)
    print(ans)