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
import { GetHardwareNameApiMixin } from "@/mixins/api/server/hardware/getName";


export default {
  mixins: [GetHardwareNameApiMixin],
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
    disabled: {
          type: Boolean,
          default: false,
    },
  },
  data() {
    return {
        english_name: 'hardware_name',
        chinese_name: '硬體名稱',
        options_array: [{ text: '未选择', value: null }],
    }
  },
  async created() {
    try {
      const response = await this.getHardwareNameApi();
      console.log("create response:", response);
      this.options_array = [{ text: '未选择', value: null }, ...response.data.map(item => ({
        text: item.hardware_name, // 使用 hardware_name 作為顯示的文字
        value: item.hardware_name  // 使用 hardware_name 作為選項的值
      }))];    } catch (error) {
      console.error("Error fetching hardware names:", error);
    }
  }
};
</script>
