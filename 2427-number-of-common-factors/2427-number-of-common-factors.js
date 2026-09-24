/**
 * @param {number} a
 * @param {number} b
 * @return {number}
 */
var commonFactors = function(a, b) {

    // we ensure a is the larger number
    if (b > a) [a, b] = [b, a];

    const gcd = (x, y) => y === 0 ? x : gcd(y, x % y);
    const GCD = gcd(a, b);

    let count = 0;
    for (let i = 1; i <= GCD; i++) {
        if (GCD % i === 0) count++;

    }

    return count;
};