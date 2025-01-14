<template>
      <base-button :btn_name="btn_name" :btn_type="btn_type" :btn_variant="btn_variant" @click="click"/>
</template>

<script>

import BaseButton from "@/components/form/button/base/BaseButton.vue";


export default {
  components: {
      'base-button': BaseButton
  },
  props: {
    btn_name: {
      type: String,
      required: true,
    },
    btn_variant: {
      type: String,
      required: true,
    },
  },

  data() {
      return {
          btn_type: 'submit',
      }
  },
  methods: {
    click() {
      this.$store.commit("crudOperate/setOperate", this.btn_name)
    }
  }
}
</script>