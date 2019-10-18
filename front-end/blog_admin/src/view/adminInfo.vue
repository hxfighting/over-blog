<template>
    <div>
        <card>
            <Form ref="formCustom" :model="formCustom" :rules="ruleCustom" :label-width="80">
                <FormItem label="头像">
                    <div class="demo-upload-list" v-for="item in uploadList">
                    <template v-if="item.status === 'finished'">
                        <img :src="item.url">
                        <div class="demo-upload-list-cover">
                            <Icon type="ios-eye-outline" @click.native="handleView(item.name)"></Icon>
                            <Icon type="ios-trash-outline" @click.native="handleRemove(item)"></Icon>
                        </div>
                    </template>
                    <template v-else>
                        <Progress v-if="item.showProgress" :percent="item.percentage" hide-info></Progress>
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
                            :multiple="false"
                            type="drag"
                            :action="upload_url"
                            :headers="uploadHeader"
                            style="display: inline-block;width:58px;">
                        <div style="width: 58px;height:58px;line-height: 58px;">
                            <Icon type="ios-camera" size="20"></Icon>
                        </div>
                    </Upload>
                    <Modal title="头像" v-model="visible">
                        <img :src="imgName" v-if="visible" style="width: 100%">
                    </Modal>
                </FormItem>
                <FormItem label="姓名" prop="name">
                    <Input v-model="formCustom.name" :style="input_style"></Input>
                </FormItem>
                <FormItem label="邮箱" prop="email">
                    <Input v-model="formCustom.email" :style="input_style"></Input>
                </FormItem>
                <FormItem label="电话" prop="mobile">
                    <Input type="text" v-model="formCustom.mobile" :style="input_style"></Input>
                </FormItem>
                <FormItem>
                    <Button type="primary" @click="handleSubmit('formCustom')">修改</Button>
                </FormItem>
            </Form>
        </card>
    </div>
</template>

<script>
    import {getToken} from "../libs/util";
    import config from '@/config'
    import {updateInfo} from '../api/user'
    const { baseUrl,imageUrl,uploadUrl } = config
    export default {
        name: 'adminInfoPage',
        data() {
            const validateName = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请输入用户名'))
                } else {
                    callback()
                }
            }
            const validateEmail = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请输入邮箱'))
                } else {
                    callback()
                }
            }
            const validateMobile = (rule, value, callback) => {
                if (value=='') {
                    callback(new Error('请输入电话号码'))
                }else {
                    callback()
                }
            }

            return {
                uploadHeader:{Authorization:getToken()},
                formCustom: {
                    name: '',
                    mobile: '',
                    email: '',
                    avatar: ''
                },
                ruleCustom: {
                    name: [
                        {validator: validateName, trigger: 'blur'}
                    ],
                    email: [
                        {validator: validateEmail, trigger: 'blur'}
                    ],
                    mobile: [
                        {validator: validateMobile, trigger: 'blur'}
                    ]
                },
                input_style: {
                    width: 200 + 'px'
                },
                defaultList: [],
                imgName: '',
                visible: false,
                uploadList: [],
                upload_url:''
            }
        },
        methods: {
            handleSubmit(name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        updateInfo(this.formCustom).then(res=>{
                            let data = res.data;
                            if(data.code==200){
                                this.$Message.success(data.msg);
                                location.reload();
                            }else{
                                this.$Message.error(data.msg);
                            }
                        })
                    } else {
                        this.$Message.error('验证失败!')
                    }
                })
            },
            handleView (name) {
                this.imgName = imageUrl+name;
                this.visible = true;
            },
            handleRemove (file) {
                const fileList = this.$refs.upload.fileList;
                this.$refs.upload.fileList.splice(fileList.indexOf(file), 1);
            },
            handleSuccess (res, file) {
                if(res.code==200){
                    file.url = imageUrl+res.data;
                    file.name = res.data;
                    this.formCustom.avatar = imageUrl+res.data;
                }else {
                    this.$Message.error(res.msg)
                }
            },
            handleFormatError (file) {
                this.$Message.error(file.name+'格式错误，允许的格式有jpg、jpeg、png、gif');
            },
            handleMaxSize (file) {
                this.$Message.error('上传文件过大，最多2M！');
            },
            handleBeforeUpload () {
                this.uploadList.splice(0, 1);
                const check = this.uploadList.length < 2;
                if (!check) {
                    this.$Message.error('最多上传一张照片')
                }
                return check;
            }
        },
        created(){
            let avatar = this.$store.state.user.avatorImgPath,
                name = this.$store.state.user.avatarName;
            this.formCustom.avatar = avatar;
            this.formCustom.name = this.$store.state.user.userName;
            this.formCustom.email = this.$store.state.user.email;
            this.formCustom.mobile = this.$store.state.user.mobile;
            this.defaultList.push({name:name,url:avatar});
            this.upload_url = uploadUrl
        },
        mounted () {
            this.uploadList = this.$refs.upload.fileList;
        }
    }
</script>

<style scoped>
    .demo-upload-list{
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
        box-shadow: 0 1px 1px rgba(0,0,0,.2);
        margin-right: 4px;
    }
    .demo-upload-list img{
        width: 100%;
        height: 100%;
    }
    .demo-upload-list-cover{
        display: none;
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        background: rgba(0,0,0,.6);
    }
    .demo-upload-list:hover .demo-upload-list-cover{
        display: block;
    }
    .demo-upload-list-cover i{
        color: #fff;
        font-size: 20px;
        cursor: pointer;
        margin: 0 2px;
    }
</style>
