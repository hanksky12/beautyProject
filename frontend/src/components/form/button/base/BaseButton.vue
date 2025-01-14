<template>
  <b-button pill size="lg" :type="btn_type" :variant="btn_variant" @click="click">{{btn_name}} </b-button>
</template>

<script>
export default {
    name: "BaseButton",
    props: {
    btn_name: {
      type: String,
      required: true,
    },
    btn_type: {
      type: String,
      required: true,
    },
    btn_variant: {
      type: String,
      required: true,
    },
  },
  methods: {
    click() {
      this.$emit('click');
    }
  }
}
</script>
