const board = [
    ["1", "2", ".", ".", "3", ".", ".", ".", "."],
    ["4", ".", ".", "5", ".", ".", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", ".", "3"],
    ["5", ".", ".", ".", "6", ".", ".", ".", "4"],
    [".", ".", ".", "8", ".", "3", ".", ".", "5"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", ".", ".", ".", ".", ".", "2", ".", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "8"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
];
const board2 = [
    ["1", "2", ".", ".", "3", ".", ".", ".", "."],
    ["4", ".", ".", "5", ".", ".", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", ".", "3"],
    ["5", ".", ".", ".", "6", ".", ".", ".", "4"],
    [".", ".", ".", "8", ".", "3", "2", ".", "5"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", ".", ".", ".", ".", ".", "2", ".", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "8"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
];
function isValidSudoku(board) {
    const rows = new Set();
    const cols = new Set();
    const boxes = new Set();
    for (let r = 0; r < 9; r++) {
        for (let c = 0; c < 9; c++) {
            const val = board[r][c];
            if (val === ".")
                continue;
            const rowKey = `r${r}-${val}`;
            const colKey = `c${c}-${val}`;
            const box = Math.floor(r / 3) * 3 + Math.floor(c / 3);
            const boxKey = `b${box}-${val}`;
            if (rows.has(rowKey) || cols.has(colKey) || boxes.has(boxKey)) {
                return false;
            }
            rows.add(rowKey);
            cols.add(colKey);
            boxes.add(boxKey);
        }
        console.log(r, "rows -- ", rows);
        console.log(r, "cols -- ", cols);
        console.log(r, "boxes -- ", boxes);
    }
    return true;
}
console.log(isValidSudoku(board2));
