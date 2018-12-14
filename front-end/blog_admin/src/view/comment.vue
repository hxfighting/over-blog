<template>
    <div>
        <card>
            <div class="search-con search-con-top">
                <Select v-model="listData.article_id" style="width:300px" clearable filterable placeholder="请选择文章">
                    <Option v-for="item in article" :value="item.id" :key="item.id">{{ item.title }}</Option>
                </Select>
                <Button @click="handleSearch" class="search-btn" type="primary">
                    <Icon type="search"/>&nbsp;&nbsp;搜索
                </Button>
            </div>
            <Table :data="tableData" :columns="tableColumns" :loading="loading" stripe></Table>
            <div style="margin: 10px;overflow: hidden">
                <div style="float: right;">
                    <Page :total="total" show-total :page-size="listData.pageSize" @on-change="changePage"></Page>
                </div>
            </div>
            <Modal
                    v-model="replyCommentModal"
                    title="回复评论"
                    :loading="button_loading"
                    @on-ok="replyComment"
                    @on-cancel="formData.reply_content=''">
                <Input v-model="formData.reply_content" type="textarea" :rows="6" placeholder="请输入回复内容"></Input>
            </Modal>
        </card>

    </div>
</template>

<script>
    import {getCommentList, deleteComment,replyComment} from '../api/comment'

    export default {
        name: 'commentPage',
        data() {
            return {
                tableData: [],
                loading: false,
                article: [],
                replyCommentModal:false,
                button_loading:true,
                total: 0,
                listData: {
                    article_id: '',
                    pageSize: 10,
                    pageNum: 1,
                },
                formData:{
                  id:'',
                  reply_content:''
                },
                tableColumns: [
                    {
                        title: 'ID',
                        key: 'id'
                    },
                    {
                        title: '评论人',
                        key: 'user',
                        tooltip: true,
                        render: (h, params) => {
                            const row = params.row;
                            const text = row.user.name;
                            return h('span', {}, text);
                        }
                    },
                    {
                        title: '被评论人',
                        key: 'reply',
                        tooltip: true,
                        render: (h, params) => {
                            const row = params.row;
                            const text = row.replier == null ? '无' : row.replier.name;
                            return h('span', {}, text);
                        }
                    },
                    {
                        title: '评论内容',
                        key: 'content',
                        tooltip: true
                    },
                    {
                        title: '评论时间',
                        key: 'created_at',
                        sortable: true,
                        tooltip: true
                    },
                    {
                        title: '操作',
                        key: 'action',
                        width: 150,
                        align: 'center',
                        render: (h, params) => {
                            let data = params.row;
                            return h('div', [
                                h('Button', {
                                    props: {
                                        type: 'primary',
                                        size: 'small'
                                    },
                                    style: {
                                        marginRight: '5px'
                                    },
                                    on: {
                                        click: () => {
                                            this.formData.id = data.id;
                                            this.replyCommentModal = true;
                                        }
                                    }
                                }, '回复'),
                                h('Poptip', {
                                    props: {
                                        confirm: true,
                                        title: '你确定要删除吗?'
                                    },
                                    on: {
                                        'on-ok': () => {
                                            this.deleteComment(data.id)
                                        }
                                    }
                                }, [
                                    h('Button', {
                                        props: {
                                            type: 'error',
                                            size: 'small'
                                        }
                                    }, '删除')
                                ])
                            ]);
                        }
                    }
                ]
            }
        },
        methods: {
            changePage(page) {
                this.listData.pageNum = page;
                this.getCommentList();
            },
            handleSearch() {
                this.getCommentList();
            },
            //获取评论列表
            getCommentList() {
                this.loading = true;
                getCommentList(this.listData).then(res => {
                    this.loading = false;
                    let data = res.data;
                    if (data.code == 200) {
                        this.tableData = data.data.list;
                        this.total = data.data.total;
                        this.article = data.data.article;
                    } else {
                        this.tableData = [];
                        this.total = 0;
                        this.$Message.error(data.msg);
                    }
                })
            },
            //处理button的loading状态
            dealButtonLoading() {
                let _this = this
                setTimeout(function () {
                    _this.button_loading = false;
                    _this.$nextTick(() => {
                        _this.button_loading = true;
                    });
                }, 500);
            },
            //删除评论
            deleteComment(id){
                this.loading = true;
                deleteComment({id}).then(res => {
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg)
                        this.getCommentList();
                    } else {
                        this.loading = false;
                        this.$Message.error(data.msg);
                    }
                })
            },
            //回复评论
            replyComment(){
                if(this.formData.reply_content==''){
                    this.$Message.error('请输入回复内容!');
                    this.dealButtonLoading();
                    return;
                }
                replyComment(this.formData).then(res=>{
                    let data = res.data;
                    if (data.code == 200) {
                        this.replyCommentModal = false;
                        this.formData.reply_content = '';
                        this.$Message.success(data.msg);
                        this.getCommentList();
                    } else {
                        this.dealButtonLoading();
                        this.$Message.error(data.msg);
                    }
                })
            }
        },
        created() {
            this.getCommentList()
        }
    }
</script>

<style scoped>

    .search-con {
        padding: 10px 0;
    }

    .search-con button {
        margin-left: 5px;
    }
</style>
