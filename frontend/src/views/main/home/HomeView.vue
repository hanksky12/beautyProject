<template>
    <div>
        <bread-crumb :props_items="items"/>
        <br>
    </div>
</template>

<script>
import Breadcrumb from '@/components/Breadcrumb.vue'
export default {
    components: {
        'bread-crumb':Breadcrumb,
    },
    data() {
      return {
        items: [
          {
            text: '首頁',
             active: true
          },
        ]
      }
    }
}
</script>

