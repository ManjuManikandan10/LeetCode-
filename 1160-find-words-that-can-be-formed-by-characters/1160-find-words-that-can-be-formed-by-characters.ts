function countCharacters(words: string[], chars: string): number {
    return words.reduce((sum, word) => {
        const charsArray = chars.split("");
        let isValid = false;
        for(let i=0; i<word.length; i++) {
            const index = charsArray.indexOf(word.charAt(i));
            if(index !== -1) {
                charsArray.splice(index, 1);
                isValid = true;
            } else {
                isValid = false;
                break;
            }
        }
        sum = isValid ? sum += word.length : sum;
        return sum;
    }, 0);
};