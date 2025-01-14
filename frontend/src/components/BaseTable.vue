<template>
    <div className="result-table">
        <BootstrapTable
                ref="table"
                :columns="columns"
                :options="options"
                @on-check="selectedChange"
                @on-check-all="selectedChange"
                @on-uncheck="selectedChange"
                @on-uncheck-all="selectedChange"
                @on-post-body="afterRendered"
                @on-click-cell="clicCellEvent"
        />
    </div>
</template>

<script>
import BootstrapTable from 'bootstrap-table/dist/bootstrap-table-vue.esm.js'
import {createNamespacedHelpers} from 'vuex'

const {mapState: mapStateQuery} = createNamespacedHelpers('query')
const {mapGetters: mapGetterForm} = createNamespacedHelpers('form')
const {mapState: mapStateConst} = createNamespacedHelpers('const')
const {mapState: mapStateAjax} = createNamespacedHelpers('ajax')

import {PrintToastMixin} from "@/mixins/utils/send-request-to-post-processing/printToast.js";
import {GetDomainMixin} from "@/mixins/utils/handle-request/create-request/getDomain.js";

export default {
    components: {
        BootstrapTable,
    },
    mixins: [PrintToastMixin, GetDomainMixin],
    inject: ['showToast', 'showSwAl'],
    props: {
        columns: {
            type: Array,
            required: true
        },
        api_path: {
            type: String,
            required: true
        },
        api_domain: {
            type: String,
            default: undefined
        },
        row_style: {
            type: Function,
            default: function () {
                return {classes: "bg-type-yellow"}
            }
        },
        header_style: {
            type: Function,
            default: function (column) {
                return {
                    // id: {
                    //   classes: 'uppercase'
                    // },
                }[column.field]
            }
        },
        table_height: {
            type: Number,
            default: 780,
        },
        multipleSelectRow: {
            type: Boolean,
            default: false,
        },
        checkboxHeader: {
            type: Boolean,
            default: true,
        },
        click_cell_event: {
            type: Function,
            default: function () {
            }
        },
        callApis: {
            type: Object,
        },
        callFuncs: {
            type: Object,
        },
    },
    data() {
        return {
            options: {},
        }
    },

    methods: {
        getData() {
            console.log("[Table] getData")
            return this.$refs.table.getData()
        },
        queryParam(params) {
            let temp = {
                per_page: params.limit, // 必填 因為一次只要顯示一頁,所以這邊等於一頁的頁數即可 受pageSize影響
                page: (params.offset / params.limit) + 1, //必填 從offset 推算頁數
                sortOrder: params.order, //找sortOrder的參數
            }
            if (params.sort) {
                temp["sort"] = params.sort
            } //如果有點擊欄位，就會抓到欄位名稱
            temp = {...temp, ...this.getFormData}

            console.log("[Table] QueryParam")
            console.log(this.getFormData)
            console.log(temp)
            return temp
        },
        responseHandler(res) {
            console.log("[Table] response")
            if (res.code < 200 | res.code > 300) {
                this.$store.dispatch("ajax/failHandling", res)
                this.printToast(this.responseData)
            }
            console.log(res)
            return res
        },
        onLoadError(status, jqXHR) {
            console.log("on Load Table Error")
            console.log(status)
            console.log(jqXHR)
            let msg = ""
            if (jqXHR.status == 0) {
                msg = "無法連線到伺服器"
            } else {
                msg = jqXHR.responseJSON.msg
                // msg=JSON.stringify(jqXHR, null, 2)
            }
            this.showToast(`表格資料載入失敗! ${msg}`, "danger")
        },
        selectedChange() {
            // console.log("[Table] 起源 selectedChange")
            const selections = this.$refs.table.getSelections();
            // console.log(selections)
            this.$store.commit("table/setTableSelected", selections);
        },
        afterRendered() {
            // console.log("[Table] afterRendered")
            this.$store.commit("query/setQueryClick", false);
            this.selectedChange()
        },
        refresh_table() {
            console.log("[Table] refresh_table")
            this.$refs.table.refresh();
        },
        clicCellEvent(field, value, row) {
            return this.click_cell_event(field, value, row, this.callApis, this.callFuncs)
        },
    },
    created() {
        // 在 created 钩子中初始化 options ，因為url從state取得，所以不能在data中初始化
        this.options = {
            //分頁相關
            sidePagination: "server",//方式：client ,server
            cache: false,   //使用缓存，默认为 true，所以一般情况下需要设置一下这个属性
            pagination: true,  //顯示分頁
            pageList: [20, 100, 1000, 99999], //可供選擇的每頁的行數 ,選擇後更改pageSize
            pageNumber: 1,//初始頁
            pageSize: 20,//每頁筆數
            formatRecordsPerPage: (pageSize) => {
                return '&nbsp;&nbsp;每頁顯示' + pageSize + '筆';
            },
            formatShowingRows: function (fromIndex, toIndex, totalSize) {
                let currentPage = Math.ceil(fromIndex / this.pageSize)      //目前第幾頁
                let totalPageCount = Math.ceil(totalSize / this.pageSize) //總共幾頁
                return '第' + currentPage + '頁&nbsp;&nbsp;共' + totalPageCount + '頁'
            },

            //排序相關
            sortable: true, //開啟排序，會在onSort重抓參數後，自動像後端發送目前參數的request
            sortOrder: "desc",//預設 大到小

            //選擇相關
            checkboxHeader: this.checkboxHeader,//一次選擇所有
            multipleSelectRow: this.multipleSelectRow,//是否可以多選 false:可以 true:不可以
            clickToSelect: true,

            //外觀
            height: this.table_height,
            classes: "table  table-hover border-primary table-bordered table-sm  text-nowrap",
            rowStyle: this.row_style,
            headerStyle: this.header_style,
            showColumns: true,

            //Ajax定義 處理
            method: 'get',
            contentType: "application/json",
            url: this.getDomain(this.api_domain) + this.api_path,
            dataField: "data",//後端回來裝data的key,
            ajaxOptions: {
                xhrFields: {
                    withCredentials: true, // 确保带上 cookie
                },
            },
            totalField: "total",// 後端返回含有total的key：總記錄數
            queryParams: (params) => {
                return this.queryParam(params)
            },
            onLoadError: (status, jqXHR) => {
                return this.onLoadError(status, jqXHR)
            },
            responseHandler: (res) => {
                return this.responseHandler(res)
            },
        }
    },
    computed: {
        ...mapGetterForm([
            'getFormData',
        ]),
        ...mapStateQuery({
            'is_query_click': state => state.is_query_click,
        }),
        ...mapStateConst({
            'const': state => state.const,
        }),
        ...mapStateAjax({
            'responseData': state => state.responseData,
        }),
    },
    watch: {
        is_query_click: {
            handler() {
                if (this.is_query_click) {
                    this.refresh_table();
                }
            }
        },
    },
}
</script>