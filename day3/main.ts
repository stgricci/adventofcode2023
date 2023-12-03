const fs = require('fs');
const board = fs.readFileSync('data.txt', 'utf-8').split('\n');
const chars = new Map();

for (let r = 0; r < 140; r++) {
    for (let c = 0; c < 140; c++) {
        if (!board[r][c].match(/[0-9.]/)) {
            chars.set(`${r},${c}`, []);
        }
    }
}

for (let r = 0; r < board.length; r++) {
    const row = board[r];
    const regex = /\d+/g;
    let match;
    while ((match = regex.exec(row)) !== null) {
        const n = parseInt(match[0]);
        const edge = new Set();
        for (let i = r - 1; i <= r + 1; i++) {
            for (let j = match.index - 1; j <= match.index + match[0].length; j++) {
                edge.add(`${i},${j}`);
            }
        }
        for (const o of edge) {
            if (chars.has(o)) {
                chars.get(o).push(n);
            }
        }
    }
}

let sum1 = 0;
let sum2 = 0;
for (const p of chars.values()) {
    sum1 += p.reduce((a, b) => a + b, 0);
    if (p.length === 2) {
        sum2 += p.reduce((a, b) => a * b, 1);
    }
}

console.log(sum1, sum2);
