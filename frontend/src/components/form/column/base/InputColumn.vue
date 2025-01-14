<template>
       <base-slot-column
               :input_id="input_id"
               :input_group_id="input_group_id"
               :input_live_feedback_id="input_live_feedback_id"
               :label="label"
               :rules="rules"
       >
          <template v-slot:columns="{ validationContext }" >
              <b-form-input
              :id="input_id"
              v-model="localValues"
              :type="type"
              @change="handleChange"
              :state="getValidationState(validationContext)"
              :aria-describedby="input_live_feedback_id"
              :disabled="disabled"
              :min="min"
              :max="max"
              :step="step"
              ></b-form-input>
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
    type: {
      type: String,
      default: 'text',
    },
    min: {
      type: Number,
      default: 0,
    },
    max: {
      type: Number,
      default: 100,
    },
    step: {
      type: Number,
      default: 1,
    },
  },
  created() {
    // 在组件生成时触发的函数
    this.handleChange();
  },
};
</script>