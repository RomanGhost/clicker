async function sha256(message) {
    const encoder = new TextEncoder();
    const data = encoder.encode(message);
    const hashBuffer = await crypto.subtle.digest("SHA-256", data);
    return Array.from(new Uint8Array(hashBuffer))
        .map(b => b.toString(16).padStart(2, '0'))
        .join('');
}

let playerName = "Roman"
let validHashCount = 0
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
    