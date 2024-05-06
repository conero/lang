// 数字压缩
function zipNumber(num) {
    // 数字压缩: K 千, M 百万, G 十亿， T
    let dick = { 3: 'K', 6: 'M', 9: 'G', 12: 'T' };
    let lg10 = Math.log10(num);
    let x = Math.floor(lg10);
    let unit = dick[x] || null;
    if (!unit) {
        let lastX = 0;
        for (var n in dick) {
            if (n > x) {
                unit = dick[lastX];
                x = lastX;
                break
            }
            lastX = n;
        }
    }

    let xNum = num / (Math.pow(10, x));
    let intStr = parseInt(xNum);
    intStr = xNum > intStr ? `${intStr}+` : intStr;
    let data = {
        number: xNum,
        intStr,
        unit,
        _n: x,
    }
    return data;
}

// 数字压缩显示中文
function zipNumberZhCn(num) {
    // 数字压缩: K 千, M 百万, G 十亿， T
    let dick = { 1: '十', 2: '百', 3: '千', 4: '万', 8: '亿', 24: '兆', 16: '京' };
    let lg10 = Math.log10(num);
    let x = Math.floor(lg10);
    let unit = dick[x] || null;
    if (!unit) {
        let lastX = 0;
        for (var n in dick) {
            if (n > x) {
                unit = dick[lastX];
                x = lastX;
                break
            }
            lastX = n;
        }
    }

    let xNum = num / (Math.pow(10, x));
    let intStr = parseInt(xNum);
    intStr = xNum > intStr ? `${intStr}+` : intStr;
    let data = {
        number: xNum,
        intStr,
        unit,
        _n: x,
    }
    return data;
}

