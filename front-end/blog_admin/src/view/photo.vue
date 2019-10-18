<template>
    <card>
        <div class="search-con search-con-top">
            <Button @click="modal_add=true" class="search-btn" type="primary">
                <Icon type="search"/>&nbsp;&nbsp;新增照片
            </Button>
        </div>
        <Table border :data="tableData" :columns="tableColumns" stripe :loading="loading"></Table>
        <Modal v-model="modal_update" scrollable :mask-closable="false" :closable="false" title="修改照片"
               @on-ok="updatePhoto" :loading="button_loading" @on-cancel="modalCancel">
            <Form ref="formCustom" :model="formCustom" :label-width="80">
                <FormItem label="轮播图图片">
                    <Spin size="large" fix v-if="spinShow"></Spin>
                    <div class="demo-upload-list" v-if="formCustom.image_url!=null">
                        <template>
                            <img :src="formCustom.image_url">
                            <div class="demo-upload-list-cover">
                                <Icon type="ios-eye-outline" @click.native="visible=true"></Icon>
                                <Icon type="ios-trash-outline" @click.native="formCustom.image_url=null"></Icon>
                            </div>
                        </template>
                    </div>
                    <Upload
                            ref="upload"
                            :show-upload-list="false"
                            :default-file-list="defaultList"
                            :on-success="handleSuccess"
                            :format="['jpg','jpeg','png','gif']"
                            :max-size="2048"
                            :on-format-error="handleFormatError"
                            :on-exceeded-size="handleMaxSize"
                            :before-upload="handleBeforeUpload"
                            :on-progress="handleFileProcess"
                            :multiple="false"
                            type="drag"
                            :action="upload_url"
                            :headers="uploadHeader"
                            style="display: inline-block;width:58px;">
                        <div style="width: 58px;height:58px;line-height: 58px;">
                            <Icon type="ios-camera" size="20"></Icon>
                        </div>
                    </Upload>
                    <Modal title="轮播图图片" v-model="visible">
                        <img :src="formCustom.image_url" v-if="visible" style="width: 100%">
                    </Modal>
                </FormItem>
            </Form>
        </Modal>
        <Modal v-model="modal_add" scrollable :mask-closable="false" :closable="false" title="新增照片" @on-ok="addPhoto"
               :loading="button_loading" @on-cancel="modalCancel">
            <Form ref="formCustom" :model="formCustom" :label-width="80">
                <FormItem label="图片">
                    <Spin size="large" fix v-if="spinShow"></Spin>
                    <div class="demo-upload-list" v-if="formCustom.image_url!=null">
                        <template>
                            <img :src="formCustom.image_url">
                            <div class="demo-upload-list-cover">
                                <Icon type="ios-eye-outline" @click.native="visible=true"></Icon>
                                <Icon type="ios-trash-outline" @click.native="formCustom.image_url=null"></Icon>
                            </div>
                        </template>
                    </div>
                    <Upload
                            ref="upload"
                            :show-upload-list="false"
                            :default-file-list="defaultList"
                            :on-success="handleSuccess"
                            :format="['jpg','jpeg','png','gif']"
                            :max-size="2048"
                            :on-format-error="handleFormatError"
                            :on-exceeded-size="handleMaxSize"
                            :before-upload="handleBeforeUpload"
                            :on-progress="handleFileProcess"
                            :multiple="false"
                            type="drag"
                            :action="upload_url"
                            :headers="uploadHeader"
                            style="display: inline-block;width:58px;">
                        <div style="width: 58px;height:58px;line-height: 58px;">
                            <Icon type="ios-camera" size="20"></Icon>
                        </div>
                    </Upload>
                    <Modal title="图片" v-model="visible">
                        <img :src="formCustom.image_url" v-if="visible" style="width: 100%">
                    </Modal>
                </FormItem>
            </Form>
        </Modal>
    </card>
</template>

<script>
    import config from '@/config'
    import {updatePhoto, addPhoto, deletePhoto, getPhotoList} from '../api/photo'
    import {getToken} from "../libs/util";

    const {baseUrl, imageUrl, uploadUrl} = config
    export default {
        name: 'photoPage',
        data() {
            return {
                uploadHeader:{Authorization:getToken()},
                formCustom: {
                    image_url: null,
                    id: ''
                },
                input_style: {
                    width: 200 + 'px'
                },
                spinShow: false,
                defaultList: [],
                modal_add: false,
                imgName: null,
                visible: false,
                uploadList: [],
                upload_url: '',
                modal_update: false,
                loading: false,
                button_loading: true,
                tableData: [],
                tableColumns: [
                    {
                        title: 'ID',
                        key: 'id',
                        render: (h, params) => {
                            const row = params.row;
                            return h('span', {
                                style: {
                                    width: "300px",
                                    verticalAlign: "middle",
                                    textAlign: 'center'
                                }
                            }, row.image.id)
                        }
                    },
                    {
                        title: '图片',
                        key: 'image_url',
                        render: (h, params) => {
                            const row = params.row;
                            return h('img', {
                                style: {
                                    width: "300px",
                                    verticalAlign: "middle",
                                    textAlign: 'center'
                                },
                                attrs: {
                                    src: row.image_url
                                }
                            },)
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
                                            const param = params.row;
                                            this.formCustom.image_url = param.image_url;
                                            this.formCustom.id = param.image.id;
                                            this.modal_update = true;
                                        }
                                    }
                                }, '编辑'),
                                h('Poptip', {
                                    props: {
                                        confirm: true,
                                        title: '你确定要删除吗?',
                                        transfer: true
                                    },
                                    on: {
                                        'on-ok': () => {
                                            let data = params.row;
                                            this.deletePhoto(data.image.id)
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
            modalCancel() {
                this.formCustom.image_url = null;
                this.formCustom.id = null;
                this.formCustom.words = null;
            },
            handleSuccess(res, file) {
                this.spinShow = false;
                if (res.code == 200) {
                    file.url = imageUrl + res.data;
                    file.name = res.data;
                    this.formCustom.image_url = imageUrl + res.data;
                } else {
                    this.$Message.error(res.msg)
                }
            },
            handleFormatError(file) {
                this.$Message.error(file.name + '格式错误，允许的格式有jpg、jpeg、png、gif');
            },
            handleMaxSize(file) {
                this.$Message.error('上传文件过大，最多2M！');
            },
            handleBeforeUpload() {
                this.uploadList.splice(0, 1);
                const check = this.uploadList.length < 2;
                if (!check) {
                    this.$Message.error('最多上传一张照片')
                }
                return check;
            },
            handleFileProcess() {
                this.spinShow = true;
            },
            updatePhoto() {
                if (this.formCustom.image_url == null) {
                    this.$Message.error('请上传图片!');
                    this.dealButtonLoading();
                    return;
                }
                updatePhoto(this.formCustom).then(res => {
                    let data = res.data;
                    if (data.code == 200) {
                        this.modal_update = false;
                        this.$Message.success(data.msg);
                        this.modalCancel();
                        this.getPhotoList();
                    } else {
                        this.$Message.error(data.msg);
                        this.dealButtonLoading();
                    }
                })
            },
            getPhotoList() {
                this.loading = true;
                getPhotoList().then(res => {
                    this.loading = false;
                    let data = res.data;
                    if (data.code == 200) {
                        this.tableData = data.data;
                    } else {
                        this.tableData = [];
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
            addPhoto() {
                if (this.formCustom.image_url == null) {
                    this.$Message.error('请上传图片!');
                    this.dealButtonLoading();
                    return;
                }
                addPhoto(this.formCustom).then(res => {
                    let data = res.data;
                    if (data.code == 200) {
                        this.modal_add = false;
                        this.$Message.success(data.msg);
                        this.modalCancel();
                        this.getPhotoList();
                    } else {
                        this.$Message.error(data.msg);
                        this.dealButtonLoading();
                    }
                })
            },
            deletePhoto(id) {
                this.loading = true;
                deletePhoto({id}).then(res => {
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg);
                        this.getPhotoList();
                    } else {
                        this.loading = false;
                        this.$Message.error(data.msg);
                    }
                })
            }
        },
        created() {
            this.getPhotoList();
            this.upload_url = uploadUrl
        },
        mounted() {
            this.uploadList = this.$refs.upload.fileList;
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

    .demo-upload-list {
        display: inline-block;
        width: 60px;
        height: 60px;
        text-align: center;
        line-height: 60px;
        border: 1px solid transparent;
        border-radius: 4px;
        overflow: hidden;
        background: #fff;
        position: relative;
        box-shadow: 0 1px 1px rgba(0, 0, 0, .2);
        margin-right: 4px;
    }

    .demo-upload-list img {
        width: 100%;
        height: 100%;
    }

    .demo-upload-list-cover {
        display: none;
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        background: rgba(0, 0, 0, .6);
    }

    .demo-upload-list:hover .demo-upload-list-cover {
        display: block;
    }

    .demo-upload-list-cover i {
        color: #fff;
        font-size: 20px;
        cursor: pointer;
        margin: 0 2px;
    }
</style>
