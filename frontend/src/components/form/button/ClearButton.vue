<template>
  <base-button :btn_name="btn_name" :btn_type="btn_type" :btn_variant="btn_variant"/>
</template>

<script>

import BaseButton from "./base/BaseButton.vue";


export default {
  components: {
      'base-button': BaseButton
  },
    data() {
        return {
            btn_name: '清空',
            btn_type: 'reset',
            btn_variant: 'danger',
        }
    },
}
</script>
