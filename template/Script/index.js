$(function () {

    MainVue();
});


function MainVue() {
    new Vue({
        el: '#app-1',
        data: function () {
            return { visible: false }
        }
    });

    new Vue({
        el: '#app-2',
        data: {
            input: '',
            password: ''
        }
    });
}
