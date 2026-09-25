function dayOfTheWeek(day: number, month: number, year: number): string {
    const daysOfWeek = ["Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"];
    const daysInMonth = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];

    const isLeapYear = (y: number): boolean => {
        return (y % 4 === 0 && y % 100 !== 0) || (y % 400 === 0);
    };

    let totalDays = 0;

    for (let y = 1971; y < year; y++) {
        totalDays += isLeapYear(y) ? 366 : 365;
    }

    for (let m = 0; m < month - 1; m++) {
        totalDays += daysInMonth[m];
        if (m === 1 && isLeapYear(year)) {
            totalDays += 1;
        }
    }

    totalDays += day - 1;

    return daysOfWeek[totalDays % 7];
}