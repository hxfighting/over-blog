<template>
    <div>
        <card>
            <div class="search-con search-con-top">
                <Button @click="chat_add_modal=true" class="search-btn" type="primary">
                    <Icon type="search"/>&nbsp;&nbsp;新增说说
                </Button>
            </div>
            <Timeline>
                <TimelineItem v-for="item in chatData">
                    <p class="time">
                        <Tooltip content="不显示" v-if="item.is_show==0" placement="top">
                            <Icon type="md-eye-off"/>
                        </Tooltip>
                        <Tooltip content="显示" v-else placement="top">
                            <Icon type="md-eye"/>
                        </Tooltip>
                        &nbsp;{{item.created_at}}
                    </p>
                    <p class="content">{{item.content}}</p>
                    <p class="content">
                        <Button type="primary" size="small" @click="showUpdateModal(item)">编辑</Button>
                        <Poptip
                                confirm
                                title="你确定删除这条说说吗?"
                                @on-ok="deleteChat(item)">
                            <Button type="error" size="small">删除</Button>
                        </Poptip>
                    </p>
                </TimelineItem>
                <TimelineItem><span @click="getMore" style="cursor: pointer">查看更多</span></TimelineItem>
            </Timeline>
            <Spin fix v-if="loading"></Spin>
            <Modal
                    v-model="modalShow"
                    title="修改说说"
                    @on-ok="updateChat('formInline')"
                    :loading="button_loading"
                    @on-cancel="modalCancel('formInline')"
            >
                <Form ref="formInline" :model="chat" :rules="ruleValidate">
                    <FormItem prop="content" label="说说内容">
                        <Input type="textarea" v-model="chat.content" placeholder="请输入说说内容" :rows="6">
                            <Icon type="ios-person-outline" slot="prepend"></Icon>
                        </Input>
                    </FormItem>
                    <FormItem>
                        <i-switch v-model="is_show_bool" size="large">
                            <span slot="open">显示</span>
                            <span slot="close">不显示</span>
                        </i-switch>
                    </FormItem>
                </Form>
            </Modal>
            <Modal
                    v-model="chat_add_modal"
                    title="新增说说"
                    @on-ok="addChat('formInline')"
                    :loading="button_loading"
                    @on-cancel="modalCancel('formInline')"
            >
                <Form ref="formInline" :model="chat" :rules="ruleValidate">
                    <FormItem prop="content" label="说说内容">
                        <Input type="textarea" v-model="chat.content" placeholder="请输入说说内容" :rows="6">
                            <Icon type="ios-person-outline" slot="prepend"></Icon>
                        </Input>
                    </FormItem>
                    <FormItem>
                        <i-switch v-model="is_show_bool" size="large">
                            <span slot="open">显示</span>
                            <span slot="close">不显示</span>
                        </i-switch>
                    </FormItem>
                </Form>
            </Modal>

        </card>
    </div>
</template>

<script>
    import {getChatList, addChat, updateChat, deleteChat} from '../api/chat'

    export default {
        name: 'chatPage',
        data() {
            return {
                chatData: [],
                is_show_bool: true,
                update_loading: false,
                chat_add_modal: false,
                add_loading: false,
                button_loading:true,
                chat: {
                    content: '',
                    is_show: 1,
                    id: ''
                },
                loading: false,
                modalShow: false,
                page: {
                    pageNum: 1,
                    pageSize: 6
                },
                ruleValidate: {
                    content: [
                        {required: true, message: '请输入说说内容', trigger: 'blur'},
                        {type: 'string', min: 2, max: 255, message: '说说内容在2到255个字符之间', trigger: 'change'},
                    ]
                }
            }
        },
        methods: {
            modalCancel(name) {
                this.$refs[name].resetFields();
            },
            getChat() {
                this.loading = true;
                getChatList(this.page).then(res => {
                    this.loading = false;
                    let data = res.data;
                    if (data.code == 200) {
                        this.chatData = data.data;
                    } else {
                        this.$Message.error(data.msg)
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
            getMore() {
                this.page.pageSize += 3;
                this.getChat();
            },
            showUpdateModal(item) {
                this.modalShow = true;
                this.chat.id = item.id;
                this.chat.content = item.content;
                this.is_show_bool = item.is_show == 1 ? true : false;
            },
            updateChat(name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        this.chat.is_show = this.is_show_bool ? 1 : 0;
                        updateChat(this.chat).then(res => {
                            this.modalShow = false;
                            let data = res.data;
                            if (data.code == 200) {
                                this.update_loading = false;
                                this.$Message.success(data.msg);
                                this.chat.content = '';
                                this.is_show_bool = true;
                                this.getChat()
                            } else {
                                this.dealButtonLoading();
                                this.$Message.error(data.msg)
                            }
                        })
                    } else {
                        this.dealButtonLoading();
                    }
                })
            },
            addChat(name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        this.chat.is_show = this.is_show_bool ? 1 : 0;
                        addChat(this.chat).then(res => {
                            this.chat_add_modal = false;
                            let data = res.data;
                            if (data.code == 200) {
                                this.$Message.success(data.msg);
                                this.chat.content = '';
                                this.is_show_bool = true;
                                this.getChat()
                            } else {
                                this.dealButtonLoading();
                                this.$Message.error(data.msg)
                            }
                        })
                    } else {
                        this.dealButtonLoading();
                    }
                })
            },
            deleteChat(item) {
                this.loading = true;
                deleteChat({id: item.id}).then(res => {
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg);
                        this.getChat()
                    } else {
                        this.$Message.error(data.msg)
                    }
                })
            }
        },
        created() {
            this.getChat()
        }
    }
</script>

<style type="text/less" scoped>
    .time {
        font-size: 14px;
        font-weight: bold;
    }

    .content {
        padding-left: 5px;
    }

    .content button {
        margin: 2px;
    }

    .search-con {
        padding: 10px 0;

        .search {
            &-col {
                display: inline-block;
                width: 200px;
            }

            &-input {
                display: inline-block;
                width: 200px;
                margin-left: 2px;
            }

            &-btn {
                margin-left: 2px;
            }
        }
    }
</style>