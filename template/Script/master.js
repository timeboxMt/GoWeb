$(function () {

});


//ajax Get请求
function ajaxGet(url, data) {
    $.ajax({
        type: "get",
        url: url,
        data: data,
        dataType: "json",
        success: function (data) {
            return data;
        }
    });
}

//ajax Poast请求
function ajaxPost(url, data) {
    $.ajax({
        type: "post",
        url: url,
        data: data,
        dataType: "json",
        success: function (data) {
            return data;
        }
    });
}


///验证函数

//是否是手机号
function isPhone(data) {

}

//是否是邮箱
function isEmail(data) {

}