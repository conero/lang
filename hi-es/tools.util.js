/**
 * js 即用工具
 * 
 */

 /**
 * 字符串清洗
 * @param {string} str 
 * @param {array} cleanList 
 */
function cleanString(str, cleanList){
    if ('object' !== typeof cleanList){
        return str
    }
    cleanList.forEach((vque) => {
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
            str = str.replace(new RegExp(ts, 'g'), key)
        });
    });

    return str
}

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
        ['二十', ['廿']],
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

    text = cleanString(text, clearData);

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
            //@notice 使用`g`标识，可能会替换很多次
            //text = text.replace((new RegExp(oMs, 'g')), ms);
            text = text.replace((new RegExp(oMs)), ms);

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

// 测试错误
// 生公元一九七0年十月二十九日。徙居天柱城。
// 生公元一九九四年正月十三日。
// 
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
    text = cleanString(text, [
        [0, ['零', '〇', 'O', '0', 'O']], 
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
    ]);
    text = turnNumByText(text);
    //dickMap/Object `for in` key 会进行排序，遂换用数组。
    const clearData = [
        ['12-', ['腊月']],
        ['01-', ['正月']],
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
const GzSpanceYear = 60;        //干支间距年
const ZodiacSpanceYear = 12;
class GanZhi{
    __gzList = [];  //干支缓存
    __gOrder = [];
    __zOrder = [];
    __zodiac = [];
    __ganZodiacMap = {};        //干支生肖字典
    constructor(){
        // 实际的GZ顺序
        this.__gOrder = ['庚', '辛', '壬', '癸', '甲', '乙', '丙', '丁', '戊', '己'];
        this.__zOrder = ['申', '酉', '戌', '亥', '子', '丑', '寅', '卯', '辰', '巳', '午', '未'];
        // 子-鼠，丑-牛，寅-虎，卯-兔，辰-龙，巳-蛇， 午-马，未-羊，申-猴，酉-鸡，戌-狗，亥-猪。
        this.__zodiac = ['鼠', '牛', '虎', '兔', '龙', '蛇', '马', '羊', '猴', '鸡', '狗', '猪'];

        // 字典生成
        ZhiOrder.forEach((v, idx) => {
            this.__ganZodiacMap[v] = this.__zodiac[idx];
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
        return this.__gOrder;
    }

    /**
     * 支列表
     * @returns {array}
     */
    get zOrder(){
        return this.__zOrder;
    }

    /**
     * 生肖列表
     * @returns {array}
     */
    get zodiacOrder(){
        return this.__zodiac;
    }
    /**
     * 获取干支表
     * @return {array}
     */
    getGZList(){
        if(this.__gzList.length == 0){
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
            this.__gzList = que;
        }
        return this.__gzList;
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

        return `${this.__gOrder[gWord]}${this.__zOrder[zWord]}`;
    }

    /**
     * 通过干支获取对应的年份列表
     * @param {string|number} gz 年份或者甲子
     * @param {number} benchmarkYear 基准年（小于等于它的数据列表），默认当年
     * @param {number} num 生成数量，默认: 20
     */
    getGzYearList(gz, benchmarkYear, num){
        gz = gz? ('number' === typeof gz? this.getYearGz(gz): gz) : gz;
        gz = gz? gz: this.getYearGz();
        let gzList = this.getGZList();
        let checkGzIdx = gzList.indexOf(gz);

        if (checkGzIdx === -1){
            throw new Error(`${gz} 不是一个有效的干支名称！`);
            return;
        }
        let gzArr = gz.split('');
        if(gzArr && gzArr.length === 2){
            let gIdx = this.__gOrder.indexOf(gzArr[0]), 
                zIdx = this.__zOrder.indexOf(gzArr[1])
            ;
            if (gIdx > -1 && zIdx > -1){
                benchmarkYear = isNaN(benchmarkYear)? 0 : parseInt(benchmarkYear);
                benchmarkYear = benchmarkYear < 1?  (new Date()).getFullYear(): benchmarkYear;

                let stdGidx = benchmarkYear%10;
                let stdZidx = benchmarkYear%12;
                
                // 如果基准与需要的干支年一样，则其作为基准年
                let stdYear = 0;    //离基准最近的年份
                if (stdGidx === gIdx && stdZidx === zIdx){
                    stdYear = benchmarkYear;
                }else{
                    // 通过干支表求出idex索引，再有索引比对得到基准年份
                    let curGzIdx = gzList.indexOf(this.getYearGz(benchmarkYear));
                    if(curGzIdx > checkGzIdx){
                        stdYear = benchmarkYear-(curGzIdx-checkGzIdx);
                    }else{
                        stdYear = benchmarkYear+(checkGzIdx-curGzIdx)-GzSpanceYear;
                    }
                    
                }

                num = num? num: 20;

                if (stdYear > 0){
                    let nC = 0;
                    let yearList = [stdYear];
                    num -= 1;
                    while(nC < num){
                        stdYear -= GzSpanceYear;
                        yearList.push(stdYear);
                        nC += 1;
                    }

                    return yearList;
                }

            }
        }
        return null;
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
        let zhi = this.__zOrder[zWord];

        return this.__ganZodiacMap[zhi];
    }

     /**
     * 通过干支获取对应的年份列表
     * @param {string|number} zodiac 年份或者甲子
     * @param {number} benchmarkYear 基准年（小于等于它的数据列表），默认当年
     * @param {number} num 生成数量，默认: 20
     */
    getZodiacList(zodiac, benchmarkYear, num){
        zodiac = zodiac? ('number' === typeof zodiac? this.zodiac(zodiac): zodiac) : zodiac;
        zodiac = zodiac? zodiac: this.zodiac();
        let idx = this.__zodiac.indexOf(zodiac);

        if (idx === -1){
            throw new Error(`${zodiac} 不是一个有效的生肖名称！`);
            return;
        }
        
        benchmarkYear = isNaN(benchmarkYear)? 0 : parseInt(benchmarkYear);
        benchmarkYear = benchmarkYear < 1?  (new Date()).getFullYear(): benchmarkYear;

        let stdIdx = this.__zodiac.indexOf(this.zodiac(benchmarkYear));
        // 如果基准与需要的干支年一样，则其作为基准年
        let stdYear = 0;    //离基准最近的年份
        if(idx <= stdIdx){
            stdYear = benchmarkYear - (stdIdx - idx);
        }else{
            stdYear = benchmarkYear + (idx-stdIdx) - ZodiacSpanceYear;
        }

        num = num? num: 20;

        if (stdYear > 0){
            let nC = 0;
            let yearList = [stdYear];
            num -= 1;
            while(nC < num){
                stdYear -= ZodiacSpanceYear;
                yearList.push(stdYear);
                nC += 1;
            }

            return yearList;
        }
        return null;
    }

}

// --------------------------------------------------------------------------(GZ/end)------------------------------------------------------


