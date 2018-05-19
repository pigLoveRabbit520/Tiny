<template>
  <div class="app-container">
      <el-row style="margin-bottom: 20px;">
          <el-button type="primary" @click="showAddAccount">添加公众号</el-button>
      </el-row>
      <el-row>
          <el-col :span="16">
              <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
                        :data="accounts"
                        style="width: 100%">
                  <el-table-column
                          type="index"
                          label="#"
                          width="180">
                  </el-table-column>
                  <el-table-column
                          prop="wx_number"
                          label="微信号"
                          width="180">
                  </el-table-column>
                  <el-table-column
                          prop="name"
                          label="名称">
                  </el-table-column>
                  <el-table-column
                          prop="push_urls"
                          label="推送链接">
                  </el-table-column>
                  <el-table-column
                          fixed="right"
                          label="操作"
                          width="200">
                      <template slot-scope="scope">
                          <el-button @click="viewArticles(scope.row)" type="text" size="small">查看文章</el-button>
                          <el-button @click="showSetPushUrls(scope.row)" type="text" size="small">设置推送链接</el-button>
                      </template>
                  </el-table-column>
              </el-table>
          </el-col>
      </el-row>

      <el-dialog
              :title="currentAccount.name + '推送链接'"
              :visible.sync="setPushUrlsDialogVisible"
              width="25%"
              center>
          <textarea style="width: 100%;" rows="6" placeholder="一行一个" v-model="urls"></textarea>
          <span slot="footer" class="dialog-footer">
            <el-button @click="setPushUrlsDialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="updatePushUrls">确 定</el-button>
        </span>
      </el-dialog>

      <el-dialog
              title="添加公众号"
              :visible.sync="addAccountDialogVisible"
              width="25%"
              center>
          <el-form ref="form" :model="form" label-width="130px">
              <el-form-item label="公众号名称">
                  <el-input v-model="form.name"></el-input>
              </el-form-item>
              <el-form-item label="公众号微信号">
                  <el-input v-model="form.wxNum"></el-input>
              </el-form-item>
          </el-form>
          <span slot="footer" class="dialog-footer">
            <el-button @click="addAccountDialogVisible = false">取 消</el-button>
            <el-button type="primary" @click="addAccount">确 定</el-button>
        </span>
      </el-dialog>
  </div>
</template>

<script>
import { getList, updatePushUrls, addAccount } from '@/api/page'

export default {
    data() {
        return {
            list: null,
            listLoading: true,
            accounts: [],
            setPushUrlsDialogVisible: false,
            addAccountDialogVisible: false,
            currentAccount: {},
            urls: '',
            form: {
                name: '',
                wxNum: ''
            }
        }
    },
    created() {
        this.posts()
    },
    methods: {
        posts() {
            this.listLoading = true
            getList().then(response => {
                this.accounts = response.data
                for (let account of this.accounts) {
                    if (account.push_urls) {
                        account.push_urls = (JSON.parse(account.push_urls)).join("\r\n")
                    }
                }
                this.listLoading = false
            })
        },
        viewArticles(row) {
            this.$router.push({name: 'PublicAccountArticle', params: { accountName: row.name }})
        },
        showSetPushUrls(row) {
            this.currentAccount = row
            this.urls = row.push_urls
            this.setPushUrlsDialogVisible = true
        },
        updatePushUrls() {
            if (!this.currentAccount.account_id) {
                this.$message({
                    message: '没有选择公众号',
                    type: 'warning'
                });
            } else {
                let accountId = this.currentAccount.account_id
                let pushUrls = this.urls
                let that = this
                updatePushUrls({
                    accountId,
                    pushUrls
                }).then(response => {
                    this.$message({
                        message: '更新成功',
                        type: 'success'
                    });
                    this.setPushUrlsDialogVisible = false
                    that.currentAccount.push_urls = that.urls
                })
            }
        },
        showAddAccount() {
            this.addAccountDialogVisible = true
        },
        addAccount() {
            if (!this.form.name || !this.form.wxNum) {
                this.$message({
                    message: '名称或微信号未填写',
                    type: 'warning'
                });
                return
            }
            let account = {
                wxNum: this.form.wxNum,
                name: this.form.name,
            }
            addAccount(account).then(response => {
                this.addAccountDialogVisible = false
                this.accounts.push({
                    wx_number: this.form.wxNum,
                    name: this.form.name,
                })
            })
        }
    },
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

