<template>
  <input-column  :english_name='english_name'
                 :chinese_name='chinese_name'
                 :rules="rules"
                 :values="values"
                  :disabled="disabled"
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
    values: {
       type: String,
       default: null,
    },
    disabled: {
          type: Boolean,
          default: false,
    },
  },
  data() {
    return {
        english_name: 'hardware_name',
        chinese_name: '硬體名稱',
    }
  },
};
</script>
