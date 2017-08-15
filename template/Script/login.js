$(function () {
    main();
});

//初始化执行js函数
function main() {
    new Vue().$mount('#app')

    $("#btn_login").click(function () {
        $.ajax({
            type: "post",
            url: "http://127.0.0.1:9090/login",
            data: "",
            dataType: "json",
            success: function (data) {
                location.href = data;
            }
        });
    });
}