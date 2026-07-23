const input = ["sunil", "parajuli"];

function encode(input: string[]): string {
  let cypher = "";

  for (const n of input) {
    cypher += `${n.length}#${n}`;
  }

  return cypher;
}

function decode(cypher: string): string[] {
  const output: string[] = [];

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

function encodeDecode(input: string[]): string[] {
  return decode(encode(input));
}

console.log("final : ", decode(encode(input)));
