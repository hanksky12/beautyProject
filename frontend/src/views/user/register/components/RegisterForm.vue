<template>
    <common-form title="注册" @handleFormSubmit="handleFormSubmit" />
</template>

<script>

import CommonForm from '@/views/user/common-components/CommonForm.vue'
import { createNamespacedHelpers } from 'vuex'
import {RegisterApiMixin} from "@/mixins/api/user/register";
const { mapGetters:mapGetterForm } = createNamespacedHelpers('form')

export default {
      components: {
      'common-form':CommonForm
  },
    mixins: [RegisterApiMixin],
    methods: {
      async handleFormSubmit() {
          let responseData = await this.registerApi(this.getFormData)
          if(responseData.code === 200){
            console.log('註冊成功')
            this.$router.push('/login')
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