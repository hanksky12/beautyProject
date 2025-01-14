<template>
    <div>
    <bread-crumb :props_items="items"/>
    <login-form />
    </div>

</template>

<script>
import LoginForm from "@/views/user/login/components/LoginForm.vue";
import Breadcrumb from '@/components/Breadcrumb.vue'
export default {
    components: {
        'login-form': LoginForm,
        'bread-crumb':Breadcrumb,
      },
      inject: ['showToast'],
        data() {
      return {
        items: [
          {
            text: '传图管理后台',
             active: true
          },
          {
            text: '帐号',
            active: true
          },
          {
            text: '登入',
            active: true
          }
        ]
      }
    }
}
</script>

