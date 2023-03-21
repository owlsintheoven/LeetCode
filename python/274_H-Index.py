class Solution(object):
   # A scientist has an index h if h of their n papers have at least h citations each, and the other n − h papers have no more than h citations eachdef hIndex(self, citations):
   # def hIndex(self, citations):
   #     citations.sort(reverse=True)
   #     found = False
   #     for h in range(len(citations) - 1, -1, -1):
   #         for i in range(len(citations)):
   #             citation = citations[i]
   #             if i < h and citation < h:
   #                 break
   #             elif i == h and citation <= h:
   #                 found = True
   #                 break
   #         if found:
   #             break
   #     return h if found else len(citations)


    def hIndex(self, citations):
        buckets = [0 for i in range((len(citations) + 1))]
        for citation in citations:
            if citation > len(citations):
                buckets[-1] += 1
            else:
                buckets[citation] += 1
        print(buckets)
        count = 0
        for i in range(len(buckets) - 1, -1, -1):
            count += buckets[i]
            if count >= i:
                return i



if __name__ == '__main__':
    solution = Solution()
#    citations = [0]
#    citations = [1]
    citations = [3, 0, 3, 1, 5]
#    citations = [1, 3, 1]
#    citations = [11, 15]
    ans = solution.hIndex(citations)
    print(ans)

