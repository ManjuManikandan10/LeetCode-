function numPrimeArrangements(n: number): number {
    const MOD = 1_000_000_007n;

    const isPrime = (num: number): boolean => {
        if (num < 2) return false;
        for (let i = 2; i * i <= num; i++) {
            if (num % i === 0) return false;
        }
        return true;
    };

    let primeCount = 0;
    for (let i = 1; i <= n; i++) {
        if (isPrime(i)) {
            primeCount++;
        }
    }

    let nonPrimeCount = n - primeCount;

    const getFactorialMod = (count: number): bigint => {
        let result = 1n;
        for (let i = 2; i <= count; i++) {
            result = (result * BigInt(i)) % MOD;
        }
        return result;
    };

    const primeFactorial = getFactorialMod(primeCount);
    const nonPrimeFactorial = getFactorialMod(nonPrimeCount);

    return Number((primeFactorial * nonPrimeFactorial) % MOD);
}