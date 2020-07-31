/**
 * js 即用工具
 * 
 */

//@todo 需要修正，如错误解析数据
//test case: 
//    生于公元一九五六年十一月二十四日。
//    生于公元一九七八年十二月二十二日。
//    生于二0七年六月十日
//    生于公元一九九四年
//    生于公元一九八四年五月
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
        ['01', ['初一', '正月']],
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
        ['12', ['十二', '腊月']],
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


// 测试用例:  
//      turnNumByText('三十一和十五日（亦有成为一十五的），二十八日，三十岁')
//      turnNumByText('一十五日（亦有为十五的）；三十一和十五日（亦有成为一十五的），二十八日，三十岁')
/**
 * 将文本中的中文数字转阿拉伯数字
 * @param {string} text 
 */
function turnNumByText(text){
    const clearData = [
        ['cS', ['十']],
        ['cB', ['百']],
        ['cQ', ['千']],
        ['cW', ['万']],
        [1, ['一', '壹']], 
        [2, ['二', '贰']], 
        [3, ['三', '叁']], 
        [4, ['四', '肆']], 
        [5, ['五', '伍']], 
        [6, ['六', '陆']], 
        [7, ['七', '柒']], 
        [8, ['八', '捌']], 
        [9, ['九', '玖']], 
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

    // 中文转数字
    let matchQue = text.match(/[0-9]*c[SBQW][0-9]*/g);
    if (matchQue){
        matchQue.forEach((ms) => {
            let keyword = "cS", tArr = [], tArrLen = 0, sep = '.', oMs = ms;
            if(ms.indexOf(keyword) > -1){
                tArr = ms.replace(keyword, sep).split(sep);
                tArrLen = tArr.length;
                tArr.forEach((s, idx) => {
                    if(s.trim() === ''){
                        tArr[idx] = idx === 0? '1': '0';
                    }
                });
                ms = tArr.join('');
            }
            text = text.replace((new RegExp(oMs, 'g')), ms);

            // @todo 百千暂时未能实现
            // let keyword = "cB", tArr = [], tArrLen = 0, sep = '.', oMs = ms;
            // if(ms.indexOf(keyword) > -1){
            //     tArr = ms.replace(keyword, sep).split(sep);
            //     tArrLen = tArr.length;
            //     if (tArrLen === 2){
            //         ms = ms.replace(keyword, '0');
            //     }else{
            //         ms = ms.replace(keyword, '');
            //     }
            // }
            // text = text.replace((new RegExp(oMs, 'g')), ms);
        });
    }

    return text;
}


//@todo 需要修正，如错误解析数据
//test case: 
//    生于公元一九五六年十一月二十四日。
//    生于公元一九七八年十二月二十二日。
//    生于二0七年六月十日
//    生于公元一九九四年
//    生于公元一九八四年五月
//    中文单位匹配: {G 个, S 十, B 百, Q 千, W 万}
/**
 * 日期数据提取通过数字(年月日)。文本转标准数日日期: 2006-04-05.
 * 年月日缺省时补01
 * @param {string} text 
 * @return {string}
 */ 
function getDateByTextX(text){
    //中文替换英文
    text = text.replace(/\s/g, '');
    text = turnNumByText(text);
    //dickMap/Object `for in` key 会进行排序，遂换用数组。
    const clearData = [
        [0, ['零', '〇']], 
        ['01', ['初一', '正月']],
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
        ['12', ['十二', '腊月']],
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


// --------------------------------------------------------------------------(GZ/begin)------------------------------------------------------
// 干支运算
const GanOrder = ['甲', '乙', '丙', '丁', '戊', '己', '庚', '辛', '壬', '癸'];
const ZhiOrder = ['子', '丑', '寅', '卯', '辰', '巳', '午', '未', '申', '酉', '戌', '亥'];
class GanZhi{
    #_gzList = [];  //干支缓存
    #_gOrder = [];
    #_zOrder = [];
    #_zodiac = [];
    #_ganZodiacMap = {};        //干支生肖字典
    constructor(){
        // 实际的GZ顺序
        this.#_gOrder = ['庚', '辛', '壬', '癸', '甲', '乙', '丙', '丁', '戊', '己'];
        this.#_zOrder = ['申', '酉', '戌', '亥', '子', '丑', '寅', '卯', '辰', '巳', '午', '未'];
        // 子-鼠，丑-牛，寅-虎，卯-兔，辰-龙，巳-蛇， 午-马，未-羊，申-猴，酉-鸡，戌-狗，亥-猪。
        this.#_zodiac = ['鼠', '牛', '虎', '兔', '龙', '蛇', '马', '羊', '猴', '鸡', '狗', '猪'];

        // 字典生成
        ZhiOrder.forEach((v, idx) => {
            this.#_ganZodiacMap[v] = this.#_zodiac[idx];
        });
    }
    /**
     * 干支列表
     * @returns {array}
     */
    get gzList(){
        return this.getGZList();
    }

    /**
     * 干列表
     * @returns {array}
     */
    get gOrder(){
        return this.#_gOrder;
    }

    /**
     * 支列表
     * @returns {array}
     */
    get zOrder(){
        return this.#_zOrder;
    }

    /**
     * 生肖列表
     * @returns {array}
     */
    get zodiacOrder(){
        return this.#_zodiac;
    }
    /**
     * 获取干支表
     * @return {array}
     */
    getGZList(){
        if(this.#_gzList.length == 0){
            let que = [], max = 60;
            let order = 0;
            //单配单，双配双
            while(order < max){
                let gIdx = order%10,
                    zIdx = order%12
                ;

                let gz = GanOrder[gIdx] + ZhiOrder[zIdx];
                que.push(gz);
                order = que.length;
            }
            this.#_gzList = que;
        }
        return this.#_gzList;
    }

    /**
     * 获取年的干支信息
     * @param {int} year 
     */
    getYearGz(year){
        if (!year){
            year = (new Date()).getFullYear();
        }
        let gWord, zWord;
        gWord = year%10;
        zWord = year%12;

        return `${this.#_gOrder[gWord]}${this.#_zOrder[zWord]}`;
    }

    /**
     * 获取年的干支信息
     * @param {int} year 
     */
    zodiac(year){
        if (!year){
            year = (new Date()).getFullYear();
        }
        let zWord;
        zWord = year%12;
        let zhi = this.#_zOrder[zWord];

        return this.#_ganZodiacMap[zhi];
    }
}

// --------------------------------------------------------------------------(GZ/end)------------------------------------------------------


