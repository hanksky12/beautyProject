<template>
  <input-column  :english_name='english_name'
                 :chinese_name='chinese_name'
                  :rules="rules"
  />
</template>
<script>
import InputColumn from "@/components/form/column/base/InputColumn.vue";


export default {
  components: {
      'input-column': InputColumn
  },
  props:{
    rules: {
      type: String,
      default: ""
    },
  },
  data() {
    return {
        english_name: 'user_password',
        chinese_name: '密码',
    }
  },
};
</script>
