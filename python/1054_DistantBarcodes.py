class Solution(object):
    def rearrangeBarcodes(self, barcodes):
        value_count_dict = {}
        for barcode in barcodes:
            if barcode not in value_count_dict:
                value_count_dict[barcode] = 1
            else:
                value_count_dict[barcode] += 1
        tuple_list = []
        for value, count in value_count_dict.items():
            tuple_list.append((value, count))
        tuple_list.sort(key=lambda a: a[1], reverse=True)
        ref = [tuple[0] for tuple in tuple_list for _ in range(tuple[1])]
        new_barcodes = []
        for i in range(0, len(barcodes), 2):
            new_barcodes.insert(i, ref.pop(0))
        for i in range(1, len(barcodes), 2):
            new_barcodes.insert(i, ref.pop(0))
        return new_barcodes


if __name__ == '__main__':
    solution = Solution()
    barcodes = [1, 1, 1, 2, 2, 2]
#    barcodes = [1, 1, 1, 1, 2, 2, 3, 3]
#    barcodes = [2, 1, 1]
#    barcodes = [7,7,7,8,5,7,5,5,5,8]
#    barcodes = [51,83,51,40,51,40,51,40,83,40,83,83,51,40,40,51,51,51,40,40,40,83,51,51,40,51,51,40,40,51,51,40,51,51,51,40,83,40,40,83,51,51,51,40,40,40,51,51,83,83,40,51,51,40,40,40,51,40,83,40,83,40,83,40,51,51,40,51,51,51,51,40,51,83,51,83,51,51,40,51,40,51,40,51,40,40,51,51,51,40,51,83,51,51,51,40,51,51,40,40]
    ans = solution.rearrangeBarcodes(barcodes)
    print(ans)