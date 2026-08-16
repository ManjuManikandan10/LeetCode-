class Solution:
    def reversePairs(self, nums):
        def merge_sort(start, end):
            if end - start <= 1:
                return 0
            mid = (start + end) // 2
            count = merge_sort(start, mid) + merge_sort(mid, end)
            j = mid
            for i in range(start, mid):
                while j < end and nums[i] > 2 * nums[j]:
                    j += 1
                count += j - mid

            # merge two sorted halves
            temp = []
            l, r = start, mid
            while l < mid and r < end:
                if nums[l] <= nums[r]:
                    temp.append(nums[l])
                    l += 1
                else:
                    temp.append(nums[r])
                    r += 1
            while l < mid:
                temp.append(nums[l])
                l += 1
            while r < end:
                temp.append(nums[r])
                r += 1
            nums[start:end] = temp
            return count

        return merge_sort(0, len(nums))
