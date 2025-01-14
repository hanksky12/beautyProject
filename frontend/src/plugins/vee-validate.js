import Vue from 'vue';
import { ValidationProvider, extend } from 'vee-validate';
import { required} from "vee-validate/dist/rules";

//Add a rule.
extend("required", {
  ...required,
  message: "必填欄位"
});

// Register it globally
Vue.component('ValidationProvider', ValidationProvider);


