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
import {GetMouseActionNameApiMixin} from "@/mixins/api/user/mouse/getName";


export default {
  mixins: [GetMouseActionNameApiMixin],
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
        english_name: 'mouse_action_name',
        chinese_name: '滑鼠操作名稱',
        options_array: [{ text: '未选择', value: null }],
    }
  },
  async created() {
    try {
      const response = await this.getMouseActionNameApi();
      console.log("create response:", response);
      this.options_array = [{ text: '未选择', value: null }, ...response.data.map(item => ({
        text: item.mouse_action_name,
        value: item.mouse_action_name
      }))];    } catch (error) {
      console.error("Error fetching hardware names:", error);
    }
  }
};
</script>