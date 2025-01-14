<template>
      <crud-base-button :btn_name="btn_name" :btn_variant="btn_variant" />
</template>

<script>
import CrudBaseButton from "@/components/form/button/crud/CrudBaseButton.vue";

export default {
  components: {
        'crud-base-button': CrudBaseButton
  },
  data() {
        return {
            btn_name: '新增',
            btn_variant: 'outline-primary',
        }
  }
}
</script>