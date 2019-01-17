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
            <Modal
                    v-model="contact_reply"
                    title="回复"
                    :loading="button_loading"
                    scrollable
                    :mask-closable="false"
                    :closable="false"
                    @on-ok="replyToUser('formCustomTwo')"
                    @on-cancel="modalCancel('formCustomTwo')">
                <Form ref="formCustomTwo" :model="formData" :rules="ruleValidate">
                    <FormItem prop="reply_content">
                        <Input type="textarea" placeholder="请输入回复内容..." :rows="6" v-model="formData.reply_content"/>
                    </FormItem>
                </Form>
            </Modal>
        </card>
    </div>
</template>

<script>
    import {getContactList, deleteContact, replyContact} from '../api/contact'

    export default {
        name: 'contactPage',
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
                    reply_content: '',
                },
                tableColumns1: [
                    {
                        title: 'ID',
                        key: 'id'
                    },
                    {
                        title: '姓名',
                        key: 'name'
                    },
                    {
                        title: '邮箱',
                        key: 'email'
                    },
                    {
                        title: '内容',
                        key: 'content'
                    },
                    {
                        title: '留言时间',
                        key: 'created_at'
                    },
                    {
                        title: '是否回复',
                        key: 'status',
                        render: (h, params) => {
                            const row = params.row;
                            const color = row.is_reply === 1 ? 'success' : 'error';
                            const text = row.is_reply === 1 ? '√' : '×';
                            const content = row.is_reply === 1 ? '回复内容: ' + row.reply_content : '还没有回复';
                            return h('Tooltip', {
                                props: {
                                    content: content,
                                    placement: 'top'
                                }
                            }, [
                                h('Tag', {
                                    props: {
                                        color: color
                                    }
                                }, text)
                            ])
                        }
                    },
                    {
                        title: '操作',
                        key: 'action',
                        width: 150,
                        align: 'center',
                        render: (h, params) => {
                            let btns = [
                                h('Poptip', {
                                    props: {
                                        confirm: true,
                                        title: '你确定要删除吗?'
                                    },
                                    on: {
                                        'on-ok': () => {
                                            let data = params.row;
                                            this.deleteContact(data.id)
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
                            if (params.row.is_reply == 0) {
                                btns.unshift(
                                    h('Button', {
                                        props: {
                                            type: 'primary',
                                            size: 'small'
                                        },
                                        attrs: {
                                            title: '回复',
                                        },
                                        style: {
                                            marginRight: '5px'
                                        },
                                        on: {
                                            click: () => {
                                                let data = params.row;
                                                this.formData.id = data.id;
                                                this.contact_reply = true;
                                            }
                                        }
                                    }, '回复'),
                                );
                            }
                            return h('div', btns);
                        }
                    }
                ],
                ruleValidate: {
                    reply_content: [
                        {required: true, message: '请输入回复内容', trigger: 'blur'},
                        {type: 'string', min: 2, max: 255, message: '回复内容在2到255个字符之间', trigger: 'change'},
                    ],
                }
            }
        },
        methods: {
            modalCancel(name) {
                this.$refs[name].resetFields();
            },
            changePage(page) {
                this.listData.pageNum = page;
                this.getContactList();
            },
            handleClear(e) {
                if (e.target.value === '') this.insideTableData = this.value
            },
            handleSearch() {
                this.listData.pageNum = 1;
                this.getContactList();
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
            getContactList() {
                this.loading = true;
                getContactList(this.listData).then(res => {
                    this.loading = false;
                    let data = res.data;
                    if (data.code == 200) {
                        this.tableData = data.data.list;
                        this.total = data.data.total;
                    } else {
                        this.$Message.error(data.msg);
                    }
                })
            },
            deleteContact(id) {
                this.loading = true;
                deleteContact({id}).then(res => {
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg);
                        this.getContactList();
                    } else {
                        this.loading = false;
                        this.$Message.error(data.msg);
                    }
                })
            },
            replyToUser(name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        replyContact(this.formData).then(res => {
                            let data = res.data;
                            if (data.code == 200) {
                                this.contact_reply = false;
                                this.modalCancel('formCustomTwo');
                                this.$Message.success(data.msg);
                                this.getContactList();
                            } else {
                                this.dealButtonLoading();
                                this.$Message.error(data.msg);
                            }
                        });
                    } else {
                        this.dealButtonLoading();
                    }
                })
            }
        },
        created() {
            this.getContactList()
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
