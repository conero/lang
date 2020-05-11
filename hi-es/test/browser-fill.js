// 浏览器，快速填入测试数据
class __test_jc{
    static Author = 'Joshua Conero';
}


class __test extends __test_jc{
    /**
     * 基于jquery
     * @param {*} pdom 
     */
    static jq(pdom){
        if(pdom && 'object' !== typeof pdom){
            pdom = $(pdom);
        }
        if(!pdom){
            pdom = $(document);
        }
        // input
        let inputs = pdom.find('input');
        for(let i=0; i<inputs.length; i++){
            let input = $(inputs[i]),
                vtype = input.attr('type')
            ;

            if(vtype == 'hidden' || vtype == 'radio' || vtype == 'checkbox'){
                continue;
            }

            if(input.val().trim() == ''){
                input.val(1);
            }
        }

        // textarea
        let textareas = pdom.find('textarea');
        for(let i=0; i<textareas.length; i++){
            let textarea = $(textareas[i])
            ;

            if(textarea.val().trim() == ''){
                textarea.val(1);
            }
        }
    }
}