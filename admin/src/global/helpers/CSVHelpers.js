

export const convertCSVToArray = (csv) => {
    const rows = csv.split("\n");
    if (rows.length < 1) return [];

    return rows.map((r) => r.split(",").map((c) => c.trim()));
}

export const convertArrayToCSV = (arr) => {
    const headers = arr[0].join(",");
    const body = arr.slice(1)
    let columns = [];

    for (let i = 0; i < body.length; i++) {
        let col = [];
        for (let j = 0; j < body[i].length; j++) {
            col.push(body[j][i]);
        }
        columns.push(col.join(","));
    }
    console.log(headers, columns);
    return headers + "\n" + columns.join("\n");
}