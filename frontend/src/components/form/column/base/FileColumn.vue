<template>
       <base-slot-column
               :input_id="input_id"
               :input_group_id="input_group_id"
               :input_live_feedback_id="input_live_feedback_id"
               :label="label"
               :rules="rules"
       >
          <template v-slot:columns="{ validationContext }" >
              <b-form-file
                  :id="input_id"
                  v-model="localValues"
                  @input="handleChange"
                  :state="getValidationState(validationContext)"
                  :disabled="disabled"
                  :multiple="multiple"
                  :accept="accept"
                  :browse-text="'选择档案'"
                  placeholder="选择文件 或 拖放到此..."
                  drop-placeholder="拖放到此..."
                ></b-form-file>
           </template>
       </base-slot-column>
</template>

<script>

import BaseSlotColumn from "./BaseSlotColumn.vue";
import {ColumnMixin} from "@/mixins/column/column";

export default {
  components: {
      'base-slot-column':BaseSlotColumn,
  },
  mixins: [ColumnMixin],
  data() {
    return {
      localValues: this.values // 创建一个本地副本
    };
  },
  props: {
    english_name: {
      type: String,
      required: true,
    },
    chinese_name: {
      type: String,
      required: true,
    },
    rules: {
      type: String,
      default: ""
    },
    values: {
      type: [String],
      default: null,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    accept: {
        type: String,
        default: "",
    },
    multiple: {
      type: Boolean,
      default: false,
    },
  },
  created() {
    // 在组件生成时触发的函数
    this.handleChange();
  },
};
</script>
