<template>
  <input-column  :english_name='english_name'
                 :chinese_name='chinese_name'
                 :rules="rules"
                 :values="values"
                 :disabled="disabled"
                 :type="type"
                 :min="min"
                 :max="max"
                 :step="step"
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
        english_name: 'max_percent',
        chinese_name: '最大運作率',
        type:"number",
        min: 0,
        max: 100,
        step: 0.1
    }
  },
};
</script>

