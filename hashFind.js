for (let i = 0; i < 10000; i++){
    let hash = await sha256(`${playerName}_${validHashCount}_${i}`);
    if (hash.startsWith("000")) {
        console.log(i);
        clickCountHash = i - 1
        clickCount += clickCountHash
        break;
    }
}