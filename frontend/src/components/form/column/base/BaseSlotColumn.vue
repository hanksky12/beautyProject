<template>

  <validation-provider :name="input_id" :rules="rules" v-slot="validationContext">
    <b-form-group :id="input_group_id" :label="label" :label-for="input_id">
      <slot name="columns" v-bind:validationContext="validationContext"></slot>
      <b-form-invalid-feedback :id="input_live_feedback_id">
          {{ validationContext.errors[0] }}
      </b-form-invalid-feedback>
    </b-form-group>
  </validation-provider>
</template>

<script>

export default {
  //驗證表單欄位 機制
  props: {
    input_id: {
      type: String,
      required: true,
    },
    input_group_id: {
      type: String,
      required: true,
    },
    input_live_feedback_id: {
      type: String,
      required: true,
    },
    label: {
      type: String,
      required: true,
    },
    rules: {
          type: [Object, String],
          default: ""
        },
  },
};
</script>