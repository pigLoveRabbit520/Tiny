<template>
  <div class="app-container">
      <el-row>
      </el-row>

      <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
              :data="posts"
              style="width: 100%">
          <el-table-column
                  type="index"
                  label="#"
                  width="180">
          </el-table-column>
          <el-table-column
                  prop="title"
                  label="标题">
          </el-table-column>
          <el-table-column
                  prop="screenName"
                  label="作者">
          </el-table-column>
          <el-table-column
                  prop="category"
                  label="分类">
          </el-table-column>
          <el-table-column
                  prop="created"
                  label="发布日期">
          </el-table-column>]
      </el-table>

      <div style="margin-top: 10px; margin-left: -10px">
          <el-pagination
                  background
                  layout="prev, pager, next"
                  @current-change="handleCurrentChange"
                  :page-size="pageSize"
                  :current-page.sync="currentPage"
                  :total="postsTotal">
          </el-pagination>
      </div>
  </div>
</template>

<script>
import { getList } from '@/api/post'
import ElRow from 'element-ui/packages/row/src/row'

export default {
    components: { ElRow },
    data() {
        return {
            list: null,
            listLoading: true,
            postsTotal: 0,
            pageSize: 15,
            currentPage: 1,
            posts: [],
        }
    },
    created() {
        this.getPosts()
    },
    methods: {
        getPosts(params) {
            this.listLoading = true
            getList(params).then(response => {
                this.posts = response.data.posts
                this.postsTotal = response.data.total
                this.listLoading = false
            })
        },
        handleCurrentChange(val) {
            this.getPosts({
                page: val
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

