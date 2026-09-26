function evaluate(s: string, k: string[][]): string {
    const map = new Map<string, string>();

    for (const pair of k) {
        map.set(pair[0], pair[1]);
    }

    let res: string = '';

    for (let i = 0; i < s.length; i++) {
        if (s[i] === '(') {
            const ind = s.indexOf(')', i + 1);
            const t = s.slice(i + 1, ind);

            res += map.get(t) ?? '?';
            i = ind;
        } else {
            res += s[i];
        }
    }

    return res;
}