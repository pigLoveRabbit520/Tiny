<template>
    <div class="app-container">
        <el-row style="margin-bottom: 20px;">
            <el-button type="primary">添加独立页面</el-button>
        </el-row>
        <el-row>
          <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
                    :data="pages"
                    style="width: 100%">
            <el-table-column
              type="selection"
              width="55">
            </el-table-column>
            <el-table-column
              prop="title"
              label="标题">
            </el-table-column>
            <el-table-column
              prop="slug"
              label="缩略名">
            </el-table-column>
            <el-table-column
              prop="screenName"
              label="作者">
            </el-table-column>
            <el-table-column
              prop="created"
              :formatter="getDate"
              label="发布日期">
            </el-table-column>]
          </el-table>
        </el-row>
    </div>
</template>

<script>
import { getList, addPage } from '@/api/page'

export default {
    data() {
        return {
            list: null,
            listLoading: true,
            pages: [],
            setPushUrlsDialogVisible: false,
            addAccountDialogVisible: false,
            form: {
                name: '',
                wxNum: ''
            }
        }
    },
    created() {
        document.title = "独立页面管理"
        this.getPages()
    },
    methods: {
        getPages() {
            this.listLoading = true
            getList().then(response => {
                this.pages = response.data.pages
                this.listLoading = false
            })
        },
        getDate(row, col) {
            let value = row[col.property];
            let formatter = function (v) {
                return ('0' + v).substr(-2)
            }
            if (value) {
                let now = new Date()
                let date = new Date(value * 1000);
                let year = date.getFullYear();
                let month = formatter(date.getMonth() + 1);
                let day = formatter(date.getDate());
                let hours = formatter(date.getHours());
                let minutes = formatter(date.getMinutes());
                let seconds = formatter(date.getSeconds());
                if (now.getFullYear() === year) {
                     return `${month}-${day} ${hours}:${minutes}:${seconds}`
                } else {
                     return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
                }
            } else {
                return ""
            }
        }
    },
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

