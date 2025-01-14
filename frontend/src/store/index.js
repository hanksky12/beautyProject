import Vue from 'vue'
import Vuex from 'vuex'
import {moduleQuery} from "@/store/modules/query";
import {moduleTable} from "@/store/modules/table";
import {moduleAjax} from "@/store/modules/ajax";
import {moduleConst} from "@/store/modules/const";
import {moduleForm} from "@/store/modules/form";
import {moduleAuthenticate} from "@/store/modules/authenticate";
import {moduleAuthorization} from "@/store/modules/authorization"
import {moduleCrudOperate} from "@/store/modules/crud_operate";

Vue.use(Vuex)

const store= new Vuex.Store({

  modules: {
    query: moduleQuery,
    table: moduleTable,
    ajax: moduleAjax,
    const: moduleConst,
    form: moduleForm,
    authenticate: moduleAuthenticate,
    authorization:moduleAuthorization,
    crudOperate: moduleCrudOperate,
  }
}
)

export default store;