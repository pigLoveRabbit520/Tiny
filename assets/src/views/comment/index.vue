<template>
    <div class="app-container">

        <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
                :data="comments"
                style="width: 100%">
            <el-table-column
                    type="selection"
                    width="33">
            </el-table-column>
            <el-table-column
                    prop="name"
                    label="作者">
                <template slot-scope="scope">
                    <el-button v-for="cate in scope.row.cates" type="text" size="small">{{ cate.name }}</el-button>
                </template>
            </el-table-column>
            <el-table-column
                    prop="content"
                    label="内容">
            </el-table-column>
        </el-table>

        <div style="margin-top: 10px; margin-left: -10px">
            <el-pagination
                    background
                    layout="prev, pager, next"
                    @current-change="handleCurrentChange"
                    :page-size="pageSize"
                    :current-page.sync="currentPage"
                    :total="articlesTotal">
            </el-pagination>
        </div>
    </div>
</template>

<script>
import { getList } from '@/api/comment'

export default {
    data() {
        return {
            listLoading: true,
            commentsTotal: 0,
            pageSize: 15,
            currentPage: 1,
            comments: [],
        }
    },
    created() {
        document.title = "分类管理"
        this.getCategories()
    },
    methods: {
        getCategories(params) {
            this.listLoading = true
            getList(params).then(response => {
                this.comments = response.data.categories
                this.categoriesTotal = response.data.total
                this.listLoading = false
            })
        },
        handleCurrentChange(val) {
            this.getCategories({
                page: val,
                accountName: this.article.accountName
            })
        }
    }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

