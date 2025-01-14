<template>
  <select-column  :english_name='english_name'
                 :chinese_name='chinese_name'
                 :options_array='options_array'
                  :rules="rules"
                  :values="values"
                  :disabled="disabled"
  />
</template>

<script>
import SelectColumn from "@/components/form/column/base/SelectColumn.vue";

export default {
  components: {
      'select-column': SelectColumn
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
    options_array: {
      type: Array,
      default: () => [],
    },
    disabled: {
          type: Boolean,
          default: false,
    },
  },
  data() {
    return {
        english_name: 'function_operate',
        chinese_name: '功能操作',
    }
  },
};
</script>