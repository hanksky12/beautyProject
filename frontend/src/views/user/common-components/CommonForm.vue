<template>

  <div class="container d-flex justify-content-center align-items-center" style="min-height: 50vh;">
    <base-slot-form :fields="fields" :customize_columns="1" @submit="handleFormSubmit" >
      <template v-slot:title>
        {{title}}
      </template>
      <template v-for="field in fields"  v-slot:[field.slotName] >
            <component :is="field.component" :options_array="field.props.options_array" :rules="field.props.rules"></component>
      </template>
      <template v-slot:button>
        <base-button :btn_name="title" :btn_type="'submit'" :btn_variant="'primary'"/>
      </template>
    </base-slot-form>
  </div>
</template>

<script>
import UserNameColumn from "@/components/form/column/UserNameColumn.vue";
import UserPasswordColumn from "@/components/form/column/UserPasswordColumn.vue";
import BaseSlotForm from '@/components/form/BaseSlotForm.vue';
import BaseButton from "@/components/form/button/base/BaseButton.vue"

export default {
  components: {
      'base-button':BaseButton,
      'base-slot-form': BaseSlotForm,
      'user-name-column': UserNameColumn,
      'user-password-column': UserPasswordColumn,
  },
  props: {
    title: {
      type: String,
      required: true,
    },
  },
  data() {
      return {
       fields: [
        {
            slotName: 'user-name-column',
            component: 'user-name-column' ,
            props: {rules:"required"}
        },
        {
            slotName: 'user-password-column',
            component: 'user-password-column' ,
            props: {rules:"required"}
        }
      ],
      }
  },
  methods: {
    handleFormSubmit() {
      this.$emit('handleFormSubmit');
    },
  },
};
</script>

<style>
.card {
  border: 1px solid #e1e1e1; /* 設置筐線樣式 */
  border-radius: 8px; /* 可選：設置筐線的圓角 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); /* 可選：設置陰影效果 */
}
</style>