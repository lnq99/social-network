<template>
  <center class="login">
    <card class="login-card box-shadow p18">
      <h2>Login</h2>
      <el-form
        :model="model"
        :rules="rules"
        ref="form"
        @submit.native.prevent="onLogin"
      >
        <el-form-item prop="email">
          <el-input
            v-model="model.email"
            placeholder="email"
            prefix-icon="el-icon-message"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="model.password"
            placeholder="Password"
            type="password"
            prefix-icon="el-icon-lock"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button
            :loading="loading"
            class="login-button"
            type="primary"
            native-type="submit"
            block
            >Login</el-button
          >
        </el-form-item>
        <a href="https://">Forgot password?</a>
      </el-form>
    </card>
    <hr />
    <card class="register-card box-shadow p18">
      <router-link to="/signup">Don't have an account? Sign up</router-link>
    </card>
  </center>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  name: 'login',
  data() {
    return {
      model: {
        email: 'test@gmail.com',
        password: 'easypass',
      },
      loading: false,
      rules: {
        email: [
          { required: true, message: 'Email is required', trigger: 'blur' },
          { type: 'email', message: 'Email address invalid!', trigger: 'blur' },
        ],
        password: [
          { required: true, message: 'Password is required', trigger: 'blur' },
          {
            min: 8,
            message: 'Password must be at least 8 characters',
            trigger: 'blur',
          },
        ],
      },
    }
  },
  computed: {
    ...mapGetters(['isAuthenticated'])
  },
  methods: {
    ...mapActions(['login']),
    onLogin() {
      let valid = this.$refs.form.validate()
      if (!valid) {
        return
      }
      this.loading = true
      this.loginHandle(this.model)
    },
    loginHandle(model) {
      this.login(model).then(() => {
        this.loading = false
        if (this.isAuthenticated) {
          this.$message.success('Login successfull')
          this.$router.push({ name: 'NewsFeed' })
        } else if (model) {
          this.$message.error('Email or password is invalid')
        }
      })
    }
  },
  created() {
    this.loginHandle()
  }
}
</script>

<style scoped>
.login .card {
  text-align: center;
  width: 300px;
  background: var(--hl);
}
.login-button {
  width: 100%;
  margin-top: 30px;
}
h2 {
  letter-spacing: 1px;
  padding-bottom: 20px;
}
hr {
  margin-top: 20px;
  margin-bottom: 20px;
}
</style>
