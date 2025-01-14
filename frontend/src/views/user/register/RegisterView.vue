<template>
    <div>
      <bread-crumb :props_items="items"/>
      <register-form />
    </div>
</template>

<script>
import RegisterForm from "@/views/user/register/components/RegisterForm.vue";
import Breadcrumb from '@/components/Breadcrumb.vue'

export default {
    components: {
        'bread-crumb':Breadcrumb,
        'register-form': RegisterForm
      },
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
                    text: '注册',
                    active: true
                }
            ]
        }
    }
}
</script>

