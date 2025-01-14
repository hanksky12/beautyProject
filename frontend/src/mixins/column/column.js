import { createNamespacedHelpers } from 'vuex'
const { mapGetters:mapGetterForm } = createNamespacedHelpers('form')
export const ColumnMixin =  {
  data() {
    return {
        localValues: null,
    };
  },
  computed: {
    input_id() {
      return `input_${this.english_name}`;
    },
    input_group_id() {
      return `input_group_${this.english_name}`;
    },
    input_live_feedback_id() {
      return `input_live_feedback_${this.english_name}`;
    },
    label() {
      return `${this.chinese_name}:`;
    },
    ...mapGetterForm([
      'getFormData',
    ]),
  },
  methods: {
    getValidationState({ dirty, validated, valid = null }) {
      return dirty || validated ? valid : null;
    },
    handleChange() {
      this.$store.commit("form/setColumn", {[this.english_name]:this.localValues});
    },
    handleEnglishNameChange(newVal) {
      this.localValues = newVal[this.english_name];
    },
  },
  watch: {
    //將表單與formData 綁定，改變formData就可同步影響呈現
    'getFormData': {
      handler: 'handleEnglishNameChange',
      deep: true,
    },
  },
};
