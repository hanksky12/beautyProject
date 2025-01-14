<template>
       <base-slot-column
               :input_id="input_id"
               :input_group_id="input_group_id"
               :input_live_feedback_id="input_live_feedback_id"
               :label="label"
               :rules="rules"
       >
          <template v-slot:columns="{ validationContext }" >
              <b-form-group >
                <b-form-checkbox-group
                  :id="input_id"
                  v-model="localValues"
                  :disabled="disabled"
                  :options="options_array"
                  @change="handleChange"
                  :state="getValidationState(validationContext)"
                ></b-form-checkbox-group>
              </b-form-group>

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
    options_array: {
      type: [Array,null],
      required: true,
    },
    rules: {
      type: String,
      default: ""
    },
    values: {
      type: [Array,null],
      default: null,
    },
    disabled: {
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
