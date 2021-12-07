class Solution:
    def magicTower(self, nums: List[int]) -> int:
        if sum(nums) < 0:
            return -1
        count = 0
        blood = 1
        n = len(nums)

        monster = []
        for i in nums:
            if i < 0:
                monster.append(i)
            blood += i
            while(blood < 1):
                count += 1
                blood -= min(monster)
                monster.remove(min(monster))
        return count

