function dayOfYear(date: string): number {
    let months = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    let dates = date.split('-')
    let year = Number(dates[0])
    let month = Number(dates[1])
    let day = Number(dates[2])

    if(year%4 == 0) {
        months[1] = 29
    }
    if(year === 1800 || year === 1900) {
        months[1] = 28
    }
    
    let result = day;
    for(let i = 0; i<month-1; i++) {
        result += months[i]
    }
    return result

};