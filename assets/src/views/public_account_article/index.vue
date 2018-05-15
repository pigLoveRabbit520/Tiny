<template>
  <div class="app-container">
      <el-row>
          <el-col :span="18">
              <el-form ref="form" :model="article" :inline="true" label-width="80px" @submit.prevent="onSubmit">
                  <el-form-item label="公众号">
                      <el-input v-model="article.accountName"/>
                  </el-form-item>
                  <el-form-item label="标题">
                      <el-input v-model="article.title"/>
                  </el-form-item>
                  <el-form-item>
                      <el-button type="primary" @click="onSubmit">搜索</el-button>
                  </el-form-item>
              </el-form>
          </el-col>
      </el-row>

      <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
              :data="articles"
              style="width: 100%">
          <el-table-column
                  type="index"
                  label="#"
                  width="180">
          </el-table-column>
          <el-table-column
                  prop="title"
                  label="文章标题">
          </el-table-column>
          <el-table-column
                  prop="publish_datetime"
                  label="发布时间">
          </el-table-column>
          <el-table-column
                  prop="originality"
                  label="原创度">
          </el-table-column>
          <el-table-column
                  prop="name"
                  label="公众号">
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
import { getList } from '@/api/public_account_article'
import ElRow from 'element-ui/packages/row/src/row'

export default {
    components: { ElRow },
    data() {
        return {
            list: null,
            listLoading: true,
            articlesTotal: 0,
            pageSize: 15,
            currentPage: 1,
            articles: [],
            article: {
                accountName: '',
                title: '',
            }
        }
    },
    created() {
        if (this.$route.params) {
            this.article.accountName = this.$route.params.accountName;
        }
        this.fetchAccounts(this.article)
    },
    methods: {
        fetchAccounts(params) {
            this.listLoading = true
            getList(params).then(response => {
                this.articles = response.data.articles
                this.articlesTotal = response.data.total
                this.listLoading = false
            })
        },
        handleCurrentChange(val) {
            this.fetchAccounts({
                page: val,
                accountName: this.article.accountName
            })
        },
        onSubmit() {
            this.fetchAccounts(this.article)
        }
    }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

