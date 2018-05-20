<template>
    <div class="app-container">

        <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
                :data="articles"
                style="width: 100%">
            <el-table-column
                    type="selection"
                    width="22">
            </el-table-column>
            <el-table-column
                    prop="title"
                    label="名称">
            </el-table-column>
            <el-table-column
                    prop="slug"
                    label="缩略名">
            </el-table-column>
            <el-table-column
                    prop="num"
                    label="文章数量">
            </el-table-column>]
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
import { getList } from '@/api/category'

export default {
    data() {
        return {
            list: null,
            listLoading: true,
            postsTotal: 0,
            pageSize: 15,
            currentPage: 1,
            articles: [],
        }
    },
    created() {
        this.getCategories()
    },
    methods: {
        getCategories(params) {
            this.listLoading = true
            getList(params).then(response => {
                this.articles = response.data.articles
                this.articlesTotal = response.data.total
                this.listLoading = false
            })
        },
        handleCurrentChange(val) {
            this.posts({
                page: val,
                accountName: this.article.accountName
            })
        },
        onSubmit() {
            this.posts(this.article)
        }
    }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

