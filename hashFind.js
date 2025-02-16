for(let j = 0; j < 10; j++){
    for (let i = 0; ; i++){
        let hash = await sha256(`${playerName}_${validHashCount}_${i}`);
        if (hash.startsWith("000")) {
            console.log(i);
            clickCountHash = i - 1
            clickCount += clickCountHash
            clickButton()
            break;
        }
    }
    }
    