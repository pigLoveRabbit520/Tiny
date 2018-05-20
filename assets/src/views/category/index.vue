<template>
    <div class="app-container">

        <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
                :data="categories"
                style="width: 100%">
            <el-table-column
                    type="selection"
                    width="33">
            </el-table-column>
            <el-table-column
                    prop="name"
                    label="名称">
            </el-table-column>
                <el-table-column
                    prop="子分类"
                    label="子分类">
            </el-table-column>
            <el-table-column
                    prop="slug"
                    label="缩略名">
            </el-table-column>
            <el-table-column
                    prop="count"
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
            categoriesTotal: 0,
            pageSize: 15,
            currentPage: 1,
            categories: [],
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
                this.categories = response.data.categories
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

