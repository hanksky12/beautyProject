<template>
    <common-form title="登入" @handleFormSubmit="handleFormSubmit" />
</template>

<script>
import CommonForm from '@/views/user/common-components/CommonForm.vue'
import {LogInApiMixin} from "@/mixins/api/user/logIn";
import {GetCookieMixin} from "@/mixins/utils/handle-request/create-request/getCookie";
import { createNamespacedHelpers } from 'vuex'
const { mapGetters:mapGetterForm } = createNamespacedHelpers('form')

export default {
      components: {
      'common-form':CommonForm
  },
    mixins: [LogInApiMixin,GetCookieMixin],
    methods: {
      async handleFormSubmit() {
          const responseData = await this.logInApi(this.getFormData)
          if(responseData.code === 200){
            console.log('登入成功')
            this.$store.commit("authenticate/setAuthenticated", true);
            const authorizationLevel = this.getCookie("authorization_level")
            this.$store.commit("authorization/setAuthorization", authorizationLevel);
            this.$router.push('/')
          }
      },
    },
    computed:{
      ...mapGetterForm(['getFormData']),
    }
}
</script>

<style scoped>

</style>