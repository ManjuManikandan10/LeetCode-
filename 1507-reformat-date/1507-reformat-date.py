class Solution:
    def reformatDate(self, date: str) -> str:
        [day, month, year] = date.split(' ')

        month_map = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
        day = day[:-2]

        day = day.zfill(2)

        month = str(month_map.index(month) + 1).zfill(2)

        return f"{year}-{month}-{day}"

