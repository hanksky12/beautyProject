<template>
  <div class="col-sm-10  mx-auto">
   <div class="card">
     <div class="card-body">
      <div class="base-form">
          <h1><slot name="title"></slot></h1>
          <slot name="content"></slot>
        <b-container fluid="lg">
          <validation-observer ref="observer" v-slot="{ handleSubmit }">
            <b-form  @submit.stop.prevent="handleSubmit(onSubmit)" @reset="onReset" v-if="show">
              <template v-for="rowIndex in rows" >
<!--                  {{"row"+rowIndex}}-->
                  <b-row >
                    <template v-for="colIndex in columns">
<!--                        {{"column"+colIndex}}-->
                      <b-col >
                        <slot :name="fields[(rowIndex-1)*columns + colIndex-1].slotName"   v-if="((rowIndex-1)*columns + colIndex-1)<fields.length"></slot>
                      </b-col>
                    </template>
                  </b-row>
              </template>
              <b-row>
                <b-col>
                    <slot name="button"></slot>
                </b-col>
              </b-row>
            </b-form>
          </validation-observer>
        </b-container>
      <p  class="left-align"><slot name="ps_content"></slot></p>
      </div>
     </div>
    </div>
  </div>
</template>

<script>
import { ValidationObserver } from "vee-validate";
export default {
  //動態生成表單 欄、列數目   組件   驗證機制
  props: {
    fields:{
        type: Array,
        required: true,
    },
    customize_columns: {
        type: Number,
        default: 0,
      }
  },
   components: {
    'validation-observer':ValidationObserver,
  },
  data() {
    return {
      show: true,
    };
  },
  computed: {
    columns() {
        if (this.customize_columns > 0) {
            return this.customize_columns;
        }
        if (this.fields.length < 3){
            return this.fields.length;
        }
        return 3;
        },
    rows() {
        return Math.ceil(this.fields.length / this.columns);
        },
    },
  methods: {
    onSubmit() {
      console.log('BaseForm onSubmit~~~~')
      this.$emit('submit');
    },
    onReset() {
        this.$store.commit("crudOperate/resetOperate")
        // Trick to reset/clear native browser form validation state
        this.show = false
        this.$nextTick(() => {
          this.show = true
        })
        this.$emit('reset');
    },
  },
};
</script>


<style>
.card {
  border: 1px solid #e1e1e1; /* 設置筐線樣式 */
  border-radius: 8px; /* 可選：設置筐線的圓角 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 1.0); /* 可選：設置陰影效果 */

  background-color: #e9ebf3;
  color: #d96767;
}

.left-align {
    text-align: left;
}
</style>