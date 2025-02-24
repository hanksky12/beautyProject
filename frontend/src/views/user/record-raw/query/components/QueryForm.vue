<template>
    <div class="query">
    <base-slot-form :fields="fields" :customize_columns="2" @submit="handleFormSubmit" >
      <template v-slot:title>
        ~@[查询]@~
      </template>
      <template v-for="field in fields"  v-slot:[field.slotName] >
            <component :is="field.component"
            >
            </component>
      </template>
       <template v-slot:button>
        <execute-button />
        <clear-button />
      </template>
    </base-slot-form>
  </div>

</template>

<script>
import {createNamespacedHelpers} from "vuex";
const { mapState:mapStateConst} = createNamespacedHelpers('const')

import BaseSlotForm from '@/components/form/BaseSlotForm.vue';
import ClearButton from "@/components/form/button/ClearButton.vue";
import ExecuteButton from "@/components/form/button/ExecuteButton.vue";
import MaxDateColumn from "@/components/form/column/MaxDateColumn.vue";
import MinDateColumn from "@/components/form/column/MinDateColumn.vue";
import MaxTimeColumn from "@/components/form/column/MaxTimeColumn.vue";
import MinTimeColumn from "@/components/form/column/MinTimeColumn.vue";
import MouseActionNameColumn from "@/components/form/column/MouseActionNameColumn.vue";
import MaxXColumn from "@/components/form/column/MaxXColumn.vue";
import MinXColumn from "@/components/form/column/MinXColumn.vue";
import MaxYColumn from "@/components/form/column/MaxYColumn.vue";
import MinYColumn from "@/components/form/column/MinYColumn.vue";

  export default {
    components: {
            'base-slot-form':BaseSlotForm,
            'max-date-column':MaxDateColumn,
            'min-date-column':MinDateColumn,
            'max-time-column':MaxTimeColumn,
            'min-time-column':MinTimeColumn,
            'max-x-column':MaxXColumn,
            'min-x-column':MinXColumn,
            'max-y-column':MaxYColumn,
            'min-y-column':MinYColumn,
            'mouse-action-name-column':MouseActionNameColumn,
            'clear-button':ClearButton,
            'execute-button':ExecuteButton,
    },
    methods: {
      handleFormSubmit() {
          console.log('Query onSubmit~~~~~~~')
          this.$store.commit("query/setQueryClick", true);
      },
    },
    computed: {
    ...mapStateConst({
        'const': state => state.const,
    }),
    fields() {
     return [
       {
          slotName: 'max-date-column',
          component: 'max-date-column',
        },
       {
         slotName: 'max-time-column',
         component: 'max-time-column',
       },
        {
          slotName: 'min-date-column',
          component: 'min-date-column',
        },
        {
          slotName: 'min-time-column',
          component: 'min-time-column',
       },
        {
          slotName: 'max-x-column',
          component: 'max-x-column',
        },
        {
          slotName: 'min-x-column',
          component: 'min-x-column',
        },
       {
         slotName: 'max-y-column',
         component: 'max-y-column',
       },
       {
         slotName: 'min-y-column',
         component: 'min-y-column',
       },
       {
         slotName: 'mouse-action-name-column',
         component: 'mouse-action-name-column' ,
       },
      ]
    },
  }
}
</script>


<style scoped>
@import '@/components/style/query.css';
</style>