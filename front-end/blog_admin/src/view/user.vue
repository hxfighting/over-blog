<template>
    <div>
        <card>
            <div class="search-con search-con-top">
                <Input @on-change="handleClear" clearable placeholder="输入姓名或email搜索" class="hx_input"
                       v-model="listData.search"/>
                <Button @click="handleSearch" class="search-btn" type="primary">
                    <Icon type="search"/>&nbsp;&nbsp;搜索
                </Button>
            </div>
            <Table :data="tableData" :columns="tableColumns1" stripe :loading="loading"></Table>
            <div style="margin: 10px;overflow: hidden">
                <div style="float: right;">
                    <Page :total="total" show-total :page-size="listData.pageSize" @on-change="changePage"></Page>
                </div>
            </div>
        </card>
    </div>
</template>

<script>
    import {getUserList,updateUser,deleteUser} from '../api/user'
    export default {
        name: 'userPage',
        data() {
            return {
                tableData: [],
                searchValue: '',
                loading: false,
                total: 0,
                modal_add: false,
                modal_update: false,
                button_loading: true,
                contact_reply: false,
                reply_user_data: '',
                listData: {
                    search: '',
                    pageSize: 10,
                    pageNum: 1,
                },
                formData: {
                    id: '',
                    is_admin: '',
                },
                tableColumns1: [
                    {
                        title: 'ID',
                        key: 'id'
                    },
                    {
                        title: '头像',
                        key: 'avatar',
                        render: (h, params) => {
                            const row = params.row;
                            return h('Avatar', {
                                props: {
                                    src: row.avatar
                                }
                            })
                        }
                    },
                    {
                        title: '姓名',
                        key: 'name'
                    },
                    {
                        title: '登录类型',
                        key: 'type',
                        render: (h, params) => {
                            const row = params.row;
                            const color = row.type === 1 ? 'primary' : row.type === 2 ? 'success' : 'error';
                            const text = row.type === 1 ? 'QQ' : row.type === 2 ? '微信' : '微博';
                            return h('Tooltip', {
                                props: {
                                    content: text,
                                    placement:'top'
                                }
                            }, [
                                h('Tag', {
                                    props: {
                                        type: 'dot',
                                        color: color
                                    }
                                }, text)
                            ]);
                        }
                    },
                    {
                        title: '邮箱',
                        key: 'email',
                        render: (h, params) => {
                            const row = params.row;
                            const text = row.email != null ? row.email : '暂未绑定邮箱';
                            return h('Span', {
                            }, text)
                        }
                    },
                    {
                        title: '登录次数',
                        key: 'login_times'
                    },
                    {
                        title: '最后登录地址',
                        key: 'last_login_ip'
                    },
                    {
                        title: '是否是管理员',
                        key: 'is_admin',
                        render: (h, params) => {
                            const row = params.row;
                            const color = row.is_admin === 1 ? 'success' : 'error';
                            const text = row.is_admin === 1 ? '√' : '×';
                            return h('Tag', {
                                props: {
                                    color: color
                                }
                            }, text)
                        }
                    },
                    {
                        title: '创建时间',
                        key: 'created_at'
                    },
                    {
                        title: '操作',
                        key: 'action',
                        width: 150,
                        align: 'center',
                        render: (h, params) => {
                            let data = params.row;
                            let admin_text = data.is_admin === 1 ? '你确定撤销管理员吗?' : '你确定加为管理员吗?';
                            let btns = [
                                h('Poptip', {
                                    props: {
                                        confirm: true,
                                        title: admin_text
                                    },
                                    style: {
                                        marginRight: '5px'
                                    },
                                    on: {
                                        'on-ok': () => {
                                            this.formData.is_admin = data.is_admin=== 1 ? 0 : 1;
                                            this.formData.id = data.id;
                                            this.updateUser();
                                        }
                                    }
                                }, [
                                    h('Button', {
                                        props: {
                                            type: 'primary',
                                            size: 'small'
                                        }
                                    }, '编辑')
                                ]),
                                h('Poptip', {
                                    props: {
                                        confirm: true,
                                        title: '你确定要删除吗?'
                                    },
                                    on: {
                                        'on-ok': () => {
                                            this.deleteUser(data.id);
                                        }
                                    }
                                }, [
                                    h('Button', {
                                        props: {
                                            type: 'error',
                                            size: 'small'
                                        }
                                    }, '删除')
                                ])];
                            return h('div', btns);
                        }
                    }
                ]
            }
        },
        methods: {
            changePage(page) {
                this.listData.pageNum = page;
                this.getUserList();
            },
            handleClear(e) {
                if (e.target.value === '') this.insideTableData = this.value
            },
            handleSearch() {
                this.getUserList();
            },
            getUserList() {
                this.loading = true;
                getUserList(this.listData).then(res => {
                    this.loading = false;
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg);
                        this.tableData = data.data.list;
                        this.total = data.data.total;
                    } else {
                        this.$Message.error(data.msg);
                    }
                })
            },
            deleteUser(id) {
                this.loading = true;
                deleteUser({id}).then(res => {
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg);
                        this.getUserList();
                    } else {
                        this.loading = false;
                        this.$Message.error(data.msg);
                    }
                })
            },
            updateUser(){
                this.loading = true;
                updateUser(this.formData).then(res=>{
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg);
                        this.getUserList();
                    } else {
                        this.loading = false;
                        this.$Message.error(data.msg);
                    }
                })
            }
        },
        created() {
            this.getUserList()
        }
    }
</script>

<style scoped>
    .hx_input {
        display: inline-block;
        width: 200px !important;
        margin-left: 2px;
        margin-right: 3px;
    }

    .search-con {
        padding: 10px 0;
    }

    .link_add_button {
        float: right;
    }
</style>
