export const moduleConst = {
    namespaced: true,
    state: {
      const: {
        domain: process.env.VUE_APP_API_DOMAIN,
        headers:  {
              'Content-Type': 'application/json',
              'Accept-Charset': 'utf-8',
              'Accept': 'application/json'
        },
        changeOperate: {
            "新增":"insert" ,
            "修改":"update" ,
            "刪除":"delete" ,
        },
      },
    },
}