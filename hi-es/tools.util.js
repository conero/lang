/**
 * js 即用工具
 * 
 */



/**
 * 日期数据提取通过数字(年月日)。文本转标准数日日期: 2006-04-05.
 * 年月日缺省时补01
 * @param {string} text 
 * @return {string}
 */ 
function getDateByText(text){
    //中文替换英文
    text = text.replace(/\s/g, '');
    //dickMap/Object `for in` key 会进行排序，遂换用数组。
    const clearData = [
        [0, ['零', '〇']], 
        ['01', ['初一']],
        ['02', ['初二']],
        ['03', ['初三']],
        ['04', ['初四']],
        ['05', ['初五']],
        ['06', ['初六']],
        ['07', ['初七']],
        ['08', ['初八']],
        ['09', ['初九']],
        ['10', ['初十']],
        ['11', ['十一']],
        ['12', ['十二']],
        ['13', ['十三']],
        ['14', ['十四']],
        ['15', ['十五']],
        ['16', ['十六']],
        ['17', ['十七']],
        ['18', ['十八']],
        ['19', ['十九']],
        [1, ['一', '壹']], 
        [2, ['二', '贰']], 
        [3, ['三', '叁']], 
        [4, ['四', '肆']], 
        [5, ['五', '伍']], 
        [6, ['六', '陆']], 
        [7, ['七', '柒']], 
        [8, ['八', '捌']], 
        [9, ['九', '玖']], 
        ['-', ['年', '月']],
        ['', ['十']]
    ];
    clearData.forEach((vque) => {
        if('object' !== typeof vque){
            return;
        }
        let key = vque.shift();
        var v = vque[0];
        if(!v){
            return;
        }
        
        if('object' !== typeof v){
            v = `${v}`;
            v = v.trim();
            if(v === ''){
                return;
            }
            v = [v];
        }
        v.forEach(ts => {
            ts = ts.trim();
            if(ts === ''){
                return;
            }
            text = text.replace(new RegExp(ts, 'g'), key)
        });
    });

    let checkDataFormat = (_ts) => {
        if(!_ts){
            return _ts;
        }
        let que = _ts.split('-')
        let len = que.length;
        if(len > 1 && que[1].length == 1){
            que[1] = `0${que[1]}`;
        }
        if(len > 2 && que[2].length == 1){
            que[2] = `0${que[2]}`;
        }
        return que.join('-')
    };

    //日期正则匹配
    text = text.replace(/\s/g, '')
    //真正匹配
    let reg = /[0-9]{4}-[0-9]{1,2}-[0-9]{1,2}/;
    let array = text.match(reg);
    let vdate, vdateQue, queLen;
    if(array && array.length > 0){
        return checkDataFormat(array[0]);
    }

    // 模糊匹配
    reg = /[0-9-]{1,}/;
    array = text.match(reg);
    if(array && array.length > 0){
        vdate = array[0];
        vdateQue = vdate.split('-')
        queLen = vdateQue.length;
        if(queLen < 4 && queLen > 0){
            if(queLen == 1){
                vdateQue = vdateQue.concat(['01', '01']);
            }else if(queLen == 2){
                vdateQue.push('01');
            }
            return checkDataFormat(vdateQue.join('-'));
        }
    }

    return '';
}