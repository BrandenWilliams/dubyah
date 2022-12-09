const { createApp } = Vue

const cookie = document.cookie
const cookieSearch = cookie.search("jump_key")
const app = Vue.createApp({

    data() {
        return {
            isLoggedIn: Boolean,
            logout: `/api/logout`
        }
    },
    methods: {
        setLoginStatus() {
            switch (cookieSearch) {
                case 0:
                    return true
                default:
                    return false
            }
        }
    },
    mounted() {
        this.isLoggedIn = this.setLoginStatus()
    },
}).mount('#loginCheck');
