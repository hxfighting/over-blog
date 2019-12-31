<template>
    <div>
        <card>
            <Form ref="formCustom" :model="formCustom" :rules="ruleCustom" :label-width="80">
                <FormItem label="新密码" prop="password">
                    <Input type="password" v-model="formCustom.password" :style="input_style"></Input>
                </FormItem>
                <FormItem label="再次输入" prop="password_confirmation">
                    <Input type="password" v-model="formCustom.password_confirmation" :style="input_style"></Input>
                </FormItem>
                <FormItem>
                    <Button type="primary" @click="handleSubmit('formCustom')">修改</Button>
                    <Button @click="handleReset('formCustom')" style="margin-left: 8px">重置</Button>
                </FormItem>
            </Form>
        </card>
    </div>
</template>

<script>
    import {updatePassword} from '../api/user'
    export default {
        name: "resetPasswordPage",
        data() {
            const validatePass = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请输入密码'))
                } else {
                    if (this.formCustom.password_confirmation !== '') {
                        // 对第二个密码框单独验证
                        this.$refs.formCustom.validateField('password_confirmation')
                    }
                    callback()
                }
            }
            const validatePassCheck = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请再次输入密码'))
                } else if (value !== this.formCustom.password) {
                    callback(new Error('两次密码输入不一致!'))
                } else {
                    callback()
                }
            }

            return {
                formCustom: {
                    password: '',
                    password_confirmation: ''
                },
                ruleCustom: {
                    password: [
                        {validator: validatePass, trigger: 'blur'}
                    ],
                    password_confirmation: [
                        {validator: validatePassCheck, trigger: 'blur'}
                    ]
                },
                input_style: {
                    width: 200 + 'px'
                },
            }
        },
        methods:{
            handleSubmit(name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        updatePassword(this.formCustom).then(res=>{
                            let data = res.data;
                            if(data.code==200){
                                this.$Message.success(data.msg);
                                this.handleReset('formCustom')
                            }else {
                                this.$Message.error(data.msg);
                            }
                        })
                    } else {
                        this.$Message.error('验证失败!')
                    }
                })
            },
            handleReset(name) {
                this.$refs[name].resetFields()
            },
        }
    }
</script>

<style scoped>

</style>