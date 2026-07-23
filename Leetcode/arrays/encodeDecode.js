const input = ["sunil", "parajuli"];
function encode(input) {
    let cypher = "";
    for (const n of input) {
        cypher += `${n.length}#${n}`;
    }
    return cypher;
}
function decode(cypher) {
    const output = [];
    let i = 0;
    while (i < cypher.length) {
        let j = i;
        while (cypher[j] !== "#") {
            j++;
        }
        const len = Number(cypher.slice(i, j));
        const word = cypher.slice(j + 1, j + 1 + len);
        output.push(word);
        i = j + 1 + len;
    }
    return output;
}
function encodeDecode(input) {
    return [];
}
console.log("final : ", decode(encode(input)));
