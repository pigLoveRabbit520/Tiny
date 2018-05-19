<template>
    <div class="app-container">
        <el-row style="margin-bottom: 20px;">
          <el-button type="primary">添加文章</el-button>
        </el-row>

        <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
                :data="posts"
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
                    prop="screenName"
                    label="作者">
            </el-table-column>
            <el-table-column
                    prop="meta_cat"
                    label="分类">
                  <template slot-scope="scope">
                      <el-button v-for="cate in scope.row.cates" type="text" size="small">{{ cate.name }}</el-button>
                  </template>
            </el-table-column>
            <el-table-column
                    prop="created"
                    :formatter="getDate"
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

export default {
    data() {
        return {
            listLoading: true,
            postsTotal: 0,
            pageSize: 15,
            currentPage: 1,
            posts: [],
            cates: [],
        }
    },
    created() {
        document.title = "文章管理"
        this.getPosts()
    },
    methods: {
        getPosts(params) {
            this.listLoading = true
            getList(params).then(response => {
                this.posts = response.data.posts
                this.cates = response.data.cates
                this.setPostCates()
                this.postsTotal = response.data.total
                this.listLoading = false
            })
        },
        handleCurrentChange(val) {
            this.getPosts({
                page: val
            })
        },
        setPostCates() {
            let cateMap = {}
            for (let cate of this.cates) {
                cateMap[cate.mid] = cate
            }
            for (let post of this.posts) {
                post.cates = []
                if (post.meta_cat) {
                    let ids = post.meta_cat.split(",")
                    for (let id of ids) {
                        post.cates.push(cateMap[id])
                    }
                }
            }
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
    }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

