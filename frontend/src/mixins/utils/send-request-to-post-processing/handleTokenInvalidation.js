export const HandleTokenInvalidationMixin = {
    methods: {
        handleTokenInvalidation() {
        console.log("handleTokenInvalidation")
        this.$store.commit("authenticate/setAuthenticated", false);
        this.$router.push('/login');
        },
    }
};